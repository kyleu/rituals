package cmenu

import (
	"context"
	"fmt"
	"github.com/kyleu/rituals/app/lib/filter"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/app/lib/telemetry"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
)

var workspaceServices = []enum.ModelService{
	enum.ModelServiceTeam, enum.ModelServiceSprint, enum.ModelServiceEstimate, enum.ModelServiceStandup, enum.ModelServiceRetro,
}

func workspaceMenu(ctx context.Context, as *app.State, params filter.ParamSet, profile *user.Profile, logger util.Logger) (menu.Items, any, error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "workspace.menu", logger)
	defer span.Complete()
	u := profile.ID
	w := &workspace.Workspace{}
	items, errs := util.AsyncCollectMap(workspaceServices, func(item enum.ModelService) enum.ModelService {
		return item
	}, func(k enum.ModelService) (*menu.Item, error) {
		switch k {
		case enum.ModelServiceTeam:
			ret, models, err := teamMenu(ctx, u, params, as, logger)
			if err != nil {
				return nil, err
			}
			w.Teams = models
			return ret, nil
		case enum.ModelServiceSprint:
			ret, models, err := sprintMenu(ctx, u, params, as, logger)
			if err != nil {
				return nil, err
			}
			w.Sprints = models
			return ret, nil
		case enum.ModelServiceEstimate:
			ret, models, err := estimateMenu(ctx, u, params, as, logger)
			if err != nil {
				return nil, err
			}
			w.Estimates = models
			return ret, nil
		case enum.ModelServiceStandup:
			ret, models, err := standupMenu(ctx, u, params, as, logger)
			if err != nil {
				return nil, err
			}
			w.Standups = models
			return ret, nil
		case enum.ModelServiceRetro:
			ret, models, err := retroMenu(ctx, u, params, as, logger)
			if err != nil {
				return nil, err
			}
			w.Retros = models
			return ret, nil
		default:
			return nil, errors.Errorf("invalid service [%s]", k)
		}
	})
	if len(errs) > 0 {
		return nil, nil, util.ErrorMerge(maps.Values(errs)...)
	}
	ret := menu.Items{
		items[enum.ModelServiceTeam],
		items[enum.ModelServiceSprint],
		items[enum.ModelServiceEstimate],
		items[enum.ModelServiceStandup],
		items[enum.ModelServiceRetro],
	}
	return ret, w, nil
}

func teamMenu(ctx context.Context, user uuid.UUID, params filter.ParamSet, as *app.State, logger util.Logger) (*menu.Item, team.Teams, error) {
	ret := &menu.Item{Key: "teams", Title: "Teams", Description: "TODO", Icon: "team", Route: "/team"}
	t, err := as.Services.Team.GetByMember(ctx, nil, user, params.Get("team", nil, logger), logger)
	if err != nil {
		return nil, nil, err
	}
	for _, x := range t {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "team", Route: "/team/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(t))
	return ret, t, nil
}

func sprintMenu(ctx context.Context, user uuid.UUID, params filter.ParamSet, as *app.State, logger util.Logger) (*menu.Item, sprint.Sprints, error) {
	ret := &menu.Item{Key: "sprints", Title: "Sprints", Description: "TODO", Icon: "sprint", Route: "/sprint"}
	s, err := as.Services.Sprint.GetByMember(ctx, nil, user, params.Get("sprint", nil, logger), logger)
	if err != nil {
		return nil, nil, err
	}
	for _, x := range s {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "sprint", Route: "/sprint/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(s))
	return ret, s, nil
}

func estimateMenu(ctx context.Context, user uuid.UUID, params filter.ParamSet, as *app.State, logger util.Logger) (*menu.Item, estimate.Estimates, error) {
	ret := &menu.Item{Key: "estimates", Title: "Estimates", Description: "TODO", Icon: "estimate", Route: "/estimate"}
	e, err := as.Services.Estimate.GetByMember(ctx, nil, user, params.Get("estimate", nil, logger), logger)
	if err != nil {
		return nil, nil, err
	}
	for _, x := range e {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "estimate", Route: "/estimate/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(e))
	return ret, e, nil
}

func standupMenu(ctx context.Context, user uuid.UUID, params filter.ParamSet, as *app.State, logger util.Logger) (*menu.Item, standup.Standups, error) {
	ret := &menu.Item{Key: "standups", Title: "Standups", Description: "TODO", Icon: "standup", Route: "/standup"}
	u, err := as.Services.Standup.GetByMember(ctx, nil, user, params.Get("standup", nil, logger), logger)
	if err != nil {
		return nil, nil, err
	}
	for _, x := range u {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "standup", Route: "/standup/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(u))
	return ret, u, nil
}

func retroMenu(ctx context.Context, user uuid.UUID, params filter.ParamSet, as *app.State, logger util.Logger) (*menu.Item, retro.Retros, error) {
	ret := &menu.Item{Key: "retros", Title: "Retros", Description: "TODO", Icon: "retro", Route: "/retro"}
	r, err := as.Services.Retro.GetByMember(ctx, nil, user, params.Get("retro", nil, logger), logger)
	if err != nil {
		return nil, nil, err
	}
	for _, x := range r {
		kid := &menu.Item{Key: x.ID.String(), Title: x.TitleString(), Description: "", Icon: "retro", Route: "/retro/" + x.Slug}
		ret.Children = append(ret.Children, kid)
	}
	ret.Badge = fmt.Sprint(len(r))
	return ret, r, nil
}
