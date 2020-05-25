package estimate

import (
	"database/sql"

	"github.com/kyleu/rituals.dev/app/action"

	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func (s *Service) GetStoryVotes(storyID uuid.UUID, params *query.Params) (Votes, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyStory, params)
	var dtos []voteDTO
	q := query.SQLSelect("*", util.KeyVote, "story_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, storyID)
	if err != nil {
		return nil, err
	}
	return toVotes(dtos), nil
}

func (s *Service) GetEstimateVotes(estimateID uuid.UUID, params *query.Params) (Votes, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyVote, params)
	var dtos []voteDTO
	q := query.SQLSelect("v.*", "vote v join story s on v.story_id = s.id", "s.estimate_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, estimateID)
	if err != nil {
		return nil, err
	}
	return toVotes(dtos), nil
}

func (s *Service) GetVote(storyID uuid.UUID, userID uuid.UUID) (*Vote, error) {
	dto := &voteDTO{}
	q := query.SQLSelect("*", util.KeyVote, "story_id = $1 and user_id = $2", "", 0, 0)
	err := s.db.Get(dto, q, nil, storyID, userID)
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
		return nil, errors.Wrap(err, "error getting current votes for story ["+storyID.String()+"]")
	}
	curr, err := s.GetVote(storyID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current votes for story ["+storyID.String()+"]")
	}
	if curr == nil {
		q := "insert into vote (story_id, user_id, choice) values ($1, $2, $3)"
		err = s.db.Insert(q, nil, storyID, userID, choice)
		if err != nil {
			return nil, errors.Wrap(err, "error saving new vote for story ["+storyID.String()+"]")
		}

		actionContent := map[string]interface{}{"storyID": storyID, "choice": choice}
		s.actions.Post(util.SvcEstimate.Key, *estimateID, userID, action.ActVoteAdd, actionContent, "")

		return &Vote{StoryID: storyID, UserID: userID, Choice: choice}, nil
	}

	q := "update vote set choice = $1 where story_id = $2 and user_id = $3"
	err = s.db.UpdateOne(q, nil, choice, storyID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "error updating vote for story ["+storyID.String()+"]")
	}
	curr.Choice = choice

	actionContent := map[string]interface{}{"storyID": storyID, "choice": choice}
	s.actions.Post(util.SvcEstimate.Key, *estimateID, userID, action.ActVoteUpdate, actionContent, "")

	return curr, nil
}

func toVotes(dtos []voteDTO) Votes {
	ret := make(Votes, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToVote())
	}
	return ret
}
