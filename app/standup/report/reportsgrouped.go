package report

import (
	"golang.org/x/exp/slices"
	"time"
)

type Group struct {
	Day     time.Time `json:"day"`
	Reports Reports   `json:"reports"`
}

func (r Reports) Grouped() []*Group {
	var ret []*Group
	slices.SortFunc(ret, func(l *Group, r *Group) bool {
		return l.Day.UnixMilli() < r.Day.UnixMilli()
	})
	return ret
}
