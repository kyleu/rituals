package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func Health(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_, _ = w.Write([]byte("OK"))
		return "", nil
	})
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Modules"
		ctx.Breadcrumbs = append(aboutBC(ctx), web.Breadcrumb{Path: ctx.Route("modules"), Title: "modules"})
		return tmpl(templates.ModulesList(ctx, w))
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Routes"
		ctx.Breadcrumbs = append(aboutBC(ctx), web.Breadcrumb{Path: ctx.Route("routes"), Title: "routes"})
		return tmpl(templates.RoutesList(ctx, w))
	})
}

func aboutBC(ctx web.RequestContext) web.Breadcrumbs {
	return web.BreadcrumbsSimple(ctx.Route(util.KeyAbout), util.KeyAbout)
}
