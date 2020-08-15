package gql

import (
	"github.com/kyleu/npn/npngraphql"
	"github.com/kyleu/npn/npnweb"
)

func InitService(app npnweb.AppInfo) error {
	initSchema()
	_, err := npngraphql.NewService(app, QueryName, QueryFields(), MutationName, MutationFields())
	return err
}

var graphQLInitialized = false

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
