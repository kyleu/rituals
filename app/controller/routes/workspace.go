package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/rituals/app/controller/cworkspace"
)

func workspaceRoutes(r *router.Router) {
	r.GET("/team", cworkspace.TeamList)
	r.POST("/team", cworkspace.TeamCreate)
	r.GET("/team/{slug}", cworkspace.TeamDetail)
	r.POST("/team/{slug}", cworkspace.TeamSave)
	r.GET("/team/{slug}/connect", cworkspace.Socket)

	r.GET("/sprint", cworkspace.SprintList)
	r.POST("/sprint", cworkspace.SprintCreate)
	r.GET("/sprint/{slug}", cworkspace.SprintDetail)
	r.POST("/sprint/{slug}", cworkspace.SprintSave)
	r.GET("/sprint/{slug}/connect", cworkspace.Socket)

	r.GET("/estimate", cworkspace.EstimateList)
	r.POST("/estimate", cworkspace.EstimateCreate)
	r.GET("/estimate/{slug}", cworkspace.EstimateDetail)
	r.POST("/estimate/{slug}", cworkspace.EstimateSave)
	r.GET("/estimate/{slug}/connect", cworkspace.Socket)

	r.GET("/standup", cworkspace.StandupList)
	r.POST("/standup", cworkspace.StandupCreate)
	r.GET("/standup/{slug}", cworkspace.StandupDetail)
	r.POST("/standup/{slug}", cworkspace.StandupSave)
	r.GET("/standup/{slug}/connect", cworkspace.Socket)

	r.GET("/retro", cworkspace.RetroList)
	r.POST("/retro", cworkspace.RetroCreate)
	r.GET("/retro/{slug}", cworkspace.RetroDetail)
	r.POST("/retro/{slug}", cworkspace.RetroSave)
	r.GET("/retro/{slug}/connect", cworkspace.Socket)
}
