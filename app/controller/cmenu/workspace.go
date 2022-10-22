package cmenu

import (
	"context"
	"fmt"

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

func workspaceMenu(ctx context.Context, as *app.State, logger util.Logger) (menu.Items, error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "workspace.menu", logger)
	defer span.Complete()
	user := util.UUIDFromString("90000000-0000-0000-0000-000000000000")
	items, errs := util.AsyncCollectMap(workspaceServices, func(item enum.ModelService) enum.ModelService {
		return item
	}, func(k enum.ModelService) (*menu.Item, error) {
		switch k {
		case enum.ModelServiceTeam:
			return teamMenu(ctx, user, as, logger)
		case enum.ModelServiceSprint:
			return sprintMenu(ctx, user, as, logger)
		case enum.ModelServiceEstimate:
			return estimateMenu(ctx, user, as, logger)
		case enum.ModelServiceStandup:
			return standupMenu(ctx, user, as, logger)
		case enum.ModelServiceRetro:
			return retroMenu(ctx, user, as, logger)
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

func teamMenu(ctx context.Context, user *uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "ws_team", Title: "Teams", Description: "TODO", Icon: "users", Route: "/team"}
	if user == nil {
		return ret, nil
	}
	t, err := as.Services.Team.GetByOwner(ctx, nil, *user, nil, logger)
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

func sprintMenu(ctx context.Context, user *uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "ws_sprint", Title: "Sprints", Description: "TODO", Icon: "running", Route: "/sprint"}
	if user == nil {
		return ret, nil
	}
	s, err := as.Services.Sprint.GetByOwner(ctx, nil, *user, nil, logger)
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

func estimateMenu(ctx context.Context, user *uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "ws_estimate", Title: "Estimates", Description: "TODO", Icon: "ruler-horizontal", Route: "/estimate"}
	if user == nil {
		return ret, nil
	}
	e, err := as.Services.Estimate.GetByOwner(ctx, nil, *user, nil, logger)
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

func standupMenu(ctx context.Context, user *uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "ws_standup", Title: "Standups", Description: "TODO", Icon: "shoe-prints", Route: "/standup"}
	if user == nil {
		return ret, nil
	}
	u, err := as.Services.Standup.GetByOwner(ctx, nil, *user, nil, logger)
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

func retroMenu(ctx context.Context, user *uuid.UUID, as *app.State, logger util.Logger) (*menu.Item, error) {
	ret := &menu.Item{Key: "ws_retro", Title: "Retros", Description: "TODO", Icon: "glasses", Route: "/retro"}
	if user == nil {
		return ret, nil
	}
	r, err := as.Services.Retro.GetByOwner(ctx, nil, *user, nil, logger)
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
