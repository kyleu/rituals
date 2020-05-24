package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

func onAddFeedback(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	category, ok := param["category"].(string)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("can't read string from [%v]", param["category"])))
	}

	content, ok := getContent(param)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("can't read content from [%v]", param["content"])))
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
	id := getUUIDPointer(param, util.KeyID)
	if id == nil {
		return errors.WithStack(errors.New("no id provided"))
	}

	category, ok := param["category"].(string)
	if !ok {
		return errors.New("cannot read category")
	}

	content, ok := getContent(param)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("can't read content from [%v]", param["content"])))
	}

	s.logger.Debug(fmt.Sprintf("updating [%s] report for [%s]", category, userID))
	feedback, err := s.retros.UpdateFeedback(*id, category, content, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot update feedback"))
	}
	err = sendFeedbackUpdate(s, ch, feedback)
	return errors.WithStack(err)
}

func onRemoveFeedback(s *Service, ch channel, userID uuid.UUID, param string) error {
	feedbackID, err := uuid.FromString(param)
	if err != nil {
		return errors.New("invalid feedback id [" + param + "]")
	}
	s.logger.Debug(fmt.Sprintf("removing report [%s]", feedbackID))
	err = s.retros.RemoveFeedback(feedbackID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot remove feedback"))
	}
	msg := Message{Svc: util.SvcRetro.Key, Cmd: ServerCmdFeedbackRemove, Param: feedbackID}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending feedback removal notification"))
}

func sendFeedbackUpdate(s *Service, ch channel, fb *retro.Feedback) error {
	msg := Message{Svc: util.SvcRetro.Key, Cmd: ServerCmdFeedbackUpdate, Param: fb}
	err := s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending feedback update"))
}
