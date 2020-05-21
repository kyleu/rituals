package admin

import (
	"fmt"
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Estimate List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)

		params := act.ParamSetFromRequest(r)
		estimates, err := ctx.App.Estimate.List(params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminEstimateList(estimates, params, ctx, w))
	})
}

func EstimateDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		estimateID := util.GetUUIDPointer(mux.Vars(r), "id")
		if estimateID == nil {
			return "", errors.New("invalid estimate id")
		}
		sess, err := ctx.App.Estimate.GetByID(*estimateID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + estimateID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcEstimate.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		members := ctx.App.Estimate.Members.GetByModelID(*estimateID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Team.Permissions.GetByModelID(*estimateID, params.Get(util.KeyPermission, ctx.Logger))

		stories, err := ctx.App.Estimate.GetStories(*estimateID, params.Get(util.KeyStory, ctx.Logger))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcEstimate.Key, *estimateID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.SvcEstimate.Key, util.KeyDetail), "id", estimateID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminEstimateDetail(sess, members, perms, stories, actions, params, ctx, w))
	})
}

func StoryDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		storyID := util.GetUUIDPointer(mux.Vars(r), "id")
		if storyID == nil {
			return "", errors.New("invalid story id")
		}
		story, err := ctx.App.Estimate.GetStoryByID(*storyID)
		if err != nil {
			return "", err
		}
		estimateID, err := ctx.App.Estimate.GetStoryEstimateID(*storyID)
		if err != nil {
			return "", err
		}
		sess, err := ctx.App.Estimate.GetByID(*estimateID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + estimateID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcEstimate.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		votes, err := ctx.App.Estimate.GetStoryVotes(*storyID, params.Get(util.KeyVote, ctx.Logger))
		if err != nil {
			return "", err
		}
		ctx.Title = fmt.Sprint(sess.Slug, ":", story.Idx)
		bc := adminBC(ctx, util.SvcEstimate.Key, util.SvcEstimate.Plural)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.SvcEstimate.Key, util.KeyDetail), "id", story.EstimateID.String()), sess.Slug)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.KeyStory, util.KeyDetail), "id", storyID.String()), fmt.Sprint("story ", story.Idx))...)
		ctx.Breadcrumbs = bc
		return tmpl(templates.AdminStoryDetail(story, votes, params, ctx, w))
	})
}
