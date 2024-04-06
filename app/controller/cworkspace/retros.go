package cworkspace

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwretro"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.retro.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Retros"
		ps.Data = ws.Retros
		return controller.Render(w, r, as, &vwretro.RetroList{Retros: ws.Retros, Teams: ws.Teams, Sprints: ws.Sprints}, ps, "retros")
	})
}

func RetroDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.retro", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fr, err := as.Services.Workspace.LoadRetro(p, func() (team.Teams, error) {
			return ws.Teams, nil
		}, func() (sprint.Sprints, error) {
			return ws.Sprints, nil
		})
		if err != nil {
			return "", err
		}
		ps.Title = fr.Retro.TitleString()
		ps.Data = fr
		v := &vwretro.RetroWorkspace{FullRetro: fr, Teams: ws.Teams, Sprints: ws.Sprints}
		return controller.Render(w, r, as, v, ps, "retros", fr.Retro.ID.String())
	})
}

func RetroCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.retro.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(r, ps.RequestBody, ps.Username())
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateRetro(
			ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Accounts.Image(), frm.Team, frm.Sprint, ps.Logger,
		)
		if err != nil {
			return "", errors.Wrap(err, "unable to save retro")
		}
		return controller.FlashAndRedir(true, "New retro created", model.PublicWebPath(), w, ps)
	})
}

func RetroDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.retro.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fr, err := as.Services.Workspace.LoadRetro(p, nil, nil)
		if err != nil {
			return "", err
		}
		err = as.Services.Workspace.DeleteRetro(ps.Context, fr, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Retrospective deleted", "/", w, ps)
	})
}

func RetroAction(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.retro.action", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
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
		_, msg, u, err := as.Services.Workspace.ActionRetro(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, w, ps)
	})
}
