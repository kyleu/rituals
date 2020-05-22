package permission

import (
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/member"
)

type permissionDTO struct {
	K       string    `db:"k"`
	V       string    `db:"v"`
	Access  string    `db:"access"`
	Created time.Time `db:"created"`
}

func (dto *permissionDTO) ToPermission() *Permission {
	return &Permission{
		K:       dto.K,
		V:       dto.V,
		Access:  member.RoleFromString(dto.Access),
		Created: dto.Created,
	}
}

type Permission struct {
	K       string      `json:"k"`
	V       string      `json:"v"`
	Access  member.Role `json:"access"`
	Created time.Time   `json:"created"`
}

type Permissions []*Permission

func (es Permissions) FindByK(code string) Permissions {
	var ret Permissions
	for _, e := range es {
		if e.K == code {
			ret = append(ret, e)
		}
	}
	return ret
}

type Error struct {
	K       string `json:"k"`
	V       string `json:"v"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Errors []*Error

func (es Errors) ToError() error {
	if len(es) == 0 {
		return nil
	}
	var msgs = make([]string, len(es))
	for i, e := range es {
		msgs[i] = e.Message
	}
	return errors.WithStack(errors.New("permission error: " + strings.Join(msgs, ", ")))
}
