package controllers

import (
	"fmt"
	"github.com/kyleu/rituals.dev/internal/app/socket"
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Socket(w http.ResponseWriter, r *http.Request) {
	ctx := web.ExtractContext(w, r)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.Logger.Info("unable to upgrade connection to websocket")
		return
	}

	connID, err := ctx.App.Socket.Register(ctx.Profile.UserID, c)
	if err != nil {
		ctx.Logger.Warn("unable to register websocket connection")
		return
	}

	msg := socket.Message{Svc: "system", Cmd: "profile", Param: ctx.Profile.ToProfile()}
	err = ctx.App.Socket.WriteMessage(connID, &msg)
	if err != nil {
		ctx.Logger.Warn("unable to send initial profile")
	}

	err = ctx.App.Socket.ReadLoop(connID)
	if err != nil {
		ctx.Logger.Warn(fmt.Sprintf("error processing socket read loop: %+v", err))
		return
	}
}
