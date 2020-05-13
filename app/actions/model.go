package actions

import (
	"github.com/gofrs/uuid"
	"time"
)

type Action struct {
	ID       uuid.UUID `json:"id"`
	Svc      string    `json:"svc"`
	ModelID  uuid.UUID `json:"modelID"`
	AuthorID uuid.UUID `json:"authorID"`
	Act      string    `json:"act"`
	Content  string    `json:"content"`
	Note     string    `json:"note"`
	Occurred time.Time `json:"occurred"`
}

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
	return &Action{
		ID:       dto.ID,
		Svc:      dto.Svc,
		ModelID:  dto.ModelID,
		AuthorID: dto.AuthorID,
		Act:      dto.Act,
		Content:  dto.Content,
		Note:     dto.Note,
		Occurred:  dto.Occurred,
	}
}
