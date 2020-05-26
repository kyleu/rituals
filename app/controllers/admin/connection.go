package admin

import (
	"encoding/json"
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func ConnectionList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Connection List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyConnection, util.Plural(util.KeyConnection))

		p := act.ParamSetFromRequest(r)
		connections, err := ctx.App.Socket.List(p.Get(util.KeySocket, ctx.Logger))
		if err != nil {
			return eresp(err, "")
		}
		return tmpl(templates.AdminConnectionList(connections, p, ctx, w))
	})
}

func ConnectionDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		connectionID, err := act.IDFromParams(util.KeyConnection, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		connection, err := ctx.App.Socket.GetByID(*connectionID)
		if err != nil {
			return eresp(err, "")
		}
		ctx.Title = connection.ID.String()
		bc := adminBC(ctx, util.KeyConnection, util.Plural(util.KeyConnection))
		link := util.AdminLink(util.KeyConnection, util.KeyDetail)
		str := connectionID.String()
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, str), str[0:8])...)
		ctx.Breadcrumbs = bc

		msg := socket.Message{Svc: util.SvcSystem.Key, Cmd: socket.ServerCmdPong, Param: nil}
		return tmpl(templates.AdminConnectionDetail(connection, msg, ctx, w))
	})
}

func ConnectionPost(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		connectionID, err := act.IDFromParams(util.KeyConnection, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		connection, err := ctx.App.Socket.GetByID(*connectionID)
		if err != nil {
			return eresp(err, "")
		}

		_ = r.ParseForm()
		svc := r.Form.Get(util.KeySvc)
		cmd := r.Form.Get("cmd")
		paramString := r.Form.Get("param")
		var param []map[string]interface{}
		_ = json.Unmarshal([]byte(paramString), &param)
		msg := socket.Message{Svc: svc, Cmd: cmd, Param: param}
		err = ctx.App.Socket.WriteMessage(*connectionID, &msg)
		if err != nil {
			return eresp(err, "")
		}

		ctx.Title = connectionID.String()
		bc := adminBC(ctx, util.KeyConnection, util.Plural(util.KeyConnection))
		link := util.AdminLink(util.KeyConnection, util.KeyDetail)
		str := connectionID.String()
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, str), str[0:8])...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminConnectionDetail(connection, msg, ctx, w))
	})
}
