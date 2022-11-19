package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
)

type Workspace struct {
	Teams     team.Teams         `json:"teams"`
	Sprints   sprint.Sprints     `json:"sprints"`
	Estimates estimate.Estimates `json:"estimates"`
	Standups  standup.Standups   `json:"standups"`
	Retros    retro.Retros       `json:"retros"`
}

func FromAny(x any) (*Workspace, error) {
	if x == nil {
		return nil, errors.New("data is nil, not [*Workspace]")
	}
	w, ok := x.(*Workspace)
	if !ok {
		return nil, errors.Errorf("data is [%T], not [*Workspace]", x)
	}
	return w, nil
}
