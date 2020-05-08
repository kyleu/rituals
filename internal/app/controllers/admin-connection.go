package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/socket"
	"github.com/kyleu/rituals.dev/internal/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func AdminConnectionList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Connection List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.connection"), "connections")...)
		ctx.Breadcrumbs = bc

		connections, err := ctx.App.Socket.List()
		if err != nil {
			return 0, err
		}
		return templates.AdminConnectionList(connections, ctx, w)
	})
}

func AdminConnectionDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		connectionIDString := mux.Vars(r)["id"]
		connectionID, err := uuid.FromString(connectionIDString)
		if err != nil {
			return 0, err
		}
		connection, err := ctx.App.Socket.GetByID(connectionID)
		if err != nil {
			return 0, err
		}
		ctx.Title = connection.ID.String()
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.connection"), "connections")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.connection.detail", "id", connectionIDString), connectionIDString)...)
		ctx.Breadcrumbs = bc

		msg := socket.Message{Svc: util.SvcSystem, Cmd: util.ServerCmdPong, Param: nil}
		return templates.AdminConnectionDetail(connection, msg, ctx, w)
	})
}

func AdminConnectionPost(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		_ = r.ParseForm()
		connectionIDString := r.Form.Get("id")
		connectionID, err := uuid.FromString(connectionIDString)
		if err != nil {
			return 0, err
		}
		connection, err := ctx.App.Socket.GetByID(connectionID)
		if err != nil {
			return 0, err
		}

		svc := r.Form.Get("svc")
		cmd := r.Form.Get("cmd")
		paramString := r.Form.Get("param")
		var param []map[string]interface{}
		_ = json.Unmarshal([]byte(paramString), &param)
		msg := socket.Message{Svc: svc, Cmd: cmd, Param: param}
		err = ctx.App.Socket.WriteMessage(connectionID, &msg)
		if err != nil {
			return 0, err
		}

		ctx.Title = connectionID.String()
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.connection"), "connections")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.connection.detail", "id", connectionIDString), connectionIDString)...)
		ctx.Breadcrumbs = bc

		return templates.AdminConnectionDetail(connection, msg, ctx, w)
	})
}
