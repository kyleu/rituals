package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
)

type PermissionParams struct {
	Svc      util.Service
	ModelID  uuid.UUID
	Slug     string
	Title    string
	TeamID   *uuid.UUID
	SprintID *uuid.UUID
}

func check(ctx *web.RequestContext, permSvc *permission.Service, p PermissionParams) (auth.Records, permission.Errors, web.Breadcrumbs) {
	var bc web.Breadcrumbs
	bc = web.BreadcrumbsSimple(ctx.Route(p.Svc.Key+".list"), p.Svc.Plural)

	auths, currTeams, currSprints, err := authsTeamsAndSprints(ctx, p.TeamID, p.SprintID)
	if err != nil {
		return nil, permission.Errors{{Svc: "system", Provider: "error", Message: err.Error()}}, bc
	}

	var tp *permission.Params
	if p.TeamID != nil {
		tm, _ := ctx.App.Team.GetByID(*p.TeamID)
		tp = &permission.Params{ID: tm.ID, Slug: tm.Slug, Title: tm.Title, Current: currTeams}
	}

	var sp *permission.Params
	if p.SprintID != nil {
		spr, _ := ctx.App.Sprint.GetByID(*p.SprintID)
		sp = &permission.Params{ID: spr.ID, Slug: spr.Slug, Title: spr.Title, Current: currSprints}
	}

	if sp != nil {
		bc = append(web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, util.KeyKey, sp.Slug), sp.Title), bc...)
	}

	_, permErrors := permSvc.Check(ctx.App.Auth.Enabled, p.Svc, p.ModelID, auths, tp, sp)

	if len(permErrors) == 0 {
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(p.Svc.Key, util.KeyKey, p.Slug), p.Title)...)
	} else {
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(p.Svc.Key, util.KeyKey, p.Slug), p.Slug)...)
	}

	return auths, permErrors, bc
}

func authsTeamsAndSprints(ctx *web.RequestContext, tm *uuid.UUID, spr *uuid.UUID) (auth.Records, []uuid.UUID, []uuid.UUID, error) {
	auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, nil)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "unable to retrieve current auth records")
	}

	var currTeams []uuid.UUID
	if tm != nil {
		currTeams, err = ctx.App.Team.GetIdsByMember(ctx.Profile.UserID)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "unable to retrieve current teams")
		}
	}

	var currSprints []uuid.UUID
	if spr != nil {
		currSprints, err = ctx.App.Sprint.GetIdsByMember(ctx.Profile.UserID)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "unable to retrieve current sprints")
		}
	}

	return auths, currTeams, currSprints, nil
}

func permErrorTemplate(svc util.Service, errors permission.Errors, auths auth.Records, ctx web.RequestContext, w http.ResponseWriter) (string, error) {
	return tmpl(templates.PermissionErrors(svc, errors, auths, ctx, w))
}
