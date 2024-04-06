package cworkspace

import (
	"net/http"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwstandup"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.standup.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Standups"
		ps.Data = ws.Standups
		return controller.Render(w, r, as, &vwstandup.StandupList{Standups: ws.Standups, Teams: ws.Teams, Sprints: ws.Sprints}, ps, "standups")
	})
}

func StandupDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.standup", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fu, err := as.Services.Workspace.LoadStandup(p, func() (team.Teams, error) {
			return ws.Teams, nil
		}, func() (sprint.Sprints, error) {
			return ws.Sprints, nil
		})
		if err != nil {
			return "", err
		}
		ps.Title = fu.Standup.TitleString()
		ps.Data = fu
		return controller.Render(w, r, as, &vwstandup.StandupWorkspace{FullStandup: fu, Teams: ws.Teams, Sprints: ws.Sprints}, ps, "standups", fu.Standup.ID.String())
	})
}

func StandupCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.standup.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(r, ps.RequestBody, ps.Username())
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateStandup(
			ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Accounts.Image(), frm.Team, frm.Sprint, ps.Logger,
		)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New standup created", model.PublicWebPath(), w, ps)
	})
}

func StandupDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.standup.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fu, err := as.Services.Workspace.LoadStandup(p, nil, nil)
		if err != nil {
			return "", err
		}
		err = as.Services.Workspace.DeleteStandup(ps.Context, fu, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Standup deleted", "/", w, ps)
	})
}

func StandupAction(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.standup.action", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
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
		_, msg, u, err := as.Services.Workspace.ActionStandup(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, w, ps)
	})
}
