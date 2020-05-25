package standup

import (
	"time"

	"github.com/gofrs/uuid"
)

type Report struct {
	ID        uuid.UUID `json:"id"`
	StandupID uuid.UUID `json:"standupID"`
	D         string    `json:"d"`
	AuthorID  uuid.UUID `json:"authorID"`
	Content   string    `json:"content"`
	HTML      string    `json:"html"`
	Created   time.Time `json:"created"`
}

type Reports = []*Report

type reportDTO struct {
	ID        uuid.UUID `db:"id"`
	StandupID uuid.UUID `db:"standup_id"`
	D         time.Time `db:"d"`
	AuthorID  uuid.UUID `db:"author_id"`
	Content   string    `db:"content"`
	HTML      string    `db:"html"`
	Created   time.Time `db:"created"`
}

func (dto *reportDTO) ToReport() *Report {
	return &Report{
		ID:        dto.ID,
		StandupID: dto.StandupID,
		D:         dto.D.Format("2006-01-02"),
		AuthorID:  dto.AuthorID,
		Content:   dto.Content,
		HTML:      dto.HTML,
		Created:   dto.Created,
	}
}
