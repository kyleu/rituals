package report

import (
	"time"

	"golang.org/x/exp/slices"
)

type Group struct {
	Day     time.Time `json:"day"`
	Reports Reports   `json:"reports"`
}

func (r Reports) Grouped() []*Group {
	m := make(map[time.Time]Reports, len(r))
	for _, x := range r {
		curr := m[x.Day]
		curr = append(curr, x)
		m[x.Day] = curr
	}
	ret := make([]*Group, 0, len(m))
	for k, v := range m {
		ret = append(ret, &Group{Day: k, Reports: v})
	}
	slices.SortFunc(ret, func(l *Group, r *Group) bool {
		return l.Day.UnixMilli() < r.Day.UnixMilli()
	})
	return ret
}
