package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func AdminStandupList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Daily Standup List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), "standups")...)
		ctx.Breadcrumbs = bc

		standups, err := ctx.App.Standup.List()
		if err != nil {
			return 0, err
		}
		return templates.AdminStandupList(standups, ctx, w)
	})
}

func AdminStandupDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		standupIDString := mux.Vars(r)["id"]
		standupID, err := uuid.FromString(standupIDString)
		if err != nil {
			return 0, errors.Wrap(err, "invalid standup id ["+standupIDString+"]")
		}
		standup, err := ctx.App.Standup.GetByID(standupID)
		if err != nil {
			return 0, err
		}
		ctx.Title = standup.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), "standups")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup.detail", "id", standupIDString), standup.Slug)...)
		ctx.Breadcrumbs = bc

		return templates.AdminStandupDetail(standup, ctx, w)
	})
}
