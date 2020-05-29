package admin

import (
	"encoding/json"
	"github.com/kyleu/rituals.dev/app/controllers/form"
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
		connections := ctx.App.Socket.List(p.Get(util.KeySocket, ctx.Logger))
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

		msg := socket.NewMessage(util.SvcSystem, socket.ServerCmdPong, nil)
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

		frm := &form.ConnectionForm{}
		err = form.Decode(r, frm, ctx.Logger)
		if err != nil {
			return eresp(err, "")
		}

		var param []map[string]interface{}
		_ = json.Unmarshal([]byte(frm.Param), &param)
		svc := util.ServiceFromString(frm.Svc)
		msg := socket.NewMessage(svc, frm.Cmd, param)
		err = ctx.App.Socket.WriteMessage(*connectionID, msg)
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
