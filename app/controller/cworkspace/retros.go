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
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, err := as.Services.Workspace.CreateRetro(ps.Context, frm.ID, frm.Slug, frm.Title, ps.Profile.ID, frm.Name, frm.Team, frm.Sprint, ps.Logger)
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
		r, err := as.Services.Workspace.LoadRetro(ps.Context, slug, ps.Profile.ID, nil, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = r.Retro.TitleString()
		ps.Data = r
		return controller.Render(rc, as, &vworkspace.RetroWorkspace{Retro: r}, ps, "retros", r.Retro.ID.String())
	})
}
