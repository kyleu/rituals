package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminActionList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Action List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.action"), "actions")...)
		ctx.Breadcrumbs = bc

		actions, err := ctx.App.Action.List()
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminActionList(actions, ctx, w))
	})
}

func AdminActionDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		actionIDString := mux.Vars(r)["id"]
		actionID, err := uuid.FromString(actionIDString)
		if err != nil {
			return "", errors.New("invalid action id [" + actionIDString + "]")
		}
		action, err := ctx.App.Action.GetByID(actionID)
		if err != nil {
			return "", err
		}
		user, err := ctx.App.User.GetByID(action.AuthorID, false)
		if err != nil {
			return "", err
		}

		ctx.Title = user.Name
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.action"), "actions")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.action.detail", "id", actionIDString), actionIDString[0:8])...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminActionDetail(action, user, ctx, w))
	})
}
