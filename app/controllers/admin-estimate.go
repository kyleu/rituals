package controllers

import (
	"fmt"
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminEstimateList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Estimate List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), util.SvcEstimate.Key)...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		estimates, err := ctx.App.Estimate.List(params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminEstimateList(estimates, params, ctx, w))
	})
}

func AdminEstimateDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		estimateID := getUUIDPointer(mux.Vars(r), "id")
		if estimateID == nil {
			return "", errors.New("invalid estimate id")
		}
		sess, err := ctx.App.Estimate.GetByID(*estimateID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + estimateID.String() + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.estimate"), nil
		}

		params := paramSetFromRequest(r)

		members := ctx.App.Estimate.Members.GetByModelID(*estimateID, params.Get(util.KeyMember, ctx.Logger))

		stories, err := ctx.App.Estimate.GetStories(*estimateID, params.Get(util.KeyStory, ctx.Logger))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcEstimate.Key, *estimateID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), util.SvcEstimate.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate.detail", "id", estimateID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminEstimateDetail(sess, members, stories, actions, params, ctx, w))
	})
}

func AdminStoryDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		storyID := getUUIDPointer(mux.Vars(r), "id")
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
			saveSession(w, r, ctx)
			return ctx.Route("admin.estimate"), nil
		}

		params := paramSetFromRequest(r)

		votes, err := ctx.App.Estimate.GetStoryVotes(*storyID, params.Get(util.KeyVote, ctx.Logger))
		if err != nil {
			return "", err
		}
		ctx.Title = fmt.Sprint(sess.Slug, ":", story.Idx)
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), util.SvcEstimate.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate.detail", "id", story.EstimateID.String()), sess.Slug)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.story.detail", "id", storyID.String()), fmt.Sprint("story ", story.Idx))...)
		ctx.Breadcrumbs = bc
		return tmpl(templates.AdminStoryDetail(story, votes, params, ctx, w))
	})
}
