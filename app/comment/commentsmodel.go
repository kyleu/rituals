package comment

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
)

func (c Comments) GetByModel(svc enum.ModelService, id uuid.UUID) Comments {
	var ret Comments
	for _, x := range c {
		if x.Svc == svc && x.ModelID == id {
			ret = append(ret, x)
		}
	}
	return ret
}
