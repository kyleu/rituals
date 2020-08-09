package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app/member"
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
			npncore.KeyOwner:  &graphql.EnumValueConfig{Value: npncore.KeyOwner},
			npncore.KeyMember: &graphql.EnumValueConfig{Value: npncore.KeyMember},
			"observer":     &graphql.EnumValueConfig{Value: "observer"},
		},
	})

	memberProfileResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		return ctx.App.User().GetByID(p.Source.(*member.Entry).UserID, false), nil
	}

	memberType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: npncore.Title(npncore.KeyMember),
			Fields: graphql.Fields{
				npncore.WithID(npncore.KeyUser): &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyName: &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				npncore.KeyRole: &graphql.Field{
					Type: graphql.NewNonNull(memberRoleType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*member.Entry).Role.Key, nil
					},
				},
				npncore.KeyCreated: &graphql.Field{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
		},
	)
}
