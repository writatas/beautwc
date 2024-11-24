package beautwccli

import (
	"fmt"

	"github.com/beautwc/pkg"
	"github.com/beautwc/tools"
	"github.com/urfave/cli/v2"
)

func CharactersCommands() ([]cli.Command, error) {
	commands := []cli.Command{
		{
			Name:                   "characters",
			Aliases:                []string{"c"},
			Usage:                  "beautwc --characters, -c get chararacters from file(s)",
			UseShortOptionHandling: true,
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() <= 0 {
					return fmt.Errorf("-c flag requires a file")
				}
				files := ctx.Args().Slice()
				charactersCount := pkg.CharactersCount{}
				err := charactersCount.GetCharactersCount(files...)
				if err != nil {
					return err
				}
				fmt.Printf("%s:  %d\n", tools.ColorRGB("CharactersTotal", 51, 204, 255), charactersCount.TotalMatches)
				for _, perFile := range charactersCount.PerFileMatches {
					fmt.Printf("%s:  %d\n", tools.ColorRGB(perFile.Filename, 51, 204, 255), perFile.Count)
				}
				return nil
			},
		},
	}
	return commands, nil
}
