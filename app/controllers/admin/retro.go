package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.SvcRetro.Title + " List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		params := act.ParamSetFromRequest(r)
		retros := ctx.App.Retro.List(params.Get(util.SvcRetro.Key, ctx.Logger))
		return tmpl(templates.AdminRetroList(retros, params, ctx, w))
	})
}

func RetroDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		retroID, err := act.IDFromParams(util.SvcRetro.Key, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		sess, err := ctx.App.Retro.GetByID(*retroID)
		if err != nil {
			return eresp(err, "")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + retroID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcRetro.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		feedbacks := ctx.App.Retro.GetFeedback(sess.ID, params.Get(util.KeyFeedback, ctx.Logger))
		comments := ctx.App.Retro.Comments.GetByModelID(*retroID, params.Get(util.KeyComment, ctx.Logger))
		members := ctx.App.Retro.Members.GetByModelID(*retroID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Retro.Permissions.GetByModelID(*retroID, params.Get(util.KeyPermission, ctx.Logger))
		actions := ctx.App.Action.GetBySvcModel(util.SvcRetro, *retroID, params.Get(util.KeyAction, ctx.Logger))

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		link := util.AdminLink(util.SvcRetro.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, retroID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminRetroDetail(sess, feedbacks, comments, members, perms, actions, params, ctx, w))
	})
}
