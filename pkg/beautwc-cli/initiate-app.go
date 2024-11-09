package beautwccli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

type BeautwcCli struct {
	Name string
	App  cli.App
}

// initiate new beatwc app with commands
func (b *BeautwcCli) New(commands []cli.Command) error {
	app := cli.NewApp()
	app.Name = b.Name
	for _, c := range commands {
		app.Commands = append(app.Commands, &c)
	}
	b.App = *app
	return nil
}

func (b *BeautwcCli) Run() error {
	if err := b.App.Run(os.Args); err != nil {
		return fmt.Errorf("failed to run app %s: %w", b.Name, err)
	}
	return nil
}
