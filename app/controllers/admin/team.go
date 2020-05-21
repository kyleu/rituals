package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Team List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)

		params := act.ParamSetFromRequest(r)
		teams, err := ctx.App.Team.List(params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminTeamList(teams, params, ctx, w))
	})
}

func TeamDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		teamID := util.GetUUIDPointer(mux.Vars(r), "id")
		if teamID == nil {
			return "", errors.New("invalid team id")
		}
		sess, err := ctx.App.Team.GetByID(*teamID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load team [" + teamID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcTeam.Key)), nil
		}

		params := act.ParamSetFromRequest(r)
		members := ctx.App.Team.Members.GetByModelID(*teamID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Team.Permissions.GetByModelID(*teamID, params.Get(util.KeyPermission, ctx.Logger))

		sprints, err := ctx.App.Sprint.GetByTeamID(*teamID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetByTeamID(*teamID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetByTeamID(*teamID, params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetByTeamID(*teamID, params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", err
		}

		actions, err := ctx.App.Action.GetBySvcModel(util.SvcTeam.Key, *teamID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcTeam.Key, util.SvcTeam.Plural)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.SvcTeam.Key, util.KeyDetail), "id", teamID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminTeamDetail(sess, members, perms, sprints, estimates, standups, retros, actions, params, ctx, w))
	})
}
