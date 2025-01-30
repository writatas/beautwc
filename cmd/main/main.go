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
	linesCommands, err := cli.LinesCommands()
	if err != nil {
		fmt.Println("failed to retrieve line commands", err)
	}
	wordsCommands, err := cli.WordsCommands()
	if err != nil {
		fmt.Println("failed to retrieve word commands", err)
	}

	allCommands := append(bytesCommands, charactersCommands...)
	allCommands = append(allCommands, linesCommands...)
	allCommands = append(allCommands, wordsCommands...)

	bCli.New(allCommands)
	err = bCli.Run()
	if err != nil {
		fmt.Println("error occured: %w", err)
	}
}
