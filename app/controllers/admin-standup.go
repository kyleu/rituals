package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminStandupList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Daily Standup List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), "standup")...)
		ctx.Breadcrumbs = bc

		standups, err := ctx.App.Standup.List()
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminStandupList(standups, ctx, w))
	})
}

func AdminStandupDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		standupIDString := mux.Vars(r)["id"]
		standupID, err := uuid.FromString(standupIDString)
		if err != nil {
			return "", errors.Wrap(err, "invalid standup id ["+standupIDString+"]")
		}
		standup, err := ctx.App.Standup.GetByID(standupID)
		if err != nil {
			return "", err
		}
		members, err := ctx.App.Standup.Members.GetByModelID(standupID)
		if err != nil {
			return "", err
		}
		reports, err := ctx.App.Standup.GetReports(standupID)
		if err != nil {
			return "", err
		}
		ctx.Title = standup.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), "standup")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup.detail", "id", standupIDString), standup.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminStandupDetail(standup, members, reports, ctx, w))
	})
}
