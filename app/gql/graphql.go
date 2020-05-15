package gql

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/kyleu/rituals.dev/app/config"
	"log"

	"github.com/graphql-go/graphql"
)

type Service struct {
	cfg    graphql.SchemaConfig
	schema graphql.Schema
	app    *config.AppInfo
}

func NewService(app *config.AppInfo) (*Service, error) {
	// Schema
	schemaConfig := graphql.SchemaConfig{
		Query:        graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: queryFields}),
		Mutation:     graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: mutationFields}),
		Subscription: nil,
		Types:        nil,
		Directives:   nil,
		Extensions:   nil,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	svc := Service{app: app, cfg: schemaConfig, schema: schema}
	return &svc, nil
}

func (s *Service) Run(operationName string, doc string, variables map[string]interface{}) (*graphql.Result, error) {
	params := graphql.Params{
		Schema:         s.schema,
		RequestString:  doc,
		VariableValues: variables,
		OperationName:  operationName,
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		errString := ""
		for i, e := range r.Errors {
			errString += fmt.Sprintf("%s: %s", i, e.Error())
			if i < len(r.Errors) - 1 {
				errString += ", "
			}
		}
		return nil, errors.WithStack(errors.New("graphql errors [" + errString + "]"))
	}
	return r, nil
}
