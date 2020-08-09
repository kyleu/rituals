package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/rituals.dev/app/gql"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func GraphiQL(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return npncontroller.EResp(err)
		}

		ctx.Breadcrumbs = adminBC(ctx, npncore.KeyGraphiQL, npncore.KeyGraphQL)
		ctx.Title = "GraphiQL"
		return npncontroller.T(templates.GraphiQL(ctx, w))
	})
}

func GraphQLVoyagerQuery(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return npncontroller.EResp(err)
		}

		bc := adminBC(ctx, npncore.KeyGraphiQL, npncore.KeyGraphQL)
		bc = append(bc, npnweb.BreadcrumbSelf("query"))
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return npncontroller.T(templates.GraphQLVoyager(gql.QueryName, ctx, w))
	})
}

func GraphQLVoyagerMutation(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return npncontroller.EResp(err)
		}

		bc := adminBC(ctx, npncore.KeyGraphiQL, npncore.KeyGraphQL)
		bc = append(bc, npnweb.BreadcrumbSelf("mutation"))
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return npncontroller.T(templates.GraphQLVoyager(gql.MutationName, ctx, w))
	})
}
