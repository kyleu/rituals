package controllers

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		teams := app.Team(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := app.Sprint(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		estimates := app.Estimate(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := app.Standup(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := app.Retro(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcRetro.Key, ctx.Logger))

		ctx.Title = npncore.AppName
		return npncontroller.T(templates.Index(ctx, teams, sprints, estimates, standups, retros, w))
	})
}
