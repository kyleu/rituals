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
		ctx.Title = util.KeyTitle(util.KeyModules)
		ctx.Breadcrumbs = append(aboutBC(ctx), web.Breadcrumb{Path: ctx.Route(util.KeyModules), Title: util.KeyModules})
		return tmpl(templates.ModulesList(ctx, w))
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.KeyTitle(util.KeyRoutes)
		ctx.Breadcrumbs = append(aboutBC(ctx), web.Breadcrumb{Path: ctx.Route(util.KeyRoutes), Title: util.KeyRoutes})
		return tmpl(templates.RoutesList(ctx, w))
	})
}

func aboutBC(ctx web.RequestContext) web.Breadcrumbs {
	return web.BreadcrumbsSimple(ctx.Route(util.KeyAbout), util.KeyAbout)
}
