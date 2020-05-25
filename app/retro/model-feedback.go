package retro

import (
	"time"

	"github.com/gofrs/uuid"
)

type Feedback struct {
	ID       uuid.UUID `json:"id"`
	RetroID  uuid.UUID `json:"retroID"`
	Idx      uint      `json:"idx"`
	AuthorID uuid.UUID `json:"authorID"`
	Category string    `json:"category"`
	Content  string    `json:"content"`
	HTML     string    `json:"html"`
	Created  time.Time `json:"created"`
}

type Feedbacks = []*Feedback

type feedbackDTO struct {
	ID       uuid.UUID `db:"id"`
	RetroID  uuid.UUID `db:"retro_id"`
	Idx      uint      `db:"idx"`
	AuthorID uuid.UUID `db:"author_id"`
	Category string    `db:"category"`
	Content  string    `db:"content"`
	HTML     string    `db:"html"`
	Created  time.Time `db:"created"`
}

func (dto *feedbackDTO) ToFeedback() *Feedback {
	return &Feedback{
		ID:       dto.ID,
		RetroID:  dto.RetroID,
		Idx:      dto.Idx,
		AuthorID: dto.AuthorID,
		Category: dto.Category,
		Content:  dto.Content,
		HTML:     dto.HTML,
		Created:  dto.Created,
	}
}
