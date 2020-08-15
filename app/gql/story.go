package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

var (
	storyResolver   npngraphql.Callback
	storiesResolver npngraphql.Callback
	storyType       *graphql.Object
)

func initStory() {
	storyResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		id := npncore.MapGetUUID(p.Args, npncore.KeyID, ctx.Logger)
		return app.Estimate(ctx.App).GetStoryByID(*id)
	}

	storiesResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return app.Estimate(ctx.App).GetStories(p.Source.(*estimate.Session).ID, npngraphql.ParamSetFromGraphQLParams(util.KeyStory, p, ctx.Logger)), nil
	}

	storyType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: npncore.Title(util.KeyStory),
			Fields: graphql.Fields{
				npncore.WithID(npncore.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
