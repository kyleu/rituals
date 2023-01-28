package story

import (
	"strings"

	"golang.org/x/exp/slices"
)

func (s Stories) Sort() {
	slices.SortFunc(s, func(l, r *Story) bool {
		if l.Idx != r.Idx {
			return l.Idx < r.Idx
		}
		return strings.ToLower(l.Title) < strings.ToLower(r.Title)
	})
}
