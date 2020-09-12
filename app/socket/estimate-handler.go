package socket

import (
	"fmt"
	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/util"
)

func onEstimateSessionSave(s *npnconnection.Service, a *auth.Service, ch npnconnection.Channel, userID uuid.UUID, param estimateSessionSaveParams) error {
	dataSvc := ctx(s).estimates
	title := util.ServiceTitle(util.SvcEstimate, param.Title)

	choices := npndatabase.StringToArray(param.Choices)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}

	sprintID := npncore.GetUUIDFromString(param.SprintID)
	teamID := npncore.GetUUIDFromString(param.TeamID)

	curr := dataSvc.GetByID(ch.ID)
	if curr == nil {
		return errors.New("no estimate available with id [" + ch.ID.String() + "]")
	}

	sr := checkPerms(s, a, userID, curr.TeamID, curr.SprintID, ch.Svc, ch.ID)
	if sr != nil {
		return sr
	}

	teamChanged := npnconnection.DifferentPointerValues(curr.TeamID, teamID)
	sprintChanged := npnconnection.DifferentPointerValues(curr.SprintID, sprintID)

	msg := "saving estimate session [%s] with choices [%s], team [%s] and sprint [%s]"
	s.Logger.Debug(fmt.Sprintf(msg, title, npncore.OxfordComma(choices, "and"), teamID, sprintID))

	err := dataSvc.UpdateSession(ch.ID, title, choices, teamID, sprintID, userID)
	if err != nil {
		return errors.Wrap(err, "error updating estimate session")
	}

	if title != curr.Title {
		slug, err := dataSvc.Data.History.UpdateSlug(curr.ID, curr.Slug, curr.Title, title, userID)
		if err != nil {
			return errors.Wrap(err, "error updating estimate slug from ["+curr.Slug+"] to ["+slug+"]")
		}
	}

	err = sendEstimateSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending estimate session")
	}

	if teamChanged {
		tm := ctx(s).teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated estimate session")
		}
	}

	if sprintChanged {
		spr := ctx(s).sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.Wrap(err, "error sending sprint for updated estimate session")
		}
	}

	err = updatePerms(s, ch, userID, teamID, sprintID, dataSvc.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendEstimateSessionUpdate(s *npnconnection.Service, ch npnconnection.Channel) error {
	est := ctx(s).estimates.GetByID(ch.ID)
	if est == nil {
		return errors.New("cannot load estimate session [" + ch.ID.String() + "]")
	}

	err := s.WriteChannel(ch, npnconnection.NewMessage(util.SvcEstimate.Key, ServerCmdSessionUpdate, est))
	return errors.Wrap(err, "error sending estimate session")
}
