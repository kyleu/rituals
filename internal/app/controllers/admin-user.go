package controllers

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func AdminUserList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "User List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		ctx.Breadcrumbs = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user"), "users")...)
		users, err := ctx.App.User.List()
		if err != nil {
			return 0, err
		}
		return templates.AdminUserList(users, ctx, w)
	})
}

func AdminUserDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		userIDString := mux.Vars(r)["id"]
		userID, err := uuid.FromString(userIDString)
		if err != nil {
			return 0, errors.Wrap(err, "invalid user id ["+userIDString+"]")
		}
		user, err := ctx.App.User.GetByID(userID, false)
		if err != nil {
			return 0, err
		}
		estimates, err := ctx.App.Estimate.GetByMember(userID)
		if err != nil {
			return 0, err
		}
		standups, err := ctx.App.Standup.GetByMember(userID)
		if err != nil {
			return 0, err
		}
		retros, err := ctx.App.Retro.GetByMember(userID)
		if err != nil {
			return 0, err
		}

		ctx.Title = user.Name
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user"), "users")...)
		ctx.Breadcrumbs = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.user.detail", "id", userIDString), user.Name)...)
		return templates.AdminUserDetail(user, estimates, standups, retros, ctx, w)
	})
}
