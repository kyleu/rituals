package cworkspace

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app/action"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwsprint"
)

func SprintList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Sprints"
		ps.Data = w.Sprints
		return controller.Render(rc, as, &vwsprint.SprintList{Sprints: w.Sprints, Teams: w.Teams}, ps, "sprints")
	})
}

func SprintDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile.ID, ps.Profile.NameSafe(), nil, ps.Params, ps.Logger)
		fs, err := as.Services.Workspace.LoadSprint(p)
		if err != nil {
			return "", err
		}
		if fs.Self == nil {
			return "", errors.New("TODO: Register")
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = fs.Sprint.TitleString()
		ps.Data = fs
		return controller.Render(rc, as, &vwsprint.SprintWorkspace{FullSprint: fs, Teams: w.Teams}, ps, "sprints", fs.Sprint.ID.String())
	})
}

func SprintCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateSprint(ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, frm.Team, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New sprint created", model.PublicWebPath(), rc, ps)
	})
}

func SprintAction(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
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
		_, msg, u, err := as.Services.Workspace.ActionSprint(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
