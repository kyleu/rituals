package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/app/web/act"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)

		teams := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		estimates := ctx.App.Estimate.GetByMember(ctx.Profile.UserID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := ctx.App.Standup.GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := ctx.App.Retro.GetByMember(ctx.Profile.UserID, params.Get(util.SvcRetro.Key, ctx.Logger))

		ctx.Title = util.AppName
		return tmpl(templates.Index(ctx, teams, sprints, estimates, standups, retros, w))
	})
}
