package socket

import (
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) SendContentUpdate(svc string, id *uuid.UUID) error {
	if id != nil {
		err := s.WriteChannel(channel{Svc: svc, ID: *id}, &Message{Svc: svc, Cmd: ServerCmdContentUpdate, Param: nil})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing "+svc+" content update message"))
		}
	}
	return nil
}

func getContent(param map[string]interface{}) (string, bool) {
	c, ok := param["content"].(string)
	if !ok {
		return "", false
	}
	c = strings.TrimSpace(c)
	if len(c) == 0 {
		c = util.KeyNoText
	}
	return c, true
}
