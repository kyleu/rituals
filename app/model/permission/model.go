package permission

import (
	"sort"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/model/member"
)

type permissionDTO struct {
	K       string    `db:"k"`
	V       string    `db:"v"`
	Access  string    `db:"access"`
	Created time.Time `db:"created"`
}

func (dto *permissionDTO) toPermission() *Permission {
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
	Created time.Time   `json:"-"`
}

type Permissions []*Permission

func (ps Permissions) FindByK(k string) Permissions {
	var ret Permissions
	for _, e := range ps {
		if e.K == k {
			ret = append(ret, e)
		}
	}
	return ret
}

func (ps Permissions) FindByKV(k string, v string) *Permission {
	for _, e := range ps {
		if e.K == k && e.V == v {
			return e
		}
	}
	return nil
}

func (ps Permissions) Sort() Permissions {
	sort.Slice(ps, func(i, j int) bool {
		l, r := ps[i], ps[j]
		if l.K == r.K {
			return l.V < r.V
		}
		return l.K < r.K
	})
	return ps
}

func (ps Permissions) Equals(filtered Permissions) bool {
	if len(ps) != len(filtered) {
		return false
	}
	for _, c := range ps {
		if filtered.FindByKV(c.K, c.V) == nil {
			return false
		}
	}
	return true
}

type Error struct {
	Svc      string `json:"svc"`
	Provider string `json:"provider"`
	Match    string `json:"match"`
	Message  string `json:"message"`
}

type Errors []*Error

func (es Errors) ToError() error {
	if len(es) == 0 {
		return nil
	}
	var msgs = make([]string, 0, len(es))
	for _, e := range es {
		msgs = append(msgs, e.Message)
	}
	return errors.New("permission error: " + strings.Join(msgs, ", "))
}

func (es Errors) FindByProvider(p string) Errors {
	var ret Errors
	for _, e := range es {
		if e.Provider == p {
			ret = append(ret, e)
		}
	}
	return ret
}

func (es Errors) GetMatches() []string {
	ret := make([]string, 0)
	for _, e := range es {
		var x []string
		for _, m := range strings.Split(e.Match, ",") {
			if len(strings.TrimSpace(m)) > 0 {
				x = append(x, strings.TrimSpace(m))
			}
		}
		ret = append(ret, x...)
	}
	return ret
}
