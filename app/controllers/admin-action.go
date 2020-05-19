package controllers

import (
	"github.com/kyleu/rituals.dev/app/util"
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminActionList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Action List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.action"), "actions")...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		actions, err := ctx.App.Action.List(params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminActionList(actions, params, ctx, w))
	})
}

func AdminActionDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		actionIDString := mux.Vars(r)["id"]
		actionID, err := uuid.FromString(actionIDString)
		if err != nil {
			return "", errors.New("invalid action id [" + actionIDString + "]")
		}
		act, err := ctx.App.Action.GetByID(actionID)
		if err != nil {
			return "", err
		}
		if act == nil {
			ctx.Session.AddFlash("error:Can't load action [" + actionIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.action"), nil
		}
		user, err := ctx.App.User.GetByID(act.AuthorID, false)
		if err != nil {
			return "", err
		}
		if user == nil {
			ctx.Session.AddFlash("error:Can't load user [" + act.AuthorID.String() + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.action"), nil
		}

		ctx.Title = user.Name
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.action"), "actions")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.action.detail", "id", actionIDString), actionIDString[0:8])...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminActionDetail(act, user, ctx, w))
	})
}
