package cworkspace

import (
	"github.com/kyleu/rituals/views/vworkspace/vwstandup"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
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
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		fu, err := as.Services.Workspace.LoadStandup(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		if x := fu.Members.Get(fu.Standup.ID, ps.Profile.ID); x == nil {
			return "", errors.New("TODO: Register")
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = fu.Standup.TitleString()
		ps.Data = fu
		return controller.Render(rc, as, &vwstandup.StandupWorkspace{FullStandup: fu, Teams: w.Teams, Sprints: w.Sprints}, ps, "standups", fu.Standup.ID.String())
	})
}

func StandupCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateStandup(ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, frm.Team, frm.Sprint, ps.Logger)
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
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		_, msg, u, err := as.Services.Workspace.ActionStandup(ps.Context, slug, frm.GetStringOpt("action"), frm, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
