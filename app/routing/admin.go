package routing

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/npn/npncontroller/routes"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/rituals.dev/app/admin"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func adminRoutes(ai npnweb.AppInfo, r *mux.Router) *mux.Router {
	r.Path(routes.Adm()).Subrouter().Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.Home))).Name(npncore.KeyAdmin)

	r.Path(routes.Adm("enable")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.Enable))).Name(npnweb.AdminLink("enable"))

	r.Path(routes.Adm(npncore.KeySandbox)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.SandboxList))).Name(npnweb.AdminLink(npncore.KeySandbox))
	r.Path(routes.Adm(npncore.KeySandbox, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.SandboxRun))).Name(npnweb.AdminLink(npncore.KeySandbox, "run"))

	r.Path(routes.Adm(npncore.KeyTranscript)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.TranscriptList))).Name(npnweb.AdminLink(npncore.KeyTranscript))
	r.Path(routes.Adm(npncore.KeyTranscript, "{key}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.TranscriptRun))).Name(npnweb.AdminLink(npncore.KeyTranscript, "run"))

	r.Path(routes.Adm(npncore.KeyUser)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.UserList))).Name(npnweb.AdminLink(npncore.KeyUser))
	r.Path(routes.Adm(npncore.KeyUser, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.UserDetail))).Name(npnweb.AdminLink(npncore.KeyUser, npncore.KeyDetail))

	r.Path(routes.Adm(npncore.KeyAuth)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.AuthList))).Name(npnweb.AdminLink(npncore.KeyAuth))
	r.Path(routes.Adm(npncore.KeyAuth, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.AuthDetail))).Name(npnweb.AdminLink(npncore.KeyAuth, npncore.KeyDetail))

	r.Path(routes.Adm(npncore.KeyAction)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.ActionList))).Name(npnweb.AdminLink(npncore.KeyAction))
	r.Path(routes.Adm(npncore.KeyAction, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.ActionDetail))).Name(npnweb.AdminLink(npncore.KeyAction, npncore.KeyDetail))

	r.Path(routes.Adm(util.SvcTeam.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.TeamList))).Name(npnweb.AdminLink(util.SvcTeam.Key))
	r.Path(routes.Adm(util.SvcTeam.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.TeamDetail))).Name(npnweb.AdminLink(util.SvcTeam.Key, npncore.KeyDetail))

	r.Path(routes.Adm(util.SvcSprint.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.SprintList))).Name(npnweb.AdminLink(util.SvcSprint.Key))
	r.Path(routes.Adm(util.SvcSprint.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.SprintDetail))).Name(npnweb.AdminLink(util.SvcSprint.Key, npncore.KeyDetail))

	r.Path(routes.Adm(util.SvcEstimate.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.EstimateList))).Name(npnweb.AdminLink(util.SvcEstimate.Key))
	r.Path(routes.Adm(util.SvcEstimate.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.EstimateDetail))).Name(npnweb.AdminLink(util.SvcEstimate.Key, npncore.KeyDetail))
	r.Path(routes.Adm(util.KeyStory, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.StoryDetail))).Name(npnweb.AdminLink(util.KeyStory, npncore.KeyDetail))

	r.Path(routes.Adm(util.SvcStandup.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.StandupList))).Name(npnweb.AdminLink(util.SvcStandup.Key))
	r.Path(routes.Adm(util.SvcStandup.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.StandupDetail))).Name(npnweb.AdminLink(util.SvcStandup.Key, npncore.KeyDetail))

	r.Path(routes.Adm(util.SvcRetro.Key)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.RetroList))).Name(npnweb.AdminLink(util.SvcRetro.Key))
	r.Path(routes.Adm(util.SvcRetro.Key, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.RetroDetail))).Name(npnweb.AdminLink(util.SvcRetro.Key, npncore.KeyDetail))

	r.Path(routes.Adm(npncore.KeyComment)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.CommentList))).Name(npnweb.AdminLink(npncore.KeyComment))
	r.Path(routes.Adm(npncore.KeyComment, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.CommentDetail))).Name(npnweb.AdminLink(npncore.KeyComment, npncore.KeyDetail))

	r.Path(routes.Adm(npncore.KeyEmail)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.EmailList))).Name(npnweb.AdminLink(npncore.KeyEmail))
	r.Path(routes.Adm(npncore.KeyEmail, "{id}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.EmailDetail))).Name(npnweb.AdminLink(npncore.KeyEmail, npncore.KeyDetail))

	r.Path(routes.Adm(npncore.KeyMigration)).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.MigrationList))).Name(npnweb.AdminLink(npncore.KeyMigration))
	r.Path(routes.Adm(npncore.KeyMigration, "{idx}")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.HandlerFunc(admin.MigrationDetail))).Name(npnweb.AdminLink(npncore.KeyMigration, npncore.KeyDetail))

	// Dev mode sourcemaps
	r.PathPrefix(routes.Adm("src/")).Methods(http.MethodGet).Handler(routes.AddContext(r, ai, http.StripPrefix(routes.Adm("src/"), http.HandlerFunc(admin.Source)))).Name(npnweb.AdminLink("source"))

	// GraphQL
	npncontroller.RoutesSocketAdmin(ai, app.Svc(ai).Socket, r)
	npngraphql.RoutesGraphQLAdmin(ai, r)

	return r
}
