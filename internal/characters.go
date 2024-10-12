package internal

import (
	"fmt"
	"os"
)

type CharactersCount struct {
	totalMatches   int
	perFileMatches []characters
}

type characters struct {
	filename string
	count    int
}

func (c *CharactersCount) GetCharactersCount(files ...string) error {
	if len(ValidateFiles(files...)) > 0 {
		return fmt.Errorf("failed to validate files")
	}
	for _, f := range files {
		data, _ := os.ReadFile(f)
		stat, _ := os.Stat(f)
		c.totalMatches += len(string(data))
		c.perFileMatches = append(c.perFileMatches, characters{
			filename: stat.Name(),
		})
	}

	return nil
}
