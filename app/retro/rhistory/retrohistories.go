// Content managed by Project Forge, see [projectforge.md] for details.
package rhistory

import "golang.org/x/exp/slices"

type RetroHistories []*RetroHistory

func (r RetroHistories) Get(slug string) *RetroHistory {
	for _, x := range r {
		if x.Slug == slug {
			return x
		}
	}
	return nil
}

func (r RetroHistories) GetBySlugs(slugs ...string) RetroHistories {
	var ret RetroHistories
	for _, x := range r {
		if slices.Contains(slugs, x.Slug) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (r RetroHistories) Slugs() []string {
	ret := make([]string, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (r RetroHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (r RetroHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range r {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (r RetroHistories) Clone() RetroHistories {
	return slices.Clone(r)
}
