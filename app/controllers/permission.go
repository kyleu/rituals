package controllers

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
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

func CheckPerms(ctx *npnweb.RequestContext, permSvc *permission.Service, p *PermissionParams) (auth.Records, permission.Errors, npnweb.Breadcrumbs) {
	var bc npnweb.Breadcrumbs

	auths, currTeams, currSprints := authsTeamsAndSprints(ctx, p.TeamID, p.SprintID)

	var tp *permission.Params
	if p.TeamID != nil {
		tm := app.Svc(ctx.App).Team.GetByID(*p.TeamID)
		tp = &permission.Params{ID: tm.ID, Slug: tm.Slug, Title: tm.Title, Current: currTeams}
	}

	var sp *permission.Params
	if p.SprintID != nil {
		spr := app.Svc(ctx.App).Sprint.GetByID(*p.SprintID)
		sp = &permission.Params{ID: spr.ID, Slug: spr.Slug, Title: spr.Title, Current: currSprints}
	}

	if sp != nil {
		bc = npnweb.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, npncore.KeyKey, sp.Slug), sp.Title)
	} else {
		bc = npnweb.BreadcrumbsSimple(ctx.Route(p.Svc.Key+".list"), p.Svc.Plural)
	}

	if ctx.Profile.Role == npnuser.RoleAdmin {
		bc = append(bc, npnweb.BreadcrumbsSimple(ctx.Route(p.Svc.Key, npncore.KeyKey, p.Slug), p.Title)...)
		return auths, nil, bc
	}

	_, permErrors := permSvc.Check(ctx.App.Auth().Enabled(), p.Svc.Key, p.ModelID, auths, tp, sp)

	if len(permErrors) == 0 {
		bc = append(bc, npnweb.BreadcrumbsSimple(ctx.Route(p.Svc.Key, npncore.KeyKey, p.Slug), p.Title)...)
	} else {
		bc = append(bc, npnweb.BreadcrumbsSimple(ctx.Route(p.Svc.Key, npncore.KeyKey, p.Slug), p.Slug)...)
	}
	return auths, permErrors, bc
}

func authsTeamsAndSprints(ctx *npnweb.RequestContext, tm *uuid.UUID, spr *uuid.UUID) (auth.Records, []uuid.UUID, []uuid.UUID) {
	auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, nil)

	var currTeams []uuid.UUID
	if tm != nil {
		currTeams = app.Svc(ctx.App).Team.GetIdsByMember(ctx.Profile.UserID)
	}

	var currSprints []uuid.UUID
	if spr != nil {
		currSprints = app.Svc(ctx.App).Sprint.GetIdsByMember(ctx.Profile.UserID)
	}

	return auths, currTeams, currSprints
}

func PermErrorTemplate(svc util.Service, errors permission.Errors, auths auth.Records, ctx *npnweb.RequestContext, w http.ResponseWriter) (string, error) {
	return npncontroller.T(templates.PermissionErrors(svc, errors, auths, ctx, w))
}
