package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func InvitationList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Invitation List"
		bc := adminBC(ctx, util.KeyInvitation, "invitations")
		ctx.Breadcrumbs = bc

		params := act.ParamSetFromRequest(r)
		invitations, err := ctx.App.Invitation.List(params.Get(util.KeyInvitation, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminInvitationList(invitations, params, ctx, w))
	})
}

func InvitationDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Invitation.GetByKey(key)
		if err != nil {
			return "", err
		}
		ctx.Title = sess.Key
		bc := adminBC(ctx, util.KeyInvitation, "invitations")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.KeyInvitation, util.KeyDetail), util.KeyKey, key), key)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminInvitationDetail(sess, ctx, w))
	})
}
