package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"
)

var homeSections = []string{
	util.SvcTeam.Key, util.SvcSprint.Key, util.SvcEstimate.Key, util.SvcStandup.Key, util.SvcRetro.Key,
	npncore.KeyUser, npncore.KeyAuth, npncore.KeyAction, npncore.KeyComment, npncore.KeyEmail, npncore.KeyMigration,
	npncore.KeyConnection, npncore.KeySandbox, npncore.KeyRoutes, npncore.KeyModules, util.KeyTranscript, npncore.KeyGraphQL,
}

func Home(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)
		ctx.Title = "Admin"
		ctx.Breadcrumbs = npnweb.BreadcrumbsSimple("", npncore.KeyAdmin)
		countMap, recentMap, err := SectionCounts(homeSections, npnweb.ExtractRoutes(ctx.Routes), app.Database(ctx.App), app.Socket(ctx.App))
		if err != nil {
			return npncontroller.EResp(err, "error getting section counts")
		}
		return npncontroller.T(admintemplates.Home(ctx, homeSections, countMap, recentMap, params.Get(npncore.KeyAdmin, ctx.Logger), w))
	})
}

func Source(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		http.ServeFile(w, r, "./"+r.URL.Path)
		return "", nil
	})
}
