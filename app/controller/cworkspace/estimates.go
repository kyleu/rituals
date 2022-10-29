package cworkspace

import (
	"fmt"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"time"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views"
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
		frm, err := parseRequestForm(rc)
		if err != nil {
			return "", err
		}
		model := &estimate.Estimate{
			ID: frm.ID, Slug: frm.Slug, Title: frm.Title, Status: enum.SessionStatusNew,
			TeamID: frm.Team, SprintID: frm.Sprint, Owner: ps.Profile.ID, Created: time.Now(),
		}
		err = as.Services.Estimate.Create(ps.Context, nil, ps.Logger, model)
		if err != nil {
			return "", errors.Wrap(err, "unable to save estimate")
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
		e, err := as.Services.Workspace.LoadEstimate(ps.Context, slug, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = e.Estimate.TitleString()
		ps.Data = e
		return controller.Render(rc, as, &views.Debug{}, ps, "estimates", e.Estimate.ID.String())
	})
}
