package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Team List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)

		params := act.ParamSetFromRequest(r)
		teams := ctx.App.Team.List(params.Get(util.SvcTeam.Key, ctx.Logger))
		return tmpl(templates.AdminTeamList(teams, params, ctx, w))
	})
}

func TeamDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		teamID, err := act.IDFromParams(util.SvcTeam.Key, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		sess, err := ctx.App.Team.GetByID(*teamID)
		if err != nil {
			return eresp(err, "")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load team [" + teamID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcTeam.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		sprints := ctx.App.Sprint.GetByTeamID(*teamID, params.Get(util.SvcSprint.Key, ctx.Logger))
		estimates := ctx.App.Estimate.GetByTeamID(*teamID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := ctx.App.Standup.GetByTeamID(*teamID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := ctx.App.Retro.GetByTeamID(*teamID, params.Get(util.SvcRetro.Key, ctx.Logger))

		comments := ctx.App.Team.Comments.GetByModelID(*teamID, params.Get(util.KeyComment, ctx.Logger))
		actions := ctx.App.Action.GetBySvcModel(util.SvcTeam, *teamID, params.Get(util.KeyAction, ctx.Logger))
		members := ctx.App.Team.Members.GetByModelID(*teamID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Team.Permissions.GetByModelID(*teamID, params.Get(util.KeyPermission, ctx.Logger))

		bc := adminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)
		link := util.AdminLink(util.SvcTeam.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, teamID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminTeamDetail(sess, sprints, estimates, standups, retros, comments, members, perms, actions, params, ctx, w))
	})
}
