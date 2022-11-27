package cworkspace

import (
	"github.com/kyleu/rituals/views/vworkspace/vwretro"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
)

func RetroList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Retros"
		ps.Data = w.Retros
		return controller.Render(rc, as, &vwretro.RetroList{Retros: w.Retros, Teams: w.Teams, Sprints: w.Sprints}, ps, "retros")
	})
}

func RetroDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		fr, err := as.Services.Workspace.LoadRetro(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		if fr.Self == nil {
			return "", errors.New("TODO: Register")
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = fr.Retro.TitleString()
		ps.Data = fr
		v := &vwretro.RetroWorkspace{FullRetro: fr, Teams: w.Teams, Sprints: w.Sprints}
		return controller.Render(rc, as, v, ps, "retros", fr.Retro.ID.String())
	})
}

func RetroCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateRetro(ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, frm.Team, frm.Sprint, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to save retro")
		}
		return controller.FlashAndRedir(true, "New retro created", model.PublicWebPath(), rc, ps)
	})
}

func RetroAction(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		_, msg, u, err := as.Services.Workspace.ActionRetro(ps.Context, slug, frm.GetStringOpt("action"), frm, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
