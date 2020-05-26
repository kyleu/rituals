package controllers

import (
	"net/http"
	"net/url"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/user"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/web"
)

func Health(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_, _ = w.Write([]byte("OK"))
		return "", nil
	})
}

func eresp(err error, msg string) (string, error) {
	if len(msg) == 0 {
		return "", err
	}
	return "", errors.Wrap(err, msg)
}

func enew(msg string) (string, error) {
	return "", errors.New(msg)
}

type formResult struct {
	User     *user.SystemUser
	Title    string
	TeamID   *uuid.UUID
	SprintID *uuid.UUID
	Perms    permission.Permissions
}

func parseSessionForm(userID uuid.UUID, svc util.Service, form url.Values, userSvc *user.Service) (*formResult, error) {
	u, err := userSvc.GetByID(userID, true)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create user")
	}
	title := util.ServiceTitle(svc, form.Get(util.KeyTitle))
	teamID := getUUID(form, util.SvcTeam.Key)
	sprintID := getUUID(form, util.SvcSprint.Key)
	perms := parsePerms(form, teamID, sprintID)

	ret := formResult{User: u, Title: title, TeamID: teamID, SprintID: sprintID, Perms: perms}
	return &ret, nil
}
