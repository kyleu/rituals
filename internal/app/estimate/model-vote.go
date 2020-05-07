package estimate

import (
	"github.com/gofrs/uuid"
	"time"
)

type Vote struct {
	UserID  uuid.UUID
	Choice  string
	Updated time.Time
	Created time.Time
}

type voteDTO struct {
	PollID  uuid.UUID `db:"poll_id"`
	UserID  uuid.UUID `db:"user_id"`
	Choice  string    `db:"choice"`
	Updated time.Time `db:"updated"`
	Created time.Time `db:"created"`
}

func (dto voteDTO) ToVote() Vote {
	return Vote{
		UserID:  dto.UserID,
		Choice:  dto.Choice,
		Updated: dto.Updated,
		Created: dto.Created,
	}
}
