package standup

import (
	"time"

	"github.com/gofrs/uuid"
)

type Update struct {
	ID        uuid.UUID `json:"id"`
	StandupID uuid.UUID `json:"standupID"`
	D         string    `json:"d"`
	Author    uuid.UUID `json:"author"`
	Content   string    `json:"content"`
	Created   time.Time `json:"created"`
}

type updateDTO struct {
	ID        uuid.UUID `db:"id"`
	StandupID uuid.UUID `db:"standup_id"`
	D         time.Time `db:"d"`
	Author    uuid.UUID `db:"author_id"`
	Content   string    `db:"content"`
	Created   time.Time `db:"created"`
}

func (dto updateDTO) ToUpdate() Update {
	return Update{
		ID:        dto.ID,
		StandupID: dto.StandupID,
		D:         dto.D.Format("2006-01-02"),
		Author:    dto.Author,
		Content:   dto.Content,
		Created:   dto.Created,
	}
}
