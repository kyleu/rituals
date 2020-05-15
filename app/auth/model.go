package auth

import (
	"github.com/gofrs/uuid"
	"time"
)

type Record struct {
	ID      uuid.UUID  `json:"id"`
	UserID  uuid.UUID  `json:"userID"`
	K       string     `json:"k"`
	V       string     `json:"v"`
	Expires *time.Time `json:"expires"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Picture string     `json:"picture"`
	Created time.Time  `json:"created"`
}

type recordDTO struct {
	ID      uuid.UUID  `db:"id"`
	UserID  uuid.UUID  `db:"user_id"`
	K       string     `db:"k"`
	V       string     `db:"v"`
	Expires *time.Time `db:"expires"`
	Name    string     `db:"name"`
	Email   string     `db:"email"`
	Picture string     `db:"picture"`
	Created time.Time  `db:"created"`
}

func (dto *recordDTO) ToRecord() *Record {
	return &Record{
		ID:      dto.ID,
		UserID:  dto.UserID,
		K:       dto.K,
		V:       dto.V,
		Expires: dto.Expires,
		Name:    dto.Name,
		Email:   dto.Email,
		Picture: dto.Picture,
		Created: dto.Created,
	}
}

