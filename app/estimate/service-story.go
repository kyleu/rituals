package estimate

import (
	"fmt"
	"strconv"

	"github.com/kyleu/rituals.dev/app/action"

	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) GetStories(estimateID uuid.UUID, params *query.Params) ([]*Story, error) {
	var defaultOrdering = query.Orderings{{Column: "idx", Asc: true}}

	params = query.ParamsWithDefaultOrdering(util.KeyStory, params, defaultOrdering...)
	var dtos []storyDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.KeyStory, "estimate_id = $1", params.OrderByString(), params.Limit, params.Offset), estimateID)
	if err != nil {
		return nil, err
	}
	return toStories(dtos), nil
}

func (s *Service) GetStoryByID(storyID uuid.UUID) (*Story, error) {
	dto := &storyDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.KeyStory, "id = $1", "", 0, 0), storyID)
	if err != nil {
		return nil, err
	}
	return dto.ToStory(), nil
}

func (s *Service) GetStoryEstimateID(storyID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := query.SQLSelect("estimate_id", util.KeyStory, "id = $1", "", 0, 0)
	err := s.db.Get(&ret, q, storyID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Service) NewStory(estimateID uuid.UUID, title string, authorID uuid.UUID) (*Story, error) {
	id := util.UUID()

	q := `insert into story (id, estimate_id, idx, author_id, title) values (
    $1, $2, coalesce((select max(idx) + 1 from story p2 where p2.estimate_id = $3), 0), $4, $5
	)`
	_, err := s.db.Exec(q, id, estimateID, estimateID, authorID, title)
	if err != nil {
		return nil, err
	}

	ret, err := s.GetStoryByID(id)
	if err == nil && ret != nil {
		postStory(s.actions, ret, authorID, action.ActStoryAdd)
	}

	return ret, err
}

func (s *Service) UpdateStory(storyID uuid.UUID, title string, userID uuid.UUID) (*Story, error) {
	q := `update story set title = $1 where id = $2`
	_, err := s.db.Exec(q, title, storyID)
	if err != nil {
		return nil, err
	}
	story, err := s.GetStoryByID(storyID)
	if story == nil {
		return nil, errors.New("cannot load newly-updated story")
	}

	postStory(s.actions, story, userID, action.ActStoryUpdate)
	return story, err
}

func (s *Service) RemoveStory(storyID uuid.UUID, userID uuid.UUID) error {
	story, err := s.GetStoryByID(storyID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot load report ["+storyID.String()+"] for removal"))
	}
	if story == nil {
		return errors.New("cannot load story [" + storyID.String() + "] for removal")
	}

	q1 := "delete from vote where story_id = $1"
	_, err = s.db.Exec(q1, storyID)
	if err != nil {
		return err
	}

	q2 := "delete from story where id = $1"
	_, err = s.db.Exec(q2, storyID)

	postStory(s.actions, story, userID, action.ActStoryRemove)
	return err
}

func postStory(actions *action.Service, story *Story, userID uuid.UUID, act string) {
	actionContent := map[string]interface{}{"storyID": story.ID}
	actions.Post(util.SvcEstimate.Key, story.EstimateID, userID, act, actionContent, "")
}

func (s *Service) SetStoryStatus(storyID uuid.UUID, status StoryStatus, userID uuid.UUID) (bool, string, error) {
	story, err := s.GetStoryByID(storyID)
	if err != nil {
		return false, "", errors.WithStack(errors.Wrap(err, "cannot load story ["+storyID.String()+"]"))
	}
	if story.Status == status {
		return false, "", nil
	}

	finalVote := ""

	if status == StoryStatusComplete {
		votes, err := s.GetStoryVotes(storyID, nil)
		if err != nil {
			return false, finalVote, errors.WithStack(errors.Wrap(err, "cannot load story votes for ["+storyID.String()+"]"))
		}
		finalVote = calcFinalVote(votes)
	}
	q := "update story set status = $1, final_vote = $2 where id = $3"
	_, err = s.db.Exec(q, status.String(), finalVote, storyID)

	actionContent := map[string]interface{}{"storyID": storyID, "status": status, "finalVote": finalVote}
	s.actions.Post(util.SvcEstimate.Key, story.EstimateID, userID, action.ActStoryStatus, actionContent, "")

	return true, finalVote, errors.WithStack(errors.Wrap(err, "error updating story status"))
}

func calcFinalVote(votes []*Vote) string {
	choices := make([]float64, 0)
	for _, v := range votes {
		f, err := strconv.ParseFloat(v.Choice, 64)
		if err == nil {
			choices = append(choices, f)
		}
	}

	sum := float64(0)
	for _, f := range choices {
		sum += f
	}

	final := float64(0)
	if len(choices) > 0 {
		final = sum / float64(len(choices))
	}
	ret := fmt.Sprint(final)
	min := 4
	if len(ret) < min {
		return ret
	}
	return ret[0:4]
}

func toStories(dtos []storyDTO) []*Story {
	ret := make([]*Story, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToStory())
	}
	return ret
}
