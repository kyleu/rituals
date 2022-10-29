package cworkspace

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views"
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
		frm, err := parseRequestForm(rc)
		if err != nil {
			return "", err
		}
		model := &standup.Standup{
			ID: frm.ID, Slug: frm.Slug, Title: frm.Title, Status: enum.SessionStatusNew,
			TeamID: frm.Team, SprintID: frm.Sprint, Owner: ps.Profile.ID, Created: time.Now(),
		}
		err = as.Services.Standup.Create(ps.Context, nil, ps.Logger, model)
		if err != nil {
			return "", errors.Wrap(err, "unable to save standup")
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
		u, err := as.Services.Workspace.LoadStandup(ps.Context, slug, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = u.Standup.TitleString()
		ps.Data = u
		return controller.Render(rc, as, &views.Debug{}, ps, "standups", u.Standup.ID.String())
	})
}
