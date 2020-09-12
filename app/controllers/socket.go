package controllers

import (
	"fmt"
	"net/http"

	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Socket(w http.ResponseWriter, r *http.Request) {
	ctx := npnweb.ExtractContext(w, r, true)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.Logger.Info("unable to upgrade connection to websocket")
		return
	}

	connID, err := app.Svc(ctx.App).Socket.Register(ctx.Profile.ToProfile(), c)
	if err != nil {
		ctx.Logger.Warn("unable to register websocket connection")
		return
	}

	err = app.Svc(ctx.App).Socket.ReadLoop(connID)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("error processing socket read loop: %+v", err))
		return
	}
}
