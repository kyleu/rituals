package sandbox

import (
	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/web"
)

var Error = Sandbox{
	Key:         "error",
	Title:       "Error",
	Description: "An example of the error page",
	Resolve: func(ctx web.RequestContext) (interface{}, error) {
		return nil, errors.New("here's an intentional error")
	},
}
