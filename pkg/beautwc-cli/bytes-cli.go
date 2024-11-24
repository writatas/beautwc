package beautwccli

import (
	"fmt"

	"github.com/beautwc/pkg"
	"github.com/beautwc/tools"
	"github.com/urfave/cli/v2"
)

func BytesCommands() ([]cli.Command, error) {
	commands := []cli.Command{
		{
			Name:                   "bytes",
			Aliases:                []string{"b"},
			Usage:                  "beautwc --bytes, -b  get bytes from file",
			UseShortOptionHandling: true,
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:    "long",
					Usage:   "beautwc bytes, b --long, -l  get bytes from file with long prefixes for each file",
					Aliases: []string{"l"},
				},
			},
			Action: func(ctx *cli.Context) error {
				flaggedSlice := ctx.StringSlice("long")
				noFlaggedSlice := ctx.Args().Slice()
				flaggedSlice = append(flaggedSlice, noFlaggedSlice...)
				longShort := false // defaults to shortest prefix
				if ctx.IsSet("long") {
					longShort = true
				}
				if len(flaggedSlice) > 0 {
					memory, err := pkg.GetFilesSize(flaggedSlice...)
					if err != nil {
						return fmt.Errorf("failed to retrieve memory of files")
					}
					err = memory.TotalBytes.CalculateBytesAndPrefixes()
					if err != nil {
						return fmt.Errorf("failed to calculate bytes and prefixes: %w", err)
					}
					display, err := memory.TotalBytes.DisplayText(longShort)
					if err != nil {
						return fmt.Errorf("failed to display text: %w", err)
					}
					fmt.Printf(
						"%s:\n%s\n%s\n%s\n%s\n%s\n%s\n",
						tools.ColorRGB("Total Bytes for File(s)", 51, 204, 255),
						tools.ColorRGB("----------------------------------------------------", 51, 204, 255),
						tools.ColorRGB(fmt.Sprintf("%v", memory.TotalBytes.Files), 102, 153, 153),
						display.Bytes,
						display.Kilobytes,
						display.Megabytes,
						display.Gigabytes,
					)
				}
				return nil
			},
		},
	}
	return commands, nil
}
