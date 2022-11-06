// Content managed by Project Forge, see [projectforge.md] for details.
package auth

import (
	"github.com/kyleu/rituals/app/util"
)

type Service struct {
	baseURL   string
	port      uint16
	providers Providers
}

func NewService(baseURL string, port uint16, logger util.Logger) *Service {
	ret := &Service{baseURL: baseURL, port: port}
	_ = ret.load(logger)
	return ret
}

func (s *Service) LoginURL() string {
	if len(s.providers) == 1 {
		return "/auth/" + s.providers[0].ID
	}
	return defaultProfilePath
}
