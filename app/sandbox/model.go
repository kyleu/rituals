package sandbox

import "github.com/kyleu/rituals.dev/app/web"

type Resolver func(ctx *web.RequestContext) (string, interface{}, error)

type Sandbox struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Resolve     Resolver `json:"-"`
}

type Sandboxes = []*Sandbox

var AllSandboxes = Sandboxes{&Testbed, &Nightly, &Error}

func FromString(s string) *Sandbox {
	for _, t := range AllSandboxes {
		if t.Key == s {
			return t
		}
	}
	return nil
}
