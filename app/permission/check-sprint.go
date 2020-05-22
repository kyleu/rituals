package permission

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) checkSprint(svc util.Service, perms Permissions, sprintID *uuid.UUID, title string, sprints []uuid.UUID) *Error {
	sp := perms.FindByK(util.SvcSprint.Key)

	if len(sp) == 0 || sprintID == nil {
		return nil
	}

	hasSprint := false
	for _, t := range sprints {
		if t == *sprintID {
			hasSprint = true
			break
		}
	}

	if hasSprint {
		return nil
	}

	msg := fmt.Sprintf("you are not a member of [%v], this %v's sprint", title, svc.Key)
	return &Error{K: "sprint", V: sprintID.String(), Code: "sprint", Message: msg}
}
