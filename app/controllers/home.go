package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

// func Temp(w http.ResponseWriter, r *http.Request) {
// 	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
// 		s := `{
// 	"associatedApplications": [
// 		{
// 			"applicationId": "f2187a97-e0ee-4f52-8e58-ab527a84fc69"
// 		}
// 	]
// }`
// 		w.Header().Set("Content-Type", "application/json")
// 		_, _ = w.Write([]byte(s))
// 		return "", nil
// 	})
// }

func Home(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)

		teams, err := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		sprints, err := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetByMember(ctx.Profile.UserID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetByMember(ctx.Profile.UserID, params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = "Home"
		return tmpl(templates.Index(ctx, teams, sprints, estimates, standups, retros, w))
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "About " + util.AppName
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.KeyAbout), util.KeyAbout)
		return tmpl(templates.About(ctx, w))
	})
}
