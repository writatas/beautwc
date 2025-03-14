package pkg

import (
	"fmt"

	"github.com/beautwc/tools"
)

func WordsLinesCharacters(files ...string) error {
	wordsCount := WordsCount{}
	linesCount := LinesCount{}
	charactersCount := CharactersCount{}

	err := wordsCount.GetWordsCount(files...)
	if err != nil {
		return err
	}
	err = linesCount.GetLinesCount(files...)
	if err != nil {
		return err
	}
	err = charactersCount.GetCharactersCount(files...)
	if err != nil {
		return err
	}

	return nil
}

func printOutValue(name string, files []any) error {
	for _, f := range files {
		fmt.Printf("%s: %d\n", tools.ColorRGB(f.FileName, 51, 204, 255), f.Count)
	}
	return nil
}
