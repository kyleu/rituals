package cworkspace

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app/action"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwteam"
)

func TeamList(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Teams"
		ps.Data = w.Teams
		return controller.Render(rc, as, &vwteam.TeamList{Teams: w.Teams}, ps, "teams")
	})
}

func TeamDetail(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, nil, ps.Params, ps.Logger)
		ft, err := as.Services.Workspace.LoadTeam(p)
		if err != nil {
			return "", err
		}
		if ft.Self == nil {
			return "", errors.New("TODO: Register")
		}
		ps.Title = ft.Team.TitleString()
		ps.Data = ft
		return controller.Render(rc, as, &vwteam.TeamWorkspace{FullTeam: ft}, ps, "teams", ft.Team.ID.String())
	})
}

func TeamCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateTeam(ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New team created", model.PublicWebPath(), rc, ps)
	})
}

func TeamAction(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		act := action.Act(frm.GetStringOpt("action"))
		p := workspace.NewParams(ps.Context, slug, act, frm, ps.Profile, as.Services.Workspace, ps.Logger)
		_, msg, u, err := as.Services.Workspace.ActionTeam(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
