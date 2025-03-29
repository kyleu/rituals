package sandbox

import (
	"context"
	"fmt"

	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/app/util"
)

type runFn func(ctx context.Context, st *app.State, args util.ValueMap, logger util.Logger) (any, error)

type Sandbox struct {
	Key   string          `json:"key,omitempty"`
	Title string          `json:"title,omitempty"`
	Icon  string          `json:"icon,omitempty"`
	Args  util.FieldDescs `json:"args,omitempty"`
	Run   runFn           `json:"-"`
}

type Sandboxes []*Sandbox

func (s Sandboxes) Get(key string) *Sandbox {
	return lo.FindOrElse(s, nil, func(v *Sandbox) bool {
		return v.Key == key
	})
}

// $PF_SECTION_START(sandboxes)$.

var AllSandboxes = Sandboxes{testbed}

// $PF_SECTION_END(sandboxes)$

func Menu(_ context.Context) *menu.Item {
	ret := make(menu.Items, 0, len(AllSandboxes))
	lo.ForEach(AllSandboxes, func(s *Sandbox, _ int) {
		desc := fmt.Sprintf("Sandbox [%s]", s.Key)
		rt := fmt.Sprintf("/admin/sandbox/%s", s.Key)
		ret = append(ret, &menu.Item{Key: s.Key, Title: s.Title, Icon: s.Icon, Description: desc, Route: rt})
	})
	const desc = "Playgrounds for testing new features"
	return &menu.Item{Key: "sandbox", Title: "Sandboxes", Description: desc, Icon: "play", Route: "/admin/sandbox", Children: ret}
}
