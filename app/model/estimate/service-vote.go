package estimate

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/model/action"

	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func (s *Service) GetStoryVotes(storyID uuid.UUID, params *query.Params) Votes {
	params = query.ParamsWithDefaultOrdering(util.KeyVote, params)
	var dtos []voteDTO
	q := query.SQLSelect("*", util.KeyVote, "story_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, storyID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving votes for story [%v]: %+v", storyID, err))
		return nil
	}
	return toVotes(dtos)
}

func (s *Service) GetEstimateVotes(estimateID uuid.UUID, params *query.Params) Votes {
	params = query.ParamsWithDefaultOrdering(util.KeyVote, params)
	var dtos []voteDTO
	q := query.SQLSelect("v.*", "vote v join story s on v.story_id = s.id", "s.estimate_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, estimateID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving votes for estimate [%v]: %+v", estimateID, err))
		return nil
	}
	return toVotes(dtos)
}

func (s *Service) GetVote(storyID uuid.UUID, userID uuid.UUID) (*Vote, error) {
	dto := &voteDTO{}
	q := query.SQLSelectSimple("*", util.KeyVote, "story_id = $1 and user_id = $2")
	err := s.db.Get(dto, q, nil, storyID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.toVote(), nil
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
		q := query.SQLInsert(util.KeyVote, []string{util.WithDBID(util.KeyStory), util.WithDBID(util.KeyUser), util.KeyChoice}, 1)
		err = s.db.Insert(q, nil, storyID, userID, choice)
		if err != nil {
			return nil, errors.Wrap(err, "error saving new vote for story ["+storyID.String()+"]")
		}

		actionContent := map[string]interface{}{"storyID": storyID, util.KeyChoice: choice}
		s.Data.Actions.Post(s.svc, *estimateID, userID, action.ActVoteAdd, actionContent)

		return &Vote{StoryID: storyID, UserID: userID, Choice: choice}, nil
	}

	cols := []string{util.KeyChoice}
	q := query.SQLUpdate(util.KeyVote, cols, fmt.Sprintf("story_id = $%v and user_id = $%v", len(cols)+1, len(cols)+1+1))
	err = s.db.UpdateOne(q, nil, choice, storyID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "error updating vote for story ["+storyID.String()+"]")
	}
	curr.Choice = choice

	actionContent := map[string]interface{}{util.WithID(util.KeyStory): storyID, util.KeyChoice: choice}
	s.Data.Actions.Post(s.svc, *estimateID, userID, action.ActVoteUpdate, actionContent)

	return curr, nil
}

func toVotes(dtos []voteDTO) Votes {
	ret := make(Votes, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toVote())
	}
	return ret
}
