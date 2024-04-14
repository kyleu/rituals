// Package theme - Content managed by Project Forge, see [projectforge.md] for details.
package theme

import (
	"fmt"
	"image/color"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/util"
)

const (
	white, black = "#ffffff", "#000000"
	threshold    = (65535 * 3) / 2
)

var Default = func() *Theme {
	nbl := "#8fbc8f"
	if o := util.GetEnv("app_nav_color_light"); o != "" {
		nbl = o
	}
	nbd := "#234d27"
	if o := util.GetEnv("app_nav_color_dark"); o != "" {
		nbd = o
	}

	return &Theme{
		Key: "default",
		Light: &Colors{
			Border: "1px solid #dddddd", LinkDecoration: "none",
			Foreground: "#000000", ForegroundMuted: "#777777",
			Background: "#ffffff", BackgroundMuted: "#e8f2e8",
			LinkForeground: "#546c54", LinkVisitedForeground: "#394839",
			NavForeground: "#000000", NavBackground: nbl,
			MenuForeground: "#000000", MenuSelectedForeground: "#000000",
			MenuBackground: "#bcd7bb", MenuSelectedBackground: "#8fbc8f",
			ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#ff0000",
		},
		Dark: &Colors{
			Border: "1px solid #666666", LinkDecoration: "none",
			Foreground: "#ffffff", ForegroundMuted: "#777777",
			Background: "#121212", BackgroundMuted: "#142214",
			LinkForeground: "#779076", LinkVisitedForeground: "#a3b4a2",
			NavForeground: "#ffffff", NavBackground: nbd,
			MenuForeground: "#eeeeee", MenuSelectedForeground: "#ffffff",
			MenuBackground: "#19301b", MenuSelectedBackground: "#234d27",
			ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#ff0000",
		},
	}
}()

func TextColorFor(clr string) string {
	c, err := ParseHexColor(clr)
	if err != nil {
		return white
	}
	r, g, b, _ := c.RGBA()
	total := r + g + b
	if total < threshold {
		return white
	}
	return black
}

func ParseHexColor(s string) (color.RGBA, error) {
	ret := color.RGBA{A: 0xff}
	var err error
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &ret.R, &ret.G, &ret.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &ret.R, &ret.G, &ret.B)
		// Double the hex digits:
		ret.R *= 17
		ret.G *= 17
		ret.B *= 17
	default:
		err = errors.Errorf("invalid length [%d], must be 7 or 4", len(s))
	}
	return ret, err
}
