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

func EstimateList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("estimate", nil, ps.Logger).Sanitize("estimate")
		ts, err := as.Services.Estimate.GetByOwner(ps.Context, nil, *owner, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = ts
		return controller.Render(rc, as, &views.Debug{}, ps, "estimates")
	})
}

func EstimateDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		owner := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		e, err := as.Services.Workspace.LoadEstimate(ps.Context, slug, *owner, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = e
		return controller.Render(rc, as, &views.Debug{}, ps, "estimates", e.Estimate.ID.String())
	})
}
