package socket

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/permission"
)

// team
type teamSessionSaveParams struct {
	Title       string                 `json:"title"`
	Permissions permission.Permissions `json:"permissions"`
}

// sprint
type sprintSessionSaveParams struct {
	Title       string                 `json:"title"`
	StartDate   string                 `json:"startDate"`
	EndDate     string                 `json:"endDate"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

// estimate
type estimateSessionSaveParams struct {
	Title       string                 `json:"title"`
	Choices     string                 `json:"choices"`
	SprintID    string                 `json:"sprintID"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

type addStoryParams struct {
	Title string `json:"title"`
}

type updateStoryParams struct {
	StoryID uuid.UUID `json:"storyID"`
	Title   string    `json:"title"`
}

type setStoryStatusParams struct {
	StoryID uuid.UUID `json:"storyID"`
	Status  string    `json:"status"`
}

type submitVoteParams struct {
	StoryID uuid.UUID `json:"storyID"`
	Choice  string    `json:"choice"`
}

// standup
type standupSessionSaveParams struct {
	Title       string                 `json:"title"`
	SprintID    string                 `json:"sprintID"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

type addReportParams struct {
	D       string `json:"d"`
	Content string `json:"content"`
}

type editReportParams struct {
	ID      uuid.UUID `json:"id"`
	D       string    `json:"d"`
	Content string    `json:"content"`
}

// retro
type retroSessionSaveParams struct {
	Title       string                 `json:"title"`
	Categories  string                 `json:"categories"`
	SprintID    string                 `json:"sprintID"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

type addFeedbackParams struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

type editFeedbackParams struct {
	ID       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Content  string    `json:"content"`
}
