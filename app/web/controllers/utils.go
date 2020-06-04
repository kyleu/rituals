package controllers

import (
	"net/http"
	"net/url"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/user"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/app/web/act"
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
	User       *user.SystemUser
	Title      string
	MemberName string
	TeamID     *uuid.UUID
	SprintID   *uuid.UUID
	Perms      permission.Permissions
}

func parseSessionForm(userID uuid.UUID, svc util.Service, form url.Values, userSvc *user.Service) *formResult {
	u := userSvc.GetByID(userID, true)
	title := util.ServiceTitle(svc, form.Get(util.KeyTitle))
	memberName := form.Get("member-name")
	teamID := getUUID(form, util.SvcTeam.Key)
	sprintID := getUUID(form, util.SvcSprint.Key)
	perms := parsePerms(form, teamID, sprintID)

	ret := formResult{User: u, Title: title, MemberName: memberName, TeamID: teamID, SprintID: sprintID, Perms: perms}
	return &ret
}
