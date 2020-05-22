package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	memberProfileResolver Callback
	memberType            *graphql.Object
)

func initMember() {
	memberRoleType := graphql.NewEnum(graphql.EnumConfig{
		Name: "MemberRole",
		Values: graphql.EnumValueConfigMap{
			"owner":    &graphql.EnumValueConfig{Value: "owner"},
			"member":   &graphql.EnumValueConfig{Value: "member"},
			"observer": &graphql.EnumValueConfig{Value: "observer"},
		},
	})

	memberProfileResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.User.GetByID(p.Source.(*member.Entry).UserID, false)
	}

	memberType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Member",
			Fields: graphql.Fields{
				"userID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"role": &graphql.Field{
					Type: graphql.NewNonNull(memberRoleType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*member.Entry).Role.Key, nil
					},
				},
				util.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
