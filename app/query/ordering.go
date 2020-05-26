package query

import (
	"github.com/kyleu/rituals.dev/app/util"
)

type Ordering struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc"`
}

func (o Ordering) String() string {
	if o.Asc {
		return o.Column
	}
	return o.Column + "-desc"
}

type Orderings []*Ordering

var DefaultCreatedOrdering = Orderings{{Column: util.KeyCreated, Asc: false}}
var DefaultMCreatedOrdering = Orderings{{Column: "m." + util.KeyCreated, Asc: false}}
