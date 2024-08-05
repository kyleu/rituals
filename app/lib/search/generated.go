package search

import (
	"context"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/search/result"
	"github.com/kyleu/rituals/app/util"
)

func generatedSearch() []Provider {
	estimateFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("estimate", logger).WithLimit(5)
		return as.Services.Estimate.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	retroFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("retro", logger).WithLimit(5)
		return as.Services.Retro.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	sprintFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("sprint", logger).WithLimit(5)
		return as.Services.Sprint.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	standupFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("standup", logger).WithLimit(5)
		return as.Services.Standup.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	storyFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("story", logger).WithLimit(5)
		return as.Services.Story.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	teamFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("team", logger).WithLimit(5)
		return as.Services.Team.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	userFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("user", logger).WithLimit(5)
		return as.Services.User.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	return []Provider{estimateFunc, retroFunc, sprintFunc, standupFunc, storyFunc, teamFunc, userFunc}
}
