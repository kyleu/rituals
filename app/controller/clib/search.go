// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/search"
	"github.com/kyleu/rituals/views/vsearch"
)

const searchKey = "search"

func Search(rc *fasthttp.RequestCtx) {
	controller.Act(searchKey, rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := string(rc.URI().QueryArgs().Peek("q"))
		params := &search.Params{Q: q, PS: ps.Params}
		results, errs := search.Search(ps.Context, params, as, ps)
		ps.SetTitleAndData("Search Results", results)
		if q != "" {
			ps.Title = fmt.Sprintf("[%s] %s", q, ps.Title)
		}
		if len(results) == 1 && results[0].URL != "" {
			return controller.FlashAndRedir(true, "single search result found", results[0].URL, rc, ps)
		}
		ps.DefaultNavIcon = searchKey
		bc := []string{"Search||/search"}
		if q != "" {
			bc = append(bc, q)
		}
		return controller.Render(rc, as, &vsearch.Results{Params: params, Results: results, Errors: errs}, ps, bc...)
	})
}
