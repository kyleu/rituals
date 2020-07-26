package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	memberProfileResolver Callback
	memberRoleType        *graphql.Enum
	memberType            *graphql.Object
)

func initMember() {
	memberRoleType = graphql.NewEnum(graphql.EnumConfig{
		Name: "MemberRole",
		Values: graphql.EnumValueConfigMap{
			util.KeyOwner:  &graphql.EnumValueConfig{Value: util.KeyOwner},
			util.KeyMember: &graphql.EnumValueConfig{Value: util.KeyMember},
			"observer":     &graphql.EnumValueConfig{Value: "observer"},
		},
	})

	memberProfileResolver = func(p graphql.ResolveParams, ctx *web.RequestContext) (interface{}, error) {
		return ctx.App.User.GetByID(p.Source.(*member.Entry).UserID, false), nil
	}

	memberType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: util.Title(util.KeyMember),
			Fields: graphql.Fields{
				util.WithID(util.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyName: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.KeyRole: &graphql.Field{
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
