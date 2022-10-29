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
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views"
	"github.com/kyleu/rituals/views/vworkspace"
)

func RetroList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Retros"
		ps.Data = w.Retros
		return controller.Render(rc, as, &vworkspace.RetroList{Retros: w.Retros, Teams: w.Teams, Sprints: w.Sprints}, ps, "retros")
	})
}

func RetroCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc)
		if err != nil {
			return "", err
		}
		model := &retro.Retro{
			ID: frm.ID, Slug: frm.Slug, Title: frm.Title, Status: enum.SessionStatusNew,
			TeamID: frm.Team, SprintID: frm.Sprint, Owner: ps.Profile.ID, Created: time.Now(),
		}
		err = as.Services.Retro.Create(ps.Context, nil, ps.Logger, model)
		if err != nil {
			return "", errors.Wrap(err, "unable to save retro")
		}
		return controller.FlashAndRedir(true, "New retro created", fmt.Sprintf("/retro/%s", model.Slug), rc, ps)
	})
}

func RetroDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		r, err := as.Services.Workspace.LoadRetro(ps.Context, slug, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = r.Retro.TitleString()
		ps.Data = r
		return controller.Render(rc, as, &views.Debug{}, ps, "retros", r.Retro.ID.String())
	})
}
