package tools

import (
	"github.com/fatih/color"
)

func GetValueColors() (color.Color, color.Color) {
	prefix := color.New(color.FgBlue)
	value := color.New(color.FgGreen)

	return *prefix, *value
}
