package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func CommentList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Comment List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyComment, util.Plural(util.KeyComment))

		params := act.ParamSetFromRequest(r)
		comments := ctx.App.Comment.List(params.Get(util.KeyComment, ctx.Logger))
		return act.T(admintemplates.CommentList(comments, params, ctx, w))
	})
}

func CommentDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		commentID, err := act.IDFromParams(util.SvcEstimate.Key, mux.Vars(r))
		if err != nil {
			return act.EResp(err, "invalid comment id")
		}
		e := ctx.App.Comment.GetByID(*commentID)
		if e == nil {
			msg := "can't load comment [" + commentID.String() + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.KeyComment), w, r, ctx)
		}

		params := act.ParamSetFromRequest(r)

		ctx.Title = e.ID.String()
		bc := adminBC(ctx, util.KeyComment, util.Plural(util.KeyComment))
		bc = append(bc, web.BreadcrumbSelf(e.ID.String()))
		ctx.Breadcrumbs = bc

		return act.T(admintemplates.CommentDetail(e, params, ctx, w))
	})
}
