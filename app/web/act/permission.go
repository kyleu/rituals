package act

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
	"net/http"
)

type PermissionParams struct {
	Svc      util.Service
	ModelID  uuid.UUID
	Slug     string
	Title    string
	TeamID   *uuid.UUID
	SprintID *uuid.UUID
}

func CheckPerms(ctx *web.RequestContext, permSvc *permission.Service, p *PermissionParams) (auth.Records, permission.Errors, web.Breadcrumbs) {
	var bc web.Breadcrumbs

	auths, currTeams, currSprints := authsTeamsAndSprints(ctx, p.TeamID, p.SprintID)

	var tp *permission.Params
	if p.TeamID != nil {
		tm := ctx.App.Team.GetByID(*p.TeamID)
		tp = &permission.Params{ID: tm.ID, Slug: tm.Slug, Title: tm.Title, Current: currTeams}
	}

	var sp *permission.Params
	if p.SprintID != nil {
		spr := ctx.App.Sprint.GetByID(*p.SprintID)
		sp = &permission.Params{ID: spr.ID, Slug: spr.Slug, Title: spr.Title, Current: currSprints}
	}

	if sp != nil {
		bc = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, util.KeyKey, sp.Slug), sp.Title)
	} else {
		bc = web.BreadcrumbsSimple(ctx.Route(p.Svc.Key+".list"), p.Svc.Plural)
	}

	if ctx.Profile.Role == util.RoleAdmin {
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(p.Svc.Key, util.KeyKey, p.Slug), p.Title)...)
		return auths, nil, bc
	}

	_, permErrors := permSvc.Check(ctx.App.Auth.Enabled, p.Svc, p.ModelID, auths, tp, sp)

	if len(permErrors) == 0 {
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(p.Svc.Key, util.KeyKey, p.Slug), p.Title)...)
	} else {
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(p.Svc.Key, util.KeyKey, p.Slug), p.Slug)...)
	}
	return auths, permErrors, bc
}

func authsTeamsAndSprints(ctx *web.RequestContext, tm *uuid.UUID, spr *uuid.UUID) (auth.Records, []uuid.UUID, []uuid.UUID) {
	auths := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, nil)

	var currTeams []uuid.UUID
	if tm != nil {
		currTeams = ctx.App.Team.GetIdsByMember(ctx.Profile.UserID)
	}

	var currSprints []uuid.UUID
	if spr != nil {
		currSprints = ctx.App.Sprint.GetIdsByMember(ctx.Profile.UserID)
	}

	return auths, currTeams, currSprints
}

func PermErrorTemplate(svc util.Service, errors permission.Errors, auths auth.Records, ctx *web.RequestContext, w http.ResponseWriter) (string, error) {
	return T(templates.PermissionErrors(svc, errors, auths, ctx, w))
}
