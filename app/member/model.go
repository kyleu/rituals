package member

import (
	"encoding/json"
	"time"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
)

type Role struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var RoleOwner = Role{Key: util.KeyOwner, Title: "Owner", Description: "can do pretty much anything"}
var RoleMember = Role{Key: util.KeyMember, Title: "Member", Description: "can participate in the session"}
var RoleObserver = Role{Key: "observer", Title: "Observer", Description: "can only view the session"}

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

func (t *Role) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = RoleFromString(s)
	return nil
}

type entryDTO struct {
	UserID  uuid.UUID `db:"user_id"`
	Name    string    `db:"name"`
	Picture string    `db:"picture"`
	Role    string    `db:"role"`
	Created time.Time `db:"created"`
}

func (dto *entryDTO) toEntry() *Entry {
	return &Entry{
		UserID:  dto.UserID,
		Name:    dto.Name,
		Picture: dto.Picture,
		Role:    RoleFromString(dto.Role),
		Created: dto.Created,
	}
}

type Entry struct {
	UserID  uuid.UUID `json:"userID"`
	Name    string    `json:"name"`
	Picture string    `json:"picture"`
	Role    Role      `json:"role"`
	Created time.Time `json:"created"`
}

type Entries []*Entry

func (es Entries) GetName(userID uuid.UUID) string {
	for _, e := range es {
		if e.UserID == userID {
			return e.Name
		}
	}
	return "{former member}"
}
