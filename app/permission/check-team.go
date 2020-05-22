package permission

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) checkTeam(svc util.Service, perms Permissions, teamID *uuid.UUID, title string, currentTeams []uuid.UUID) *Error {
	tp := perms.FindByK(util.SvcTeam.Key)

	if len(tp) == 0 || teamID == nil {
		return nil
	}

	hasTeam := false
	for _, t := range currentTeams {
		if t == *teamID {
			hasTeam = true
			break
		}
	}

	if hasTeam {
		return nil
	}

	msg := fmt.Sprintf("you are not a member of [%v], this %v's team", title, svc.Key)
	return &Error{K: "team", V: teamID.String(), Code: "team", Message: msg}
}
