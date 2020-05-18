package sprint

import (
	"strings"
	"time"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID      uuid.UUID  `json:"id"`
	Slug    string     `json:"slug"`
	Title   string     `json:"title"`
	Owner   uuid.UUID  `json:"owner"`
	EndDate *time.Time `json:"endDate"`
	Created time.Time  `json:"created"`
}

func NewSession(title string, slug string, userID uuid.UUID, endDate *time.Time) Session {
	return Session{
		ID:      util.UUID(),
		Slug:    slug,
		Title:   strings.TrimSpace(title),
		Owner:   userID,
		EndDate: endDate,
		Created: time.Time{},
	}
}

type sessionDTO struct {
	ID      uuid.UUID  `db:"id"`
	Slug    string     `db:"slug"`
	Title   string     `db:"title"`
	Owner   uuid.UUID  `db:"owner"`
	EndDate *time.Time `db:"end_date"`
	Created time.Time  `db:"created"`
}

func (dto *sessionDTO) ToSession() *Session {
	return &Session{
		ID:      dto.ID,
		Slug:    dto.Slug,
		Title:   dto.Title,
		Owner:   dto.Owner,
		EndDate: dto.EndDate,
		Created: dto.Created,
	}
}
