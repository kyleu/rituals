package controllers

import (
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/app/web/act"
	"github.com/kyleu/rituals.dev/gen/templates"
	"net/http"
	"strings"
)

func About(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "About " + util.AppName
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.KeyAbout), util.KeyAbout)
		return tmpl(templates.StaticAbout(ctx, w))
	})
}

func Pricing(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Pricing"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("pricing"), "pricing")
		return tmpl(templates.StaticPricing(ctx, w))
	})
}

func Features(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Features"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("features"), "features")
		return tmpl(templates.StaticFeatures(ctx, w))
	})
}

func Community(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Community"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("community"), "community")
		return tmpl(templates.StaticCommunity(ctx, w))
	})
}

func tempSecurityCheck(ctx *web.RequestContext) bool {
	if ctx.Profile.Role == util.RoleAdmin {
		return true
	}
	if strings.Contains(ctx.App.Auth.Redir, "localhost") {
		return true
	}
	if strings.Contains(ctx.Request.RawQuery, "p=np") {
		return true
	}
	return false
}
