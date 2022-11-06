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

func StandupList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Standups"
		ps.Data = w.Standups
		return controller.Render(rc, as, &vworkspace.StandupList{Standups: w.Standups, Teams: w.Teams, Sprints: w.Sprints}, ps, "standups")
	})
}

func StandupCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, err := as.Services.Workspace.CreateStandup(ps.Context, frm.ID, frm.Slug, frm.Title, ps.Profile.ID, frm.Name, frm.Team, frm.Sprint, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New standup created", fmt.Sprintf("/standup/%s", model.Slug), rc, ps)
	})
}

func StandupDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		u, err := as.Services.Workspace.LoadStandup(ps.Context, slug, ps.Profile.ID, nil, ps.Logger)
		if err != nil {
			return "", err
		}
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = u.Standup.TitleString()
		ps.Data = u
		return controller.Render(rc, as, &vworkspace.StandupWorkspace{Standup: u, Teams: w.Teams, Sprints: w.Sprints}, ps, "standups", u.Standup.ID.String())
	})
}
