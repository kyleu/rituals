package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminAuthList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Auth List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.auth"), "auths")...)
		ctx.Breadcrumbs = bc

		users, err := ctx.App.Auth.List()
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminAuthList(users, ctx, w))
	})
}

func AdminAuthDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		authIDString := mux.Vars(r)["id"]
		authID, err := uuid.FromString(authIDString)
		if err != nil {
			return "", errors.New("invalid auth id [" + authIDString + "]")
		}
		auth, err := ctx.App.Auth.GetByID(authID)
		if err != nil {
			return "", err
		}
		if auth == nil {
			ctx.Session.AddFlash("error:Can't load auth [" + authIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.auth"), nil
		}

		user, err := ctx.App.User.GetByID(auth.UserID, false)
		if err != nil {
			return "", err
		}
		if user == nil {
			ctx.Session.AddFlash("error:Can't load user [" + auth.UserID.String() + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.auth"), nil
		}

		ctx.Title = user.Name
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.auth"), "auths")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.auth.detail", "id", authIDString), authIDString[0:8])...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminAuthDetail(auth, user, ctx, w))
	})
}
