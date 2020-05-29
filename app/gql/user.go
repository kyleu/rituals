package gql

import (
	"emperror.dev/errors"
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	userResolver  Callback
	usersResolver Callback
)

func initUser() {
	userResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		id := util.MapGetUUID(p.Args, util.KeyID, ctx.Logger)
		ret, err := ctx.App.User.GetByID(*id, false)
		if err != nil {
			return nil, errors.Wrap(err, "error retrieving user with id ["+id.String()+"]")
		}
		if ret == nil {
			return nil, nil
		}
		return ret.ToProfile().ToProfile(), nil
	}

	usersResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		curr := ctx.App.User.List(paramSetFromGraphQLParams(util.KeyUser, params, ctx.Logger))
		ret := make([]util.Profile, 0, len(curr))
		for _, u := range curr {
			ret = append(ret, u.ToProfile().ToProfile())
		}
		return ret, nil
	}
}
