package cworkspace

import (
	"github.com/kyleu/rituals/views/vworkspace/vwsprint"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
)

func SprintList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Sprints"
		ps.Data = w.Sprints
		return controller.Render(rc, as, &vwsprint.SprintList{Sprints: w.Sprints, Teams: w.Teams}, ps, "sprints")
	})
}

func SprintDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		s, err := as.Services.Workspace.LoadSprint(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = s.Sprint.TitleString()
		ps.Data = s
		return controller.Render(rc, as, &vwsprint.SprintWorkspace{Sprint: s, Teams: w.Teams}, ps, "sprints", s.Sprint.ID.String())
	})
}
