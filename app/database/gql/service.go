package gql

import (
	"context"
	"fmt"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/web"
	"logur.dev/logur"

	"github.com/graphql-go/graphql"
)

type Service struct {
	Logger logur.Logger
	cfg    graphql.SchemaConfig
	schema graphql.Schema
	app    *config.AppInfo
}

func NewService(app *config.AppInfo) (*Service, error) {
	logger := logur.WithFields(app.Logger, map[string]interface{}{util.KeyService: util.KeyGraphQL})

	initSchema()

	// Schema
	schemaConfig := graphql.SchemaConfig{
		Query:        graphql.NewObject(graphql.ObjectConfig{Name: QueryName, Fields: queryFields()}),
		Mutation:     graphql.NewObject(graphql.ObjectConfig{Name: MutationName, Fields: mutationFields()}),
		Subscription: nil,
		Types:        nil,
		Directives:   nil,
		Extensions:   nil,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new schema")
	}
	svc := Service{Logger: logger, cfg: schemaConfig, schema: schema, app: app}

	logger.Debug("initialized GraphQL service")
	return &svc, nil
}

func (s *Service) Run(operationName string, doc string, variables map[string]interface{},  ctx *web.RequestContext) (*graphql.Result, error) {
	params := graphql.Params{
		Schema:         s.schema,
		RequestString:  doc,
		VariableValues: variables,
		OperationName:  operationName,
		Context:        context.WithValue(context.Background(), util.ContextKey, ctx),
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		errString := ""
		for i, e := range r.Errors {
			errString += fmt.Sprintf("%v: %v@%v", i, e.Message, e.Path)
			if i < len(r.Errors)-1 {
				errString += ", "
			}
		}
		return nil, errors.New("graphql errors [" + errString + "]")
	}
	return r, nil
}

func ctxF(f func(p graphql.ResolveParams,  ctx *web.RequestContext) (interface{}, error)) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		c, ok := p.Context.Value(util.ContextKey).(*web.RequestContext)
		if !ok {
			return nil, errors.New("no ctx in GraphQL resolve params")
		}
		return f(p, c)
	}
}

func initSchema() {
	if !graphQLInitialized {
		graphQLInitialized = true

		initMember()
		initComment()
		initHistory()

		initEstimate()
		initStandup()
		initRetro()

		initSprint()
		initTeam()

		initPermission()
		initAction()

		initProfile()
		initUser()

		initSandbox()
	}
}
