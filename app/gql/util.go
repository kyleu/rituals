package gql

import (
	"fmt"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/web"
	"logur.dev/logur"
)

var graphQLInitialized = false

type Callback func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error)

func paramSetFromGraphQLParams(key string, params graphql.ResolveParams, logger logur.Logger) *query.Params {
	orderings := make(query.Orderings, 0)
	o, ok := params.Args["orders"]
	if ok {
		for _, x := range o.([]interface{}) {
			m := x.(map[string]interface{})
			col := m["col"].(string)
			var defaultOrdering = query.Orderings{{Column: col, Asc: m["asc"].(bool)}}
			orderings = append(orderings, defaultOrdering...)
		}
	}

	limit := 0
	l, ok := params.Args["limit"]
	if ok {
		limit = l.(int)
	}

	offset := 0
	x, ok := params.Args["offset"]
	if ok {
		offset = x.(int)
	}

	ret := &query.Params{Key: key, Orderings: orderings, Limit: limit, Offset: offset}
	return ret.Filtered(logger)
}

func paramKeyString(p graphql.ResolveParams) (string, error) {
	arg, ok := p.Args[util.KeyKey]
	if !ok {
		return "", errors.New(fmt.Sprintf("parameter %v is not present", util.KeyKey))
	}
	ret, ok := arg.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("parameter %v is not a string: %v", util.KeyKey, arg))
	}
	return ret, nil
}

func paramUUID(p graphql.ResolveParams, key string) (*uuid.UUID, error) {
	arg, ok := p.Args[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("parameter %v is not present", key))
	}
	ret, ok := arg.(uuid.UUID)
	if !ok {
		return nil, errors.New(fmt.Sprintf("parameter %v is not a UUID: %v", key, arg))
	}
	return &ret, nil
}

func ErrorResponseJSON(logger logur.Logger, errors ...error) *graphql.Result {
	var errs = make([]gqlerrors.FormattedError, len(errors))

	for i, err := range errors {
		logger.Warn(fmt.Sprintf("error running GraphQL: %+v", err))

		errs[i] = gqlerrors.FormattedError{Message: err.Error()}
	}

	return &graphql.Result{
		Errors: errs,
	}
}
