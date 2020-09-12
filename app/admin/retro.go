package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = util.SvcRetro.Title + " List"
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		params := npnweb.ParamSetFromRequest(r)
		retros := app.Svc(ctx.App).Retro.List(params.Get(util.SvcRetro.Key, ctx.Logger))
		return npncontroller.T(admintemplates.RetroList(retros, params, ctx, w))
	})
}

func RetroDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		retroID, err := npnweb.IDFromParams(util.SvcRetro.Key, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		sess := app.Svc(ctx.App).Retro.GetByID(*retroID)
		if sess == nil {
			msg := "can't load retro [" + retroID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(util.SvcRetro.Key), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		feedbacks := app.Svc(ctx.App).Retro.GetFeedback(sess.ID, params.Get(util.KeyFeedback, ctx.Logger))

		data := app.Svc(ctx.App).Retro.Data.GetData(*retroID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := npncontroller.AdminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		bc = append(bc, npnweb.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.RetroDetail(sess, feedbacks, data, params, ctx, w))
	})
}
