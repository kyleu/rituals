package standup

import (
	"github.com/kyleu/npn/npncore"
	"time"

	"github.com/gofrs/uuid"
)

type Report struct {
	ID        uuid.UUID `json:"id"`
	StandupID uuid.UUID `json:"standupID"`
	D         string    `json:"d"`
	UserID    uuid.UUID `json:"userID"`
	Content   string    `json:"content"`
	HTML      string    `json:"html"`
	Created   time.Time `json:"created"`
}

type Reports = []*Report

type reportDTO struct {
	ID        uuid.UUID `db:"id"`
	StandupID uuid.UUID `db:"standup_id"`
	D         time.Time `db:"d"`
	UserID    uuid.UUID `db:"user_id"`
	Content   string    `db:"content"`
	HTML      string    `db:"html"`
	Created   time.Time `db:"created"`
}

func (dto *reportDTO) toReport() *Report {
	return &Report{
		ID:        dto.ID,
		StandupID: dto.StandupID,
		D:         npncore.ToYMD(&dto.D),
		UserID:    dto.UserID,
		Content:   dto.Content,
		HTML:      dto.HTML,
		Created:   dto.Created,
	}
}
