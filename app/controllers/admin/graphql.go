package admin

import (
	"encoding/json"
	"github.com/kyleu/rituals.dev/app/util"
	"io"
	"io/ioutil"
	"net/http"

	"emperror.dev/errors"
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/gql"
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
			return graphQLResponse(w, gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "cannot read JSON body for GraphQL")))
		}
		err = r.Body.Close()
		if err != nil {
			return graphQLResponse(w, gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "cannot close body for GraphQL")))
		}

		var req map[string]interface{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			return graphQLResponse(w, gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "error decoding JSON body for GraphQL")))
		}

		op := util.MapGetString(req, "operationName", ctx.Logger)
		query := util.MapGetString(req, "query", ctx.Logger)
		v := util.MapGetMap(req, "variables", ctx.Logger)

		res, err := graphQLService.Run(op, query, v, ctx)
		if err != nil {
			return graphQLResponse(w, gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "error running GraphQL")))
		}

		return graphQLResponse(w, res)
	})
}

func graphQLResponse(w http.ResponseWriter, res *graphql.Result) (string, error) {
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return eresp(err, "error encoding GraphQL results")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(b)
	if err != nil {
		return eresp(err, "error writing GraphQL response")
	}
	return "", nil
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
