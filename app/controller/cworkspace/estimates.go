package cworkspace

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app/action"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwestimate"
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
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile.ID, ps.Profile.NameSafe(), nil, ps.Params, ps.Logger)
		fe, err := as.Services.Workspace.LoadEstimate(p)
		if err != nil {
			return "", err
		}
		if fe.Self == nil {
			return "", errors.New("TODO: Register")
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = fe.Estimate.TitleString()
		ps.Data = fe
		page := &vwestimate.EstimateWorkspace{FullEstimate: fe, Teams: w.Teams, Sprints: w.Sprints}
		return controller.Render(rc, as, page, ps, "estimates", fe.Estimate.ID.String())
	})
}

func EstimateCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateEstimate(ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, frm.Team, frm.Sprint, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New estimate created", model.PublicWebPath(), rc, ps)
	})
}

func EstimateAction(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		act := action.Act(frm.GetStringOpt("action"))
		p := workspace.NewParams(ps.Context, slug, act, frm, ps.Profile.ID, as.Services.Workspace, ps.Logger)
		_, msg, u, err := as.Services.Workspace.ActionEstimate(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
