package controllers

import (
	"github.com/kyleu/rituals.dev/app/util"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminInvitationList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Invitation List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invitation"), "invitations")...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		invitations, err := ctx.App.Invitation.List(params.Get(util.KeyInvitation, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminInvitationList(invitations, params, ctx, w))
	})
}

func AdminInvitationDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sess, err := ctx.App.Invitation.GetByKey(key)
		if err != nil {
			return "", err
		}
		ctx.Title = sess.Key
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invitation"), util.KeyInvitation)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invitation.detail", "key", key), key)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminInvitationDetail(sess, ctx, w))
	})
}
