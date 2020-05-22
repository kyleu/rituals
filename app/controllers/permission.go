package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func check(
	ctx *web.RequestContext, permSvc *permission.Service, svc util.Service,
	modelID uuid.UUID, slug string, title string,
	teamID *uuid.UUID, sprintID *uuid.UUID) (permission.Errors, web.Breadcrumbs) {
	var tmTitle, sprTitle string

	var tm *team.Session
	if teamID != nil {
		tm, _ = ctx.App.Team.GetByID(*teamID)
		tmTitle = tm.Title
	}

	var spr *sprint.Session
	if sprintID != nil {
		spr, _ = ctx.App.Sprint.GetByID(*sprintID)
		sprTitle = spr.Title
	}

	var bc web.Breadcrumbs
	if spr == nil {
		bc = web.BreadcrumbsSimple(ctx.Route(svc.Key+".list"), svc.Key)
	} else {
		bc = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, util.KeyKey, spr.Slug), spr.Title)
	}

	auths, currTeams, currSprints, err := authsTeamsAndSprints(ctx, teamID, sprintID)
	if err != nil {
		return permission.Errors{{K: "system", V: "00000000-0000-0000-0000-000000000000", Code: "error", Message: err.Error()}}, bc
	}

	permErrors := permSvc.Check(svc, modelID, auths, teamID, tmTitle, currTeams, sprintID, sprTitle, currSprints)

	if len(permErrors) == 0 {
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(svc.Key, util.KeyKey, slug), title)...)
	} else {
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(svc.Key, util.KeyKey, slug), slug)...)
	}

	return permErrors, bc
}

func authsTeamsAndSprints(ctx *web.RequestContext, tm *uuid.UUID, spr *uuid.UUID) (auth.Records, []uuid.UUID, []uuid.UUID, error) {
	auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, nil)
	if err != nil {
		return nil, nil, nil, errors.WithStack(errors.Wrap(err, "unable to retrieve current auth records"))
	}

	var currTeams []uuid.UUID
	if tm != nil {
		currTeams, err = ctx.App.Team.GetIdsByMember(ctx.Profile.UserID)
		if err != nil {
			return nil, nil, nil, errors.WithStack(errors.Wrap(err, "unable to retrieve current teams"))
		}
	}

	var currSprints []uuid.UUID
	if spr != nil {
		currSprints, err = ctx.App.Sprint.GetIdsByMember(ctx.Profile.UserID)
		if err != nil {
			return nil, nil, nil, errors.WithStack(errors.Wrap(err, "unable to retrieve current sprints"))
		}
	}

	return auths, currTeams, currSprints, nil
}

func permErrorTemplate(svc util.Service, errors permission.Errors, ctx web.RequestContext, w http.ResponseWriter) (string, error) {
	return tmpl(templates.PermissionErrors(svc, errors, ctx, w))
}
