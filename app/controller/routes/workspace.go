package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals/app/controller/cworkspace"
)

func workspaceRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/team", cworkspace.TeamList)
	makeRoute(r, http.MethodPost, "/team", cworkspace.TeamCreate)
	makeRoute(r, http.MethodGet, "/team/{slug}", cworkspace.TeamDetail)
	makeRoute(r, http.MethodPost, "/team/{slug}", cworkspace.TeamAction)
	makeRoute(r, http.MethodGet, "/team/{slug}/delete", cworkspace.TeamDelete)
	makeRoute(r, http.MethodGet, "/team/{id}/connect", cworkspace.TeamSocket)

	makeRoute(r, http.MethodGet, "/sprint", cworkspace.SprintList)
	makeRoute(r, http.MethodPost, "/sprint", cworkspace.SprintCreate)
	makeRoute(r, http.MethodGet, "/sprint/{slug}", cworkspace.SprintDetail)
	makeRoute(r, http.MethodPost, "/sprint/{slug}", cworkspace.SprintAction)
	makeRoute(r, http.MethodGet, "/sprint/{slug}/delete", cworkspace.SprintDelete)
	makeRoute(r, http.MethodGet, "/sprint/{id}/connect", cworkspace.SprintSocket)

	makeRoute(r, http.MethodGet, "/estimate", cworkspace.EstimateList)
	makeRoute(r, http.MethodPost, "/estimate", cworkspace.EstimateCreate)
	makeRoute(r, http.MethodGet, "/estimate/{slug}", cworkspace.EstimateDetail)
	makeRoute(r, http.MethodPost, "/estimate/{slug}", cworkspace.EstimateAction)
	makeRoute(r, http.MethodGet, "/estimate/{slug}/delete", cworkspace.EstimateDelete)
	makeRoute(r, http.MethodGet, "/estimate/{id}/connect", cworkspace.EstimateSocket)

	makeRoute(r, http.MethodGet, "/standup", cworkspace.StandupList)
	makeRoute(r, http.MethodPost, "/standup", cworkspace.StandupCreate)
	makeRoute(r, http.MethodGet, "/standup/{slug}", cworkspace.StandupDetail)
	makeRoute(r, http.MethodPost, "/standup/{slug}", cworkspace.StandupAction)
	makeRoute(r, http.MethodGet, "/standup/{slug}/delete", cworkspace.StandupDelete)
	makeRoute(r, http.MethodGet, "/standup/{id}/connect", cworkspace.StandupSocket)

	makeRoute(r, http.MethodGet, "/retro", cworkspace.RetroList)
	makeRoute(r, http.MethodPost, "/retro", cworkspace.RetroCreate)
	makeRoute(r, http.MethodGet, "/retro/{slug}", cworkspace.RetroDetail)
	makeRoute(r, http.MethodPost, "/retro/{slug}", cworkspace.RetroAction)
	makeRoute(r, http.MethodGet, "/retro/{slug}/delete", cworkspace.RetroDelete)
	makeRoute(r, http.MethodGet, "/retro/{id}/connect", cworkspace.RetroSocket)
}
