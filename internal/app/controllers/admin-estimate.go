package controllers

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func AdminEstimateList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Estimate List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), "estimates")...)
		ctx.Breadcrumbs = bc

		estimates, err := ctx.App.Estimate.List()
		if err != nil {
			return 0, err
		}
		return templates.AdminEstimateList(estimates, ctx, w)
	})
}

func AdminEstimateDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		estimateIDString := mux.Vars(r)["id"]
		estimateID, err := uuid.FromString(estimateIDString)
		if err != nil {
			return 0, errors.Wrap(err, "invalid estimate id ["+estimateIDString+"]")
		}
		estimate, err := ctx.App.Estimate.GetByID(estimateID)
		if err != nil {
			return 0, err
		}
		polls, err := ctx.App.Estimate.GetPolls(estimateID)
		if err != nil {
			return 0, err
		}
		members, err := ctx.App.Estimate.GetMembers(estimateID)
		if err != nil {
			return 0, err
		}

		ctx.Title = estimate.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), "estimates")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate.detail", "id", estimateIDString), estimate.Slug)...)
		ctx.Breadcrumbs = bc

		return templates.AdminEstimateDetail(estimate, polls, members, ctx, w)
	})
}

func AdminPollDetail(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		pollIDString := mux.Vars(r)["id"]
		pollID, err := uuid.FromString(pollIDString)
		if err != nil {
			return 0, errors.Wrap(err, "invalid poll id ["+pollIDString+"]")
		}
		poll, err := ctx.App.Estimate.GetPollByID(pollID)
		if err != nil {
			return 0, err
		}
		estimate, err := ctx.App.Estimate.GetByID(poll.EstimateID)
		if err != nil {
			return 0, err
		}
		votes, err := ctx.App.Estimate.GetPollVotes(pollID)
		if err != nil {
			return 0, err
		}
		ctx.Title = fmt.Sprintf("%v:%v", estimate.Slug, poll.Idx)
		bc := web.BreadcrumbsSimple(ctx.Route("admin.home"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate"), "estimates")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.estimate.detail", "id", poll.EstimateID.String()), estimate.Slug)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.poll.detail", "id", pollIDString), fmt.Sprintf("poll %v", poll.Idx))...)
		ctx.Breadcrumbs = bc
		return templates.AdminPollDetail(poll, votes, ctx, w)
	})
}
