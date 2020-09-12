package admin

import (
	"fmt"
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Estimate List"
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)

		params := npnweb.ParamSetFromRequest(r)
		estimates := app.Svc(ctx.App).Estimate.List(params.Get(util.SvcEstimate.Key, ctx.Logger))
		return npncontroller.T(admintemplates.EstimateList(estimates, params, ctx, w))
	})
}

func EstimateDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		estimateID, err := npnweb.IDFromParams(util.SvcEstimate.Key, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		sess := app.Svc(ctx.App).Estimate.GetByID(*estimateID)
		if sess == nil {
			msg := "can't load estimate [" + estimateID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(util.SvcEstimate.Key), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		stories := app.Svc(ctx.App).Estimate.GetStories(*estimateID, params.Get(util.KeyStory, ctx.Logger))

		data := app.Svc(ctx.App).Estimate.Data.GetData(*estimateID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := npncontroller.AdminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)
		bc = append(bc, npnweb.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.EstimateDetail(sess, stories, data, params, ctx, w))
	})
}

func StoryDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		storyID, err := npnweb.IDFromParams(util.KeyStory, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		story, err := app.Svc(ctx.App).Estimate.GetStoryByID(*storyID)
		if err != nil {
			return npncontroller.EResp(err)
		}
		estimateID, err := app.Svc(ctx.App).Estimate.GetStoryEstimateID(*storyID)
		if err != nil {
			return npncontroller.EResp(err)
		}
		sess := app.Svc(ctx.App).Estimate.GetByID(*estimateID)
		if sess == nil {
			msg := "can't load estimate [" + estimateID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(util.SvcEstimate.Key), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		votes := app.Svc(ctx.App).Estimate.GetStoryVotes(*storyID, params.Get(util.KeyVote, ctx.Logger))
		ctx.Title = fmt.Sprint(sess.Slug, ":", story.Idx)
		bc := npncontroller.AdminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)
		el := npnweb.AdminLink(util.SvcEstimate.Key, npncore.KeyDetail)
		bc = append(bc, npnweb.BreadcrumbsSimple(ctx.Route(el, npncore.KeyID, story.EstimateID.String()), sess.Slug)...)
		bc = append(bc, npnweb.BreadcrumbSelf(fmt.Sprint("story ", story.Idx)))
		ctx.Breadcrumbs = bc
		return npncontroller.T(admintemplates.StoryDetail(story, votes, params, ctx, w))
	})
}
