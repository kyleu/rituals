package routing

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller/routes"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/rituals.dev/app/admin"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func adminRoutes(app npnweb.AppInfo, r *mux.Router) *mux.Router {
	r.Path(adm()).Subrouter().Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.Home))).Name(npncore.KeyAdmin)

	r.Path(adm("enable")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.Enable))).Name(npnweb.AdminLink("enable"))

	r.Path(adm(npncore.KeySandbox)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.SandboxList))).Name(npnweb.AdminLink(npncore.KeySandbox))
	r.Path(adm(npncore.KeySandbox, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.SandboxRun))).Name(npnweb.AdminLink(npncore.KeySandbox, "run"))

	r.Path(adm(util.KeyTranscript)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.TranscriptList))).Name(npnweb.AdminLink(util.KeyTranscript))
	r.Path(adm(util.KeyTranscript, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.TranscriptRun))).Name(npnweb.AdminLink(util.KeyTranscript, "run"))

	r.Path(adm(npncore.KeyUser)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.UserList))).Name(npnweb.AdminLink(npncore.KeyUser))
	r.Path(adm(npncore.KeyUser, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.UserDetail))).Name(npnweb.AdminLink(npncore.KeyUser, npncore.KeyDetail))

	r.Path(adm(npncore.KeyAuth)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.AuthList))).Name(npnweb.AdminLink(npncore.KeyAuth))
	r.Path(adm(npncore.KeyAuth, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.AuthDetail))).Name(npnweb.AdminLink(npncore.KeyAuth, npncore.KeyDetail))

	r.Path(adm(npncore.KeyAction)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.ActionList))).Name(npnweb.AdminLink(npncore.KeyAction))
	r.Path(adm(npncore.KeyAction, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.ActionDetail))).Name(npnweb.AdminLink(npncore.KeyAction, npncore.KeyDetail))

	r.Path(adm(util.SvcTeam.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.TeamList))).Name(npnweb.AdminLink(util.SvcTeam.Key))
	r.Path(adm(util.SvcTeam.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.TeamDetail))).Name(npnweb.AdminLink(util.SvcTeam.Key, npncore.KeyDetail))

	r.Path(adm(util.SvcSprint.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.SprintList))).Name(npnweb.AdminLink(util.SvcSprint.Key))
	r.Path(adm(util.SvcSprint.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.SprintDetail))).Name(npnweb.AdminLink(util.SvcSprint.Key, npncore.KeyDetail))

	r.Path(adm(util.SvcEstimate.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.EstimateList))).Name(npnweb.AdminLink(util.SvcEstimate.Key))
	r.Path(adm(util.SvcEstimate.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.EstimateDetail))).Name(npnweb.AdminLink(util.SvcEstimate.Key, npncore.KeyDetail))
	r.Path(adm(util.KeyStory, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.StoryDetail))).Name(npnweb.AdminLink(util.KeyStory, npncore.KeyDetail))

	r.Path(adm(util.SvcStandup.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.StandupList))).Name(npnweb.AdminLink(util.SvcStandup.Key))
	r.Path(adm(util.SvcStandup.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.StandupDetail))).Name(npnweb.AdminLink(util.SvcStandup.Key, npncore.KeyDetail))

	r.Path(adm(util.SvcRetro.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.RetroList))).Name(npnweb.AdminLink(util.SvcRetro.Key))
	r.Path(adm(util.SvcRetro.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.RetroDetail))).Name(npnweb.AdminLink(util.SvcRetro.Key, npncore.KeyDetail))

	r.Path(adm(npncore.KeyComment)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.CommentList))).Name(npnweb.AdminLink(npncore.KeyComment))
	r.Path(adm(npncore.KeyComment, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.CommentDetail))).Name(npnweb.AdminLink(npncore.KeyComment, npncore.KeyDetail))

	r.Path(adm(npncore.KeyEmail)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.EmailList))).Name(npnweb.AdminLink(npncore.KeyEmail))
	r.Path(adm(npncore.KeyEmail, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.EmailDetail))).Name(npnweb.AdminLink(npncore.KeyEmail, npncore.KeyDetail))

	r.Path(adm(npncore.KeyMigration)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.MigrationList))).Name(npnweb.AdminLink(npncore.KeyMigration))
	r.Path(adm(npncore.KeyMigration, "{idx}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.MigrationDetail))).Name(npnweb.AdminLink(npncore.KeyMigration, npncore.KeyDetail))

	r.Path(adm(npncore.KeyConnection)).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.ConnectionList))).Name(npnweb.AdminLink(npncore.KeyConnection))
	r.Path(adm(npncore.KeyConnection)).Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.ConnectionPost))).Name(npnweb.AdminLink(npncore.KeyConnection, "post"))
	r.Path(adm(npncore.KeyConnection, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.ConnectionDetail))).Name(npnweb.AdminLink(npncore.KeyConnection, npncore.KeyDetail))

	// Dev mode sourcemaps
	r.PathPrefix(adm("src/")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.StripPrefix(adm("src/"), http.HandlerFunc(admin.Source)))).Name(npnweb.AdminLink("source"))

	// GraphQL
	graphql := r.Path(adm(npncore.KeyGraphQL)).Subrouter()
	graphql.Methods(http.MethodPost).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.GraphQLRun))).Name(npnweb.AdminLink(npncore.KeyGraphQL))
	graphql.Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.GraphiQL))).Name(npnweb.AdminLink(npncore.KeyGraphiQL))
	r.Path(adm(npncore.KeyVoyager, "query")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.GraphQLVoyagerQuery))).Name(npnweb.AdminLink(npncore.KeyVoyager, "query"))
	r.Path(adm(npncore.KeyVoyager, "mutation")).Methods(http.MethodGet).Handler(routes.AddContext(r, app, http.HandlerFunc(admin.GraphQLVoyagerMutation))).Name(npnweb.AdminLink(npncore.KeyVoyager, "mutation"))

	return r
}

func adm(params ...string) string {
	params = append([]string{npncore.KeyAdmin}, params...)
	return routes.Path(params...)
}
