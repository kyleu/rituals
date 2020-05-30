package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func InvitationList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Invitation List"
		bc := adminBC(ctx, util.KeyInvitation, util.Plural(util.KeyInvitation))
		ctx.Breadcrumbs = bc

		params := act.ParamSetFromRequest(r)
		invitations := ctx.App.Invitation.List(params.Get(util.KeyInvitation, ctx.Logger))
		return tmpl(templates.AdminInvitationList(invitations, params, ctx, w))
	})
}

func InvitationDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Invitation.GetByKey(key)
		if err != nil {
			return eresp(err, "")
		}
		ctx.Title = sess.Key
		bc := adminBC(ctx, util.KeyInvitation, util.Plural(util.KeyInvitation))
		link := util.AdminLink(util.KeyInvitation, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyKey, key), key)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminInvitationDetail(sess, ctx, w))
	})
}
