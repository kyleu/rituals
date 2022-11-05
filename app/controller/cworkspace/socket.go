package cworkspace

import (
	"fmt"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
)

var upgrader = websocket.FastHTTPUpgrader{
	EnableCompression: true,
}

func Socket(rc *fasthttp.RequestCtx) {
	controller.Act("socket", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		err := upgrader.Upgrade(rc, func(conn *websocket.Conn) {
			connID, err := as.Services.Socket.Register(ps.Profile, conn)
			if err != nil {
				ps.Logger.Warn("unable to register websocket connection")
				return
			}
			joined, err := as.Services.Socket.Join(connID.ID, "TODO")
			if err != nil {
				ps.Logger.Error(fmt.Sprintf("error processing socket join (%v): %+v", joined, err))
				return
			}
			err = as.Services.Socket.ReadLoop(connID.ID, nil)
			if err != nil {
				ps.Logger.Error(fmt.Sprintf("error processing socket read loop: %+v", err))
				return
			}
		})
		if err != nil {
			ps.Logger.Warn("unable to upgrade connection to websocket")
			return "", err
		}
		return "", nil
	})
}
