package controllers

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/gql"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

var svc *gql.Service

func GraphQLHome(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return "", err
		}

		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("graphql"), "graphql")...)
		ctx.Breadcrumbs = bc

		ctx.Title = "GraphiQL"
		return tmpl(templates.GraphiQL(ctx, w))
	})
}

func GraphQLRun(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		err := prepareService(ctx.App)
		if err != nil {
			return "", err
		}
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			return graphQLResponse(w, errorResponseJSON(errors.WithStack(errors.Wrap(err, "cannot read JSON body for GraphQL"))))
		}
		err = r.Body.Close()
		if err != nil {
			return graphQLResponse(w, errorResponseJSON(errors.WithStack(errors.Wrap(err, "cannot close body for GraphQL"))))
		}

		var req map[string]interface{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			return graphQLResponse(w, errorResponseJSON(errors.WithStack(errors.Wrap(err, "error decoding JSON body for GraphQL"))))
		}
		op := ""
		opParam, ok := req["operationName"]
		if ok {
			op = opParam.(string)
		}
		query := ""
		queryParam, ok := req["query"]
		if ok {
			query = queryParam.(string)
		}

		var v map[string]interface{}
		variables, ok := req["variables"]
		if ok {
			v = variables.(map[string]interface{})
		}

		res, err := svc.Run(op, query, v)
		if err != nil {
			return graphQLResponse(w, errorResponseJSON(errors.WithStack(errors.Wrap(err, "error running GraphQL"))))
		}

		return graphQLResponse(w, res)
	})
}

func graphQLResponse(w http.ResponseWriter, res *graphql.Result) (string, error) {
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", errors.WithStack(errors.Wrap(err, "error encoding GraphQL results"))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(b)
	if err != nil {
		return "", errors.WithStack(errors.Wrap(err, "error writing GraphQL response"))
	}
	return "", nil
}

func prepareService(app *config.AppInfo) error {
	if svc == nil {
		s, err := gql.NewService(app)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "unable to initialize GraphQL schema"))
		}
		svc = s
	}
	return nil
}

func errorResponseJSON(errors ...error) *graphql.Result {
	var errs []gqlerrors.FormattedError
	for _, err := range errors {
		errs = append(errs, gqlerrors.FormattedError{Message: err.Error()})
	}
	return &graphql.Result{
		Errors:     errs,
	}
}
