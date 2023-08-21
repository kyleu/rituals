package story

import (
	"cmp"
	"slices"
	"strings"
)

func (s Stories) Sort() {
	slices.SortFunc(s, func(l, r *Story) int {
		if l.Idx != r.Idx {
			return cmp.Compare(l.Idx, r.Idx)
		}
		return cmp.Compare(strings.ToLower(l.Title), strings.ToLower(r.Title))
	})
}
