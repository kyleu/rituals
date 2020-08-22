package socket

import (
	"fmt"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/rituals.dev/app/session"
	"strings"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

type StoryStatusChange struct {
	StoryID   uuid.UUID      `json:"storyID"`
	Status    session.Status `json:"status"`
	FinalVote string         `json:"finalVote"`
}

func onAddStory(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param addStoryParams) error {
	param.Title = strings.TrimSpace(param.Title)
	if len(param.Title) == 0 {
		param.Title = "Untitled " + npncore.Title(util.KeyStory)
	}
	s.Logger.Debug(fmt.Sprintf("adding story [%s]", param.Title))

	story, err := estimates(s).NewStory(ch.ID, param.Title, userID)
	if err != nil {
		return errors.Wrap(err, "cannot save new story")
	}
	err = sendStoryUpdate(s, ch, story)
	return errors.Wrap(err, "error sending story update")
}

func onUpdateStory(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param updateStoryParams) error {
	param.Title = strings.TrimSpace(param.Title)
	if len(param.Title) == 0 {
		param.Title = "Untitled " + npncore.Title(util.KeyStory)
	}
	st, err := estimates(s).UpdateStory(param.StoryID, param.Title, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update story")
	}
	err = sendStoryUpdate(s, ch, st)
	return err
}

func onRemoveStory(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, storyID uuid.UUID) error {
	s.Logger.Debug(fmt.Sprintf("removing report [%s]", storyID))
	err := estimates(s).RemoveStory(storyID, userID)
	if err != nil {
		return errors.Wrap(err, "cannot remove story")
	}
	err = s.WriteChannel(ch, npnconnection.NewMessage(util.SvcEstimate.Key, ServerCmdStoryRemove, storyID))
	return errors.Wrap(err, "error sending story removal notification")
}

func onSetStoryStatus(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, params setStoryStatusParams) error {
	status := session.StatusFromString(params.Status)
	changed, finalVote, err := estimates(s).SetStoryStatus(params.StoryID, status, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update status of story ["+params.StoryID.String()+"]")
	}

	if changed {
		param := StoryStatusChange{StoryID: params.StoryID, Status: status, FinalVote: finalVote}
		err := s.WriteChannel(ch, npnconnection.NewMessage(util.SvcEstimate.Key, ServerCmdStoryStatusChange, param))
		return errors.Wrap(err, "error sending story update")
	}

	return nil
}

func onSubmitVote(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param submitVoteParams) error {
	vote, err := estimates(s).UpdateVote(param.StoryID, userID, param.Choice)
	if err != nil {
		return errors.Wrap(err, "cannot update vote")
	}

	err = s.WriteChannel(ch, npnconnection.NewMessage(util.SvcEstimate.Key, ServerCmdVoteUpdate, vote))
	return errors.Wrap(err, "error sending story update")
}

func sendStoryUpdate(s *npnconnection.Service, ch npnconnection.Channel, story *estimate.Story) error {
	err := s.WriteChannel(ch, npnconnection.NewMessage(util.SvcEstimate.Key, ServerCmdStoryUpdate, story))
	return errors.Wrap(err, "error sending story update")
}