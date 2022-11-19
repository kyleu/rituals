package cworkspace

import (
	"github.com/kyleu/rituals/views/vworkspace/vwteam"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
)

func TeamList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Teams"
		ps.Data = w.Teams
		return controller.Render(rc, as, &vwteam.TeamList{Teams: w.Teams}, ps, "teams")
	})
}

func TeamDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		t, err := as.Services.Workspace.LoadTeam(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = t.Team.TitleString()
		ps.Data = t
		return controller.Render(rc, as, &vwteam.TeamWorkspace{Team: t}, ps, "teams", t.Team.ID.String())
	})
}
