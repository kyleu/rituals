// Content managed by Project Forge, see [projectforge.md] for details.
package theme

import (
	"github.com/kyleu/rituals/app/util"
)

var ThemeDefault = func() *Theme {
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
			MenuForeground: "#000000", MenuBackground: "#bcd7bb", MenuSelectedBackground: "#8fbc8f", MenuSelectedForeground: "#000000",
			ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#ff0000",
		},
		Dark: &Colors{
			Border: "1px solid #666666", LinkDecoration: "none",
			Foreground: "#ffffff", ForegroundMuted: "#777777",
			Background: "#121212", BackgroundMuted: "#142214",
			LinkForeground: "#779076", LinkVisitedForeground: "#a3b4a2",
			NavForeground: "#ffffff", NavBackground: nbd,
			MenuForeground: "#eeeeee", MenuBackground: "#19301b", MenuSelectedBackground: "#234d27", MenuSelectedForeground: "#ffffff",
			ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#ff0000",
		},
	}
}()
