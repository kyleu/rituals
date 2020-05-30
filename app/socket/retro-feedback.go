package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

func onAddFeedback(s *Service, ch Channel, userID uuid.UUID, param addFeedbackParams) error {
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("adding [%s] feedback for [%s]", param.Category, userID))
	fb, err := s.retros.NewFeedback(ch.ID, param.Category, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot save new feedback")
	}
	err = sendFeedbackUpdate(s, ch, fb)
	return errors.Wrap(err, "error sending feedback")
}

func onEditFeedback(s *Service, ch Channel, userID uuid.UUID, param editFeedbackParams) error {
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("updating [%s] feedback for [%s]", param.Category, userID))
	feedback, err := s.retros.UpdateFeedback(param.ID, param.Category, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update feedback")
	}
	err = sendFeedbackUpdate(s, ch, feedback)
	return err
}

func onRemoveFeedback(s *Service, ch Channel, userID uuid.UUID, feedbackID uuid.UUID) error {
	s.Logger.Debug(fmt.Sprintf("removing report [%s]", feedbackID))
	err := s.retros.RemoveFeedback(feedbackID, userID)
	if err != nil {
		return errors.Wrap(err, "cannot remove feedback")
	}
	err = s.WriteChannel(ch, NewMessage(util.SvcRetro, ServerCmdFeedbackRemove, feedbackID))
	return errors.Wrap(err, "error sending feedback removal notification")
}

func sendFeedbackUpdate(s *Service, ch Channel, fb *retro.Feedback) error {
	err := s.WriteChannel(ch, NewMessage(util.SvcRetro, ServerCmdFeedbackUpdate, fb))
	return errors.Wrap(err, "error sending feedback update")
}
