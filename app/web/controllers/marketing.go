package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/app/web/act"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func About(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "About " + util.AppName
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.KeyAbout), util.KeyAbout)
		return act.T(templates.StaticAbout(ctx, w))
	})
}

func Pricing(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Pricing"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("pricing"), "pricing")
		return act.T(templates.Pricing(ctx, w))
	})
}

func Features(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Features"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("features"), "features")
		return act.T(templates.Features(ctx, w))
	})
}

func Community(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Community"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("community"), "community")
		return act.T(templates.Community(ctx, w))
	})
}
