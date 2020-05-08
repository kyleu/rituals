package estimate

import (
	"time"

	"github.com/gofrs/uuid"
)

type StoryStatus struct {
	Key string `json:"key"`
}

var StoryStatusPending = StoryStatus{Key: "pending"}
var StoryStatusActive = StoryStatus{Key: "active"}
var StoryStatusComplete = StoryStatus{Key: "complete"}
var StoryStatusDeleted = StoryStatus{Key: "deleted"}

var AllStoryStatuses = []StoryStatus{StoryStatusPending, StoryStatusActive, StoryStatusComplete, StoryStatusDeleted}

func StoryStatusFromString(s string) StoryStatus {
	for _, t := range AllStoryStatuses {
		if t.String() == s {
			return t
		}
	}
	return StoryStatusPending
}

func (t StoryStatus) String() string {
	return t.Key
}

type Story struct {
	ID         uuid.UUID   `json:"id"`
	EstimateID uuid.UUID   `json:"estimateID"`
	Idx        uint        `json:"idx"`
	Author     uuid.UUID   `json:"author"`
	Title      string      `json:"title"`
	Status     StoryStatus `json:"status"`
	FinalVote  string      `json:"finalVote"`
	Created    time.Time   `json:"created"`
}

type storyDTO struct {
	ID         uuid.UUID `db:"id"`
	EstimateID uuid.UUID `db:"estimate_id"`
	Idx        uint      `db:"idx"`
	Author     uuid.UUID `db:"author_id"`
	Title      string    `db:"title"`
	Status     string    `db:"status"`
	FinalVote  string    `db:"final_vote"`
	Created    time.Time `db:"created"`
}

func (dto storyDTO) ToStory() Story {
	return Story{
		ID:         dto.ID,
		EstimateID: dto.EstimateID,
		Idx:        dto.Idx,
		Author:     dto.Author,
		Title:      dto.Title,
		Status:     StoryStatusFromString(dto.Status),
		FinalVote:  dto.FinalVote,
		Created:    dto.Created,
	}
}
