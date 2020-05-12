package socket

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

func onAddFeedback(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	category, ok := param["category"].(string)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("can't read string from [%v]", param["category"])))
	}

	c, ok := param["content"].(string)
	if !ok {
		return errors.WithStack(errors.New("cannot read content"))
	}
	content := strings.TrimSpace(c)
	if len(content) == 0 {
		content = "-no text-"
	}

	s.logger.Debug(fmt.Sprintf("adding [%s] feedback for [%s]", category, userID))
	fb, err := s.retros.NewFeedback(ch.ID, category, content, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot save new feedback"))
	}
	err = sendFeedbackUpdate(s, ch, fb)
	return errors.WithStack(errors.Wrap(err, "error sending feedback"))
}

func onEditFeedback(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	i, ok := param["id"].(string)
	if !ok {
		return errors.WithStack(errors.New("Cannot read [%T] as string"))
	}
	id, err := uuid.FromString(i)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("cannot parse uuid [%v]: %+v", i, err))
	}

	category, ok := param["category"].(string)
	if !ok {
		return errors.New("cannot read category")
	}

	c, ok := param["content"].(string)
	if !ok {
		return errors.WithStack(errors.Wrap(err, "cannot read feedback content"))
	}
	content := strings.TrimSpace(c)
	if len(content) == 0 {
		content = "-no text-"
	}

	s.logger.Debug(fmt.Sprintf("updating [%s] report for [%s]", category, userID))
	feedback, err := s.retros.UpdateFeedback(id, category, content)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot update feedback"))
	}
	err = sendFeedbackUpdate(s, ch, feedback)
	return errors.WithStack(err)
}

func sendFeedbackUpdate(s *Service, ch channel, fb *retro.Feedback) error {
	msg := Message{Svc: util.SvcRetro, Cmd: util.ServerCmdFeedbackUpdate, Param: fb}
	err := s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending feedback update"))
}
