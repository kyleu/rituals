package controllers

import (
	"net/url"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/user"
	"github.com/kyleu/rituals.dev/app/util"
)

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

func getUUID(m url.Values, key string) *uuid.UUID {
	retString := m.Get(key)
	var retID *uuid.UUID
	if retString != "" {
		s, err := uuid.FromString(retString)
		if err == nil {
			retID = &s
		}
	}
	return retID
}

func parsePerm(form map[string][]string, key string, ret permission.Permissions) permission.Permissions {
	t, ok := form["perm-"+key]
	if ok && len(t) > 0 && t[0] == "true" {
		var emails []string
		emailArray := form["perm-"+key+"-email"]
		for _, e := range emailArray {
			emails = append(emails, strings.TrimSpace(e))
		}
		ret = append(ret, &permission.Permission{K: key, V: strings.Join(emails, ","), Access: member.RoleMember})
	}
	return ret
}

func parsePerms(form map[string][]string, teamID *uuid.UUID, sprintID *uuid.UUID) permission.Permissions {
	var ret permission.Permissions
	if teamID != nil {
		ret = parsePerm(form, util.SvcTeam.Key, ret)
	}
	if sprintID != nil {
		ret = parsePerm(form, util.SvcSprint.Key, ret)
	}
	for _, prv := range auth.AllProviders {
		ret = parsePerm(form, prv.Key, ret)
	}
	return ret
}
