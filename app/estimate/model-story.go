package estimate

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type StoryStatus struct {
	Key string `json:"key"`
}

var StoryStatusPending = StoryStatus{Key: "pending"}
var StoryStatusActive = StoryStatus{Key: "active"}
var StoryStatusComplete = StoryStatus{Key: "complete"}

var AllStoryStatuses = []StoryStatus{StoryStatusPending, StoryStatusActive, StoryStatusComplete}

func StoryStatusFromString(s string) StoryStatus {
	for _, t := range AllStoryStatuses {
		if t.Key == s {
			return t
		}
	}
	return StoryStatusPending
}

func (t *StoryStatus) String() string {
	return t.Key
}

func (t StoryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *StoryStatus) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = StoryStatusFromString(s)
	return nil
}

type Story struct {
	ID         uuid.UUID   `json:"id"`
	EstimateID uuid.UUID   `json:"estimateID"`
	Idx        uint        `json:"idx"`
	UserID     uuid.UUID   `json:"userID"`
	Title      string      `json:"title"`
	Status     StoryStatus `json:"status"`
	FinalVote  string      `json:"finalVote"`
	Created    time.Time   `json:"created"`
}

type Stories = []*Story

type storyDTO struct {
	ID         uuid.UUID `db:"id"`
	EstimateID uuid.UUID `db:"estimate_id"`
	Idx        uint      `db:"idx"`
	UserID     uuid.UUID `db:"user_id"`
	Title      string    `db:"title"`
	Status     string    `db:"status"`
	FinalVote  string    `db:"final_vote"`
	Created    time.Time `db:"created"`
}

func (dto *storyDTO) toStory() *Story {
	return &Story{
		ID:         dto.ID,
		EstimateID: dto.EstimateID,
		Idx:        dto.Idx,
		UserID:     dto.UserID,
		Title:      dto.Title,
		Status:     StoryStatusFromString(dto.Status),
		FinalVote:  dto.FinalVote,
		Created:    dto.Created,
	}
}
