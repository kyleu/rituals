package estimate

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/kyleu/rituals.dev/app/query"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

type Status struct {
	Key string `json:"key"`
}

var StatusNew = Status{Key: "new"}
var StatusActive = Status{Key: "active"}
var StatusComplete = Status{Key: "complete"}
var StatusDeleted = Status{Key: "deleted"}

var AllStatuses = []Status{StatusNew, StatusActive, StatusComplete, StatusDeleted}

var DefaultChoices = []string{"0", "0.5", "1", "2", "3", "5", "8", "13", "100", "?"}

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

type SessionOptions struct {
	Foo string `json:"foo"`
}

func (o *SessionOptions) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func optionsFromDB(x string) SessionOptions {
	return SessionOptions{Foo: x}
}

func choicesFromDB(s string) []string {
	ret := query.StringToArray(s)
	if len(ret) == 0 {
		return DefaultChoices
	}
	return ret
}

type Session struct {
	ID       uuid.UUID      `json:"id"`
	Slug     string         `json:"slug"`
	Title    string         `json:"title"`
	SprintID *uuid.UUID     `json:"sprintID"`
	Owner    uuid.UUID      `json:"owner"`
	Status   Status         `json:"status"`
	Choices  []string       `json:"choices"`
	Options  SessionOptions `json:"options"`
	Created  time.Time      `json:"created"`
}

func NewSession(title string, slug string, userID uuid.UUID, sprintID *uuid.UUID) Session {
	return Session{
		ID:       util.UUID(),
		Slug:     slug,
		Title:    strings.TrimSpace(title),
		SprintID: sprintID,
		Owner:    userID,
		Status:   StatusNew,
		Choices:  nil,
		Options:  SessionOptions{Foo: ""},
		Created:  time.Time{},
	}
}

type sessionDTO struct {
	ID       uuid.UUID  `db:"id"`
	Slug     string     `db:"slug"`
	Title    string     `db:"title"`
	SprintID *uuid.UUID `db:"sprint_id"`
	Owner    uuid.UUID  `db:"owner"`
	Status   string     `db:"status"`
	Choices  string     `db:"choices"`
	Options  string     `db:"options"`
	Created  time.Time  `db:"created"`
}

func (dto *sessionDTO) ToSession() *Session {
	return &Session{
		ID:       dto.ID,
		Slug:     dto.Slug,
		Title:    dto.Title,
		SprintID: dto.SprintID,
		Owner:    dto.Owner,
		Status:   statusFromString(dto.Status),
		Choices:  choicesFromDB(dto.Choices),
		Options:  optionsFromDB(dto.Options),
		Created:  dto.Created,
	}
}
