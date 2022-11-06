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

func TeamList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Teams"
		ps.Data = w.Teams
		return controller.Render(rc, as, &vworkspace.TeamList{Teams: w.Teams}, ps, "teams")
	})
}

func TeamCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, err := as.Services.Workspace.CreateTeam(ps.Context, frm.ID, frm.Slug, frm.Title, ps.Profile.ID, frm.Name, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New team created", fmt.Sprintf("/team/%s", model.Slug), rc, ps)
	})
}

func TeamDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		t, err := as.Services.Workspace.LoadTeam(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = t.Team.TitleString()
		ps.Data = t
		return controller.Render(rc, as, &vworkspace.TeamWorkspace{Team: t}, ps, "teams", t.Team.ID.String())
	})
}
