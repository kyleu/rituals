package permission

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/util"
)

type Params struct {
	ID      uuid.UUID
	Slug    string
	Title   string
	Current []uuid.UUID
}

func (s *Service) Check(authEnabled bool, svc util.Service, modelID uuid.UUID, auths auth.Records, teamP *Params, sprintP *Params) (Permissions, Errors) {
	perms := s.GetByModelID(modelID, nil)

	var ret Errors

	authResult := s.checkAuths(authEnabled, svc, perms, auths)
	if authResult != nil {
		ret = append(ret, authResult...)
	}

	teamResult := s.checkModel(util.SvcTeam.Key, svc, perms, teamP)
	if teamResult != nil {
		ret = append(ret, teamResult)
	}

	sprintResult := s.checkModel(util.SvcSprint.Key, svc, perms, sprintP)
	if sprintResult != nil {
		ret = append(ret, sprintResult)
	}

	return perms, ret
}
