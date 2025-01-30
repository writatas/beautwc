package beautwccli

import (
	"fmt"

	"github.com/beautwc/pkg"
	"github.com/beautwc/tools"
	"github.com/urfave/cli/v2"
)

func LinesCommands() ([]cli.Command, error) {
	commands := []cli.Command{
		{
			Name:                   "lines",
			Aliases:                []string{"l"},
			Usage:                  "beautwc --lines -l get characters from file(s)",
			UseShortOptionHandling: true,
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() <= 0 {
					return fmt.Errorf("-l flag requires a file")
				}
				files := ctx.Args().Slice()
				linesCount := pkg.LinesCount{}
				err := linesCount.GetLinesCount(files...)
				if err != nil {
					return err
				}
				fmt.Printf("%s: %d\n", tools.ColorRGB("LinesTotal", 51, 204, 255), linesCount.TotalLines)
				for _, perFile := range linesCount.PerFileLines {
					fmt.Printf("%s: %d\n", tools.ColorRGB(perFile.Filename, 51, 204, 255), perFile.Count)
				}
				return nil
			},
		},
	}

	return commands, nil
}
