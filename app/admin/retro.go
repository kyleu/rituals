package admin

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

)

func RetroList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = util.SvcRetro.Title + " List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		params := npnweb.ParamSetFromRequest(r)
		retros := app.Retro(ctx.App).List(params.Get(util.SvcRetro.Key, ctx.Logger))
		return npncontroller.T(admintemplates.RetroList(retros, params, ctx, w))
	})
}

func RetroDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		retroID, err := npnweb.IDFromParams(util.SvcRetro.Key, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		sess := app.Retro(ctx.App).GetByID(*retroID)
		if sess == nil {
			msg := "can't load retro [" + retroID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(util.SvcRetro.Key), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		feedbacks := app.Retro(ctx.App).GetFeedback(sess.ID, params.Get(util.KeyFeedback, ctx.Logger))

		data := app.Retro(ctx.App).Data.GetData(*retroID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		bc = append(bc, npnweb.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.RetroDetail(sess, feedbacks, data, params, ctx, w))
	})
}
