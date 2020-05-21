package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/gql"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func GraphiQL(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return "", err
		}

		ctx.Breadcrumbs = adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		ctx.Title = "GraphiQL"
		return tmpl(templates.GraphiQL(ctx, w))
	})
}

func GraphQLVoyagerQuery(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return "", err
		}

		bc := adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.KeyVoyager, "query")), util.KeyVoyager)...)
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return tmpl(templates.GraphQLVoyager(gql.QueryName, ctx, w))
	})
}

func GraphQLVoyagerMutation(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return "", err
		}

		bc := adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(util.KeyVoyager, "mutation")), util.KeyVoyager)...)
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return tmpl(templates.GraphQLVoyager(gql.MutationName, ctx, w))
	})
}
