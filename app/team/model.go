package team

import (
	"github.com/kyleu/rituals.dev/app/session"
	"strings"
	"time"

	"github.com/kyleu/npn/npncore"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID      uuid.UUID      `json:"id"`
	Slug    string         `json:"slug"`
	Title   string         `json:"title"`
	Status  session.Status `json:"status"`
	Owner   uuid.UUID      `json:"owner"`
	Created time.Time      `json:"created"`
}

type Sessions []*Session

func NewSession(title string, slug string, userID uuid.UUID) Session {
	return Session{
		ID:      npncore.UUID(),
		Slug:    slug,
		Title:   strings.TrimSpace(title),
		Status:  session.StatusNew,
		Owner:   userID,
		Created: time.Time{},
	}
}

type sessionDTO struct {
	ID      uuid.UUID `db:"id"`
	Slug    string    `db:"slug"`
	Title   string    `db:"title"`
	Status    string     `db:"status"`
	Owner   uuid.UUID `db:"owner"`
	Created time.Time `db:"created"`
}

func (dto *sessionDTO) toSession() *Session {
	return &Session{
		ID:      dto.ID,
		Slug:    dto.Slug,
		Title:   dto.Title,
		Status:  session.StatusFromString(dto.Status),
		Owner:   dto.Owner,
		Created: dto.Created,
	}
}
