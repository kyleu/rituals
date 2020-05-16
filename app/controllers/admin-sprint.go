package controllers

import (
	"github.com/kyleu/rituals.dev/app/util"
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminSprintList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Sprint List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint"), "sprint")...)
		ctx.Breadcrumbs = bc

		sprints, err := ctx.App.Sprint.List()
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminSprintList(sprints, ctx, w))
	})
}

func AdminSprintDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		sprintIDString := mux.Vars(r)["id"]
		sprintID, err := uuid.FromString(sprintIDString)
		if err != nil {
			return "", errors.New("invalid sprint id [" + sprintIDString + "]")
		}
		sprint, err := ctx.App.Sprint.GetByID(sprintID)
		if err != nil {
			return "", err
		}
		if sprint == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + sprintIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.sprint"), nil
		}
		members, err := ctx.App.Sprint.Members.GetByModelID(sprintID)
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetBySprint(sprintID, 0)
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetBySprint(sprintID, 0)
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetBySprint(sprintID, 0)
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcSprint.Key, sprintID)
		if err != nil {
			return "", err
		}

		ctx.Title = sprint.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint"), "sprint")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint.detail", "id", sprintIDString), sprint.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminSprintDetail(sprint, members, estimates, standups, retros, actions, ctx, w))
	})
}
