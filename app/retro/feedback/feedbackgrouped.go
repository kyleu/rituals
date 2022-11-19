package feedback

import (
	"golang.org/x/exp/slices"
)

type Group struct {
	Category  string    `json:"category"`
	Feedbacks Feedbacks `json:"feedbacks"`
}

func (f Feedbacks) Grouped(initial []string) []*Group {
	m := make(map[string]Feedbacks, len(f))
	for _, x := range f {
		curr := m[x.Category]
		curr = append(curr, x)
		m[x.Category] = curr
	}
	ret := make([]*Group, 0, len(initial)+len(m))
	for _, k := range initial {
		v := m[k]
		ret = append(ret, &Group{Category: k, Feedbacks: v})
	}
	for k, v := range m {
		if !slices.Contains(initial, k) {
			ret = append(ret, &Group{Category: k, Feedbacks: v})
		}
	}
	return ret
}
