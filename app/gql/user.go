package gql

import (
	"emperror.dev/errors"
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	userArgs      graphql.FieldConfigArgument
	userResolver  Callback
	usersResolver Callback
)

func initUser() {
	userArgs = graphql.FieldConfigArgument{
		util.KeyID: &graphql.ArgumentConfig{
			Type: scalarUUID,
		},
	}

	userResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		id, err := paramUUID(p, util.KeyID)
		if err != nil {
			return nil, err
		}
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
		curr, err := ctx.App.User.List(paramSetFromGraphQLParams(util.KeyUser, params, ctx.Logger))
		ret := make([]util.Profile, len(curr))
		if err != nil {
			return nil, errors.Wrap(err, "error retrieving users")
		}
		for i, u := range curr {
			ret[i] = u.ToProfile().ToProfile()
		}
		return ret, nil
	}
}
