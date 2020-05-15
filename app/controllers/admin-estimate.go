package controllers

import (
	"fmt"
	"github.com/kyleu/rituals.dev/app/util"
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminEstimateList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Estimate List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), "estimate")...)
		ctx.Breadcrumbs = bc

		estimates, err := ctx.App.Estimate.List()
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminEstimateList(estimates, ctx, w))
	})
}

func AdminEstimateDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		estimateIDString := mux.Vars(r)["id"]
		estimateID, err := uuid.FromString(estimateIDString)
		if err != nil {
			return "", errors.New("invalid estimate id [" + estimateIDString + "]")
		}
		estimate, err := ctx.App.Estimate.GetByID(estimateID)
		if err != nil {
			return "", err
		}
		if estimate == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + estimateIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.estimate"), nil
		}

		members, err := ctx.App.Estimate.Members.GetByModelID(estimateID)
		if err != nil {
			return "", err
		}
		stories, err := ctx.App.Estimate.GetStories(estimateID)
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcEstimate, estimateID)
		if err != nil {
			return "", err
		}

		ctx.Title = estimate.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), "estimate")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate.detail", "id", estimateIDString), estimate.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminEstimateDetail(estimate, members, stories, actions, ctx, w))
	})
}

func AdminStoryDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		storyIDString := mux.Vars(r)["id"]
		storyID, err := uuid.FromString(storyIDString)
		if err != nil {
			return "", errors.New("invalid story id [" + storyIDString + "]")
		}
		story, err := ctx.App.Estimate.GetStoryByID(storyID)
		if err != nil {
			return "", err
		}
		estimateID, err := ctx.App.Estimate.GetStoryEstimateID(storyID)
		if err != nil {
			return "", err
		}
		estimate, err := ctx.App.Estimate.GetByID(*estimateID)
		if err != nil {
			return "", err
		}
		if estimate == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + estimateID.String() + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.estimate"), nil
		}

		votes, err := ctx.App.Estimate.GetStoryVotes(storyID)
		if err != nil {
			return "", err
		}
		ctx.Title = fmt.Sprint(estimate.Slug, ":", story.Idx)
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), "estimate")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate.detail", "id", story.EstimateID.String()), estimate.Slug)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.story.detail", "id", storyIDString), fmt.Sprint("story ", story.Idx))...)
		ctx.Breadcrumbs = bc
		return tmpl(templates.AdminStoryDetail(story, votes, ctx, w))
	})
}
