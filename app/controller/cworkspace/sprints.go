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
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, err := as.Services.Workspace.CreateSprint(ps.Context, frm.ID, frm.Slug, frm.Title, ps.Profile.ID, frm.Name, frm.Team, ps.Logger)
		if err != nil {
			return "", err
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
		s, err := as.Services.Workspace.LoadSprint(ps.Context, slug, ps.Profile.ID, nil, ps.Logger)
		if err != nil {
			return "", err
		}
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = s.Sprint.TitleString()
		ps.Data = s
		return controller.Render(rc, as, &vworkspace.SprintWorkspace{Sprint: s, Teams: w.Teams}, ps, "sprints", s.Sprint.ID.String())
	})
}
