package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminTeamList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Team List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.team"), util.SvcTeam.Key)...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		teams, err := ctx.App.Team.List(params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminTeamList(teams, params, ctx, w))
	})
}

func AdminTeamDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		teamID := getUUIDPointer(mux.Vars(r), "id")
		if teamID == nil {
			return "", errors.New("invalid team id")
		}
		sess, err := ctx.App.Team.GetByID(*teamID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load team [" + teamID.String() + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.team"), nil
		}

		params := paramSetFromRequest(r)
		members := ctx.App.Team.Members.GetByModelID(*teamID, params.Get(util.KeyMember, ctx.Logger))

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
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.team"), util.SvcTeam.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.team.detail", "id", teamID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminTeamDetail(sess, members, sprints, estimates, standups, retros, actions, params, ctx, w))
	})
}
