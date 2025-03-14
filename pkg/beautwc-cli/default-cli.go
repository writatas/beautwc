package beautwccli

import (
	"fmt"

	"github.com/beautwc/pkg"
	"github.com/urfave/cli/v2"
)

func PrintAll(files ...string) error {
	wordsCount := pkg.WordsCount{}
	linesCount := pkg.LinesCount{}
	charactersCount := pkg.CharactersCount{}
	memoryBytes, err := pkg.GetFilesSize(files...)
	if err != nil {
		return err
	}
	err = memoryBytes.TotalBytes.CalculateBytesAndPrefixes()
	if err != nil {
		return err
	}
	err = wordsCount.GetWordsCount(files...)
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
	pkg.PrintDefaultValue(&memoryBytes.TotalBytes)
	pkg.PrintDefaultValue(&wordsCount)
	pkg.PrintDefaultValue(&linesCount)
	pkg.PrintDefaultValue(&charactersCount)
	return nil
}

func DefaultCommand() ([]cli.Command, error) {
	commands := []cli.Command{
		{
			Name: "default",
			Action: func(ctx *cli.Context) error {
				args := ctx.Args().Slice()
				if len(args) == 0 {
					return fmt.Errorf("file paths required")
				}
				err := PrintAll(args...)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	return commands, nil
}
