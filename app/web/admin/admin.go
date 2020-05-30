package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

var homeSections = []string{
	util.KeyUser, util.KeyAuth, util.KeyAction, util.KeyInvitation,
	util.SvcTeam.Key, util.SvcSprint.Key, util.SvcEstimate.Key, util.SvcStandup.Key, util.SvcRetro.Key,
	util.KeyConnection, util.KeySandbox, util.KeyRoutes, util.KeyModules, util.KeyGraphQL,
}

func Home(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		ctx.Title = "Admin"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.AdminLink()), util.KeyAdmin)
		countMap, recentMap, err := SectionCounts(homeSections, util.ExtractRoutes(ctx.Routes), ctx.App.Database, ctx.App.Socket)
		if err != nil {
			return eresp(err, "error getting section counts")
		}
		return tmpl(templates.AdminHome(ctx, homeSections, countMap, recentMap, params.Get(util.KeyAdmin, ctx.Logger), w))
	})
}

func Source(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		http.ServeFile(w, r, "./" + r.URL.Path)
		return "", nil
	})
}
