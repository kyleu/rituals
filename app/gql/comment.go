package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app/comment"
)

var (
	commentUserResolver npngraphql.Callback
	commentType         *graphql.Object
)

func initComment() {
	commentUserResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return ctx.App.User().GetByID(p.Source.(*comment.Comment).UserID, false), nil
	}

	commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: npncore.Title(npncore.KeyComment),
			Fields: graphql.Fields{
				npncore.KeyID: &graphql.Field{
					Type: graphql.NewNonNull(npngraphql.ScalarUUID),
				},
				npncore.KeySvc: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.WithID(npncore.KeyModel): &graphql.Field{
					Type: graphql.NewNonNull(npngraphql.ScalarUUID),
				},
				npncore.WithID(npncore.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(npngraphql.ScalarUUID),
				},
				npncore.KeyAct: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyContent: &graphql.Field{
					Type: graphql.String,
				},
				npncore.KeyHTML: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
