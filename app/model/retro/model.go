package retro

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/kyleu/rituals.dev/app/database/query"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
)

type Status struct {
	Key string `json:"key"`
}

var StatusNew = Status{Key: "new"}
var StatusDeleted = Status{Key: "deleted"}

var AllStatuses = []Status{StatusNew, StatusDeleted}

func statusFromString(s string) Status {
	for _, t := range AllStatuses {
		if t.Key == s {
			return t
		}
	}
	return StatusNew
}

func (t *Status) String() string {
	return t.Key
}

func (t Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Status) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*t = statusFromString(s)
	return nil
}

var DefaultCategories = []string{"good", "bad", "improve"}

func categoriesFromDB(s string) []string {
	ret := query.StringToArray(s)
	if len(ret) == 0 {
		return DefaultCategories
	}
	return ret
}

type Session struct {
	ID         uuid.UUID  `json:"id"`
	Slug       string     `json:"slug"`
	Title      string     `json:"title"`
	TeamID     *uuid.UUID `json:"teamID"`
	SprintID   *uuid.UUID `json:"sprintID"`
	Owner      uuid.UUID  `json:"owner"`
	Status     Status     `json:"status"`
	Categories []string   `json:"categories"`
	Created    time.Time  `json:"created"`
}

type Sessions []*Session

func NewSession(title string, slug string, userID uuid.UUID, categories []string, teamID *uuid.UUID, sprintID *uuid.UUID) Session {
	return Session{
		ID:         util.UUID(),
		Slug:       slug,
		Title:      strings.TrimSpace(title),
		TeamID:     teamID,
		SprintID:   sprintID,
		Owner:      userID,
		Status:     StatusNew,
		Categories: categories,
		Created:    time.Now(),
	}
}

type sessionDTO struct {
	ID         uuid.UUID  `db:"id"`
	Slug       string     `db:"slug"`
	Title      string     `db:"title"`
	TeamID     *uuid.UUID `db:"team_id"`
	SprintID   *uuid.UUID `db:"sprint_id"`
	Owner      uuid.UUID  `db:"owner"`
	Status     string     `db:"status"`
	Categories string     `db:"categories"`
	Created    time.Time  `db:"created"`
}

func (dto *sessionDTO) ToSession() *Session {
	return &Session{
		ID:         dto.ID,
		Slug:       dto.Slug,
		Title:      dto.Title,
		TeamID:     dto.TeamID,
		SprintID:   dto.SprintID,
		Owner:      dto.Owner,
		Status:     statusFromString(dto.Status),
		Categories: categoriesFromDB(dto.Categories),
		Created:    dto.Created,
	}
}
