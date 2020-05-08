package standup

import (
	"github.com/kyleu/rituals.dev/app/util"
	"strings"
	"time"

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
		if t.String() == s {
			return t
		}
	}
	return StatusNew
}

func (t Status) String() string {
	return t.Key
}

type Session struct {
	ID      uuid.UUID `json:"id"`
	Slug    string    `json:"slug"`
	Title   string    `json:"title"`
	Owner   uuid.UUID `json:"owner"`
	Status  Status    `json:"status"`
	Created time.Time `json:"created"`
}

func NewSession(title string, slug string, userID uuid.UUID) Session {
	return Session{
		ID:      util.UUID(),
		Slug:    slug,
		Title:   strings.TrimSpace(title),
		Owner:   userID,
		Status:  StatusNew,
		Created: time.Time{},
	}
}

type sessionDTO struct {
	ID      uuid.UUID `db:"id"`
	Slug    string    `db:"slug"`
	Title   string    `db:"title"`
	Owner   uuid.UUID `db:"owner"`
	Status  string    `db:"status"`
	Created time.Time `db:"created"`
}

func (dto sessionDTO) ToSession() Session {
	return Session{
		ID:      dto.ID,
		Slug:    dto.Slug,
		Title:   dto.Title,
		Owner:   dto.Owner,
		Status:  statusFromString(dto.Status),
		Created: dto.Created,
	}
}
