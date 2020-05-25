package retro

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/markdown"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/util"
)

var defaultFeedbackOrdering = query.Orderings{{Column: "category", Asc: true}, {Column: "idx", Asc: true}, {Column: util.KeyCreated, Asc: false}}

func (s *Service) GetFeedback(retroID uuid.UUID, params *query.Params) (Feedbacks, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyFeedback, params, defaultFeedbackOrdering...)
	var dtos []feedbackDTO
	q := query.SQLSelect("*", util.KeyFeedback, "retro_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, retroID)
	if err != nil {
		return nil, err
	}
	return toFeedbacks(dtos), nil
}

func (s *Service) GetFeedbackByID(feedbackID uuid.UUID) (*Feedback, error) {
	dto := &feedbackDTO{}
	q := query.SQLSelect("*", util.KeyFeedback, "id = $1", "", 0, 0)
	err := s.db.Get(dto, q, nil, feedbackID)
	if err != nil {
		return nil, err
	}
	return dto.ToFeedback(), nil
}

func (s *Service) GetFeedbackRetroID(feedbackID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := query.SQLSelect("retro_id", util.KeyFeedback, "id = $1", "", 0, 0)
	err := s.db.Get(&ret, q, nil, feedbackID)
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
	err := s.db.Insert(q, nil, id, retroID, retroID, category, authorID, category, content, html)
	if err != nil {
		return nil, err
	}

	actionContent := map[string]interface{}{"feedbackID": id}
	s.actions.Post(util.SvcRetro.Key, retroID, authorID, action.ActFeedbackAdd, actionContent, "")

	return s.GetFeedbackByID(id)
}

func (s *Service) UpdateFeedback(feedbackID uuid.UUID, category string, content string, userID uuid.UUID) (*Feedback, error) {
	html := markdown.ToHTML(content)

	q := `update feedback set category = $1, content = $2, html = $3 where id = $4`
	err := s.db.UpdateOne(q, nil, category, content, html, feedbackID)
	if err != nil {
		return nil, err
	}

	fb, err := s.GetFeedbackByID(feedbackID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load feedback ["+feedbackID.String()+"] for update")
	}
	if fb == nil {
		return nil, errors.New("cannot load newly-updated feedback")
	}

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.actions.Post(util.SvcRetro.Key, fb.RetroID, userID, action.ActFeedbackUpdate, actionContent, "")

	return s.GetFeedbackByID(feedbackID)
}

func (s *Service) RemoveFeedback(feedbackID uuid.UUID, userID uuid.UUID) error {
	feedback, err := s.GetFeedbackByID(feedbackID)
	if err != nil {
		return errors.Wrap(err, "cannot load feedback ["+feedbackID.String()+"] for removal")
	}
	if feedback == nil {
		return errors.New("cannot load feedback [" + feedbackID.String() + "] for removal")
	}

	q := "delete from feedback where id = $1"
	_, err = s.db.Delete(q, nil, feedbackID)

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.actions.Post(util.SvcRetro.Key, feedback.RetroID, userID, action.ActFeedbackRemove, actionContent, "")

	return err
}

func toFeedbacks(dtos []feedbackDTO) Feedbacks {
	ret := make(Feedbacks, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToFeedback())
	}
	return ret
}
