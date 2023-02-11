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

func AvailablePermissions(accounts user.Accounts, ts team.Teams, ss sprint.Sprints) util.Permissions {
	ret := make(util.Permissions, 0, len(ts)+len(ss)+len(accounts))
	for _, t := range ts {
		ret = append(ret, &util.Permission{Key: "team", Value: t.ID.String() + "|" + t.TitleString()})
	}
	for _, s := range ss {
		ret = append(ret, &util.Permission{Key: "sprint", Value: s.ID.String() + "|" + s.TitleString()})
	}
	for _, a := range accounts {
		ret = append(ret, &util.Permission{Key: a.Provider, Value: a.Email})
	}
	return ret
}

func CheckPermissions(perms util.Permissions, accounts user.Accounts, tf func() (team.Teams, error), sf func() (sprint.Sprints, error)) (bool, string) {
	if ok, msg := checkTeamPermissions(perms.TeamPerms().Values(), tf); !ok {
		return false, msg
	}
	if ok, msg := checkSprintPermissions(perms.SprintPerms().Values(), sf); !ok {
		return false, msg
	}
	if ok, msg := checkAuthPermissions(perms.AuthPerms(), accounts); !ok {
		return false, msg
	}
	return true, "ok"
}

func checkTeamPermissions(idStrings []string, tf func() (team.Teams, error)) (bool, string) {
	if len(idStrings) == 0 {
		return true, "ok"
	}
	ids := make([]uuid.UUID, 0, len(idStrings))
	for _, x := range idStrings {
		id := util.UUIDFromString(x)
		if id == nil {
			return false, "invalid team ID [" + x + "]"
		}
		ids = append(ids, *id)
	}
	ts, err := tf()
	if err != nil {
		return false, err.Error()
	}
	for _, id := range ids {
		if ts.Get(id) != nil {
			return true, "ok"
		}
	}
	tNames := util.StringArrayOxfordComma(idStrings, "or")
	return false, fmt.Sprintf("to join, you must be a member of one of the following teams: %s", tNames)
}

func checkSprintPermissions(idStrings []string, sf func() (sprint.Sprints, error)) (bool, string) {
	if len(idStrings) == 0 {
		return true, "ok"
	}
	ids := make([]uuid.UUID, 0, len(idStrings))
	for _, x := range idStrings {
		id := util.UUIDFromString(x)
		if id == nil {
			return false, "invalid sprint ID [" + x + "]"
		}
		ids = append(ids, *id)
	}
	ss, err := sf()
	if err != nil {
		return false, err.Error()
	}
	for _, id := range ids {
		if ss.Get(id) != nil {
			return true, "ok"
		}
	}
	tNames := util.StringArrayOxfordComma(idStrings, "or")
	return false, fmt.Sprintf("to join, you must be a member of one of the following sprints: %s", tNames)
}

func checkAuthPermissions(perms util.Permissions, accounts user.Accounts) (bool, string) {
	if len(perms) == 0 {
		return true, "ok"
	}
	ret := make([]string, 0, len(perms))
	for _, perm := range perms {
		curr := accounts.GetByProvider(perm.Key)
		for _, a := range curr {
			if strings.HasSuffix(a.Email, perm.Value) {
				return true, "ok"
			}
		}
		if perm.Value == "" {
			ret = append(ret, fmt.Sprintf("[%s]", perm.Key))
		} else {
			ret = append(ret, fmt.Sprintf("[%s] with an email matching [%s]", perm.Key, perm.Value))
		}
	}
	return false, fmt.Sprintf("you must be signed in to %s", util.StringArrayOxfordComma(ret, "or"))
}
