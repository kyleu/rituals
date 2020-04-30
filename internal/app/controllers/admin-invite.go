package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func AdminInviteList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Invitation List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invite"), "invites")...)
		ctx.Breadcrumbs = bc

		invites, err := ctx.App.Invite.List()
		if err != nil {
			return 0, err
		}
		return templates.AdminInviteList(invites, ctx, w)
	})
}

func AdminInviteDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		key := mux.Vars(r)["key"]
		invite, err := ctx.App.Invite.GetByKey(key)
		if err != nil {
			return 0, err
		}
		ctx.Title = invite.Key
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invite"), "invites")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.invite.detail", "key", key), key)...)
		ctx.Breadcrumbs = bc

		return templates.AdminInviteDetail(invite, ctx, w)
	})
}
