package tools

import (
	"github.com/gookit/color"
)

func ColorRGB(val string, r uint8, g uint8, b uint8) string {
	valColored := color.RGB(r, g, b).Basic().Render(val)
	return valColored
}
