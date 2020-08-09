package controllers

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"net/http"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func About(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "About " + npncore.AppName
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyAbout)}
		return npncontroller.T(templates.StaticAbout(ctx, w))
	})
}

func Pricing(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Pricing"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("pricing")}
		return npncontroller.T(templates.Pricing(ctx, w))
	})
}

func Features(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Features"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("features")}
		return npncontroller.T(templates.Features(ctx, w))
	})
}

func Community(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Community"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf("community")}
		return npncontroller.T(templates.Community(ctx, w))
	})
}
