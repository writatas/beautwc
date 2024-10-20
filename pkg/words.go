package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordsCount struct {
	TotalWords   int
	WordsPerFile []Words
}

type Words struct {
	Filename string
	Count    int
}

func (w *WordsCount) GetWordsCount(files ...string) error {
	if len(ValidateFiles(files...)) > 0 {
		return fmt.Errorf("failed to validate files")
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			return fmt.Errorf("failed to open file: %s", f)
		}
		defer file.Close()
		stat, _ := os.Stat(f)

		scanner := bufio.NewScanner(file)
		var wordsPerFile int
		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Fields(line)
			w.TotalWords += len(words)
			wordsPerFile += len(words)
		}
		w.WordsPerFile = append(w.WordsPerFile, Words{
			Filename: stat.Name(),
			Count:    wordsPerFile,
		})
	}

	return nil
}
