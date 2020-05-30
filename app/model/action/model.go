package action

import (
	"encoding/json"
	"github.com/kyleu/rituals.dev/app/util"
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

const (
	ActCreate  = "create"
	ActConnect = "connect"
	ActUpdate  = "update"

	ActMemberAdd   = "member-add"
	ActPermissions = "permissions"

	ActContentAdd = "content-add"

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
	ID       uuid.UUID    `json:"id"`
	Svc      util.Service `json:"svc"`
	ModelID  uuid.UUID    `json:"modelID"`
	UserID uuid.UUID    `json:"userID"`
	Act      string       `json:"act"`
	Content  interface{}  `json:"content"`
	Note     string       `json:"note"`
	Created  time.Time    `json:"created"`
}

func (a *Action) ContentJSON() (string, error) {
	bytes, err := json.MarshalIndent(a.Content, "", "  ")

	if err != nil {
		return "", errors.Wrap(err, "error marshalling action content")
	}

	return string(bytes), nil
}

type Actions []*Action

type actionDTO struct {
	ID       uuid.UUID `db:"id"`
	Svc      string    `db:"svc"`
	ModelID  uuid.UUID `db:"model_id"`
	UserID uuid.UUID `db:"user_id"`
	Act      string    `db:"act"`
	Content  string    `db:"content"`
	Note     string    `db:"note"`
	Created  time.Time `db:"created"`
}

func (dto *actionDTO) ToAction() *Action {
	var param interface{}
	_ = json.Unmarshal([]byte(dto.Content), &param)

	return &Action{
		ID:       dto.ID,
		Svc:      util.ServiceFromString(dto.Svc),
		ModelID:  dto.ModelID,
		UserID: dto.UserID,
		Act:      dto.Act,
		Content:  param,
		Note:     dto.Note,
		Created:  dto.Created,
	}
}
