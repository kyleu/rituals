package member

import (
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
)

func (s *Service) Register(modelID uuid.UUID, userID uuid.UUID, memberName string, role Role) *Entry {
	dto, err := s.Get(modelID, userID)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting existing member for user [%v] and model [%v]: %+v", modelID, userID, err))
		return nil
	}

	if dto == nil {
		q := npndatabase.SQLInsert(s.tableName, []string{s.colName, npncore.WithDBID(npncore.KeyUser), npncore.KeyName, "picture", npncore.KeyRole}, 1)
		err = s.db.Insert(q, nil, modelID, userID, memberName, "", role.String())
		if err != nil {
			s.logger.Error(fmt.Sprintf("error inserting member for user [%v] and model [%v]: %+v", modelID, userID, err))
		}
		dto, err = s.Get(modelID, userID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error retrieving updated member for user [%v] and model [%v]: %+v", modelID, userID, err))
		}

		s.actions.Post(s.svc, modelID, userID, action.ActMemberAdd, nil)
	}

	return dto
}

func (s *Service) RegisterRef(modelID *uuid.UUID, userID uuid.UUID, memberName string, role Role) *Entry {
	if modelID == nil {
		return nil
	}
	return s.Register(*modelID, userID, memberName, role)
}
