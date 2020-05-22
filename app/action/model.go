package action

import (
	"encoding/json"
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

const (
	ActCreate  = "create"
	ActConnect = "connect"
	ActUpdate  = "update"

	ActMemberAdd     = "member-add"
	ActPermissionAdd = "permission-add"

	ActContentAdd    = "content-add"
	ActContentUpdate = "content-update"
	ActContentRemove = "content-remove"

	ActFeedbackAdd    = "feedback-add"
	ActFeedbackUpdate = "feedback-update"
	ActFeedbackRemove = "feedback-remove"

	ActReportAdd    = "report-add"
	ActReportUpdate = "report-update"
	ActReportRemove = "report-remove"

	ActStoryAdd    = "story-add"
	ActStoryUpdate = "story-update"
	ActStoryRemove = "story-remove"
	ActStoryStatus = "story-status"

	ActVoteAdd    = "vote-add"
	ActVoteUpdate = "vote-update"
)

type Action struct {
	ID       uuid.UUID   `json:"id"`
	Svc      string      `json:"svc"`
	ModelID  uuid.UUID   `json:"modelID"`
	AuthorID uuid.UUID   `json:"authorID"`
	Act      string      `json:"act"`
	Content  interface{} `json:"content"`
	Note     string      `json:"note"`
	Occurred time.Time   `json:"occurred"`
}

func (a *Action) ContentJSON() (string, error) {
	bytes, err := json.MarshalIndent(a.Content, "", "  ")

	if err != nil {
		return "", errors.WithStack(errors.Wrap(err, "error marshalling action content"))
	}

	return string(bytes), nil
}

type Actions []*Action

type actionDTO struct {
	ID       uuid.UUID `db:"id"`
	Svc      string    `db:"svc"`
	ModelID  uuid.UUID `db:"model_id"`
	AuthorID uuid.UUID `db:"author_id"`
	Act      string    `db:"act"`
	Content  string    `db:"content"`
	Note     string    `db:"note"`
	Occurred time.Time `db:"occurred"`
}

func (dto *actionDTO) ToAction() *Action {
	var param interface{}
	_ = json.Unmarshal([]byte(dto.Content), &param)

	return &Action{
		ID:       dto.ID,
		Svc:      dto.Svc,
		ModelID:  dto.ModelID,
		AuthorID: dto.AuthorID,
		Act:      dto.Act,
		Content:  param,
		Note:     dto.Note,
		Occurred: dto.Occurred,
	}
}
