package retro

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/markdown"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/util"
)

var defaultFeedbackOrdering = []*query.Ordering{{Column: "category", Asc: true}, {Column: "idx", Asc: true}, {Column: "created", Asc: false}}

func (s *Service) GetFeedback(retroID uuid.UUID, params *query.Params) ([]*Feedback, error) {
	params = query.ParamsWithDefaultOrdering("feddback", params, defaultFeedbackOrdering...)
	var dtos []feedbackDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", "feedback", "retro_id = $1", params.OrderByString(), params.Limit, params.Offset), retroID)
	if err != nil {
		return nil, err
	}
	return toFeedbacks(dtos), nil
}

func (s *Service) GetFeedbackByID(feedbackID uuid.UUID) (*Feedback, error) {
	dto := &feedbackDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", "feedback", "id = $1", "", 0, 0), feedbackID)
	if err != nil {
		return nil, err
	}
	return dto.ToFeedback(), nil
}

func (s *Service) GetFeedbackRetroID(feedbackID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := query.SQLSelect("retro_id", "feedback", "id = $1", "", 0, 0)
	err := s.db.Get(&ret, q, feedbackID)
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
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "cannot load feedback ["+feedbackID.String()+"] for update"))
	}
	if fb == nil {
		return nil, errors.New("cannot load newly-updated feedback")
	}

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.actions.Post(util.SvcRetro.Key, fb.RetroID, userID, "update-feedback", actionContent, "")

	return s.GetFeedbackByID(feedbackID)
}

func (s *Service) RemoveFeedback(feedbackID uuid.UUID, userID uuid.UUID) error {
	feedback, err := s.GetFeedbackByID(feedbackID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot load report ["+feedbackID.String()+"] for removal"))
	}
	if feedback == nil {
		return errors.New("cannot load feedback [" + feedbackID.String() + "] for removal")
	}

	q := "delete from feedback where id = $1"
	_, err = s.db.Exec(q, feedbackID)

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.actions.Post(util.SvcRetro.Key, feedback.RetroID, userID, "remove-feedback", actionContent, "")

	return err
}

func toFeedbacks(dtos []feedbackDTO) []*Feedback {
	ret := make([]*Feedback, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToFeedback())
	}
	return ret
}
