package thirdparty

import (
	"github.com/fatih/color"
)

func GetValueColors() (color.Color, color.Color) {
	prefixColored := color.New(color.FgBlue)
	valueColored := color.New(color.FgGreen)
	return *prefixColored, *valueColored
}
