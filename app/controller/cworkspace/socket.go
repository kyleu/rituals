package cworkspace

import (
	"net/http"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

func TeamSocket(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.team.socket", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(w, r, util.KeyTeam, as, ps)
	})
}

func SprintSocket(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.sprint.socket", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(w, r, util.KeySprint, as, ps)
	})
}

func EstimateSocket(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.estimate.socket", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(w, r, util.KeyEstimate, as, ps)
	})
}

func StandupSocket(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.standup.socket", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(w, r, util.KeyStandup, as, ps)
	})
}

func RetroSocket(w http.ResponseWriter, r *http.Request) {
	controller.Act("workspace.retro.socket", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(w, r, util.KeyRetro, as, ps)
	})
}

func Socket(w http.ResponseWriter, r *http.Request, svc string, as *app.State, ps *cutil.PageState) (string, error) {
	id, err := cutil.PathUUID(r, "id")
	if err != nil {
		return "", err
	}
	h := as.Services.Workspace.SocketHandler
	sockID, err := as.Services.Socket.Upgrade(ps.Context, w, r, svc+":"+id.String(), ps.User, ps.Profile, ps.Accounts, h, ps.Logger)
	if err != nil {
		ps.Logger.Warnf("unable to upgrade connection to WebSocket: %s", err.Error())
		return "", err
	}
	return "", as.Services.Socket.ReadLoop(ps.Context, sockID, ps.Logger)
}
