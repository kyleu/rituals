package member

import (
	"github.com/gofrs/uuid"
	"time"
)

type Entry struct {
	UserID  uuid.UUID `db:"user_id"`
	Name    string    `db:"name"`
	Role    string    `db:"role"`
	Created time.Time `db:"created"`
}
