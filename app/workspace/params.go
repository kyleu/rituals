package workspace

import (
	"context"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/util"
)

type Params struct {
	Ctx      context.Context
	Slug     string
	Act      action.Act
	Frm      util.ValueMap
	Profile  *user.Profile
	Accounts user.Accounts
	ConnIDs  []uuid.UUID
	Svc      *Service
	Logger   util.Logger
}

func NewParams(
	ctx context.Context, slug string, act action.Act, frm util.ValueMap,
	profile *user.Profile, accounts user.Accounts, svc *Service,
	logger util.Logger, except ...uuid.UUID,
) *Params {
	return &Params{Ctx: ctx, Slug: slug, Act: act, Frm: frm, Profile: profile, Accounts: accounts, ConnIDs: except, Svc: svc, Logger: logger}
}
