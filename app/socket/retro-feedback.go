package socket

import (
	"fmt"

	"github.com/kyleu/npn/npnconnection"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

func onAddFeedback(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param addFeedbackParams) error {
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("adding [%s] feedback for [%s]", param.Category, userID))
	fb, err := ctx(s).retros.NewFeedback(ch.ID, param.Category, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot save new feedback")
	}
	err = sendFeedbackUpdate(s, ch, fb)
	return errors.Wrap(err, "error sending feedback")
}

func onEditFeedback(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param editFeedbackParams) error {
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("updating [%s] feedback for [%s]", param.Category, userID))
	feedback, err := ctx(s).retros.UpdateFeedback(param.ID, param.Category, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update feedback")
	}
	err = sendFeedbackUpdate(s, ch, feedback)
	return err
}

func onRemoveFeedback(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, feedbackID uuid.UUID) error {
	s.Logger.Debug(fmt.Sprintf("removing report [%s]", feedbackID))
	err := ctx(s).retros.RemoveFeedback(feedbackID, userID)
	if err != nil {
		return errors.Wrap(err, "cannot remove feedback")
	}
	err = s.WriteChannel(ch, npnconnection.NewMessage(util.SvcRetro.Key, ServerCmdFeedbackRemove, feedbackID))
	return errors.Wrap(err, "error sending feedback removal notification")
}

func sendFeedbackUpdate(s *npnconnection.Service, ch npnconnection.Channel, fb *retro.Feedback) error {
	err := s.WriteChannel(ch, npnconnection.NewMessage(util.SvcRetro.Key, ServerCmdFeedbackUpdate, fb))
	return errors.Wrap(err, "error sending feedback update")
}
