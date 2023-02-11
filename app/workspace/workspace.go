package workspace

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
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

type LoadParams struct {
	Ctx      context.Context
	Slug     string
	Profile  *user.Profile
	Accounts user.Accounts
	Tx       *sqlx.Tx
	Params   filter.ParamSet
	Logger   util.Logger
}

func NewLoadParams(
	ctx context.Context, slug string, profile *user.Profile, accts user.Accounts, tx *sqlx.Tx, params filter.ParamSet, logger util.Logger,
) *LoadParams {
	return &LoadParams{Ctx: ctx, Slug: slug, Profile: profile, Accounts: accts, Tx: tx, Params: params, Logger: logger}
}
