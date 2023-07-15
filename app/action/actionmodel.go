package action

import (
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func NewAction(svc enum.ModelService, model uuid.UUID, user uuid.UUID, act string, content util.ValueMap, note string) *Action {
	return &Action{ID: util.UUID(), Svc: svc, ModelID: model, UserID: user, Act: act, Content: content, Note: note, Created: time.Now()}
}

func (a Actions) GetByModel(svc enum.ModelService, id uuid.UUID) Actions {
	return lo.FilterMap(a, func(x *Action, _ int) (*Action, bool) {
		return x, x.Svc == svc && x.ModelID == id
	})
}
