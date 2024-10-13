package pkg

import (
	"fmt"
	"os"
)

type TotalMemory struct {
	TotalBytes    Bytes
	MemoryPerFile []Bytes
}

type Prefixes struct {
	BytesShort     string
	KilobytesShort string
	MegaBytesShort string
	GigabytesShort string
	BytesLong      string
	KilobytesLong  string
	GigabytesLong  string
	MegaBytesLong  string
}

type Bytes struct {
	Prefixes  Prefixes
	Bytes     float32
	Kilobytes float32
	Megabytes float32
	Gigabytes float32
}

type DisplayBytes struct {
	Bytes     string
	Kilobytes string
	Megabytes string
	Gigabytes string
}

func (d *Bytes) CalculateBytesAndPrefixes() error {
	bytes := float64(d.Bytes)
	if bytes == 0 {
		return fmt.Errorf("no bytes provided")
	}

	d.Kilobytes = float32(bytes / 1024)
	d.Megabytes = float32(bytes / (1024 * 1024))
	d.Gigabytes = float32(bytes / (1024 * 1024 * 1024))

	prefixes := Prefixes{
		BytesShort:     "B: ",
		KilobytesShort: "KB: ",
		MegaBytesShort: "MB: ",
		GigabytesShort: "GB: ",
		BytesLong:      "Bytes: ",
		KilobytesLong:  "kilobytes: ",
		MegaBytesLong:  "Megabytes: ",
		GigabytesLong:  "Gigabytes: ",
	}
	d.Prefixes = prefixes

	return nil
}

func (d *Bytes) DisplayText(long bool) (DisplayBytes, error) {
	if d.Bytes == float32(0) {
		return DisplayBytes{}, fmt.Errorf("there are no bytes to display: %f", d.Bytes)
	}
	displayText := DisplayBytes{}
	switch long {
	case true:
		displayText.Bytes = fmt.Sprintf("%s%d", d.Prefixes.BytesLong, int(d.Bytes))
		displayText.Bytes = fmt.Sprintf("%s%.4f", d.Prefixes.KilobytesLong, d.Kilobytes)
		displayText.Bytes = fmt.Sprintf("%s%.4f", d.Prefixes.MegaBytesLong, d.Megabytes)
		displayText.Bytes = fmt.Sprintf("%s%.4f", d.Prefixes.GigabytesLong, d.Gigabytes)
	default:
		displayText.Bytes = fmt.Sprintf("%s%d", d.Prefixes.BytesShort, int(d.Bytes))
		displayText.Bytes = fmt.Sprintf("%s%.4f", d.Prefixes.KilobytesShort, d.Kilobytes)
		displayText.Bytes = fmt.Sprintf("%s%.4f", d.Prefixes.MegaBytesShort, d.Megabytes)
		displayText.Bytes = fmt.Sprintf("%s%.4f", d.Prefixes.GigabytesShort, d.Gigabytes)
	}

	return displayText, nil
}

func GetFilesSize(files ...string) (TotalMemory, error) {
	totalMemory := TotalMemory{}
	totalMemory.TotalBytes = Bytes{}
	if len(ValidateFiles(files...)) > 0 {
		return totalMemory, fmt.Errorf("failed to validate files")
	}

	for _, f := range files {
		stat, _ := os.Stat(f)
		bytes := stat.Size()
		totalMemory.TotalBytes.Bytes += float32(bytes)
		perFileMemory := Bytes{Bytes: float32(bytes)}
		err := perFileMemory.CalculateBytesAndPrefixes()
		if err != nil {
			return totalMemory, fmt.Errorf("failed to calculate bytes: %q", err)
		}
		totalMemory.MemoryPerFile = append(totalMemory.MemoryPerFile, perFileMemory)
	}
	err := totalMemory.TotalBytes.CalculateBytesAndPrefixes()
	if err != nil {
		return totalMemory, fmt.Errorf("failed to calculate bytes: %q", err)
	}

	return totalMemory, nil
}
