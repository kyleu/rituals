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

func StandupList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("standup", nil, ps.Logger).Sanitize("standup")
		ts, err := as.Services.Standup.GetByOwner(ps.Context, nil, *owner, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = ts
		return controller.Render(rc, as, &views.Debug{}, ps, "standups")
	})
}

func StandupDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		u, err := as.Services.Workspace.LoadStandup(ps.Context, slug, *owner, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = u
		return controller.Render(rc, as, &views.Debug{}, ps, "standups", u.Standup.ID.String())
	})
}
