package controllers

import (
	"github.com/kyleu/rituals.dev/internal/app/util"
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Home"
		return templates.Index(ctx, w)
	})
}

var upgrader = websocket.Upgrader{}

func Socket(w http.ResponseWriter, r *http.Request) {
	ctx := web.ExtractContext(r)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.Logger.Info("Unable to upgrade connection to websocket")
		return
	}
	defer func() {
		_ = c.Close()
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		ctx.Logger.Debug("Received message on websocket: " + string(message))
		err = c.WriteMessage(mt, message)
		if err != nil {
			ctx.Logger.Warn("Unable to write to websocket")
			break
		}
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "About " + util.AppName
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("about"), "about")
		return templates.About(ctx, w)
	})
}
