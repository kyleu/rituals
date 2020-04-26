package retro

import (
	"github.com/gofrs/uuid"
	"time"
)

type Status struct {
	Key string
}

var StatusPending = Status{Key: "pending"}
var StatusRedeemed = Status{Key: "redeemed"}
var StatusDeleted = Status{Key: "deleted"}

var AllStatuses = []Status{StatusPending, StatusRedeemed, StatusDeleted}

func statusFromString(s string) Status {
	for _, t := range AllStatuses {
		if t.String() == s {
			return t
		}
	}
	return StatusPending
}

func (t Status) String() string {
	return t.Key
}

type Session struct {
	ID       uuid.UUID
	Slug     string
	Password string
	Title    string
	Owner    uuid.UUID
	Status   Status
	Created  time.Time
}

type sessionDTO struct {
	ID       uuid.UUID `db:"id"`
	Slug     string    `db:"slug"`
	Password string    `db:"password"`
	Title    string    `db:"title"`
	Owner    uuid.UUID `db:"owner"`
	Status   string    `db:"status"`
	Created  time.Time `db:"created"`
}

func (dto sessionDTO) ToSession() Session {
	return Session{
		ID:       dto.ID,
		Slug:     dto.Slug,
		Password: dto.Password,
		Title:    dto.Title,
		Owner:    dto.Owner,
		Status:   statusFromString(dto.Status),
		Created:  dto.Created,
	}
}
