package cworkspace

import (
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views"
)

func SprintList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("sprint", nil, ps.Logger).Sanitize("sprint")
		ts, err := as.Services.Sprint.GetByOwner(ps.Context, nil, *owner, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = ts
		return controller.Render(rc, as, &views.Debug{}, ps, "sprints")
	})
}

func SprintDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		s, err := as.Services.Workspace.LoadSprint(ps.Context, slug, *owner, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = s
		return controller.Render(rc, as, &views.Debug{}, ps, "sprints", s.Sprint.ID.String())
	})
}
