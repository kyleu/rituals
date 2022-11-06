package gql

import (
	_ "embed"

	"github.com/kyleu/rituals/app/util"
)

import (
	"github.com/kyleu/rituals/app/lib/graphql"
)

//go:embed schema.graphql
var schemaString string

type Schema struct {
	svc *graphql.Service
	sch string
}

func NewSchema(svc *graphql.Service) *Schema {
	ret := &Schema{svc: svc, sch: schemaString}
	err := ret.svc.RegisterStringSchema(util.AppKey, util.AppName, ret.sch, ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (s *Schema) Hello() string {
	return "Howdy!"
}