package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminInviteList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Invitation List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invite"), "invites")...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		invites, err := ctx.App.Invite.List(params.Get("invite"))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminInviteList(invites, params, ctx, w))
	})
}

func AdminInviteDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sess, err := ctx.App.Invite.GetByKey(key)
		if err != nil {
			return "", err
		}
		ctx.Title = sess.Key
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invite"), "invites")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invite.detail", "key", key), key)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminInviteDetail(sess, ctx, w))
	})
}
