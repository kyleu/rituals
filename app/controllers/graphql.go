package controllers

import (
	"emperror.dev/errors"
	"encoding/json"
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
			return "", errors.WithStack(errors.Wrap(err, "cannot read JSON body for GraphQL"))
		}
		println(1)
		err = r.Body.Close()
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot close body for GraphQL"))
		}

		var req map[string]interface{}
		println(2)
		err = json.Unmarshal(body, &req)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error decoding JSON body for GraphQL"))
		}
		println(3)
		operationName := req["operationName"]
		query := req["query"]
		variables := req["variables"]
		println(4)

		o := operationName.(string)
		println(5)
		q := query.(string)
		println(6)
		var v map[string]interface{}
		if variables != nil {
			v = variables.(map[string]interface{})
		}
		println(7)

		r, err := svc.Run(o, q, v)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error running GraphQL"))
		}

		b, err := json.MarshalIndent(r.Data, "", "  ")
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
	})
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
