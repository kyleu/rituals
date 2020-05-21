package controllers

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
	"net/http"
)

func authsAndTeams(ctx web.RequestContext, tm *uuid.UUID) ([]*auth.Record, []uuid.UUID, error) {
	auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, nil)
	if err != nil {
		return nil, nil, errors.WithStack(errors.Wrap(err, "unable to retrieve current auth records"))
	}

	var currTeams []uuid.UUID
	if tm != nil {
		currTeams, err = ctx.App.Team.GetIdsByMember(ctx.Profile.UserID)
		if err != nil {
			return nil, nil, errors.WithStack(errors.Wrap(err, "unable to retrieve current teams"))
		}
	}

	return auths, currTeams, nil
}

func permErrorTemplate(errors permission.Errors, ctx web.RequestContext, w http.ResponseWriter) (string, error) {
	return tmpl(templates.PermissionErrors(errors, ctx, w))
}
