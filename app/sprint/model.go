package sprint

import (
	"strings"
	"time"

	"github.com/kyleu/npn/npncore"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID        uuid.UUID  `json:"id"`
	Slug      string     `json:"slug"`
	Title     string     `json:"title"`
	TeamID    *uuid.UUID `json:"teamID"`
	Owner     uuid.UUID  `json:"owner"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
	Created   time.Time  `json:"created"`
}

type Sessions []*Session

func NewSession(title string, slug string, userID uuid.UUID, teamID *uuid.UUID, startDate *time.Time, endDate *time.Time) Session {
	return Session{
		ID:        npncore.UUID(),
		Slug:      slug,
		Title:     strings.TrimSpace(title),
		TeamID:    teamID,
		Owner:     userID,
		StartDate: startDate,
		EndDate:   endDate,
		Created:   time.Time{},
	}
}

type sessionDTO struct {
	ID        uuid.UUID  `db:"id"`
	Slug      string     `db:"slug"`
	Title     string     `db:"title"`
	TeamID    *uuid.UUID `db:"team_id"`
	Owner     uuid.UUID  `db:"owner"`
	Status    string     `db:"status"`
	StartDate *time.Time `db:"start_date"`
	EndDate   *time.Time `db:"end_date"`
	Created   time.Time  `db:"created"`
}

func (dto *sessionDTO) toSession() *Session {
	return &Session{
		ID:        dto.ID,
		Slug:      dto.Slug,
		Title:     dto.Title,
		TeamID:    dto.TeamID,
		Owner:     dto.Owner,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
		Created:   dto.Created,
	}
}
