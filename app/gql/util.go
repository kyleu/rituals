package gql

import (
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"logur.dev/logur"
)

var graphQLInitialized = false

type Callback func(p graphql.ResolveParams, ctx *npnweb.RequestContext) (interface{}, error)

func paramSetFromGraphQLParams(key string, params graphql.ResolveParams, logger logur.Logger) *npncore.Params {
	orderings := make(npncore.Orderings, 0)
	o, ok := params.Args["orders"]
	if ok {
		for _, x := range o.([]interface{}) {
			m := x.(map[string]interface{})
			col := npncore.MapGetString(m, "col", logger)
			asc := npncore.MapGetBool(m, "asc", logger)
			var defaultOrdering = npncore.Orderings{{Column: col, Asc: asc}}
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

	ret := &npncore.Params{Key: key, Orderings: orderings, Limit: limit, Offset: offset}
	return ret.Filtered(logger)
}

func ErrorResponseJSON(logger logur.Logger, errors ...error) *graphql.Result {
	var errs = make([]gqlerrors.FormattedError, 0, len(errors))

	for _, err := range errors {
		logger.Warn(fmt.Sprintf("error running GraphQL: %+v", err))
		errs = append(errs, gqlerrors.FormattedError{Message: err.Error()})
	}

	return &graphql.Result{
		Errors: errs,
	}
}
