package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.SvcRetro.Title + " List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		params := act.ParamSetFromRequest(r)
		retros := ctx.App.Retro.List(params.Get(util.SvcRetro.Key, ctx.Logger))
		return tmpl(admintemplates.RetroList(retros, params, ctx, w))
	})
}

func RetroDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		retroID, err := act.IDFromParams(util.SvcRetro.Key, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		sess := ctx.App.Retro.GetByID(*retroID)
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + retroID.String() + "]")
			act.SaveSession(w, r, &ctx)
			return ctx.Route(util.AdminLink(util.SvcRetro.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		feedbacks := ctx.App.Retro.GetFeedback(sess.ID, params.Get(util.KeyFeedback, ctx.Logger))

		data := ctx.App.Retro.Data.GetData(*retroID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		link := util.AdminLink(util.SvcRetro.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, retroID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(admintemplates.RetroDetail(sess, feedbacks, data, params, ctx, w))
	})
}
