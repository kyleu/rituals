package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) updatePerms(ch Channel, userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID, permSvc *permission.Service, perms permission.Permissions) error {
	filtered := make(permission.Permissions, 0)
	for _, p := range perms {
		skipTeam := p.K == util.SvcTeam.Key && teamID == nil
		skipSprint := p.K == util.SvcSprint.Key && sprintID == nil
		if (!skipTeam) && (!skipSprint) {
			filtered = append(filtered, p)
		}
	}
	filtered.Sort()

	curr := permSvc.GetByModelID(ch.ID, nil)

	if curr.Equals(filtered) {
		return nil
	}

	final, err := permSvc.SetAll(ch.ID, filtered, userID)
	if err != nil {
		return errors.Wrap(err, "unable to set permissions")
	}

	err = sendPermissionsUpdate(s, ch, final)
	return errors.Wrap(err, "unable to send permissions update")
}

func sendPermissionsUpdate(s *Service, ch Channel, perms permission.Permissions) error {
	err := s.WriteChannel(ch, NewMessage(ch.Svc, ServerCmdPermissionsUpdate, perms))
	if err != nil {
		return errors.Wrap(err, "error writing permission update message")
	}

	return err
}

func (s *Service) check(userID uuid.UUID, auths auth.Records, teamID *uuid.UUID, sprintID *uuid.UUID, svc util.Service, modelID uuid.UUID) (permission.Permissions, permission.Errors) {
	var currTeams []uuid.UUID
	if teamID != nil {
		currTeams = s.teams.GetIdsByMember(userID)
	}

	var currSprints []uuid.UUID
	if sprintID != nil {
		currSprints = s.sprints.GetIdsByMember(userID)
	}

	var tp *permission.Params
	if teamID != nil {
		tm := s.teams.GetByID(*teamID)
		tp = &permission.Params{ID: tm.ID, Slug: tm.Slug, Title: tm.Title, Current: currTeams}
	}

	var sp *permission.Params
	if sprintID != nil {
		spr := s.sprints.GetByID(*sprintID)
		sp = &permission.Params{ID: spr.ID, Slug: spr.Slug, Title: spr.Title, Current: currSprints}
	}

	perms, e := dataFor(s, svc).Permissions.Check(s.auths.Enabled, svc, modelID, auths, tp, sp)
	return perms, e
}

func (s *Service) checkPerms(userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID, svc util.Service, modelID uuid.UUID) error {
	auths := s.auths.GetByUserID(userID, nil)
	_, permErrors := s.check(userID, auths, teamID, sprintID, svc, modelID)
	if len(permErrors) > 0 {
		return errors.New("permission violation")
	}
	return nil
}
