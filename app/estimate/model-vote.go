package estimate

import (
	"time"

	"github.com/gofrs/uuid"
)

type Vote struct {
	StoryID uuid.UUID `json:"storyID"`
	UserID  uuid.UUID `json:"userID"`
	Choice  string    `json:"choice"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type voteDTO struct {
	StoryID uuid.UUID `db:"story_id"`
	UserID  uuid.UUID `db:"user_id"`
	Choice  string    `db:"choice"`
	Updated time.Time `db:"updated"`
	Created time.Time `db:"created"`
}

func (dto *voteDTO) ToVote() Vote {
	return Vote{
		StoryID: dto.StoryID,
		UserID:  dto.UserID,
		Choice:  dto.Choice,
		Updated: dto.Updated,
		Created: dto.Created,
	}
}
