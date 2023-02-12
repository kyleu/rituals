package workspace

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func CheckPermissions(
	key string, perms util.Permissions, accounts user.Accounts,
	teamID *uuid.UUID, tf func() (team.Teams, error),
	sprintID *uuid.UUID, sf func() (sprint.Sprints, error),
) (bool, string) {
	if tf != nil && perms.TeamPerm() != nil && teamID != nil {
		if ok, msg := checkTeamPermissions(key, *teamID, tf); !ok {
			return false, msg
		}
	}
	if sf != nil && perms.SprintPerm() != nil && sprintID != nil {
		if ok, msg := checkSprintPermissions(key, *sprintID, sf); !ok {
			return false, msg
		}
	}
	if ok, msg := checkAuthPermissions(key, perms.AuthPerms(), accounts); !ok {
		return false, msg
	}
	return true, KeyOK
}

func checkTeamPermissions(key string, teamID uuid.UUID, tf func() (team.Teams, error)) (bool, string) {
	ts, err := tf()
	if err != nil {
		return false, err.Error()
	}
	if ts.Get(teamID) != nil {
		return true, KeyOK
	}
	return false, fmt.Sprintf("to join this %s, you must be a member of team [%s]", key, teamID.String())
}

func checkSprintPermissions(key string, sprintID uuid.UUID, sf func() (sprint.Sprints, error)) (bool, string) {
	ss, err := sf()
	if err != nil {
		return false, err.Error()
	}
	if ss.Get(sprintID) != nil {
		return true, KeyOK
	}
	return false, fmt.Sprintf("to join this %s, you must be a member of sprint [%s]", key, sprintID.String())
}

func checkAuthPermissions(key string, perms util.Permissions, accounts user.Accounts) (bool, string) {
	if len(perms) == 0 {
		return true, KeyOK
	}
	ret := make([]string, 0, len(perms))
	for _, perm := range perms {
		curr := accounts.GetByProvider(perm.Key)
		for _, a := range curr {
			if perm.Value == "*" || strings.HasSuffix(a.Email, perm.Value) {
				return true, KeyOK
			}
		}
		if perm.Value == "" || perm.Value == "*" {
			ret = append(ret, fmt.Sprintf("[%s]", perm.Key))
		} else {
			ret = append(ret, fmt.Sprintf("[%s] with an email matching [%s]", perm.Key, perm.Value))
		}
	}
	return false, fmt.Sprintf("you must be signed in to %s to join this %s", util.StringArrayOxfordComma(ret, "or"), key)
}
