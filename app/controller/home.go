// Package controller - $PF_IGNORE$
package controller

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views"
)

func Home(rc *fasthttp.RequestCtx) {
	Act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		teams, err := as.Services.Team.GetByMember(ps.Context, nil, ps.Profile.ID, params.Get("team", nil, ps.Logger), ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve teams")
		}
		sprints, err := as.Services.Sprint.GetByMember(ps.Context, nil, ps.Profile.ID, params.Get("sprint", nil, ps.Logger), ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve sprints")
		}
		estimates, err := as.Services.Estimate.GetByMember(ps.Context, nil, ps.Profile.ID, params.Get("estimate", nil, ps.Logger), ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve estimates")
		}
		standups, err := as.Services.Standup.GetByMember(ps.Context, nil, ps.Profile.ID, params.Get("standup", nil, ps.Logger), ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve standups")
		}
		retros, err := as.Services.Retro.GetByMember(ps.Context, nil, ps.Profile.ID, params.Get("retro", nil, ps.Logger), ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve retros")
		}
		ps.Data = util.ValueMap{"teams": teams, "sprints": sprints, "estimates": estimates, "standups": standups, "retros": retros}
		return Render(rc, as, &views.Home{Teams: teams, Sprints: sprints, Estimates: estimates, Standups: standups, Retros: retros}, ps)
	})
}
