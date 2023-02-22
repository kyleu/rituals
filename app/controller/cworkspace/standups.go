package cworkspace

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwstandup"
)

func StandupList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Standups"
		ps.Data = w.Standups
		return controller.Render(rc, as, &vwstandup.StandupList{Standups: w.Standups, Teams: w.Teams, Sprints: w.Sprints}, ps, "standups")
	})
}

func StandupDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fu, err := as.Services.Workspace.LoadStandup(p, func() (team.Teams, error) {
			return w.Teams, nil
		}, func() (sprint.Sprints, error) {
			return w.Sprints, nil
		})
		if err != nil {
			return "", err
		}
		if fu.Self == nil {
			return "", errors.New("TODO: Register")
		}
		ps.Title = fu.Standup.TitleString()
		ps.Data = fu
		return controller.Render(rc, as, &vwstandup.StandupWorkspace{FullStandup: fu, Teams: w.Teams, Sprints: w.Sprints}, ps, "standups", fu.Standup.ID.String())
	})
}

func StandupCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Username())
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateStandup(
			ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Accounts.Image(), frm.Team, frm.Sprint, ps.Logger,
		)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New standup created", model.PublicWebPath(), rc, ps)
	})
}

func StandupAction(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		act := action.Act(frm.GetStringOpt("action"))
		p := workspace.NewParams(ps.Context, slug, act, frm, ps.Profile, ps.Accounts, as.Services.Workspace, ps.Logger)
		_, msg, u, err := as.Services.Workspace.ActionStandup(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
