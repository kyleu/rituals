package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
)

var (
	userResolver  Callback
	usersResolver Callback
)

func initUser() {
	userResolver = func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		id := npncore.MapGetUUID(p.Args, npncore.KeyID, ctx.Logger)
		ret := ctx.App.User().GetByID(*id, false)
		if ret == nil {
			return nil, nil
		}
		return ret.ToProfile().ToProfile(), nil
	}

	usersResolver = func(params graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error) {
		curr := ctx.App.User().List(paramSetFromGraphQLParams(npncore.KeyUser, params, ctx.Logger))
		ret := make([]npnuser.Profile, 0, len(curr))
		for _, u := range curr {
			ret = append(ret, u.ToProfile().ToProfile())
		}
		return ret, nil
	}
}
