package comment

import (
	"time"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
)

type commentDTO struct {
	ID         uuid.UUID  `db:"id"`
	Svc        string     `db:"svc"`
	ModelID    uuid.UUID  `db:"model_id"`
	TargetType string     `db:"target_type"`
	TargetID   *uuid.UUID `db:"target_id"`
	UserID     uuid.UUID  `db:"user_id"`
	Content    string     `db:"content"`
	HTML       string     `db:"html"`
	Created    time.Time  `db:"created"`
}

func (dto *commentDTO) toComment() *Comment {
	return &Comment{
		ID:         dto.ID,
		Svc:        util.ServiceFromString(dto.Svc),
		ModelID:    dto.ModelID,
		TargetType: dto.TargetType,
		TargetID:   dto.TargetID,
		UserID:     dto.UserID,
		Content:    dto.Content,
		HTML:       dto.HTML,
		Created:    dto.Created,
	}
}

type Comment struct {
	ID         uuid.UUID    `json:"id"`
	Svc        util.Service `json:"-"`
	ModelID    uuid.UUID    `json:"-"`
	TargetType string       `json:"targetType"`
	TargetID   *uuid.UUID   `json:"targetID"`
	UserID     uuid.UUID    `json:"userID"`
	Content    string       `json:"content"`
	HTML       string       `json:"html"`
	Created    time.Time    `json:"created"`
}

type Comments []*Comment
