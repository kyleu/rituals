// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/rituals/app/controller/clib"
)

func themeRoutes(r *router.Router) {
	r.GET("/theme", clib.ThemeList)
	r.GET("/theme/{key}", clib.ThemeEdit)
	r.POST("/theme/{key}", clib.ThemeSave)
	r.GET("/theme/{key}/remove", clib.ThemeRemove)
}
