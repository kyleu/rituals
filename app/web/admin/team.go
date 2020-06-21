package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Team List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)

		params := act.ParamSetFromRequest(r)
		teams := ctx.App.Team.List(params.Get(util.SvcTeam.Key, ctx.Logger))
		return act.T(admintemplates.TeamList(teams, params, ctx, w))
	})
}

func TeamDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		teamID, err := act.IDFromParams(util.SvcTeam.Key, mux.Vars(r))
		if err != nil {
			return act.EResp(err)
		}
		sess := ctx.App.Team.GetByID(*teamID)
		if sess == nil {
			msg := "can't load team [" + teamID.String() + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.SvcTeam.Key), w, r, ctx)
		}

		params := act.ParamSetFromRequest(r)

		sprints := ctx.App.Sprint.GetByTeamID(*teamID, params.Get(util.SvcSprint.Key, ctx.Logger))
		estimates := ctx.App.Estimate.GetByTeamID(*teamID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := ctx.App.Standup.GetByTeamID(*teamID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := ctx.App.Retro.GetByTeamID(*teamID, params.Get(util.SvcRetro.Key, ctx.Logger))

		data := ctx.App.Team.Data.GetData(*teamID, params, ctx.Logger)

		bc := adminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)
		bc = append(bc, web.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return act.T(admintemplates.TeamDetail(sess, sprints, estimates, standups, retros, data, params, ctx, w))
	})
}
