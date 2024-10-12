package internal

import (
	"bufio"
	"fmt"
	"os"
)

type LinesCount struct {
	totalLines   int
	perFileLines []lines
}

type lines struct {
	filename string
	count    int
}

func (l *LinesCount) GetLinesCount(files ...string) error {
	if len(ValidateFiles(files...)) > 0 {
		return fmt.Errorf("failed to vailidate files")
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
			l.totalLines += 1
			perFileLines += 1
		}
		l.perFileLines = append(l.perFileLines, lines{
			filename: stat.Name(),
			count:    perFileLines,
		})
	}

	return nil
}
