package cworkspace

import (
	"github.com/kyleu/rituals/views/vworkspace/vwestimate"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
)

func EstimateList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Estimates"
		ps.Data = w.Estimates
		return controller.Render(rc, as, &vwestimate.EstimateList{Estimates: w.Estimates, Teams: w.Teams, Sprints: w.Sprints}, ps, "estimates")
	})
}

func EstimateDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		e, err := as.Services.Workspace.LoadEstimate(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = e.Estimate.TitleString()
		ps.Data = e
		return controller.Render(rc, as, &vwestimate.EstimateWorkspace{Estimate: e, Teams: w.Teams, Sprints: w.Sprints}, ps, "estimates", e.Estimate.ID.String())
	})
}
