package cworkspace

import (
	"github.com/kyleu/rituals/views/vworkspace/vwstandup"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
)

func StandupList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Standups"
		ps.Data = w.Standups
		return controller.Render(rc, as, &vwstandup.StandupList{Standups: w.Standups, Teams: w.Teams, Sprints: w.Sprints}, ps, "standups")
	})
}

func StandupDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		u, err := as.Services.Workspace.LoadStandup(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = u.Standup.TitleString()
		ps.Data = u
		return controller.Render(rc, as, &vwstandup.StandupWorkspace{Standup: u, Teams: w.Teams, Sprints: w.Sprints}, ps, "standups", u.Standup.ID.String())
	})
}
