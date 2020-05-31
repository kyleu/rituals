package admin

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web/act"
	"logur.dev/logur"

	"emperror.dev/errors"
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/database/gql"
	"github.com/kyleu/rituals.dev/app/web"
)

var graphQLService *gql.Service

func GraphQLRun(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return eresp(err, "")
		}
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			e := gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "cannot read JSON body for GraphQL"))
			return graphQLResponse(w, e, ctx.Logger)
		}
		err = r.Body.Close()
		if err != nil {
			e := gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "cannot close body for GraphQL"))
			return graphQLResponse(w, e, ctx.Logger)
		}

		var req map[string]interface{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			e := gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "error decoding JSON body for GraphQL"))
			return graphQLResponse(w, e, ctx.Logger)
		}

		op := util.MapGetString(req, "operationName", ctx.Logger)
		query := util.MapGetString(req, "query", ctx.Logger)
		v := util.MapGetMap(req, "variables", ctx.Logger)

		res, err := graphQLService.Run(op, query, v, ctx)
		if err != nil {
			e := gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "error running GraphQL"))
			return graphQLResponse(w, e, ctx.Logger)
		}

		return graphQLResponse(w, res, ctx.Logger)
	})
}

func graphQLResponse(w http.ResponseWriter, res *graphql.Result, logger logur.Logger) (string, error) {
	return act.RespondJSON(w, res, logger)
}

func prepareService(app *config.AppInfo) error {
	if graphQLService == nil {
		s, err := gql.NewService(app)
		if err != nil {
			return errors.Wrap(err, "unable to initialize GraphQL schema")
		}
		graphQLService = s
	}
	return nil
}
