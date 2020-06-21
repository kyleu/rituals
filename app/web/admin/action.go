package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func ActionList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Action List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyAction, util.Plural(util.KeyAction))
		params := act.ParamSetFromRequest(r)
		actions := ctx.App.Action.List(params.Get(util.KeyAction, ctx.Logger))
		return act.T(admintemplates.ActionList(actions, params, ctx, w))
	})
}

func ActionDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		actionID, err := act.IDFromParams(util.KeyAction, mux.Vars(r))
		if err != nil {
			return act.EResp(err)
		}
		a := ctx.App.Action.GetByID(*actionID)
		if a == nil {
			msg := "can't load action [" + actionID.String() + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.KeyAction), w, r, ctx)
		}
		user := ctx.App.User.GetByID(a.UserID, false)
		if user == nil {
			msg := "can't load user [" + a.UserID.String() + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.KeyUser), w, r, ctx)
		}

		ctx.Title = user.Name
		bc := adminBC(ctx, util.KeyAction, util.Plural(util.KeyAction))
		s := actionID.String()
		bc = append(bc, web.BreadcrumbSelf(s[0:8]))
		ctx.Breadcrumbs = bc

		return act.T(admintemplates.ActionDetail(a, user, ctx, w))
	})
}
