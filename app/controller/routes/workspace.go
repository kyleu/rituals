package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/rituals/app/controller/cworkspace"
)

func workspaceRoutes(r *router.Router) {
	r.GET("/team", cworkspace.TeamList)
	r.POST("/team", cworkspace.TeamCreate)
	r.GET("/team/{slug}", cworkspace.TeamDetail)
	r.POST("/team/{slug}", cworkspace.TeamAction)
	r.GET("/team/{id}/connect", cworkspace.TeamSocket)

	r.GET("/sprint", cworkspace.SprintList)
	r.POST("/sprint", cworkspace.SprintCreate)
	r.GET("/sprint/{slug}", cworkspace.SprintDetail)
	r.POST("/sprint/{slug}", cworkspace.SprintAction)
	r.GET("/sprint/{id}/connect", cworkspace.SprintSocket)

	r.GET("/estimate", cworkspace.EstimateList)
	r.POST("/estimate", cworkspace.EstimateCreate)
	r.GET("/estimate/{slug}", cworkspace.EstimateDetail)
	r.POST("/estimate/{slug}", cworkspace.EstimateAction)
	r.GET("/estimate/{id}/connect", cworkspace.EstimateSocket)

	r.GET("/standup", cworkspace.StandupList)
	r.POST("/standup", cworkspace.StandupCreate)
	r.GET("/standup/{slug}", cworkspace.StandupDetail)
	r.POST("/standup/{slug}", cworkspace.StandupAction)
	r.GET("/standup/{id}/connect", cworkspace.StandupSocket)

	r.GET("/retro", cworkspace.RetroList)
	r.POST("/retro", cworkspace.RetroCreate)
	r.GET("/retro/{slug}", cworkspace.RetroDetail)
	r.POST("/retro/{slug}", cworkspace.RetroAction)
	r.GET("/retro/{id}/connect", cworkspace.RetroSocket)
}
