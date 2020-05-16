package estimate

import (
	"database/sql"
	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func (s *Service) GetStoryVotes(storyID uuid.UUID) ([]*Vote, error) {
	var dtos []voteDTO
	err := s.db.Select(&dtos, "select * from vote where story_id = $1", storyID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Vote, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToVote())
	}
	return ret, nil
}

func (s *Service) GetEstimateVotes(estimateID uuid.UUID) ([]*Vote, error) {
	var dtos []voteDTO
	err := s.db.Select(&dtos, "select v.* from vote v join story s on v.story_id = s.id where s.estimate_id = $1", estimateID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Vote, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToVote())
	}
	return ret, nil
}

func (s *Service) GetVote(storyID uuid.UUID, userID uuid.UUID) (*Vote, error) {
	dto := &voteDTO{}
	err := s.db.Get(dto, "select * from vote v where v.story_id = $1 and v.user_id = $2", storyID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToVote(), nil
}

func (s *Service) UpdateVote(storyID uuid.UUID, userID uuid.UUID, choice string) (*Vote, error) {
	estimateID, err := s.GetStoryEstimateID(storyID)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error getting current votes for story ["+storyID.String()+"]"))
	}
	curr, err := s.GetVote(storyID, userID)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error getting current votes for story ["+storyID.String()+"]"))
	}
	if curr == nil {
		q := "insert into vote (story_id, user_id, choice) values ($1, $2, $3)"
		_, err = s.db.Exec(q, storyID, userID, choice)
		if err != nil {
			return nil, errors.WithStack(errors.Wrap(err, "error saving new vote for story ["+storyID.String()+"]"))
		}

		actionContent := map[string]interface{}{"storyID": storyID, "choice": choice}
		s.actions.Post(util.SvcEstimate.Key, *estimateID, userID, "add-vote", actionContent, "")

		return &Vote{StoryID: storyID, UserID: userID, Choice: choice}, nil
	} else {
		q := "update vote set choice = $1 where story_id = $2 and user_id = $3"
		_, err := s.db.Exec(q, choice, storyID, userID)
		if err != nil {
			return nil, errors.WithStack(errors.Wrap(err, "error updating vote for story ["+storyID.String()+"]"))
		}
		curr.Choice = choice

		actionContent := map[string]interface{}{"storyID": storyID, "choice": choice}
		s.actions.Post(util.SvcEstimate.Key, *estimateID, userID, "update-vote", actionContent, "")

		return curr, nil
	}
}
