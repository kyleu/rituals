package cworkspace

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

func TeamSocket(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.socket", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(rc, util.KeyTeam, as, ps)
	})
}

func SprintSocket(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.socket", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(rc, util.KeySprint, as, ps)
	})
}

func EstimateSocket(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.estimate.socket", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(rc, util.KeyEstimate, as, ps)
	})
}

func StandupSocket(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.standup.socket", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(rc, util.KeyStandup, as, ps)
	})
}

func RetroSocket(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.retro.socket", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Socket(rc, util.KeyRetro, as, ps)
	})
}

func Socket(rc *fasthttp.RequestCtx, svc string, as *app.State, ps *cutil.PageState) (string, error) {
	id, err := cutil.RCRequiredUUID(rc, "id")
	if err != nil {
		return "", err
	}
	err = as.Services.Socket.Upgrade(ps.Context, rc, svc+":"+id.String(), ps.User, ps.Profile, ps.Accounts, ps.Logger)
	if err != nil {
		ps.Logger.Warnf("unable to upgrade connection to WebSocket: %s", err.Error())
		return "", err
	}
	return "", nil
}
