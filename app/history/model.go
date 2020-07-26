package history

import (
	"time"

	"github.com/gofrs/uuid"
)

type entryDTO struct {
	Slug      string    `db:"slug"`
	ModelID   uuid.UUID `db:"model_id"`
	ModelName string    `db:"model_name"`
	Created   time.Time `db:"created"`
}

func (dto *entryDTO) toEntry() *Entry {
	return &Entry{
		Slug:      dto.Slug,
		ModelID:   dto.ModelID,
		ModelName: dto.ModelName,
		Created:   dto.Created,
	}
}

type Entry struct {
	Slug      string    `json:"slug"`
	ModelID   uuid.UUID `json:"modelId"`
	ModelName string    `json:"modelName"`
	Created   time.Time `json:"created"`
}

type Entries []*Entry
