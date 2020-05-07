package controllers

import (
	"context"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/config"

	"github.com/gorilla/mux"
	"github.com/sagikazarmark/ocmux"
)

const routesKey = "routes"
const infoKey = "info"

func BuildRouter(info *config.AppInfo) (*mux.Router, error) {
	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Methods(http.MethodGet).Path("/").Handler(addContext(r, info, http.HandlerFunc(Home))).Name("home")
	r.Methods(http.MethodGet).Path("/s").Handler(addContext(r, info, http.HandlerFunc(Socket))).Name("websocket")

	profile := r.Path("/profile").Subrouter()
	profile.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(Profile))).Name("profile")
	profile.Methods(http.MethodPost).Handler(addContext(r, info, http.HandlerFunc(ProfileSave))).Name("profile.save")

	settings := r.Path("/settings").Subrouter()
	settings.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(Settings))).Name("settings")
	settings.Methods(http.MethodPost).Handler(addContext(r, info, http.HandlerFunc(SettingsSave))).Name("settings.save")

	// Sandbox
	sandbox := r.Path("/sandbox").Subrouter()
	sandbox.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(SandboxList))).Name("sandbox")
	r.Path("/sandbox/{key}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(SandboxForm))).Name("sandbox.run")

	// Join
	join := r.Path("/join").Subrouter()
	join.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(JoinForm))).Name("join.form")
	join.Methods(http.MethodPost).Handler(addContext(r, info, http.HandlerFunc(JoinPost))).Name("join.post")
	r.Path("/join/{key}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(JoinGet))).Name("join.get")
	r.Path("/new").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(SessionNew))).Name("session.new")

	// Estimate
	estimate := r.Path("/estimate").Subrouter()
	estimate.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(EstimateList))).Name(util.SvcEstimate + ".list")
	r.Path("/estimate/new").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(EstimateNewForm))).Name(util.SvcEstimate + ".new.form")
	r.Path("/estimate/new").Methods(http.MethodPost).Handler(addContext(r, info, http.HandlerFunc(EstimateNew))).Name(util.SvcEstimate + ".new")
	r.Path("/estimate/{key}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(EstimateWorkspace))).Name(util.SvcEstimate)

	// Standup
	standup := r.Path("/standup").Subrouter()
	standup.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(StandupList))).Name(util.SvcStandup + ".list")
	r.Path("/standup/new").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(StandupNewForm))).Name(util.SvcStandup + ".new.form")
	r.Path("/standup/new").Methods(http.MethodPost).Handler(addContext(r, info, http.HandlerFunc(StandupNew))).Name(util.SvcStandup + ".new")
	r.Path("/standup/{key}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(StandupWorkspace))).Name(util.SvcStandup)

	// Retro
	retro := r.Path("/retro").Subrouter()
	retro.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(RetroList))).Name(util.SvcRetro + ".list")
	r.Path("/retro/new").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(RetroNewForm))).Name(util.SvcRetro + ".new.form")
	r.Path("/retro/new").Methods(http.MethodPost).Handler(addContext(r, info, http.HandlerFunc(RetroNew))).Name(util.SvcRetro + ".new")
	r.Path("/retro/{key}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(RetroWorkspace))).Name(util.SvcRetro)

	// Admin
	admin := r.Path("/admin").Subrouter()
	admin.Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminHome))).Name("admin.home")
	r.Path("/admin/user").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminUserList))).Name("admin.user")
	r.Path("/admin/user/{id}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminUserDetail))).Name("admin.user.detail")
	r.Path("/admin/invite").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminInviteList))).Name("admin.invite")
	r.Path("/admin/invite/{key}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminInviteDetail))).Name("admin.invite.detail")
	r.Path("/admin/estimate").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminEstimateList))).Name("admin.estimate")
	r.Path("/admin/estimate/{id}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminEstimateDetail))).Name("admin.estimate.detail")
	r.Path("/admin/poll/{id}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminPollDetail))).Name("admin.poll.detail")
	r.Path("/admin/standup").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminStandupList))).Name("admin.standup")
	r.Path("/admin/standup/{id}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminStandupDetail))).Name("admin.standup.detail")
	r.Path("/admin/retro").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminRetroList))).Name("admin.retro")
	r.Path("/admin/retro/{id}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminRetroDetail))).Name("admin.retro.detail")
	r.Path("/admin/connection").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminConnectionList))).Name("admin.connection")
	r.Path("/admin/connection").Methods(http.MethodPost).Handler(addContext(r, info, http.HandlerFunc(AdminConnectionPost))).Name("admin.connection.post")
	r.Path("/admin/connection/{id}").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(AdminConnectionDetail))).Name("admin.connection.detail")

	// Utils
	_ = r.Path("/utils").Subrouter()
	r.Path("/about").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(About))).Name("about")
	r.Path("/health").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(Health))).Name("health")
	r.Path("/modules").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(Modules))).Name("modules")
	r.Path("/routes").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(Routes))).Name("routes")

	// Assets
	r.Path("/favicon.ico").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(Favicon))).Name("favicon")
	r.PathPrefix("/assets").Methods(http.MethodGet).Handler(addContext(r, info, http.HandlerFunc(Static))).Name("assets")

	r.PathPrefix("").Handler(addContext(r, info, http.HandlerFunc(NotFound)))

	return r, nil
}

func addContext(router *mux.Router, info *config.AppInfo, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer InternalServerError(router, info, w, r)
		ctx := context.WithValue(r.Context(), routesKey, router)
		ctx = context.WithValue(ctx, infoKey, info)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
