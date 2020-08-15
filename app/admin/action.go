package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/gorilla/mux"
)

func ActionList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Action List"
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, npncore.KeyAction, npncore.Plural(npncore.KeyAction))
		params := npnweb.ParamSetFromRequest(r)
		actions := app.Action(ctx.App).List(params.Get(npncore.KeyAction, ctx.Logger))
		return npncontroller.T(admintemplates.ActionList(actions, params, ctx, w))
	})
}

func ActionDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		actionID, err := npnweb.IDFromParams(npncore.KeyAction, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		a := app.Action(ctx.App).GetByID(*actionID)
		if a == nil {
			msg := "can't load action [" + actionID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyAction), w, r, ctx)
		}
		user := ctx.App.User().GetByID(a.UserID, false)
		if user == nil {
			msg := "can't load user [" + a.UserID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyUser), w, r, ctx)
		}

		ctx.Title = user.Name
		bc := npncontroller.AdminBC(ctx, npncore.KeyAction, npncore.Plural(npncore.KeyAction))
		s := actionID.String()
		bc = append(bc, npnweb.BreadcrumbSelf(s[0:8]))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.ActionDetail(a, user, ctx, w))
	})
}
