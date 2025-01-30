package beautwccli

import (
	"fmt"

	"github.com/beautwc/pkg"
	"github.com/beautwc/tools"
	"github.com/urfave/cli/v2"
)

func WordsCommands() ([]cli.Command, error) {
	commands := []cli.Command{
		{
			Name:                   "words",
			Aliases:                []string{"w"},
			Usage:                  "beautwc --words -w get word count from file(s)",
			UseShortOptionHandling: true,
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() <= 0 {
					return fmt.Errorf("-w flag requires a file")
				}
				files := ctx.Args().Slice()
				wordsCount := pkg.WordsCount{}
				err := wordsCount.GetWordsCount(files...)
				if err != nil {
					return err
				}
				fmt.Printf("%s: %d\n", tools.ColorRGB("WordsTotal", 51, 204, 255), wordsCount.TotalWords)
				for _, perFile := range wordsCount.WordsPerFile {
					fmt.Printf("%s: %d\n", tools.ColorRGB(perFile.Filename, 51, 204, 255), perFile.Count)
				}
				return nil
			},
		},
	}
	return commands, nil
}
