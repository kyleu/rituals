package cworkspace

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace"
)

func EstimateList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Estimates"
		ps.Data = w.Estimates
		return controller.Render(rc, as, &vworkspace.EstimateList{Estimates: w.Estimates, Teams: w.Teams, Sprints: w.Sprints}, ps, "estimates")
	})
}

func EstimateCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, err := as.Services.Workspace.CreateEstimate(ps.Context, frm.ID, frm.Slug, frm.Title, ps.Profile.ID, frm.Name, frm.Team, frm.Sprint, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New estimate created", fmt.Sprintf("/estimate/%s", model.Slug), rc, ps)
	})
}

func EstimateDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		e, err := as.Services.Workspace.LoadEstimate(ps.Context, slug, ps.Profile.ID, nil, ps.Logger)
		if err != nil {
			return "", err
		}
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = e.Estimate.TitleString()
		ps.Data = e
		return controller.Render(rc, as, &vworkspace.EstimateWorkspace{Estimate: e, Teams: w.Teams, Sprints: w.Sprints}, ps, "estimates", e.Estimate.ID.String())
	})
}
