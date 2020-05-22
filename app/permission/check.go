package permission

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) Check(
	svc util.Service, modelID uuid.UUID, auths auth.Records,
	teamID *uuid.UUID, teamName string, currentTeams []uuid.UUID,
	sprintID *uuid.UUID, sprintName string, sprints []uuid.UUID) Errors {
	perms := s.GetByModelID(modelID, nil)

	var ret Errors

	authResult := s.checkAuths(svc, perms, auths)
	if authResult != nil {
		ret = append(ret, authResult...)
	}

	teamResult := s.checkTeam(svc, perms, teamID, teamName, currentTeams)
	if teamResult != nil {
		ret = append(ret, teamResult)
	}

	sprintResult := s.checkSprint(svc, perms, sprintID, sprintName, sprints)
	if sprintResult != nil {
		ret = append(ret, sprintResult)
	}

	return ret
}
