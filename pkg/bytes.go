package pkg

import (
	"fmt"
	"os"
)

type Prefixes struct {
	bytesShort     string
	kilobytesShort string
	megaBytesShort string
	gigabytesShort string
	bytesLong      string
	kilobytesLong  string
	gigabytesLong  string
	megaBytesLong  string
}

type Bytes struct {
	prefixes  Prefixes
	bytes     float32
	kilobytes float32
	megabytes float32
	gigabytes float32
}

func (d *Bytes) calculateBytesAndPrefixes() error {
	bytes := float64(d.bytes)
	if bytes == 0 {
		return fmt.Errorf("no bytes provided")
	}

	d.kilobytes = float32(bytes / 1024)
	d.megabytes = float32(bytes / (1024 * 1024))
	d.gigabytes = float32(bytes / (1024 * 1024 * 1024))

	prefixes := Prefixes{
		bytesShort:     "B: ",
		kilobytesShort: "KB: ",
		megaBytesShort: "MB: ",
		gigabytesShort: "GB: ",
		bytesLong:      "Bytes: ",
		kilobytesLong:  "kilobytes: ",
		megaBytesLong:  "Megabytes: ",
		gigabytesLong:  "Gigabytes: ",
	}
	d.prefixes = prefixes

	return nil
}

func GetFilesSize(totalMemory *Bytes, memoryPerFile map[string]Bytes, files ...string) error {

	if len(ValidateFiles(files...)) > 0 {
		return fmt.Errorf("failed to validate files")
	}

	for _, f := range files {
		stat, _ := os.Stat(f)
		bytes := stat.Size()
		totalMemory.bytes += float32(bytes)
		perFile := Bytes{bytes: float32(bytes)}
		err := perFile.calculateBytesAndPrefixes()
		if err != nil {
			return fmt.Errorf("failed to calculate bytes: %q", err)
		}
		memoryPerFile[stat.Name()] = perFile
	}

	return nil
}
