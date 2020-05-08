package estimate

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) NewStory(estimateID uuid.UUID, title string, authorID uuid.UUID) (*Story, error) {
	id := util.UUID()

	sql := `insert into story (id, estimate_id, idx, author_id, title) values (
    $1, $2, coalesce((select max(idx) + 1 from story p2 where p2.estimate_id = $3), 0), $4, $5
	)`
	_, err := s.db.Exec(sql, id, estimateID, estimateID, authorID, title)
	if err != nil {
		return nil, err
	}

	return s.GetStoryByID(id)
}

func (s *Service) GetStories(id uuid.UUID) ([]Story, error) {
	var dtos []storyDTO
	err := s.db.Select(&dtos, "select * from story where estimate_id = $1 order by idx", id)
	if err != nil {
		return nil, err
	}
	ret := make([]Story, 0)
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
	ret := dto.ToStory()
	return &ret, nil
}

func (s *Service) GetStoryEstimateID(storyID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	err := s.db.Get(&ret, "select estimate_id from story where id = $1", storyID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Service) SetStoryStatus(storyID uuid.UUID, status StoryStatus) (bool, error) {
	story, err := s.GetStoryByID(storyID)
	if err != nil {
		return false, errors.WithStack(errors.Wrap(err, "cannot load story ["+storyID.String()+"]"))
	}
	if story.Status == status {
		return false, nil
	}
	q := "update story set status = $1 where id = $2"
	_, err = s.db.Exec(q, status.String(), storyID)
	return true, errors.WithStack(errors.Wrap(err, "error updating story status"))
}
