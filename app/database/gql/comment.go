package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	commentUserResolver Callback
	commentType         *graphql.Object
)

func initComment() {
	commentUserResolver = func(p graphql.ResolveParams,  ctx *web.RequestContext) (interface{}, error) {
		return ctx.App.User.GetByID(p.Source.(*comment.Comment).UserID, false), nil
	}

	commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.Title(util.KeyComment),
			Fields: graphql.Fields{
				util.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				util.KeySvc: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.WithID(util.KeyModel): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				util.WithID(util.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(scalarUUID),
				},
				util.KeyAct: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyContent: &graphql.Field{
					Type: graphql.String,
				},
				util.KeyHTML: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
