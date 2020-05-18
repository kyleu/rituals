package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

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

		params := paramSetFromRequest(r)
		sprints, err := ctx.App.Sprint.List(params.Get("sprint"))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminSprintList(sprints, params, ctx, w))
	})
}

func AdminSprintDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		sprintIDString := mux.Vars(r)["id"]
		sprintID, err := uuid.FromString(sprintIDString)
		if err != nil {
			return "", errors.New("invalid sprint id [" + sprintIDString + "]")
		}
		sess, err := ctx.App.Sprint.GetByID(sprintID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + sprintIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.sprint"), nil
		}

		params := paramSetFromRequest(r)

		members, err := ctx.App.Sprint.Members.GetByModelID(sprintID, params.Get("member"))
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetBySprint(sprintID, params.Get("estimate"))
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetBySprint(sprintID, params.Get("standup"))
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetBySprint(sprintID, params.Get("retro"))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcSprint.Key, sprintID, params.Get("action"))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint"), "sprint")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint.detail", "id", sprintIDString), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminSprintDetail(sess, members, estimates, standups, retros, actions, params, ctx, w))
	})
}
