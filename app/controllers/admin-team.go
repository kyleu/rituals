package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
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
		teamIDString := mux.Vars(r)["id"]
		teamID, err := uuid.FromString(teamIDString)
		if err != nil {
			return "", errors.New("invalid team id [" + teamIDString + "]")
		}
		sess, err := ctx.App.Team.GetByID(teamID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load team [" + teamIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.team"), nil
		}

		params := paramSetFromRequest(r)
		members, err := ctx.App.Team.Members.GetByModelID(teamID, params.Get(util.KeyMember, ctx.Logger))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcTeam.Key, teamID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.team"), util.SvcTeam.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.team.detail", "id", teamIDString), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminTeamDetail(sess, members, actions, params, ctx, w))
	})
}
