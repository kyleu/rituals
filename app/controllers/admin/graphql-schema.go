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
			return eresp(err, "")
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
			return eresp(err, "")
		}

		bc := adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		link := util.AdminLink(util.KeyVoyager, "query")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link), util.KeyVoyager)...)
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return tmpl(templates.GraphQLVoyager(gql.QueryName, ctx, w))
	})
}

func GraphQLVoyagerMutation(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return eresp(err, "")
		}

		bc := adminBC(ctx, util.KeyGraphiQL, util.KeyGraphQL)
		link := util.AdminLink(util.KeyVoyager, "mutation")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link), util.KeyVoyager)...)
		ctx.Breadcrumbs = bc
		ctx.Title = "GraphQL Voyager"
		return tmpl(templates.GraphQLVoyager(gql.MutationName, ctx, w))
	})
}
