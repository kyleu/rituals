package estimate

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Slugify(ctx context.Context, id uuid.UUID, n string, o string, h *ehistory.Service, tx *sqlx.Tx, logger util.Logger) string {
	n = util.Slugify(n)
	if n == o {
		return n
	}
	if curr, _ := s.GetBySlug(ctx, tx, n, logger); curr != nil {
		if curr.ID == id {
			return n
		}
		return s.Slugify(ctx, id, n+"-"+util.RandomString(4), o, h, tx, logger)
	}
	if hist, _ := h.Get(ctx, tx, n, logger); hist != nil {
		if hist.EstimateID != id {
			return s.Slugify(ctx, id, n+"-"+util.RandomString(4), o, h, tx, logger)
		}
		_ = h.Delete(ctx, tx, n, logger)
	}
	return n
}

func (e *Estimate) PublicWebPath() string {
	return "/estimate/" + e.Slug
}

func (e *Estimate) IconSafe() string {
	if _, ok := util.SVGLibrary[e.Icon]; !ok {
		return util.KeyTeam
	}
	return e.Icon
}
