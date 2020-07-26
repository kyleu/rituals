package estimate

import (
	"fmt"

	"github.com/kyleu/rituals.dev/app/action"

	"github.com/kyleu/rituals.dev/app/database/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) GetStories(estimateID uuid.UUID, params *query.Params) Stories {
	var defaultOrdering = query.Orderings{{Column: util.KeyIdx, Asc: true}}

	params = query.ParamsWithDefaultOrdering(util.KeyStory, params, defaultOrdering...)
	var dtos []storyDTO
	q := query.SQLSelect("*", util.KeyStory, "estimate_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, estimateID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving stories for estimate [%v]: %+v", estimateID, err))
		return nil
	}
	return toStories(dtos)
}

func (s *Service) GetStoryByID(storyID uuid.UUID) (*Story, error) {
	dto := &storyDTO{}
	q := query.SQLSelectSimple("*", util.KeyStory, util.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, storyID)
	if err != nil {
		return nil, err
	}
	return dto.toStory(), nil
}

func (s *Service) GetStoryEstimateID(storyID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := query.SQLSelectSimple(util.WithDBID(s.svc.Key), util.KeyStory, util.KeyID+" = $1")
	err := s.db.Get(&ret, q, nil, storyID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Service) NewStory(estimateID uuid.UUID, title string, userID uuid.UUID) (*Story, error) {
	id := util.UUID()

	q := `insert into story (id, estimate_id, idx, user_id, title, final_vote) values (
    $1, $2, coalesce((select max(idx) + 1 from story p2 where p2.estimate_id = $3), 0), $4, $5, ''
	)`
	err := s.db.Insert(q, nil, id, estimateID, estimateID, userID, title)
	if err != nil {
		return nil, err
	}

	ret, err := s.GetStoryByID(id)
	if err == nil && ret != nil {
		s.postStory(ret, userID, action.ActStoryAdd)
	}

	return ret, err
}

func (s *Service) UpdateStory(storyID uuid.UUID, title string, userID uuid.UUID) (*Story, error) {
	q := query.SQLUpdate(util.KeyStory, []string{util.KeyTitle}, util.KeyID+" = $2")
	err := s.db.UpdateOne(q, nil, title, storyID)
	if err != nil {
		return nil, err
	}
	story, err := s.GetStoryByID(storyID)
	if story == nil {
		return nil, errors.New("cannot load newly-updated story")
	}

	s.postStory(story, userID, action.ActStoryUpdate)
	return story, err
}

func (s *Service) RemoveStory(storyID uuid.UUID, userID uuid.UUID) error {
	story, err := s.GetStoryByID(storyID)
	if err != nil {
		return errors.Wrap(err, "cannot load report ["+storyID.String()+"] for removal")
	}
	if story == nil {
		return errors.New("cannot load story [" + storyID.String() + "] for removal")
	}

	q1 := query.SQLDelete(util.KeyVote, "story_id = $1")
	_, err = s.db.Delete(q1, nil, -1, storyID)
	if err != nil {
		return err
	}

	q2 := query.SQLDelete(util.KeyStory, util.KeyID+" = $1")
	err = s.db.DeleteOne(q2, nil, storyID)

	s.postStory(story, userID, action.ActStoryRemove)
	return err
}

func (s *Service) postStory(story *Story, userID uuid.UUID, act string) {
	actionContent := map[string]interface{}{"storyID": story.ID}
	s.Data.Actions.Post(s.svc, story.EstimateID, userID, act, actionContent)
}

func (s *Service) SetStoryStatus(storyID uuid.UUID, status StoryStatus, userID uuid.UUID) (bool, string, error) {
	story, err := s.GetStoryByID(storyID)
	if err != nil {
		return false, "", errors.Wrap(err, "cannot load story ["+storyID.String()+"]")
	}
	if story.Status == status {
		return false, "", nil
	}

	finalVote := ""

	if status == StoryStatusComplete {
		votes := s.GetStoryVotes(storyID, nil)
		vr := CalculateVoteResult(votes)
		finalVote = vr.FinalVote
	}
	cols := []string{"status", "final_vote"}
	q := query.SQLUpdate(util.KeyStory, cols, fmt.Sprintf("%v = $%v", util.KeyID, len(cols)+1))
	err = s.db.UpdateOne(q, nil, status.String(), finalVote, storyID)

	actionContent := map[string]interface{}{"storyID": storyID, util.KeyStatus: status, "finalVote": finalVote}
	s.Data.Actions.Post(s.svc, story.EstimateID, userID, action.ActStoryStatus, actionContent)

	return true, finalVote, errors.Wrap(err, "error updating story status")
}

func toStories(dtos []storyDTO) Stories {
	ret := make(Stories, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toStory())
	}
	return ret
}
