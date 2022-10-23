package cmenu

import (
	"context"
	"fmt"
	"github.com/kyleu/rituals/app/lib/user"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/app/lib/telemetry"
	"github.com/kyleu/rituals/app/util"
)

var workspaceServices = []enum.ModelService{
	enum.ModelServiceTeam, enum.ModelServiceSprint, enum.ModelServiceEstimate, enum.ModelServiceStandup, enum.ModelServiceRetro,
}

func workspaceMenu(ctx context.Context, as *app.State, profile *user.Profile, logger util.Logger) (menu.Items, error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "workspace.menu", logger)
	defer span.Complete()
	u := profile.ID
	items, errs := util.AsyncCollectMap(workspaceServices, func(item enum.ModelService) enum.ModelService {
		return item
	}, func(k enum.ModelService) (*menu.Item, error) {
		switch k {
		case enum.ModelServiceTeam:
			return teamMenu(ctx, u, as, logger)
		case enum.ModelServiceSprint:
			return sprintMenu(ctx, u, as, logger)
		case enum.ModelServiceEstimate:
			return estimateMenu(ctx, u, as, logger)
		case enum.ModelServiceStandup:
			return standupMenu(ctx, u, as, logger)
		case enum.ModelServiceRetro:
			return retroMenu(ctx, u, as, logger)
		default:
			return nil, errors.Errorf("invalid service [%s]", k)
		}
	})
	if len(errs) > 0 {
		return nil, util.ErrorMerge(maps.Values(errs)...)
	}
	ret := menu.Items{
		items[enum.ModelServiceTeam],
		items[enum.ModelServiceSprint],
		items[enum.ModelServiceEstimate],
		items[enum.ModelServiceStandup],
		items[enum.ModelServiceRetro],
	}
	return ret, nil
}

func teamMenu(ctx context.Context, user uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "teams", Title: "Teams", Description: "TODO", Icon: "users", Route: "/team"}
	t, err := as.Services.Team.GetByOwner(ctx, nil, user, nil, logger)
	if err != nil {
		return nil, err
	}
	for _, x := range t {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "users", Route: "/team/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(t))
	return ret, nil
}

func sprintMenu(ctx context.Context, user uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "sprints", Title: "Sprints", Description: "TODO", Icon: "running", Route: "/sprint"}
	s, err := as.Services.Sprint.GetByOwner(ctx, nil, user, nil, logger)
	if err != nil {
		return nil, err
	}
	for _, x := range s {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "running", Route: "/sprint/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(s))
	return ret, nil
}

func estimateMenu(ctx context.Context, user uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "estimates", Title: "Estimates", Description: "TODO", Icon: "ruler-horizontal", Route: "/estimate"}
	e, err := as.Services.Estimate.GetByOwner(ctx, nil, user, nil, logger)
	if err != nil {
		return nil, err
	}
	for _, x := range e {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "ruler-horizontal", Route: "/estimate/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(e))
	return ret, nil
}

func standupMenu(ctx context.Context, user uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "standups", Title: "Standups", Description: "TODO", Icon: "shoe-prints", Route: "/standup"}
	u, err := as.Services.Standup.GetByOwner(ctx, nil, user, nil, logger)
	if err != nil {
		return nil, err
	}
	for _, x := range u {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "shoe-prints", Route: "/standup/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(u))
	return ret, nil
}

func retroMenu(ctx context.Context, user uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "retros", Title: "Retros", Description: "TODO", Icon: "glasses", Route: "/retro"}
	r, err := as.Services.Retro.GetByOwner(ctx, nil, user, nil, logger)
	if err != nil {
		return nil, err
	}
	for _, x := range r {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "glasses", Route: "/retro/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(r))
	return ret, nil
}
