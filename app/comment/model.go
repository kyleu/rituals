package comment

import (
	"github.com/kyleu/rituals.dev/app/util"
	"time"

	"github.com/gofrs/uuid"
)

type commentDTO struct {
	ID       uuid.UUID `db:"id"`
	Svc      string    `db:"svc"`
	ModelID  uuid.UUID `db:"model_id"`
	AuthorID uuid.UUID `db:"author_id"`
	Content  string    `db:"content"`
	HTML     string    `db:"html"`
	Created  time.Time `db:"created"`
}

func (dto *commentDTO) ToComment() *Comment {
	return &Comment{
		ID:       dto.ID,
		Svc:      util.ServiceFromString(dto.Svc),
		ModelID:  dto.ModelID,
		AuthorID: dto.AuthorID,
		Content:  dto.Content,
		HTML:     dto.HTML,
		Created:  dto.Created,
	}
}

type Comment struct {
	ID       uuid.UUID    `json:"id"`
	Svc      util.Service `json:"-"`
	ModelID  uuid.UUID    `json:"-"`
	AuthorID uuid.UUID    `json:"authorID"`
	Content  string       `json:"content"`
	HTML     string       `json:"html"`
	Created  time.Time    `json:"created"`
}

type Comments []*Comment
