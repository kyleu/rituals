// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/lib/search/result"
	"github.com/kyleu/rituals/app/lib/telemetry"
	"github.com/kyleu/rituals/app/util"
)

type Provider func(context.Context, *app.State, *Params, util.Logger) (result.Results, error)

func Search(ctx context.Context, as *app.State, params *Params, logger util.Logger) (result.Results, []error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "search", logger)
	defer span.Complete()

	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	testFunc := func(ctx context.Context, as *app.State, p *Params, logger util.Logger) (result.Results, error) {
		return result.Results{{URL: "/search?q=test", Title: "Test Result", Icon: "star", Matches: nil}}, nil
	}
	allProviders = append(allProviders, testFunc)
	// $PF_SECTION_END(search_functions)$

	allProviders = append(allProviders, generatedSearch()...)
	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	params.Q = strings.TrimSpace(params.Q)

	results, errs := util.AsyncCollect(allProviders, func(item Provider) (result.Results, error) {
		return item(ctx, as, params, logger)
	})

	ret := make(result.Results, 0, len(results)*len(results))
	for _, x := range results {
		ret = append(ret, x...)
	}

	ret.Sort()
	return ret, errs
}
