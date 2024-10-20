package pkg

import (
	"fmt"
	"os"
)

type CharactersCount struct {
	TotalMatches   int
	PerFileMatches []Characters
}

type Characters struct {
	Filename string
	Count    int
}

func (c *CharactersCount) GetCharactersCount(files ...string) error {
	if len(ValidateFiles(files...)) > 0 {
		return fmt.Errorf("failed to validate files")
	}
	for _, f := range files {
		data, _ := os.ReadFile(f)
		stat, _ := os.Stat(f)
		c.TotalMatches += len(string(data))
		c.PerFileMatches = append(c.PerFileMatches, Characters{
			Filename: stat.Name(),
			Count:    len(string(data)),
		})
	}

	return nil
}
