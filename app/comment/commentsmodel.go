package comment

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
)

func (c Comments) GetByModel(svc enum.ModelService, id uuid.UUID) Comments {
	return lo.Filter(c, func(x *Comment, _ int) bool {
		return x.Svc == svc && x.ModelID == id
	})
}
