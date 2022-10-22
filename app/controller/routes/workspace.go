// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"
	"github.com/kyleu/rituals/app/controller/cworkspace"
)

func workspaceRoutes(r *router.Router) {
	r.GET("/team", cworkspace.TeamList)
	r.GET("/team/{slug}", cworkspace.TeamDetail)
	r.GET("/sprint", cworkspace.SprintList)
	r.GET("/sprint/{slug}", cworkspace.SprintDetail)
	r.GET("/estimate", cworkspace.EstimateList)
	r.GET("/estimate/{slug}", cworkspace.EstimateDetail)
	r.GET("/standup", cworkspace.StandupList)
	r.GET("/standup/{slug}", cworkspace.StandupDetail)
	r.GET("/retro", cworkspace.RetroList)
	r.GET("/retro/{slug}", cworkspace.RetroDetail)
}
