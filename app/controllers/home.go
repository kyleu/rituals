package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		sprints, err := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, 5)
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetByMember(ctx.Profile.UserID, 5)
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetByMember(ctx.Profile.UserID, 5)
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetByMember(ctx.Profile.UserID, 5)
		if err != nil {
			return "", err
		}

		ctx.Title = "Home"
		return tmpl(templates.Index(ctx, sprints, estimates, standups, retros, w))
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "About " + util.AppName
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("about"), "about")
		return tmpl(templates.About(ctx, w))
	})
}
