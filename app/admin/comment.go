package admin

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

)

func CommentList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Comment List"
		ctx.Breadcrumbs = adminBC(ctx, npncore.KeyComment, npncore.Plural(npncore.KeyComment))

		params := npnweb.ParamSetFromRequest(r)
		comments := app.Comment(ctx.App).List(params.Get(npncore.KeyComment, ctx.Logger))
		return npncontroller.T(admintemplates.CommentList(comments, params, ctx, w))
	})
}

func CommentDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		commentID, err := npnweb.IDFromParams(util.SvcEstimate.Key, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err, "invalid comment id")
		}
		e := app.Comment(ctx.App).GetByID(*commentID)
		if e == nil {
			msg := "can't load comment [" + commentID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyComment), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		ctx.Title = e.ID.String()
		bc := adminBC(ctx, npncore.KeyComment, npncore.Plural(npncore.KeyComment))
		bc = append(bc, npnweb.BreadcrumbSelf(e.ID.String()))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.CommentDetail(e, params, ctx, w))
	})
}
