package estimate

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/util"
)

func (s *Service) NewPoll(estimateID uuid.UUID, title string, authorID uuid.UUID) (*Poll, error) {
	id := util.UUID()

	sql := `insert into poll (id, estimate_id, idx, author_id, title) values (
    $1, $2, (select max(idx) + 1 from poll p2 where p2.estimate_id = $3), $4, $5
	)`;
	_, err := s.db.Exec(sql, id, estimateID, estimateID, authorID, title)
	if err != nil {
		return nil, err
	}

	return s.GetPollByID(id)
}

func (s *Service) GetPolls(id uuid.UUID) ([]Poll, error) {
	var dtos []pollDTO
	err := s.db.Select(&dtos, "select * from poll where estimate_id = $1 order by idx", id)
	if err != nil {
		return nil, err
	}
	ret := make([]Poll, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToPoll())
	}
	return ret, nil
}

func (s *Service) GetPollByID(pollID uuid.UUID) (*Poll, error) {
	dto := &pollDTO{}
	err := s.db.Get(dto, "select * from poll where id = $1", pollID)
	if err != nil {
		return nil, err
	}
	ret := dto.ToPoll()
	return &ret, nil
}

func (s *Service) GetPollEstimateID(pollID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	err := s.db.Get(&ret, "select estimate_id from poll where id = $1", pollID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
