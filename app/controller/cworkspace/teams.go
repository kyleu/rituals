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

func TeamList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("team", nil, ps.Logger).Sanitize("team")
		ts, err := as.Services.Team.GetByOwner(ps.Context, nil, *owner, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = ts
		return controller.Render(rc, as, &views.Debug{}, ps, "teams")
	})
}

func TeamDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		t, err := as.Services.Workspace.LoadTeam(ps.Context, slug, *owner, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = t
		return controller.Render(rc, as, &views.Debug{}, ps, "teams", t.Team.ID.String())
	})
}
