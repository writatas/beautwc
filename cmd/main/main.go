package main

import (
	"fmt"
	cli "github.com/beautwc/pkg/beautwc-cli"
)

func main() {
	bCli := cli.BeautwcCli{Name: "beautwc"}
	bytesCommands, err := cli.BytesCommands()
	if err != nil {
		fmt.Println("failed to retrieve bytes commands: %w", err)
	}

	bCli.New(bytesCommands)
	err = bCli.Run()
	if err != nil {
		fmt.Println("error occured: %w", err)
	}
}
