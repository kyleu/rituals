package sandbox

import (
	"github.com/kyleu/rituals.dev/app/web"
)

var Gallery = Sandbox{
	Key:         "gallery",
	Title:       "Gallery",
	Description: "An HTML demo showing available components",
	Resolve: func(ctx web.RequestContext) (interface{}, error) {
		return nil, nil
	},
}
