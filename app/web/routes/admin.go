package routes

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/web/admin"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/config"

	"github.com/gorilla/mux"
)

func adminRoutes(app *config.AppInfo, r *mux.Router) *mux.Router {
	r.Path(adm()).Subrouter().Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.Home))).Name(util.KeyAdmin)

	r.Path(adm("enable")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.Enable))).Name(util.AdminLink("enable"))

	r.Path(adm(util.KeySandbox)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.SandboxList))).Name(util.AdminLink(util.KeySandbox))
	r.Path(adm(util.KeySandbox, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.SandboxRun))).Name(util.AdminLink(util.KeySandbox, "run"))

	r.Path(adm(util.KeyTranscript)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.TranscriptList))).Name(util.AdminLink(util.KeyTranscript))
	r.Path(adm(util.KeyTranscript, "{key}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.TranscriptRun))).Name(util.AdminLink(util.KeyTranscript, "run"))

	r.Path(adm(util.KeyModules)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.Modules))).Name(util.AdminLink(util.KeyModules))
	r.Path(adm(util.KeyRoutes)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.Routes))).Name(util.AdminLink(util.KeyRoutes))

	r.Path(adm(util.KeyUser)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.UserList))).Name(util.AdminLink(util.KeyUser))
	r.Path(adm(util.KeyUser, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.UserDetail))).Name(util.AdminLink(util.KeyUser, util.KeyDetail))

	r.Path(adm(util.KeyAuth)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.AuthList))).Name(util.AdminLink(util.KeyAuth))
	r.Path(adm(util.KeyAuth, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.AuthDetail))).Name(util.AdminLink(util.KeyAuth, util.KeyDetail))

	r.Path(adm(util.KeyAction)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.ActionList))).Name(util.AdminLink(util.KeyAction))
	r.Path(adm(util.KeyAction, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.ActionDetail))).Name(util.AdminLink(util.KeyAction, util.KeyDetail))

	r.Path(adm(util.SvcTeam.Key)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.TeamList))).Name(util.AdminLink(util.SvcTeam.Key))
	r.Path(adm(util.SvcTeam.Key, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.TeamDetail))).Name(util.AdminLink(util.SvcTeam.Key, util.KeyDetail))

	r.Path(adm(util.SvcSprint.Key)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.SprintList))).Name(util.AdminLink(util.SvcSprint.Key))
	r.Path(adm(util.SvcSprint.Key, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.SprintDetail))).Name(util.AdminLink(util.SvcSprint.Key, util.KeyDetail))

	r.Path(adm(util.SvcEstimate.Key)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.EstimateList))).Name(util.AdminLink(util.SvcEstimate.Key))
	r.Path(adm(util.SvcEstimate.Key, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.EstimateDetail))).Name(util.AdminLink(util.SvcEstimate.Key, util.KeyDetail))
	r.Path(adm(util.KeyStory, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.StoryDetail))).Name(util.AdminLink(util.KeyStory, util.KeyDetail))

	r.Path(adm(util.SvcStandup.Key)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.StandupList))).Name(util.AdminLink(util.SvcStandup.Key))
	r.Path(adm(util.SvcStandup.Key, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.StandupDetail))).Name(util.AdminLink(util.SvcStandup.Key, util.KeyDetail))

	r.Path(adm(util.SvcRetro.Key)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.RetroList))).Name(util.AdminLink(util.SvcRetro.Key))
	r.Path(adm(util.SvcRetro.Key, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.RetroDetail))).Name(util.AdminLink(util.SvcRetro.Key, util.KeyDetail))

	r.Path(adm(util.KeyConnection)).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.ConnectionList))).Name(util.AdminLink(util.KeyConnection))
	r.Path(adm(util.KeyConnection)).Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(admin.ConnectionPost))).Name(util.AdminLink(util.KeyConnection, "post"))
	r.Path(adm(util.KeyConnection, "{id}")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.ConnectionDetail))).Name(util.AdminLink(util.KeyConnection, util.KeyDetail))

	// Dev mode sourcemaps
	r.PathPrefix(adm("src/")).Methods(http.MethodGet).Handler(addContext(r, app, http.StripPrefix(adm("src/"), http.HandlerFunc(admin.Source)))).Name(util.AdminLink("source"))

	// GraphQL
	graphql := r.Path(adm(util.KeyGraphQL)).Subrouter()
	graphql.Methods(http.MethodPost).Handler(addContext(r, app, http.HandlerFunc(admin.GraphQLRun))).Name(util.AdminLink(util.KeyGraphQL))
	graphql.Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.GraphiQL))).Name(util.AdminLink(util.KeyGraphiQL))
	r.Path(adm(util.KeyVoyager, "query")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.GraphQLVoyagerQuery))).Name(util.AdminLink(util.KeyVoyager, "query"))
	r.Path(adm(util.KeyVoyager, "mutation")).Methods(http.MethodGet).Handler(addContext(r, app, http.HandlerFunc(admin.GraphQLVoyagerMutation))).Name(util.AdminLink(util.KeyVoyager, "mutation"))

	return r
}

func adm(params ...string) string {
	params = append([]string{util.KeyAdmin}, params...)
	return p(params...)
}
