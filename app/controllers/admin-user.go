package controllers

import (
	"net/http"

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

		users, err := ctx.App.User.List()
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminUserList(users, ctx, w))
	})
}

func AdminUserDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		userIDString := mux.Vars(r)["id"]
		userID, err := uuid.FromString(userIDString)
		if err != nil {
			return "", errors.New("invalid user id [" + userIDString + "]")
		}
		user, err := ctx.App.User.GetByID(userID, false)
		if err != nil {
			return "", err
		}
		if user == nil {
			ctx.Session.AddFlash("error:Can't load user [" + userIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.user"), nil
		}

		auths, err := ctx.App.Auth.GetByUserID(userID, 0)
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetByMember(userID, 0)
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetByMember(userID, 0)
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetByMember(userID, 0)
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetByAuthor(userID)
		if err != nil {
			return "", err
		}

		ctx.Title = user.Name
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user"), "users")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user.detail", "id", userIDString), user.Name)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminUserDetail(user, auths, estimates, standups, retros, actions, ctx, w))
	})
}
