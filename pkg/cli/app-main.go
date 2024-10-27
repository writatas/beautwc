package pkg

import (
	"flag"
	"fmt"
	"os"
)

type Argument struct {
	Cmd    string
	SubCmd []string
	flags  []string
}

func ParseCommands(args Argument) error {
	return nil
}
