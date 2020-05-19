package controllers

import (
	"context"
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/config"

	"github.com/gorilla/mux"
	"github.com/sagikazarmark/ocmux"
)

const (
	routesKey = "routes"
  infoKey = "info"
)

func BuildRouter(app *config.AppInfo) (*mux.Router, error) {
	r := mux.NewRouter()
	r.Use(ocmux.Middleware())

	// Home
	r.Methods(http.MethodGet).Path("/").Handler(addContext(r, app, http.HandlerFunc(Home))).Name("home")
	r.Methods(http.MethodGet).Path("/s").Handler(addContext(r, app, http.HandlerFunc(Socket))).Name("websocket")

	// Profile
	profile := r.Path("/profile").Subrouter()
	profile.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(Profile))).Name("profile")
	profile.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(ProfileSave))).Name("profile.save")

	// Sandbox
	sandbox := r.Path("/sandbox").Subrouter()
	sandbox.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(SandboxList))).Name("sandbox")
	r.Path("/sandbox/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(SandboxRun))).Name("sandbox.run")

	// Auth
	_ = r.Path("/auth").Subrouter()
	r.Path("/auth/callback/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AuthCallback))).Name("auth.callback")
	r.Path("/auth/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AuthSubmit))).Name("auth.submit")

	// Join
	join := r.Path("/join").Subrouter()
	r.Path("/join/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(JoinGet))).Name("join.get")
	join.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(JoinPost))).Name("join.post")

	// Team
	team := r.Path("/team").Subrouter()
	team.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(TeamList))).Name(util.SvcTeam.Key + ".list")
	team.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(TeamNew))).Name(util.SvcTeam.Key + ".new")
	r.Path("/team/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(TeamWorkspace))).Name(util.SvcTeam.Key)

	// Sprint
	sprint := r.Path("/sprint").Subrouter()
	sprint.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(SprintList))).Name(util.SvcSprint.Key + ".list")
	sprint.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(SprintNew))).Name(util.SvcSprint.Key + ".new")
	r.Path("/sprint/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(SprintWorkspace))).Name(util.SvcSprint.Key)

	// Estimate
	estimate := r.Path("/estimate").Subrouter()
	estimate.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(EstimateList))).Name(util.SvcEstimate.Key + ".list")
	estimate.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(EstimateNew))).Name(util.SvcEstimate.Key + ".new")
	r.Path("/estimate/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(EstimateWorkspace))).Name(util.SvcEstimate.Key)

	// Standup
	standup := r.Path("/standup").Subrouter()
	standup.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(StandupList))).Name(util.SvcStandup.Key + ".list")
	standup.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(StandupNew))).Name(util.SvcStandup.Key + ".new")
	r.Path("/standup/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(StandupWorkspace))).Name(util.SvcStandup.Key)

	// Retro
	retro := r.Path("/retro").Subrouter()
	retro.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(RetroList))).Name(util.SvcRetro.Key + ".list")
	retro.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(RetroNew))).Name(util.SvcRetro.Key + ".new")
	r.Path("/retro/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(RetroWorkspace))).Name(util.SvcRetro.Key)

	// Admin
	admin := r.Path("/admin").Subrouter()
	admin.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminHome))).Name("admin")

	r.Path("/admin/user").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminUserList))).Name("admin.user")
	r.Path("/admin/user/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminUserDetail))).Name("admin.user.detail")

	r.Path("/admin/auth").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminAuthList))).Name("admin.auth")
	r.Path("/admin/auth/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminAuthDetail))).Name("admin.auth.detail")

	r.Path("/admin/action").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminActionList))).Name("admin.action")
	r.Path("/admin/action/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminActionDetail))).Name("admin.action.detail")

	r.Path("/admin/invitation").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminInvitationList))).Name("admin.invitation")
	r.Path("/admin/invitation/{key}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminInvitationDetail))).Name("admin.invitation.detail")

	r.Path("/admin/team").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminTeamList))).Name("admin.team")
	r.Path("/admin/team/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminTeamDetail))).Name("admin.team.detail")

	r.Path("/admin/sprint").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminSprintList))).Name("admin.sprint")
	r.Path("/admin/sprint/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminSprintDetail))).Name("admin.sprint.detail")

	r.Path("/admin/estimate").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminEstimateList))).Name("admin.estimate")
	r.Path("/admin/estimate/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminEstimateDetail))).Name("admin.estimate.detail")
	r.Path("/admin/story/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminStoryDetail))).Name("admin.story.detail")

	r.Path("/admin/standup").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminStandupList))).Name("admin.standup")
	r.Path("/admin/standup/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminStandupDetail))).Name("admin.standup.detail")

	r.Path("/admin/retro").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminRetroList))).Name("admin.retro")
	r.Path("/admin/retro/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminRetroDetail))).Name("admin.retro.detail")

	r.Path("/admin/connection").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminConnectionList))).Name("admin.connection")
	r.Path("/admin/connection").Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(AdminConnectionPost))).Name("admin.connection.post")
	r.Path("/admin/connection/{id}").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(AdminConnectionDetail))).Name("admin.connection.detail")

	// GraphQL
	graphql := r.Path("/admin/graphql").Subrouter()
	graphql.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(GraphQLHome))).Name("graphql")
	graphql.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(GraphQLRun))).Name("graphql.run")

	// Utils
	_ = r.Path("/utils").Subrouter()
	r.Path("/about").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(About))).Name("about")
	r.Path("/health").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(Health))).Name("health")
	r.Path("/modules").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(Modules))).Name("modules")
	r.Path("/routes").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(Routes))).Name("routes")

	// Assets
	r.Path("/favicon.ico").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(Favicon))).Name("favicon")
	r.PathPrefix("/assets").Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(Static))).Name("assets")

	r.PathPrefix("").Handler(addContext(r, app, http.HandlerFunc(NotFound)))

	return r, nil
}

func addContext(router *mux.Router, info *config.AppInfo, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer internalServerError(router, info, w, r)
		ctx := context.WithValue(r.Context(), routesKey, router)
		ctx = context.WithValue(ctx, infoKey, info)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
