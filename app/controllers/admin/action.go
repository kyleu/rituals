package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func ActionList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Action List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyAction, "actions")
		params := act.ParamSetFromRequest(r)
		actions, err := ctx.App.Action.List(params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminActionList(actions, params, ctx, w))
	})
}

func ActionDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		actionID := util.GetUUIDPointer(mux.Vars(r), "id")
		if actionID == nil {
			return "", errors.New("invalid action id")
		}
		a, err := ctx.App.Action.GetByID(*actionID)
		if err != nil {
			return "", err
		}
		if a == nil {
			ctx.Session.AddFlash("error:Can't load action [" + actionID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.KeyAction)), nil
		}
		user, err := ctx.App.User.GetByID(a.AuthorID, false)
		if err != nil {
			return "", err
		}
		if user == nil {
			ctx.Session.AddFlash("error:Can't load user [" + a.AuthorID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.KeyAction)), nil
		}

		ctx.Title = user.Name
		bc := adminBC(ctx, util.KeyAction, "actions")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.KeyAction, util.KeyDetail), "id", actionID.String()), actionID.String()[0:8])...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminActionDetail(a, user, ctx, w))
	})
}
