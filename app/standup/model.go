package standup

import (
	"github.com/kyleu/rituals.dev/app/session"
	"strings"
	"time"

	"github.com/kyleu/npn/npncore"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID       uuid.UUID      `json:"id"`
	Slug     string         `json:"slug"`
	Title    string         `json:"title"`
	Status   session.Status `json:"status"`
	TeamID   *uuid.UUID     `json:"teamID"`
	SprintID *uuid.UUID     `json:"sprintID"`
	Owner    uuid.UUID      `json:"owner"`
	Created  time.Time      `json:"created"`
}

type Sessions []*Session

func NewSession(title string, slug string, userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID) Session {
	return Session{
		ID:       npncore.UUID(),
		Slug:     slug,
		Title:    strings.TrimSpace(title),
		Status:   session.StatusNew,
		Owner:    userID,
		TeamID:   teamID,
		SprintID: sprintID,
		Created:  time.Time{},
	}
}

type sessionDTO struct {
	ID       uuid.UUID  `db:"id"`
	Slug     string     `db:"slug"`
	Title    string     `db:"title"`
	Status   string     `db:"status"`
	TeamID   *uuid.UUID `db:"team_id"`
	SprintID *uuid.UUID `db:"sprint_id"`
	Owner    uuid.UUID  `db:"owner"`
	Created  time.Time  `db:"created"`
}

func (dto *sessionDTO) toSession() *Session {
	return &Session{
		ID:       dto.ID,
		Slug:     dto.Slug,
		Title:    dto.Title,
		Status:   session.StatusFromString(dto.Status),
		TeamID:   dto.TeamID,
		SprintID: dto.SprintID,
		Owner:    dto.Owner,
		Created:  dto.Created,
	}
}
