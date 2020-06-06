package sandbox

import (
	"github.com/kyleu/rituals.dev/app/web"
)

type Sandbox struct {
	Key         string                                            `json:"key"`
	Title       string                                            `json:"title"`
	Description string                                            `json:"description,omitempty"`
	Resolve     func(ctx *web.RequestContext) (interface{}, error) `json:"-"`
}

type Sandboxes = []*Sandbox

var AllSandboxes = Sandboxes{&Testbed, &Error}

func FromString(s string) *Sandbox {
	for _, t := range AllSandboxes {
		if t.Key == s {
			return t
		}
	}
	return nil
}
