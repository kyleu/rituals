package admin

import (
	"encoding/json"
	"fmt"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"io"
	"io/ioutil"
	"net/http"

	"logur.dev/logur"

	"emperror.dev/errors"
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/gql"
)

var graphQLService *gql.Service

func GraphQLRun(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return npncontroller.EResp(err)
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

		op := npncore.MapGetString(req, "operationName", ctx.Logger)
		query := npncore.MapGetString(req, "query", ctx.Logger)
		v := mapGetMap(req, "variables", ctx.Logger)

		res, err := graphQLService.Run(op, query, v, ctx)
		if err != nil {
			e := gql.ErrorResponseJSON(graphQLService.Logger, errors.Wrap(err, "error running GraphQL"))
			return graphQLResponse(w, e, ctx.Logger)
		}

		return graphQLResponse(w, res, ctx.Logger)
	})
}

func graphQLResponse(w http.ResponseWriter, res *graphql.Result, logger logur.Logger) (string, error) {
	return npncontroller.RespondJSON(w, "", res, logger)
}

func prepareService(app npnweb.AppInfo) error {
	if graphQLService == nil {
		s, err := gql.NewService(app)
		if err != nil {
			return errors.Wrap(err, "unable to initialize GraphQL schema")
		}
		graphQLService = s
	}
	return nil
}

func mapGetMap(m map[string]interface{}, key string, logger logur.Logger) map[string]interface{} {
	retEntry := npncore.GetEntry(m, key, logger)
	ret, ok := retEntry.(map[string]interface{})
	if !ok {
		logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not map[string]interface{}", key, retEntry))
		return nil
	}
	return ret
}
