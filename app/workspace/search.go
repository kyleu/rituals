package workspace

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/search/result"
	"github.com/kyleu/rituals/app/lib/user"
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
	ret := make(result.Results, 0, len(res)*len(res))
	for _, x := range res {
		ret = append(ret, x...)
	}
	return ret, nil
}

func (s *Service) SearchTeams(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	var ret result.Results
	for _, x := range w.Teams {
		res := result.NewResult(util.KeyTeam, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyTeam, x, x, q)
		if len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}

func (s *Service) SearchSprints(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	var ret result.Results
	for _, x := range w.Sprints {
		res := result.NewResult(util.KeySprint, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeySprint, x, x, q)
		if len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}

func (s *Service) SearchEstimates(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	var ret result.Results
	for _, x := range w.Estimates {
		res := result.NewResult(util.KeyEstimate, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyEstimate, x, x, q)
		if len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}

func (s *Service) SearchStories(ctx context.Context, q string, prms filter.ParamSet, w *Workspace, logger util.Logger) (result.Results, error) {
	var ret result.Results
	curr, err := s.st.GetByEstimateIDs(ctx, nil, prms.Get("story", nil, logger), logger, w.Estimates.IDs()...)
	if err != nil {
		return nil, errors.Wrap(err, "can't load stories")
	}
	for _, x := range curr {
		res := result.NewResult(util.KeyStory, x.ID.String(), x.PublicWebPath(""), x.TitleString(), util.KeyStory, x, x, q)
		if len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}

func (s *Service) SearchStandups(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	var ret result.Results
	for _, x := range w.Standups {
		res := result.NewResult(util.KeyStandup, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyStandup, x, x, q)
		if len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}

func (s *Service) SearchRetros(_ context.Context, q string, _ filter.ParamSet, w *Workspace, _ util.Logger) (result.Results, error) {
	var ret result.Results
	for _, x := range w.Retros {
		res := result.NewResult(util.KeyRetro, x.ID.String(), x.PublicWebPath(), x.TitleString(), util.KeyRetro, x, x, q)
		if len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}
