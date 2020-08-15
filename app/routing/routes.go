package routing

import (
	"github.com/kyleu/npn/npnasset"
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncontroller/routes"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app/controllers"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
	"github.com/sagikazarmark/ocmux"
)

func BuildRouter(ai npnweb.AppInfo) (*mux.Router, error) {
	npncontroller.InitMime()

	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// route := func(rt *mux.Route, act func(w http.ResponseWriter, r *http.Request), names ...string) {
	// 	rt.Handler(routes.AddContext(r, ai, http.HandlerFunc(act))).Name(routes.Name(names...))
	// }

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Home))).Name(routes.Name("home"))
	r.Path(routes.Path("health")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Health))).Name(routes.Name("health"))
	r.Path(routes.Path("s")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Socket))).Name(routes.Name("websocket"))

	// Team
	team := r.Path(routes.Path(util.SvcTeam.Key)).Subrouter()
	team.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.TeamList))).Name(routes.Name(util.SvcTeam.Key, "list"))
	team.Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.TeamNew))).Name(routes.Name(util.SvcTeam.Key, "new"))
	r.Path(routes.Path(util.SvcTeam.Key, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.TeamWorkspace))).Name(routes.Name(util.SvcTeam.Key))
	r.Path(routes.Path(util.SvcTeam.Key, "{key}", npncore.KeyExport, "{fmt}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.TeamExport))).Name(routes.Name(util.SvcTeam.Key, npncore.KeyExport))

	// Sprint
	sprint := r.Path(routes.Path(util.SvcSprint.Key)).Subrouter()
	sprint.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.SprintList))).Name(routes.Name(util.SvcSprint.Key, "list"))
	sprint.Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.SprintNew))).Name(routes.Name(util.SvcSprint.Key, "new"))
	r.Path(routes.Path(util.SvcSprint.Key, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.SprintWorkspace))).Name(routes.Name(util.SvcSprint.Key))
	r.Path(routes.Path(util.SvcSprint.Key, "{key}", npncore.KeyExport, "{fmt}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.SprintExport))).Name(routes.Name(util.SvcSprint.Key, npncore.KeyExport))

	// Estimate
	estimate := r.Path(routes.Path(util.SvcEstimate.Key)).Subrouter()
	estimate.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.EstimateList))).Name(routes.Name(util.SvcEstimate.Key, "list"))
	estimate.Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.EstimateNew))).Name(routes.Name(util.SvcEstimate.Key, "new"))
	r.Path(routes.Path(util.SvcEstimate.Key, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.EstimateWorkspace))).Name(routes.Name(util.SvcEstimate.Key))
	r.Path(routes.Path(util.SvcEstimate.Key, "{key}", npncore.KeyExport, "{fmt}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.EstimateExport))).Name(routes.Name(util.SvcEstimate.Key, npncore.KeyExport))

	// Standup
	standup := r.Path(routes.Path(util.SvcStandup.Key)).Subrouter()
	standup.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.StandupList))).Name(routes.Name(util.SvcStandup.Key, "list"))
	standup.Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.StandupNew))).Name(routes.Name(util.SvcStandup.Key, "new"))
	r.Path(routes.Path(util.SvcStandup.Key, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.StandupWorkspace))).Name(routes.Name(util.SvcStandup.Key))
	r.Path(routes.Path(util.SvcStandup.Key, "{key}", npncore.KeyExport, "{fmt}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.StandupExport))).Name(routes.Name(util.SvcStandup.Key, npncore.KeyExport))

	// Retro
	retro := r.Path(routes.Path(util.SvcRetro.Key)).Subrouter()
	retro.Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.RetroList))).Name(routes.Name(util.SvcRetro.Key, "list"))
	retro.Methods(http.MethodPost).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.RetroNew))).Name(routes.Name(util.SvcRetro.Key, "new"))
	r.Path(routes.Path(util.SvcRetro.Key, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.RetroWorkspace))).Name(routes.Name(util.SvcRetro.Key))
	r.Path(routes.Path(util.SvcRetro.Key, "{key}", npncore.KeyExport, "{fmt}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.RetroExport))).Name(routes.Name(util.SvcRetro.Key, npncore.KeyExport))

	// Admin
	r = adminRoutes(ai, r)

	// Static
	r.Path(routes.Path(npncore.KeyAbout)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.About))).Name(routes.Name(npncore.KeyAbout))
	r.Path(routes.Path("pricing")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Pricing))).Name(routes.Name("pricing"))
	r.Path(routes.Path("features")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Features))).Name(routes.Name("features"))
	r.Path(routes.Path("community")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Community))).Name(routes.Name("community"))

	// Assets
	r.Path(routes.Path("favicon.ico")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Favicon))).Name(routes.Name("favicon"))
	r.Path(routes.Path("robots.txt")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.RobotsTxt))).Name(routes.Name("robots"))
	r.Path(routes.Path("qr")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.QRCode))).Name(routes.Name("qr"))
	r.PathPrefix(routes.Path("vendor")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(npnasset.VendorAsset))).Name(routes.Name("vendor"))
	r.PathPrefix(routes.Path("assets")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(controllers.Static))).Name(routes.Name("assets"))

	npncontroller.RoutesProfile(ai, r)
	npncontroller.RoutesAuth(ai, r)
	npncontroller.RoutesUtil(ai, r)

	return r, nil
}
