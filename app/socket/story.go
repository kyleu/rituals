package socket

import (
	"fmt"

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
	title := util.ServiceTitle(param["title"].(string))
	s.logger.Debug(fmt.Sprintf("adding story [%s]", title))

	story, err := s.estimates.NewStory(ch.ID, title, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot save new story"))
	}
	err = sendStoryUpdate(s, ch, story)
	return errors.WithStack(errors.Wrap(err, "error sending story update"))
}

func onUpdateStory(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	storyID := getUUIDPointer(param, util.KeyID)
	if storyID == nil {
		return errors.New("invalid story id")
	}

	title := util.ServiceTitle(param["title"].(string))
	st, err := s.estimates.UpdateStory(*storyID, title, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot update story"))
	}
	err = sendStoryUpdate(s, ch, st)
	return errors.WithStack(err)
}

func onRemoveStory(s *Service, ch channel, userID uuid.UUID, param string) error {
	storyID, err := uuid.FromString(param)
	if err != nil {
		return errors.New("invalid story id [" + param + "]")
	}
	s.logger.Debug(fmt.Sprintf("removing report [%s]", storyID))
	err = s.estimates.RemoveStory(storyID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot remove story"))
	}
	msg := Message{Svc: util.SvcEstimate.Key, Cmd: ServerCmdStoryRemove, Param: storyID}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending story removal notification"))
}

func onSetStoryStatus(s *Service, ch channel, userID uuid.UUID, m map[string]interface{}) error {
	storyIDString := m["storyID"].(string)
	storyID, err := uuid.FromString(storyIDString)
	if err != nil {
		return errors.WithStack(errors.New("invalid story id [" + storyIDString + "]"))
	}

	statusString := m["status"].(string)
	status := estimate.StoryStatusFromString(statusString)
	changed, finalVote, err := s.estimates.SetStoryStatus(storyID, status, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot update status of story ["+storyIDString+"]"))
	}

	if changed {
		param := StoryStatusChange{StoryID: storyID, Status: status, FinalVote: finalVote}
		msg := Message{Svc: util.SvcEstimate.Key, Cmd: ServerCmdStoryStatusChange, Param: param}
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

	msg := Message{Svc: util.SvcEstimate.Key, Cmd: ServerCmdVoteUpdate, Param: vote}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending story update"))
}

func sendStoryUpdate(s *Service, ch channel, story *estimate.Story) error {
	msg := Message{Svc: util.SvcEstimate.Key, Cmd: ServerCmdStoryUpdate, Param: story}
	err := s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending story update"))
}
