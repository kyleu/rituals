package retro

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/markdown"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) NewFeedback(retroID uuid.UUID, category string, content string, authorID uuid.UUID) (*Feedback, error) {
	id := util.UUID()
	html := markdown.ToHTML(content)

	sql := `insert into feedback (id, retro_id, idx, author_id, category, content, html) values (
    $1, $2, coalesce((select max(idx) + 1 from feedback p2 where p2.retro_id = $3 and p2.category = $4), 0), $5, $6, $7, $8
	)`
	_, err := s.db.Exec(sql, id, retroID, retroID, category, authorID, category, content, html)
	if err != nil {
		return nil, err
	}

	return s.GetFeedbackByID(id)
}

func (s *Service) UpdateFeedback(feedbackID uuid.UUID, category string, content string) (*Feedback, error) {
	html := markdown.ToHTML(content)

	sql := `update feedback set category = $1, content = $2, html = $3 where id = $4`
	_, err := s.db.Exec(sql, category, content, html, feedbackID)
	if err != nil {
		return nil, err
	}

	return s.GetFeedbackByID(feedbackID)
}

func (s *Service) GetFeedback(retroID uuid.UUID) ([]Feedback, error) {
	var dtos []feedbackDTO
	err := s.db.Select(&dtos, "select * from feedback where retro_id = $1 order by category, idx, created", retroID)
	if err != nil {
		return nil, err
	}
	ret := make([]Feedback, 0, len(dtos))
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
	ret := dto.ToFeedback()
	return &ret, nil
}

func (s *Service) GetFeedbackRetroID(feedbackID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	err := s.db.Get(&ret, "select retro_id from feedback where id = $1", feedbackID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
