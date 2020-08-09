package email

import (
	"time"

	"github.com/kyleu/npn/npndatabase"

	"github.com/gofrs/uuid"
)

type Email struct {
	ID         string      `json:"id"`
	Recipients []string    `json:"recipients"`
	Subject    string      `json:"subject"`
	Data       interface{} `json:"data"`
	Plain      string      `json:"plain"`
	HTML       string      `json:"html"`
	UserID     uuid.UUID   `json:"userID"`
	Status     string      `json:"status"`
	Created    time.Time   `json:"created"`
}

type Emails []*Email

type emailDTO struct {
	ID         string    `db:"id"`
	Recipients string    `db:"recipients"`
	Subject    string    `db:"subject"`
	Data       string    `db:"data"`
	Plain      string    `db:"plain"`
	HTML       string    `db:"html"`
	UserID     uuid.UUID `db:"user_id"`
	Status     string    `db:"status"`
	Created    time.Time `db:"created"`
}

func (dto *emailDTO) toEmail() *Email {
	return &Email{
		ID:         dto.ID,
		Recipients: npndatabase.StringToArray(dto.Recipients),
		Subject:    dto.Subject,
		Data:       dto.Data,
		Plain:      dto.Plain,
		HTML:       dto.HTML,
		UserID:     dto.UserID,
		Status:     dto.Status,
		Created:    dto.Created,
	}
}
