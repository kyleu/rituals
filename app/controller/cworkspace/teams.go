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
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views"
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
		frm, err := parseRequestForm(rc)
		if err != nil {
			return "", err
		}
		model := &team.Team{ID: frm.ID, Slug: frm.Slug, Title: frm.Title, Status: enum.SessionStatusNew, Owner: ps.Profile.ID, Created: time.Now()}
		err = as.Services.Team.Create(ps.Context, nil, ps.Logger, model)
		if err != nil {
			return "", errors.Wrap(err, "unable to save team")
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
		t, err := as.Services.Workspace.LoadTeam(ps.Context, slug, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = t.Team.TitleString()
		ps.Data = t
		return controller.Render(rc, as, &views.Debug{}, ps, "teams", t.Team.ID.String())
	})
}
