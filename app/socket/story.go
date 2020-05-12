package socket

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

type StoryStatusChange struct {
	StoryID   uuid.UUID            `json:"storyID"`
	Status    estimate.StoryStatus `json:"status"`
	FinalVote string               `json:"finalVote"`
}

func onAddStory(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := strings.TrimSpace(param["title"].(string))
	if title == "" {
		title = "Untitled"
	}
	s.logger.Debug(fmt.Sprintf("adding story [%s]", title))

	story, err := s.estimates.NewStory(ch.ID, title, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot save new story"))
	}
	err = sendStoryUpdate(s, ch, story)
	return errors.WithStack(errors.Wrap(err, "error sending story update"))
}

func onUpdateStory(s *Service) error {
	s.logger.Debug("TODO: update story")
	return nil
}

func onSetStoryStatus(s *Service, ch channel, m map[string]interface{}) error {
	storyIDString := m["storyID"].(string)
	storyID, err := uuid.FromString(storyIDString)
	if err != nil {
		return errors.WithStack(errors.New("invalid story [" + storyIDString + "]"))
	}
	statusString := m["status"].(string)
	status := estimate.StoryStatusFromString(statusString)
	changed, finalVote, err := s.estimates.SetStoryStatus(storyID, status)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot update status of story ["+storyIDString+"]"))
	}

	if changed {
		param := StoryStatusChange{StoryID: storyID, Status: status, FinalVote: finalVote}
		msg := Message{Svc: util.SvcEstimate, Cmd: util.ServerCmdStoryStatusChange, Param: param}
		err := s.WriteChannel(ch, &msg)
		return errors.WithStack(errors.Wrap(err, "error sending story update"))
	}

	return nil
}

func onSubmitVote(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	storyIDString := param["storyID"].(string)
	storyID, err := uuid.FromString(storyIDString)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot parse story ["+storyIDString+"]"))
	}
	choice := param["choice"].(string)

	vote, err := s.estimates.UpdateVote(storyID, userID, choice)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot update vote"))
	}

	msg := Message{Svc: util.SvcEstimate, Cmd: util.ServerCmdVoteUpdate, Param: vote}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending story update"))
}

func sendStoryUpdate(s *Service, ch channel, story *estimate.Story) error {
	msg := Message{Svc: util.SvcEstimate, Cmd: util.ServerCmdStoryUpdate, Param: story}
	err := s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending story update"))
}
