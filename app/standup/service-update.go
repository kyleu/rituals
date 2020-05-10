package standup

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
	"time"
)

func (s *Service) NewUpdate(standupID uuid.UUID, d time.Time, content string, authorID uuid.UUID) (*Update, error) {
	id := util.UUID()

	sql := `insert into standup_update (id, standup_id, d, author_id, content) values (
    $1, $2, $3, $4, $5
	)`
	_, err := s.db.Exec(sql, id, standupID, d, authorID, content)
	if err != nil {
		return nil, err
	}

	return s.GetUpdateByID(id)
}

func (s *Service) GetUpdates(standupID uuid.UUID) ([]Update, error) {
	var dtos []updateDTO
	err := s.db.Select(&dtos, "select * from standup_update where standup_id = $1 order by d desc", standupID)
	if err != nil {
		return nil, err
	}
	ret := make([]Update, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToUpdate())
	}
	return ret, nil
}

func (s *Service) GetUpdateByID(updateID uuid.UUID) (*Update, error) {
	dto := &updateDTO{}
	err := s.db.Get(dto, "select * from standup_update where id = $1", updateID)
	if err != nil {
		return nil, err
	}
	ret := dto.ToUpdate()
	return &ret, nil
}

func (s *Service) GetUpdateStandupID(updateID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	err := s.db.Get(&ret, "select standup_id from standup_update where id = $1", updateID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
