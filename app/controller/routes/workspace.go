// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/rituals/app/controller"
)

func workspaceRoutes(r *router.Router) {
	r.GET("/team", controller.Home)
	r.GET("/team/{slug}", controller.Home)
	r.GET("/sprint", controller.Home)
	r.GET("/sprint/{slug}", controller.Home)
	r.GET("/estimate", controller.Home)
	r.GET("/estimate/{slug}", controller.Home)
	r.GET("/standup", controller.Home)
	r.GET("/standup/{slug}", controller.Home)
	r.GET("/retro", controller.Home)
	r.GET("/retro/{slug}", controller.Home)
}
