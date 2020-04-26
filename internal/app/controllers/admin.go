package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Admin"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		return templates.AdminHome(ctx, w)
	})
}
