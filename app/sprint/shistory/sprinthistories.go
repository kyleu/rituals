// Content managed by Project Forge, see [projectforge.md] for details.
package shistory

import "golang.org/x/exp/slices"

type SprintHistories []*SprintHistory

func (s SprintHistories) Get(slug string) *SprintHistory {
	for _, x := range s {
		if x.Slug == slug {
			return x
		}
	}
	return nil
}

func (s SprintHistories) GetBySlugs(slugs ...string) SprintHistories {
	var ret SprintHistories
	for _, x := range s {
		if slices.Contains(slugs, x.Slug) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s SprintHistories) Slugs() []string {
	ret := make([]string, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (s SprintHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (s SprintHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s SprintHistories) Clone() SprintHistories {
	return slices.Clone(s)
}
