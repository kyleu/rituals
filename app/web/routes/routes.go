package routes

import (
	"net/http"
	"strings"

	"github.com/kyleu/rituals.dev/app/web/controllers"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/config"

	"github.com/gorilla/mux"
	"github.com/sagikazarmark/ocmux"
)

func BuildRouter(app *config.AppInfo) (*mux.Router, error) {
	initMime()

	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Home))).Name(n("home"))
	r.Path(p("health")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Health))).Name(n("health"))
	r.Path(p("s")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Socket))).Name(n("websocket"))

	// Profile
	profile := r.Path(p(util.KeyProfile)).Subrouter()
	profile.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Profile))).Name(n(util.KeyProfile))
	profile.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.ProfileSave))).Name(n(util.KeyProfile, "save"))
	r.Path(p(util.KeyProfile, "pic", "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.ProfilePic))).Name(n(util.KeyProfile, "pic"))

	// Auth
	_ = r.Path(p(util.KeyAuth)).Subrouter()
	r.Path(p(util.KeyAuth, "callback", "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.AuthCallback))).Name(n(util.KeyAuth, "callback"))
	r.Path(p(util.KeyAuth, "signout", "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.AuthSignout))).Name(n(util.KeyAuth, "signout"))
	r.Path(p(util.KeyAuth, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.AuthSubmit))).Name(n(util.KeyAuth, "submit"))

	// Team
	team := r.Path(p(util.SvcTeam.Key)).Subrouter()
	team.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.TeamList))).Name(n(util.SvcTeam.Key, "list"))
	team.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.TeamNew))).Name(n(util.SvcTeam.Key, "new"))
	r.Path(p(util.SvcTeam.Key, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.TeamWorkspace))).Name(n(util.SvcTeam.Key))

	// Sprint
	sprint := r.Path(p(util.SvcSprint.Key)).Subrouter()
	sprint.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SprintList))).Name(n(util.SvcSprint.Key, "list"))
	sprint.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.SprintNew))).Name(n(util.SvcSprint.Key, "new"))
	r.Path(p(util.SvcSprint.Key, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.SprintWorkspace))).Name(n(util.SvcSprint.Key))

	// Estimate
	estimate := r.Path(p(util.SvcEstimate.Key)).Subrouter()
	estimate.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.EstimateList))).Name(n(util.SvcEstimate.Key, "list"))
	estimate.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.EstimateNew))).Name(n(util.SvcEstimate.Key, "new"))
	r.Path(p(util.SvcEstimate.Key, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.EstimateWorkspace))).Name(n(util.SvcEstimate.Key))

	// Standup
	standup := r.Path(p(util.SvcStandup.Key)).Subrouter()
	standup.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.StandupList))).Name(n(util.SvcStandup.Key, "list"))
	standup.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.StandupNew))).Name(n(util.SvcStandup.Key, "new"))
	r.Path(p(util.SvcStandup.Key, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.StandupWorkspace))).Name(n(util.SvcStandup.Key))

	// Retro
	retro := r.Path(p(util.SvcRetro.Key)).Subrouter()
	retro.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.RetroList))).Name(n(util.SvcRetro.Key, "list"))
	retro.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(controllers.RetroNew))).Name(n(util.SvcRetro.Key, "new"))
	r.Path(p(util.SvcRetro.Key, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.RetroWorkspace))).Name(n(util.SvcRetro.Key))

	// Admin
	r = adminRoutes(app, r)

	// Static
	r.Path(p(util.KeyAbout)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.About))).Name(n(util.KeyAbout))
	r.Path(p("pricing")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Pricing))).Name(n("pricing"))
	r.Path(p("features")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Features))).Name(n("features"))
	r.Path(p("community")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Community))).Name(n("community"))

	// Assets
	r.Path(p("favicon.ico")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Favicon))).Name(n("favicon"))
	r.PathPrefix(p("assets")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(controllers.Static))).Name(n("assets"))

	r.PathPrefix("").Handler(addContext(r, app, http.HandlerFunc(controllers.NotFound)))

	return r, nil
}

func p(params ...string) string {
	ret := ""
	for _, p := range params {
		ret = ret + "/" + p
	}
	return ret
}

func n(params ...string) string {
	return strings.Join(params, ".")
}
