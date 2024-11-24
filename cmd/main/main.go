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
	charactersCommands, err := cli.CharactersCommands()
	if err != nil {
		fmt.Println("faled to retrieve character commands", err)
	}

	allCommands := append(bytesCommands, charactersCommands...)

	bCli.New(allCommands)
	err = bCli.Run()
	if err != nil {
		fmt.Println("error occured: %w", err)
	}
}
