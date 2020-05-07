package estimate

import (
	"github.com/gofrs/uuid"
	"time"
)

type PollStatus struct {
	Key string `json:"key"`
}

var PollStatusPending = PollStatus{Key: "pending"}
var PollStatusActive = PollStatus{Key: "active"}
var PollStatusComplete = PollStatus{Key: "complete"}
var PollStatusDeleted = PollStatus{Key: "deleted"}

var AllPollStatuses = []PollStatus{PollStatusPending, PollStatusActive, PollStatusComplete, PollStatusDeleted}

func pollStatusFromString(s string) PollStatus {
	for _, t := range AllPollStatuses {
		if t.String() == s {
			return t
		}
	}
	return PollStatusPending
}

func (t PollStatus) String() string {
	return t.Key
}

type Poll struct {
	ID         uuid.UUID  `json:"id"`
	EstimateID uuid.UUID  `json:"estimateID"`
	Idx        uint       `json:"idx"`
	Author     uuid.UUID  `json:"author"`
	Title      string     `json:"title"`
	Status     PollStatus `json:"status"`
	FinalVote  string     `json:"finalVote"`
	Created    time.Time  `json:"created"`
}

type pollDTO struct {
	ID         uuid.UUID `db:"id"`
	EstimateID uuid.UUID `db:"estimate_id"`
	Idx        uint      `db:"idx"`
	Author     uuid.UUID `db:"author_id"`
	Title      string    `db:"title"`
	Status     string    `db:"status"`
	FinalVote  string    `db:"final_vote"`
	Created    time.Time `db:"created"`
}

func (dto pollDTO) ToPoll() Poll {
	return Poll{
		ID:         dto.ID,
		EstimateID: dto.EstimateID,
		Idx:        dto.Idx,
		Author:     dto.Author,
		Title:      dto.Title,
		Status:     pollStatusFromString(dto.Status),
		FinalVote:  dto.FinalVote,
		Created:    dto.Created,
	}
}

