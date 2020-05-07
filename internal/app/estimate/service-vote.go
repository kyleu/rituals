package estimate

import "github.com/gofrs/uuid"

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
