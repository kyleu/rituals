package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Team List"
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)

		params := npnweb.ParamSetFromRequest(r)
		teams := app.Svc(ctx.App).Team.List(params.Get(util.SvcTeam.Key, ctx.Logger))
		return npncontroller.T(admintemplates.TeamList(teams, params, ctx, w))
	})
}

func TeamDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		teamID, err := npnweb.IDFromParams(util.SvcTeam.Key, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		sess := app.Svc(ctx.App).Team.GetByID(*teamID)
		if sess == nil {
			msg := "can't load team [" + teamID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(util.SvcTeam.Key), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		sprints := app.Svc(ctx.App).Sprint.GetByTeamID(*teamID, params.Get(util.SvcSprint.Key, ctx.Logger))
		estimates := app.Svc(ctx.App).Estimate.GetByTeamID(*teamID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := app.Svc(ctx.App).Standup.GetByTeamID(*teamID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := app.Svc(ctx.App).Retro.GetByTeamID(*teamID, params.Get(util.SvcRetro.Key, ctx.Logger))

		data := app.Svc(ctx.App).Team.Data.GetData(*teamID, params, ctx.Logger)

		bc := npncontroller.AdminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)
		bc = append(bc, npnweb.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.TeamDetail(sess, sprints, estimates, standups, retros, data, params, ctx, w))
	})
}
