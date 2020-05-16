package retro

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/markdown"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) GetFeedback(retroID uuid.UUID) ([]*Feedback, error) {
	var dtos []feedbackDTO
	err := s.db.Select(&dtos, "select * from feedback where retro_id = $1 order by category, idx, created", retroID)
	if err != nil {
		return nil, err
	}
	ret := make([]*Feedback, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToFeedback())
	}
	return ret, nil
}

func (s *Service) GetFeedbackByID(feedbackID uuid.UUID) (*Feedback, error) {
	dto := &feedbackDTO{}
	err := s.db.Get(dto, "select * from feedback where id = $1", feedbackID)
	if err != nil {
		return nil, err
	}
	return dto.ToFeedback(), nil
}

func (s *Service) GetFeedbackRetroID(feedbackID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	err := s.db.Get(&ret, "select retro_id from feedback where id = $1", feedbackID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Service) NewFeedback(retroID uuid.UUID, category string, content string, authorID uuid.UUID) (*Feedback, error) {
	id := util.UUID()
	html := markdown.ToHTML(content)

	q := `insert into feedback (id, retro_id, idx, author_id, category, content, html) values (
    $1, $2, coalesce((select max(idx) + 1 from feedback p2 where p2.retro_id = $3 and p2.category = $4), 0), $5, $6, $7, $8
	)`
	_, err := s.db.Exec(q, id, retroID, retroID, category, authorID, category, content, html)
	if err != nil {
		return nil, err
	}

	actionContent := map[string]interface{}{"feedbackID": id}
	s.actions.Post(util.SvcRetro.Key, retroID, authorID, "add-feedback", actionContent, "")

	return s.GetFeedbackByID(id)
}

func (s *Service) UpdateFeedback(feedbackID uuid.UUID, category string, content string, userID uuid.UUID) (*Feedback, error) {
	html := markdown.ToHTML(content)

	q := `update feedback set category = $1, content = $2, html = $3 where id = $4`
	_, err := s.db.Exec(q, category, content, html, feedbackID)
	if err != nil {
		return nil, err
	}

	fb, err := s.GetFeedbackByID(feedbackID)
	if fb == nil {
		return nil, errors.New("cannot load newly-updated feedback")
	}

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.actions.Post(util.SvcRetro.Key, fb.RetroID, userID, "update-feedback", actionContent, "")

	return s.GetFeedbackByID(feedbackID)
}

func (s *Service) RemoveFeedback(feedbackID uuid.UUID, userID uuid.UUID) error {
	feedback, err := s.GetFeedbackByID(feedbackID)
	if feedback == nil {
		return errors.New("cannot load feedback [" + feedbackID.String() + "] for removal")
	}

	q := "delete from feedback where id = $1"
	_, err = s.db.Exec(q, feedbackID)

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.actions.Post(util.SvcRetro.Key, feedback.RetroID, userID, "remove-feedback", actionContent, "")

	return err
}
