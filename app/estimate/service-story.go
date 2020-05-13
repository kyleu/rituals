package estimate

import (
	"fmt"
	"strconv"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) NewStory(estimateID uuid.UUID, title string, authorID uuid.UUID) (*Story, error) {
	id := util.UUID()

	q := `insert into story (id, estimate_id, idx, author_id, title) values (
    $1, $2, coalesce((select max(idx) + 1 from story p2 where p2.estimate_id = $3), 0), $4, $5
	)`
	_, err := s.db.Exec(q, id, estimateID, estimateID, authorID, title)
	if err != nil {
		return nil, err
	}

	return s.GetStoryByID(id)
}

func (s *Service) GetStories(estimateID uuid.UUID) ([]*Story, error) {
	var dtos []storyDTO
	err := s.db.Select(&dtos, "select * from story where estimate_id = $1 order by idx", estimateID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Story, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToStory())
	}
	return ret, nil
}

func (s *Service) GetStoryByID(storyID uuid.UUID) (*Story, error) {
	dto := &storyDTO{}
	err := s.db.Get(dto, "select * from story where id = $1", storyID)
	if err != nil {
		return nil, err
	}
	return dto.ToStory(), nil
}

func (s *Service) GetStoryEstimateID(storyID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	err := s.db.Get(&ret, "select estimate_id from story where id = $1", storyID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Service) UpdateStory(_ uuid.UUID, storyID uuid.UUID, title string) (*Story, error) {
	q := `update story set title = $1 where id = $2`
	_, err := s.db.Exec(q, title, storyID)
	if err != nil {
		return nil, err
	}
	return s.GetStoryByID(storyID)
}

func (s *Service) RemoveStory(storyID uuid.UUID, _ uuid.UUID) error {
	q1 := "delete from vote where story_id = $1"
	_, err := s.db.Exec(q1, storyID)
	if err != nil {
		return err
	}

	q2 := "delete from story where id = $1"
	_, err = s.db.Exec(q2, storyID)
	return err
}

func (s *Service) SetStoryStatus(storyID uuid.UUID, status StoryStatus) (bool, string, error) {
	story, err := s.GetStoryByID(storyID)
	if err != nil {
		return false, "", errors.WithStack(errors.Wrap(err, "cannot load story ["+storyID.String()+"]"))
	}
	if story.Status == status {
		return false, "", nil
	}

	finalVote := ""
	if status == StoryStatusComplete {
		votes, err := s.GetStoryVotes(storyID)
		if err != nil {
			return false, finalVote, errors.WithStack(errors.Wrap(err, "cannot load story votes for ["+storyID.String()+"]"))
		}
		finalVote = calcFinalVote(votes)
	}
	q := "update story set status = $1, final_vote = $2 where id = $3"
	_, err = s.db.Exec(q, status.String(), finalVote, storyID)
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
	if len(ret) < 4 {
		return ret
	}
	return ret[0:4]
}
