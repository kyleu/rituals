package member

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
)

func (s *Service) Register(modelID uuid.UUID, userID uuid.UUID, role Role) *Entry {
	dto, err := s.Get(modelID, userID)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting existing member for user [%v] and model [%v]: %+v", modelID, userID, err))
		return nil
	}

	if dto == nil {
		q := query.SQLInsert(s.tableName, []string{s.colName, "user_id", "name", "role"}, 1)
		err = s.db.Insert(q, nil, modelID, userID, "", role.String())
		if err != nil {
			s.logger.Error(fmt.Sprintf("error inserting member for user [%v] and model [%v]: %+v", modelID, userID, err))
		}
		dto, err = s.Get(modelID, userID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error retrieving updated member for user [%v] and model [%v]: %+v", modelID, userID, err))
		}

		s.actions.Post(s.svc, modelID, userID, action.ActMemberAdd, nil, "")
	}

	return dto
}

func (s *Service) RegisterRef(modelID *uuid.UUID, userID uuid.UUID, role Role) *Entry {
	if modelID == nil {
		return nil
	}
	return s.Register(*modelID, userID, role)
}
