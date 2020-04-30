package estimate

import "github.com/gofrs/uuid"

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

func (s *Service) GetPollByID(id uuid.UUID) (*Poll, error) {
	dto := &pollDTO{}
	err := s.db.Get(dto, "select * from poll where id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := dto.ToPoll()
	return &ret, nil
}

func (s *Service) GetPollVotes(id uuid.UUID) ([]Vote, error) {
	var dtos []voteDTO
	err := s.db.Select(&dtos, "select * from vote where poll_id = $1", id)
	if err != nil {
		return nil, err
	}
	ret := make([]Vote, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToVote())
	}
	return ret, nil
}
