package util

import "fmt"

type Permission struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p Permission) String() string {
	if p.Value == "" {
		return p.Key
	}
	return p.Key + ":" + p.Value
}

type Permissions []*Permission

func (p Permissions) Values() []string {
	ret := make([]string, 0, len(p))
	for _, x := range p {
		ret = append(ret, x.Value)
	}
	return ret
}

func (p Permissions) TeamPerm() *Permission {
	for _, x := range p {
		if x.Key == KeyTeam {
			return x
		}
	}
	return nil
}

func (p Permissions) SprintPerm() *Permission {
	for _, x := range p {
		if x.Key == KeySprint {
			return x
		}
	}
	return nil
}

func (p Permissions) AuthPerms() Permissions {
	var ret Permissions
	for _, x := range p {
		if x.Key != KeyTeam && x.Key != KeySprint {
			ret = append(ret, x)
		}
	}
	return ret
}

func (p Permissions) Diff(tgt Permissions) Diffs {
	if len(p) != len(tgt) {
		return Diffs{{Path: "length", Old: fmt.Sprint(len(p)), New: fmt.Sprint(len(tgt))}}
	}
	for _, s := range p {
		if t := tgt.Get(s.Key, s.Value); t == nil {
			return Diffs{{Path: "missing", Old: s.String(), New: ""}}
		}
	}
	return nil
}

func (p Permissions) Get(key string, value string) *Permission {
	for _, x := range p {
		if x.Key == key && x.Value == value {
			return x
		}
	}
	return nil
}
