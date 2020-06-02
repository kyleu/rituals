package socket

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/util"
)

type addCommentParams struct {
	TargetType string     `json:"targetType"`
	TargetID   *uuid.UUID `json:"targetID"`
	Content    string     `json:"content"`
}

type updateCommentParams struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}

func onAddComment(s *Service, ch Channel, userID uuid.UUID, param addCommentParams) error {
	param.Content = strings.TrimSpace(param.Content)
	if param.Content == "" {
		return errors.New("add some content")
	}
	if param.TargetType == "root" {
		param.TargetType = ""
	}
	s.Logger.Debug(fmt.Sprintf("adding comment [%s] for [%v:%v]", param.Content, param.TargetType, param.TargetID))

	dataSvc := dataFor(s, ch.Svc)
	c, err := dataSvc.Comments.Add(ch.Svc, ch.ID, param.TargetType, param.TargetID, param.Content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot save new story")
	}
	err = sendCommentUpdate(s, ch, c)
	return errors.Wrap(err, "error sending story update")
}

func onUpdateComment(s *Service, ch Channel, userID uuid.UUID, param updateCommentParams) error {
	param.Content = strings.TrimSpace(param.Content)
	if param.Content == "" {
		return errors.New("add some content")
	}
	s.Logger.Debug(fmt.Sprintf("updating comment [%s]: %v", param.ID, param.Content))

	dataSvc := dataFor(s, ch.Svc)
	c, err := dataSvc.Comments.Update(param.ID, param.Content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update comment")
	}
	err = sendCommentUpdate(s, ch, c)
	return errors.Wrap(err, "error sending comment update")
}

func onRemoveComment(s *Service, ch Channel, userID uuid.UUID, commentID uuid.UUID) error {
	s.Logger.Debug(fmt.Sprintf("removing report [%s]", commentID))
	err := s.RemoveComment(commentID, userID)
	if err != nil {
		return errors.Wrap(err, "cannot remove comment")
	}
	err = s.WriteChannel(ch, NewMessage(util.SvcSystem, ServerCmdCommentRemove, commentID))
	return errors.Wrap(err, "error sending comment removal notification")
}

func sendCommentUpdate(s *Service, ch Channel, comment *comment.Comment) error {
	err := s.WriteChannel(ch, NewMessage(util.SvcSystem, ServerCmdCommentUpdate, comment))
	return errors.Wrap(err, "error sending comment update")
}
