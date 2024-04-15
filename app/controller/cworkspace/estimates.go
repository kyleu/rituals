package cworkspace

import (
	"net/http"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwestimate"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.estimate.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Estimates"
		ps.Data = ws.Estimates
		return controller.Render(r, as, &vwestimate.EstimateList{Estimates: ws.Estimates, Teams: ws.Teams, Sprints: ws.Sprints}, ps, "estimates")
	})
}

func EstimateDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.estimate", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fe, err := as.Services.Workspace.LoadEstimate(p, func() (team.Teams, error) {
			return ws.Teams, nil
		}, func() (sprint.Sprints, error) {
			return ws.Sprints, nil
		})
		if err != nil {
			return "", err
		}
		ps.Title = fe.Estimate.TitleString()
		ps.Data = fe
		page := &vwestimate.EstimateWorkspace{FullEstimate: fe, Teams: ws.Teams, Sprints: ws.Sprints}
		return controller.Render(r, as, page, ps, "estimates", fe.Estimate.ID.String())
	})
}

func EstimateCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.estimate.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(r, ps.RequestBody, ps.Username())
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateEstimate(
			ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Accounts.Image(), frm.Team, frm.Sprint, ps.Logger,
		)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New estimate created", model.PublicWebPath(), ps)
	})
}

func EstimateDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.estimate.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fe, err := as.Services.Workspace.LoadEstimate(p, nil, nil)
		if err != nil {
			return "", err
		}
		err = as.Services.Workspace.DeleteEstimate(ps.Context, fe, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Estimate deleted", "/", ps)
	})
}

func EstimateAction(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.estimate.action", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", err
		}
		act := action.Act(frm.GetStringOpt("action"))
		p := workspace.NewParams(ps.Context, slug, act, frm, ps.Profile, ps.Accounts, as.Services.Workspace, ps.Logger)
		_, msg, u, err := as.Services.Workspace.ActionEstimate(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, ps)
	})
}
