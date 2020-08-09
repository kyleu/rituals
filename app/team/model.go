package team

import (
	"github.com/kyleu/npn/npncore"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID      uuid.UUID `json:"id"`
	Slug    string    `json:"slug"`
	Title   string    `json:"title"`
	Owner   uuid.UUID `json:"owner"`
	Created time.Time `json:"created"`
}

type Sessions []*Session

func NewSession(title string, slug string, userID uuid.UUID) Session {
	return Session{
		ID:      npncore.UUID(),
		Slug:    slug,
		Title:   strings.TrimSpace(title),
		Owner:   userID,
		Created: time.Time{},
	}
}

type sessionDTO struct {
	ID      uuid.UUID `db:"id"`
	Slug    string    `db:"slug"`
	Title   string    `db:"title"`
	Owner   uuid.UUID `db:"owner"`
	Created time.Time `db:"created"`
}

func (dto *sessionDTO) toSession() *Session {
	return &Session{
		ID:      dto.ID,
		Slug:    dto.Slug,
		Title:   dto.Title,
		Owner:   dto.Owner,
		Created: dto.Created,
	}
}
