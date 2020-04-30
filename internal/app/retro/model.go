package retro

import (
	"time"

	"github.com/gofrs/uuid"
)

type Status struct {
	Key string `json:"key"`
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
	ID       uuid.UUID `json:"id"`
	Slug     string    `json:"slug"`
	Password string    `json:"password"`
	Title    string    `json:"title"`
	Owner    uuid.UUID `json:"owner"`
	Status   Status    `json:"status"`
	Created  time.Time `json:"created"`
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
