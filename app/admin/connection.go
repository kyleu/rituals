package admin

import (
	"encoding/json"
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

type connectionForm struct {
	Svc   string `mapstructure:"svc"`
	Cmd   string `mapstructure:"cmd"`
	Param string `mapstructure:"param"`
}

func ConnectionList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Connection List"
		ctx.Breadcrumbs = adminBC(ctx, npncore.KeyConnection, npncore.Plural(npncore.KeyConnection))

		p := npnweb.ParamSetFromRequest(r)
		connections := app.Socket(ctx.App).List(p.Get(npncore.KeySocket, ctx.Logger))
		return npncontroller.T(admintemplates.ConnectionList(connections, p, ctx, w))
	})
}

func ConnectionDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		connectionID, err := npnweb.IDFromParams(npncore.KeyConnection, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		connection := app.Socket(ctx.App).GetByID(*connectionID)
		ctx.Title = connection.ID.String()
		bc := adminBC(ctx, npncore.KeyConnection, npncore.Plural(npncore.KeyConnection))
		str := connectionID.String()
		bc = append(bc, npnweb.BreadcrumbSelf(str[0:8]))
		ctx.Breadcrumbs = bc

		msg := socket.NewMessage(util.SvcSystem, socket.ServerCmdPong, nil)
		return npncontroller.T(admintemplates.ConnectionDetail(connection, msg, ctx, w))
	})
}

func ConnectionPost(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		connectionID, err := npnweb.IDFromParams(npncore.KeyConnection, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		connection := app.Socket(ctx.App).GetByID(*connectionID)

		frm := &connectionForm{}
		err = npnweb.Decode(r, frm, ctx.Logger)
		if err != nil {
			return npncontroller.EResp(err)
		}

		var param []map[string]interface{}
		_ = json.Unmarshal([]byte(frm.Param), &param)
		svc := util.ServiceFromString(frm.Svc)
		msg := socket.NewMessage(svc, frm.Cmd, param)
		err = app.Socket(ctx.App).WriteMessage(*connectionID, msg)
		if err != nil {
			return npncontroller.EResp(err)
		}

		ctx.Title = connectionID.String()
		bc := adminBC(ctx, npncore.KeyConnection, npncore.Plural(npncore.KeyConnection))
		str := connectionID.String()
		bc = append(bc, npnweb.BreadcrumbSelf(str[0:8]))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.ConnectionDetail(connection, msg, ctx, w))
	})
}
