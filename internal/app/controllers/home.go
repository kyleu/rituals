package controllers

import (
	"github.com/kyleu/rituals.dev/internal/app/util"
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Home"
		return templates.Index(ctx, w)
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "About " + util.AppName
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("about"), "about")
		return templates.About(ctx, w)
	})
}
