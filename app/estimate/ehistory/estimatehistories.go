// Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import "golang.org/x/exp/slices"

type EstimateHistories []*EstimateHistory

func (e EstimateHistories) Get(slug string) *EstimateHistory {
	for _, x := range e {
		if x.Slug == slug {
			return x
		}
	}
	return nil
}

func (e EstimateHistories) GetBySlugs(slugs ...string) EstimateHistories {
	var ret EstimateHistories
	for _, x := range e {
		if slices.Contains(slugs, x.Slug) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (e EstimateHistories) Slugs() []string {
	ret := make([]string, 0, len(e)+1)
	for _, x := range e {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (e EstimateHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (e EstimateHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range e {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (e EstimateHistories) Clone() EstimateHistories {
	return slices.Clone(e)
}
