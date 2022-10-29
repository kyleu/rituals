// Content managed by Project Forge, see [projectforge.md] for details.
package user

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/util"
)

var (
	table         = "user"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "picture", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID      uuid.UUID  `db:"id"`
	Name    string     `db:"name"`
	Picture string     `db:"picture"`
	Created time.Time  `db:"created"`
	Updated *time.Time `db:"updated"`
}

func (d *dto) ToUser() *User {
	if d == nil {
		return nil
	}
	return &User{
		ID:      d.ID,
		Name:    d.Name,
		Picture: d.Picture,
		Created: d.Created,
		Updated: d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToUsers() Users {
	ret := make(Users, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToUser())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
