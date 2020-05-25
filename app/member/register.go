package member

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
)

func (s *Service) Register(modelID uuid.UUID, userID uuid.UUID) *Entry {
	dto, err := s.Get(modelID, userID)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting existing member for user [%v] and model [%v]: %+v", modelID, userID, err))
		return nil
	}

	if dto == nil {
		q := fmt.Sprintf(`insert into %s (%s, user_id, name, role) values ($1, $2, '', 'member')`, s.tableName, s.colName)
		err = s.db.Insert(q, nil, modelID, userID)
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

func (s *Service) RegisterRef(modelID *uuid.UUID, userID uuid.UUID) *Entry {
	if modelID == nil {
		return nil
	}
	return s.Register(*modelID, userID)
}
