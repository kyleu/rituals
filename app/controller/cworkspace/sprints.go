package cworkspace

import (
	"net/http"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/vworkspace/vwsprint"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.sprint.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Title = "Sprints"
		ps.Data = ws.Sprints
		return controller.Render(w, r, as, &vwsprint.SprintList{Sprints: ws.Sprints, Teams: ws.Teams}, ps, "sprints")
	})
}

func SprintDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.sprint", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fs, err := as.Services.Workspace.LoadSprint(p, func() (team.Teams, error) {
			return ws.Teams, nil
		})
		if err != nil {
			return "", err
		}
		ps.Title = fs.Sprint.TitleString()
		ps.Data = fs
		return controller.Render(w, r, as, &vwsprint.SprintWorkspace{FullSprint: fs, Teams: ws.Teams}, ps, "sprints", fs.Sprint.ID.String())
	})
}

func SprintCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.sprint.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(r, ps.RequestBody, ps.Username())
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateSprint(
			ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Accounts.Image(), frm.Team, ps.Logger,
		)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New sprint created", model.PublicWebPath(), w, ps)
	})
}

func SprintDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.sprint.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.PathString(r, "slug", false)
		if err != nil {
			return "", err
		}
		p := workspace.NewLoadParams(ps.Context, slug, ps.Profile, ps.Accounts, nil, ps.Params, ps.Logger)
		fs, err := as.Services.Workspace.LoadSprint(p, nil)
		if err != nil {
			return "", err
		}
		err = as.Services.Workspace.DeleteSprint(ps.Context, fs, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Sprint deleted", "/", w, ps)
	})
}

func SprintAction(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.sprint.action", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
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
		_, msg, u, err := as.Services.Workspace.ActionSprint(p)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, w, ps)
	})
}
