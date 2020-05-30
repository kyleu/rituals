package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	storyResolver   Callback
	storiesResolver Callback
	storyType       *graphql.Object
)

func initStory() {
	storyResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		id := util.MapGetUUID(p.Args, util.KeyID, ctx.Logger)
		return ctx.App.Estimate.GetStoryByID(*id)
	}

	storiesResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.GetStories(p.Source.(*estimate.Session).ID, paramSetFromGraphQLParams(util.KeyStory, p, ctx.Logger)), nil
	}

	storyType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.Title(util.KeyStory),
			Fields: graphql.Fields{
				util.WithID(util.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
