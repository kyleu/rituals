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
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		r, err := as.Services.Workspace.LoadRetro(ps.Context, slug, ps.Profile.ID, nil, ps.Params, ps.Logger)
		if err != nil {
			return "", err
		}
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = r.Retro.TitleString()
		ps.Data = r
		return controller.Render(rc, as, &vwretro.RetroWorkspace{Retro: r, Teams: w.Teams, Sprints: w.Sprints}, ps, "retros", r.Retro.ID.String())
	})
}
