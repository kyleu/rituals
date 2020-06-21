package admin

import (
	"fmt"
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Estimate List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)

		params := act.ParamSetFromRequest(r)
		estimates := ctx.App.Estimate.List(params.Get(util.SvcEstimate.Key, ctx.Logger))
		return act.T(admintemplates.EstimateList(estimates, params, ctx, w))
	})
}

func EstimateDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		estimateID, err := act.IDFromParams(util.SvcEstimate.Key, mux.Vars(r))
		if err != nil {
			return act.EResp(err)
		}
		sess := ctx.App.Estimate.GetByID(*estimateID)
		if sess == nil {
			msg := "can't load estimate [" + estimateID.String() + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.SvcEstimate.Key), w, r, ctx)
		}

		params := act.ParamSetFromRequest(r)

		stories := ctx.App.Estimate.GetStories(*estimateID, params.Get(util.KeyStory, ctx.Logger))

		data := ctx.App.Estimate.Data.GetData(*estimateID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)
		bc = append(bc, web.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return act.T(admintemplates.EstimateDetail(sess, stories, data, params, ctx, w))
	})
}

func StoryDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		storyID, err := act.IDFromParams(util.KeyStory, mux.Vars(r))
		if err != nil {
			return act.EResp(err)
		}
		story, err := ctx.App.Estimate.GetStoryByID(*storyID)
		if err != nil {
			return act.EResp(err)
		}
		estimateID, err := ctx.App.Estimate.GetStoryEstimateID(*storyID)
		if err != nil {
			return act.EResp(err)
		}
		sess := ctx.App.Estimate.GetByID(*estimateID)
		if sess == nil {
			msg := "can't load estimate [" + estimateID.String() + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.SvcEstimate.Key), w, r, ctx)
		}

		params := act.ParamSetFromRequest(r)

		votes := ctx.App.Estimate.GetStoryVotes(*storyID, params.Get(util.KeyVote, ctx.Logger))
		ctx.Title = fmt.Sprint(sess.Slug, ":", story.Idx)
		bc := adminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)
		el := util.AdminLink(util.SvcEstimate.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(el, util.KeyID, story.EstimateID.String()), sess.Slug)...)
		bc = append(bc, web.BreadcrumbSelf(fmt.Sprint("story ", story.Idx)))
		ctx.Breadcrumbs = bc
		return act.T(admintemplates.StoryDetail(story, votes, params, ctx, w))
	})
}
