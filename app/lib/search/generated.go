// Package search - Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"

	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/lib/search/result"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
)

func generatedSearch() []Provider {
	estimateFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("estimate", logger).WithLimit(5)
		models, err := as.Services.Estimate.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *estimate.Estimate, _ int) *result.Result {
			return result.NewResult("Estimate", m.String(), m.WebPath(), m.TitleString(), "estimate", m, m, params.Q)
		}), nil
	}
	retroFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("retro", logger).WithLimit(5)
		models, err := as.Services.Retro.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *retro.Retro, _ int) *result.Result {
			return result.NewResult("Retro", m.String(), m.WebPath(), m.TitleString(), "retro", m, m, params.Q)
		}), nil
	}
	sprintFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("sprint", logger).WithLimit(5)
		models, err := as.Services.Sprint.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *sprint.Sprint, _ int) *result.Result {
			return result.NewResult("Sprint", m.String(), m.WebPath(), m.TitleString(), "sprint", m, m, params.Q)
		}), nil
	}
	standupFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("standup", logger).WithLimit(5)
		models, err := as.Services.Standup.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *standup.Standup, _ int) *result.Result {
			return result.NewResult("Standup", m.String(), m.WebPath(), m.TitleString(), "standup", m, m, params.Q)
		}), nil
	}
	storyFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("story", logger).WithLimit(5)
		models, err := as.Services.Story.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *story.Story, _ int) *result.Result {
			return result.NewResult("Story", m.String(), m.WebPath(), m.TitleString(), "story", m, m, params.Q)
		}), nil
	}
	teamFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("team", logger).WithLimit(5)
		models, err := as.Services.Team.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *team.Team, _ int) *result.Result {
			return result.NewResult("Team", m.String(), m.WebPath(), m.TitleString(), "team", m, m, params.Q)
		}), nil
	}
	userFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("user", logger).WithLimit(5)
		models, err := as.Services.User.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *user.User, _ int) *result.Result {
			return result.NewResult("User", m.String(), m.WebPath(), m.TitleString(), "profile", m, m, params.Q)
		}), nil
	}
	return []Provider{estimateFunc, retroFunc, sprintFunc, standupFunc, storyFunc, teamFunc, userFunc}
}
