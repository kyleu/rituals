package retro

import (
	"strings"
	"time"

	"github.com/kyleu/rituals.dev/app/session"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"github.com/gofrs/uuid"
)

var DefaultCategories = []string{"good", "bad", "improve"}

func categoriesFromDB(s string) []string {
	ret := npndatabase.StringToArray(s)
	if len(ret) == 0 {
		return DefaultCategories
	}
	return ret
}

type Session struct {
	ID         uuid.UUID      `json:"id"`
	Slug       string         `json:"slug"`
	Title      string         `json:"title"`
	Status     session.Status `json:"status"`
	TeamID     *uuid.UUID     `json:"teamID"`
	SprintID   *uuid.UUID     `json:"sprintID"`
	Owner      uuid.UUID      `json:"owner"`
	Categories []string       `json:"categories"`
	Created    time.Time      `json:"created"`
}

type Sessions []*Session

func NewSession(title string, slug string, userID uuid.UUID, categories []string, teamID *uuid.UUID, sprintID *uuid.UUID) Session {
	return Session{
		ID:         npncore.UUID(),
		Slug:       slug,
		Title:      strings.TrimSpace(title),
		Status:     session.StatusNew,
		TeamID:     teamID,
		SprintID:   sprintID,
		Owner:      userID,
		Categories: categories,
		Created:    time.Now(),
	}
}

type sessionDTO struct {
	ID         uuid.UUID  `db:"id"`
	Slug       string     `db:"slug"`
	Title      string     `db:"title"`
	Status     string     `db:"status"`
	TeamID     *uuid.UUID `db:"team_id"`
	SprintID   *uuid.UUID `db:"sprint_id"`
	Owner      uuid.UUID  `db:"owner"`
	Categories string     `db:"categories"`
	Created    time.Time  `db:"created"`
}

func (dto *sessionDTO) toSession() *Session {
	return &Session{
		ID:         dto.ID,
		Slug:       dto.Slug,
		Title:      dto.Title,
		Status:     session.StatusFromString(dto.Status),
		TeamID:     dto.TeamID,
		SprintID:   dto.SprintID,
		Owner:      dto.Owner,
		Categories: categoriesFromDB(dto.Categories),
		Created:    dto.Created,
	}
}
