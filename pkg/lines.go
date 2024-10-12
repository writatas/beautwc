package pkg

import (
	"bufio"
	"fmt"
	"os"
)

type LinesCount struct {
	TotalLines   int
	PerFileLines []Lines
}

type Lines struct {
	Filename string
	Count    int
}

func (l *LinesCount) GetLinesCount(files ...string) error {
	errs := ValidateFiles(files...)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
		return fmt.Errorf("failed to vailidate files: %#v", errs)
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			return fmt.Errorf("failed to open file: %s", f)
		}
		defer file.Close()
		stat, _ := os.Stat(f)

		scanner := bufio.NewScanner(file)
		var perFileLines int
		for scanner.Scan() {
			l.TotalLines += 1
			perFileLines += 1
		}
		l.PerFileLines = append(l.PerFileLines, Lines{
			Filename: stat.Name(),
			Count:    perFileLines,
		})
	}

	return nil
}
