package sandbox

import (
	"github.com/kyleu/rituals.dev/app/web"
)

type Sandbox struct {
	Key         string                                            `json:"key"`
	Title       string                                            `json:"title"`
	Description string                                            `json:"description,omitempty"`
	Resolve     func(ctx web.RequestContext) (interface{}, error) `json:"-"`
}

var Gallery = Sandbox{
	Key:         "gallery",
	Title:       "Gallery",
	Description: "An HTML demo showing available components",
	Resolve: func(ctx web.RequestContext) (interface{}, error) {
		return nil, nil
	},
}

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx web.RequestContext) (interface{}, error) {
		return "Testbed!", nil
	},
}

var AllSandboxes = []*Sandbox{&Gallery, &Testbed}

func FromString(s string) *Sandbox {
	for _, t := range AllSandboxes {
		if t.Key == s {
			return t
		}
	}
	return nil
}
