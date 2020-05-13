package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminRetroList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Retrospective List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro"), "retro")...)
		ctx.Breadcrumbs = bc

		retros, err := ctx.App.Retro.List()
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminRetroList(retros, ctx, w))
	})
}

func AdminRetroDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		retroIDString := mux.Vars(r)["id"]
		retroID, err := uuid.FromString(retroIDString)
		if err != nil {
			return "", errors.New("invalid retro id [" + retroIDString + "]")
		}
		retro, err := ctx.App.Retro.GetByID(retroID)
		if err != nil {
			return "", err
		}
		members, err := ctx.App.Retro.Members.GetByModelID(retroID)
		if err != nil {
			return "", err
		}

		ctx.Title = retro.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro"), "retro")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro.detail", "id", retroIDString), retro.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminRetroDetail(retro, members, ctx, w))
	})
}
