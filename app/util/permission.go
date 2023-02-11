package util

type Permission struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Permissions []*Permission

func (p Permissions) Values() []string {
	ret := make([]string, 0, len(p))
	for _, x := range p {
		ret = append(ret, x.Value)
	}
	return ret
}

func (p Permissions) TeamPerms() Permissions {
	var ret Permissions
	for _, x := range p {
		if x.Key == KeyTeam {
			ret = append(ret, x)
		}
	}
	return ret
}

func (p Permissions) SprintPerms() Permissions {
	var ret Permissions
	for _, x := range p {
		if x.Key == KeySprint {
			ret = append(ret, x)
		}
	}
	return ret
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
