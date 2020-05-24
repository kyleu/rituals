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
			Type: graphql.String,
		},
	}

	userResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		idString, err := paramString(p, util.KeyID)
		if err != nil {
			return nil, err
		}
		id := util.GetUUIDFromString(idString)
		if id == nil {
			return nil, errors.WithStack(errors.New("invalid user id [" + idString + "]"))
		}
		ret, err := ctx.App.User.GetByID(*id, false)
		if err != nil {
			return nil, errors.WithStack(errors.Wrap(err, "error retrieving user with id ["+idString+"]"))
		}
		if ret == nil {
			return nil, nil
		}
		return ret.ToProfile(), nil
	}

	usersResolver = func(params graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		curr, err := ctx.App.User.List(paramSetFromGraphQLParams(util.KeyUser, params, ctx.Logger))
		ret := make([]util.Profile, len(curr))
		if err != nil {
			return nil, errors.WithStack(errors.Wrap(err, "error retrieving users"))
		}
		for i, u := range curr {
			ret[i] = u.ToProfile().ToProfile()
		}
		return ret, nil
	}
}
