package member

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Role struct {
	Key string `json:"key"`
}

var RoleOwner = Role{Key: "owner"}
var RoleMember = Role{Key: "member"}
var RoleObserver = Role{Key: "observer"}

var AllRoles = []Role{RoleOwner, RoleMember, RoleObserver}

func RoleFromString(s string) Role {
	for _, t := range AllRoles {
		if t.Key == s {
			return t
		}
	}
	return RoleObserver
}

func (t *Role) String() string {
	return t.Key
}

func (t Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

type entryDTO struct {
	UserID  uuid.UUID `db:"user_id"`
	Name    string    `db:"name"`
	Role    string    `db:"role"`
	Created time.Time `db:"created"`
}

func (dto *entryDTO) ToEntry() *Entry {
	return &Entry{
		UserID:  dto.UserID,
		Name:    dto.Name,
		Role:    RoleFromString(dto.Role),
		Created: dto.Created,
	}
}

type Entry struct {
	UserID  uuid.UUID `json:"userID"`
	Name    string    `json:"name"`
	Role    Role      `json:"role"`
	Created time.Time `json:"created"`
}

type Entries []*Entry
