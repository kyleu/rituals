package workspace

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/util"
)

type FullStandup struct {
	Standup     *standup.Standup               `json:"standup"`
	Histories   uhistory.StandupHistories      `json:"histories"`
	Members     umember.StandupMembers         `json:"members"`
	Permissions upermission.StandupPermissions `json:"permissions"`
}

func (s *Service) LoadStandup(ctx context.Context, slug string, user uuid.UUID, logger util.Logger) (*FullStandup, error) {
	bySlug, err := s.u.GetBySlug(ctx, nil, slug, nil, logger)
	if err != nil {
		return nil, err
	}
	if len(bySlug) == 0 {
		id := util.UUIDFromString(slug)
		if id == nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", slug)
		}
		u, err := s.u.Get(ctx, nil, *id, logger)
		if err != nil {
			return nil, errors.Errorf("no estimate found with slug [%s]", slug)
		}
		bySlug = standup.Standups{u}
	}
	u := bySlug[0]

	hist, err := s.uh.GetByStandupID(ctx, nil, u.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	members, err := s.um.GetByStandupID(ctx, nil, u.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	perms, err := s.up.GetByStandupID(ctx, nil, u.ID, nil, logger)
	if err != nil {
		return nil, err
	}

	return &FullStandup{Standup: u, Histories: hist, Members: members, Permissions: perms}, nil
}
