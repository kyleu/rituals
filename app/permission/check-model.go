package permission

import (
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) checkModel(provider string, svc string, perms Permissions, spr *Params) *Error {
	sp := perms.FindByK(util.SvcSprint.Key)

	if len(sp) == 0 || spr == nil {
		return nil
	}

	hasSprint := false
	for _, t := range spr.Current {
		if t == spr.ID {
			hasSprint = true
			break
		}
	}

	if hasSprint {
		return nil
	}

	return &Error{Svc: svc, Provider: provider, Match: spr.Slug, Message: spr.Title}
}
