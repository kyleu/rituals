// Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import "golang.org/x/exp/slices"

type TeamHistories []*TeamHistory

func (t TeamHistories) Get(slug string) *TeamHistory {
	for _, x := range t {
		if x.Slug == slug {
			return x
		}
	}
	return nil
}

func (t TeamHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.Slug)
	}
	return ret
}

func (t TeamHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range t {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (t TeamHistories) Clone() TeamHistories {
	return slices.Clone(t)
}
