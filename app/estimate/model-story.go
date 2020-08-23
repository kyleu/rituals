package estimate

import (
	"time"

	"github.com/kyleu/rituals.dev/app/session"

	"github.com/gofrs/uuid"
)

type Story struct {
	ID         uuid.UUID      `json:"id"`
	EstimateID uuid.UUID      `json:"estimateID"`
	Idx        uint           `json:"idx"`
	UserID     uuid.UUID      `json:"userID"`
	Title      string         `json:"title"`
	Status     session.Status `json:"status"`
	FinalVote  string         `json:"finalVote"`
	Created    time.Time      `json:"created"`
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
		Status:     session.StatusFromString(dto.Status),
		FinalVote:  dto.FinalVote,
		Created:    dto.Created,
	}
}
