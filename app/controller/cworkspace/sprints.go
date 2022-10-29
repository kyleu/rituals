package cworkspace

import (
	"fmt"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"time"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views"
)

func SprintList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Sprints"
		ps.Data = w.Sprints
		return controller.Render(rc, as, &vworkspace.SprintList{Sprints: w.Sprints, Teams: w.Teams}, ps, "sprints")
	})
}

func SprintCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc)
		if err != nil {
			return "", err
		}
		model := &sprint.Sprint{
			ID: frm.ID, Slug: frm.Slug, Title: frm.Title, Status: enum.SessionStatusNew, TeamID: frm.Team, Owner: ps.Profile.ID, Created: time.Now(),
		}
		err = as.Services.Sprint.Create(ps.Context, nil, ps.Logger, model)
		if err != nil {
			return "", errors.Wrap(err, "unable to save sprint")
		}
		return controller.FlashAndRedir(true, "New sprint created", fmt.Sprintf("/sprint/%s", model.Slug), rc, ps)
	})
}

func SprintDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		s, err := as.Services.Workspace.LoadSprint(ps.Context, slug, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = s.Sprint.TitleString()
		ps.Data = s
		return controller.Render(rc, as, &views.Debug{}, ps, "sprints", s.Sprint.ID.String())
	})
}
