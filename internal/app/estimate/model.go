package estimate

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"strings"
	"time"
)

type Status struct {
	Key string
}

var StatusNew = Status{Key: "new"}
var StatusActive = Status{Key: "active"}
var StatusComplete = Status{Key: "complete"}
var StatusDeleted = Status{Key: "deleted"}

var AllStatuses = []Status{StatusNew, StatusActive, StatusComplete, StatusDeleted}

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

type SessionOptions struct {
	Foo string
}

func (o SessionOptions) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func optionsFromDB(x string) SessionOptions {
	return SessionOptions{Foo: x}
}
func choicesFromDB(s string) []string {
	ret := util.StringToArray(s)
	if len(ret) == 0 {
		return []string{"0", "0.5", "1", "2", "3", "5", "8", "13", "100", "?"}
	}
	return ret
}

type Session struct {
	ID       uuid.UUID
	Slug     string
	Password string
	Title    string
	Owner    uuid.UUID
	Status   Status
	Choices  []string
	Options  SessionOptions
	Created  time.Time
}

func NewSession(title string, slug string, userID uuid.UUID) Session {
	return Session{
		ID:      util.UUID(),
		Slug:    slug,
		Title:   strings.TrimSpace(title),
		Owner:   userID,
		Status:  StatusNew,
		Choices: nil,
		Options: SessionOptions{Foo: "TODO"},
		Created: time.Time{},
	}
}

type sessionDTO struct {
	ID       uuid.UUID `db:"id"`
	Slug     string    `db:"slug"`
	Password string    `db:"password"`
	Title    string    `db:"title"`
	Owner    uuid.UUID `db:"owner"`
	Status   string    `db:"status"`
	Choices  string    `db:"choices"`
	Options  string    `db:"options"`
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
		Choices:  choicesFromDB(dto.Choices),
		Options:  optionsFromDB(dto.Options),
		Created:  dto.Created,
	}
}

type PollStatus struct {
	Key string
}

var PollStatusPending = PollStatus{Key: "pending"}
var PollStatusDeleted = PollStatus{Key: "deleted"}

var AllPollStatuses = []PollStatus{PollStatusPending, PollStatusDeleted}

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
	ID         uuid.UUID
	EstimateID uuid.UUID
	Idx        uint
	Author     uuid.UUID
	Title      string
	Status     PollStatus
	FinalVote  string
	Created    time.Time
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

type Vote struct {
	PollID  uuid.UUID
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
		PollID:  dto.PollID,
		UserID:  dto.UserID,
		Choice:  dto.Choice,
		Updated: dto.Updated,
		Created: dto.Created,
	}
}
