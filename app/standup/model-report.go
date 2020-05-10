package standup

import (
	"time"

	"github.com/gofrs/uuid"
)

type Report struct {
	ID        uuid.UUID `json:"id"`
	StandupID uuid.UUID `json:"standupID"`
	D         string    `json:"d"`
	Author    uuid.UUID `json:"author"`
	Content   string    `json:"content"`
	HTML   string       `json:"html"`
	Created   time.Time `json:"created"`
}

type reportDTO struct {
	ID        uuid.UUID `db:"id"`
	StandupID uuid.UUID `db:"standup_id"`
	D         time.Time `db:"d"`
	Author    uuid.UUID `db:"author_id"`
	Content   string    `db:"content"`
	HTML      string    `db:"html"`
	Created   time.Time `db:"created"`
}

func (dto reportDTO) ToReport() Report {
	return Report{
		ID:        dto.ID,
		StandupID: dto.StandupID,
		D:         dto.D.Format("2006-01-02"),
		Author:    dto.Author,
		Content:   dto.Content,
		HTML:      dto.HTML,
		Created:   dto.Created,
	}
}
