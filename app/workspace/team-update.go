package workspace

import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
)

func updateTeam(t string, src *uuid.UUID, tgt *uuid.UUID, modelID uuid.UUID, modelTitle string, path string, userID uuid.UUID, p *Params) error {
	if src != tgt {
		if src != nil {
			param := map[string]any{"type": t, "id": modelID}
			err := p.Svc.send(enum.ModelServiceTeam, *src, action.ActChildRemove, param, &userID, p.Logger)
			if err != nil {
				return err
			}
		}
		if tgt != nil {
			param := map[string]any{"type": t, "id": modelID, "title": modelTitle, "path": path}
			err := p.Svc.send(enum.ModelServiceTeam, *tgt, action.ActChildAdd, param, &userID, p.Logger)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
