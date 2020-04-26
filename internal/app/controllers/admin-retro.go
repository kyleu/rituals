package controllers

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func AdminRetroList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Retrospective List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		ctx.Breadcrumbs = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro"), "retros")...)
		retros, err := ctx.App.Retro.List()
		if err != nil {
			return 0, err
		}
		return templates.AdminRetroList(retros, ctx, w)
	})
}

func AdminRetroDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		retroIDString := mux.Vars(r)["id"]
		retroID, err := uuid.FromString(retroIDString)
		if err != nil {
			return 0, errors.Wrap(err, "invalid retro id ["+retroIDString+"]")
		}
		retro, err := ctx.App.Retro.GetByID(retroID)
		if err != nil {
			return 0, err
		}
		ctx.Title = retro.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro"), "retros")...)
		ctx.Breadcrumbs = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro.detail", "id", retroIDString), retro.Slug)...)
		return templates.AdminRetroDetail(retro, ctx, w)
	})
}
