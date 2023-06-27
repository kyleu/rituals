package workspace

import (
	"context"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/search/result"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

type searchFn func(ctx context.Context, q string, prms filter.ParamSet, w *Workspace, logger util.Logger) (result.Results, error)

func (s *Service) Search(ctx context.Context, q string, params filter.ParamSet, _ *user.Profile, data any, logger util.Logger) (result.Results, error) {
	w, err := FromAny(data)
	if err != nil {
		return nil, errors.Errorf("invalid search data of type [%T]", data)
	}

	fns := []searchFn{s.SearchTeams, s.SearchSprints, s.SearchEstimates, s.SearchStandups, s.SearchRetros, s.SearchStories}
	res, errs := util.AsyncCollect(fns, func(fn searchFn) (result.Results, error) {
		return fn(ctx, q, params, w, logger)
	})
	if len(errs) > 0 {
		return nil, util.ErrorMerge(errs...)
	}
	return lo.FlatMap(res, func(x result.Results, _ int) []*result.Result {
		return x
	}), nil
}

func (s *Service) SearchTeams(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	return lo.FilterMap(w.Teams, func(x *team.Team, _ int) (*result.Result, bool) {
		res := result.NewResult(util.KeyTeam, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyTeam, x, x, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}

func (s *Service) SearchSprints(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	return lo.FilterMap(w.Sprints, func(x *sprint.Sprint, _ int) (*result.Result, bool) {
		res := result.NewResult(util.KeySprint, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeySprint, x, x, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}

func (s *Service) SearchEstimates(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	return lo.FilterMap(w.Estimates, func(x *estimate.Estimate, _ int) (*result.Result, bool) {
		res := result.NewResult(util.KeyEstimate, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyEstimate, x, x, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}

func (s *Service) SearchStories(ctx context.Context, q string, prms filter.ParamSet, w *Workspace, logger util.Logger) (result.Results, error) {
	curr, err := s.st.GetByEstimateIDs(ctx, nil, prms.Get("story", nil, logger), logger, w.Estimates.IDs()...)
	if err != nil {
		return nil, errors.Wrap(err, "can't load stories")
	}
	return lo.FilterMap(curr, func(x *story.Story, _ int) (*result.Result, bool) {
		res := result.NewResult(util.KeyStory, x.ID.String(), x.PublicWebPath(""), x.TitleString(), util.KeyStory, x, x, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}

func (s *Service) SearchStandups(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	return lo.FilterMap(w.Standups, func(x *standup.Standup, _ int) (*result.Result, bool) {
		res := result.NewResult(util.KeyStandup, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyStandup, x, x, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}

func (s *Service) SearchRetros(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	return lo.FilterMap(w.Retros, func(x *retro.Retro, _ int) (*result.Result, bool) {
		res := result.NewResult(util.KeyRetro, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyRetro, x, x, q)
		if len(res.Matches) > 0 {
			return res, true
		}
		return nil, false
	}), nil
}
