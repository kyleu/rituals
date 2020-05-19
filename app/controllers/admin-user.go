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

func AdminUserList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "User List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user"), "users")...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		users, err := ctx.App.User.List(params.Get(util.KeyUser, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminUserList(users, params, ctx, w))
	})
}

func AdminUserDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		userIDString := mux.Vars(r)["id"]
		userID, err := uuid.FromString(userIDString)
		if err != nil {
			return "", errors.New("invalid user id [" + userIDString + "]")
		}
		u, err := ctx.App.User.GetByID(userID, false)
		if err != nil {
			return "", err
		}
		if u == nil {
			ctx.Session.AddFlash("error:Can't load user [" + userIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.user"), nil
		}

		params := paramSetFromRequest(r)

		auths, err := ctx.App.Auth.GetByUserID(userID, params.Get(util.KeyAuth, ctx.Logger))
		if err != nil {
			return "", err
		}
		teams, err := ctx.App.Team.GetByMember(userID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		sprints, err := ctx.App.Sprint.GetByMember(userID, params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetByMember(userID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetByMember(userID, params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetByMember(userID, params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetByAuthor(userID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = u.Name
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user"), "users")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user.detail", "id", userIDString), u.Name)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminUserDetail(u, auths, teams, sprints, estimates, standups, retros, actions, params, ctx, w))
	})
}
