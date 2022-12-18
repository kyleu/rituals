// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/search/result"
	"github.com/kyleu/rituals/app/util"
)

//nolint:gocognit
func generatedSearch() []Provider {
	estimateFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Estimate.Search(ctx, params.Q, nil, params.PS.Get("estimate", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("estimate", m.String(), m.WebPath(), m.String(), "estimate", m, m, params.Q))
		}
		return res, nil
	}
	retroFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Retro.Search(ctx, params.Q, nil, params.PS.Get("retro", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("retro", m.String(), m.WebPath(), m.String(), "retro", m, m, params.Q))
		}
		return res, nil
	}
	sprintFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Sprint.Search(ctx, params.Q, nil, params.PS.Get("sprint", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("sprint", m.String(), m.WebPath(), m.String(), "sprint", m, m, params.Q))
		}
		return res, nil
	}
	standupFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Standup.Search(ctx, params.Q, nil, params.PS.Get("standup", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("standup", m.String(), m.WebPath(), m.String(), "standup", m, m, params.Q))
		}
		return res, nil
	}
	storyFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Story.Search(ctx, params.Q, nil, params.PS.Get("story", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("story", m.String(), m.WebPath(), m.String(), "story", m, m, params.Q))
		}
		return res, nil
	}
	teamFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Team.Search(ctx, params.Q, nil, params.PS.Get("team", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("team", m.String(), m.WebPath(), m.String(), "team", m, m, params.Q))
		}
		return res, nil
	}
	userFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.User.Search(ctx, params.Q, nil, params.PS.Get("user", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("user", m.String(), m.WebPath(), m.String(), "profile", m, m, params.Q))
		}
		return res, nil
	}
	return []Provider{estimateFunc, retroFunc, sprintFunc, standupFunc, storyFunc, teamFunc, userFunc}
}
