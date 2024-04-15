package cworkspace

import (
	"net/http"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwteam"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.team.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Teams"
		ps.Data = ws.Teams
		return controller.Render(r, as, &vwteam.TeamList{Teams: ws.Teams}, ps, "teams")
	})
}

func TeamDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.team", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		ft, err := as.Services.Workspace.LoadTeam(p)
		if err != nil {
			return "", err
		}
		ps.Title = ft.Team.TitleString()
		ps.Data = ft
		return controller.Render(r, as, &vwteam.TeamWorkspace{FullTeam: ft}, ps, "teams", ft.Team.ID.String())
	})
}

func TeamCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.team.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(r, ps.RequestBody, ps.Username())
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateTeam(
			ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Accounts.Image(), ps.Logger,
		)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New team created", model.PublicWebPath(), ps)
	})
}

func TeamDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.team.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		ft, err := as.Services.Workspace.LoadTeam(p)
		if err != nil {
			return "", err
		}
		err = as.Services.Workspace.DeleteTeam(ps.Context, ft, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Team deleted", "/", ps)
	})
}

func TeamAction(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.team.action", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", err
		}
		act := action.Act(frm.GetStringOpt("action"))
		p := workspace.NewParams(ps.Context, slug, act, frm, ps.Profile, ps.Accounts, as.Services.Workspace, ps.Logger)
		_, msg, u, err := as.Services.Workspace.ActionTeam(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, ps)
	})
}
