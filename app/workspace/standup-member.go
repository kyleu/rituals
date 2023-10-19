package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func standupMemberUpdate(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	if !fu.Admin() {
		return nil, "", "", errors.New("you do not have permission to update this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := p.Frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fu.Members.Get(fu.Standup.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this standup", userID.String())
	}
	curr.Role = enum.AllMemberStatuses.Get(role, nil)
	err := p.Svc.um.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActMemberUpdate, curr, &fu.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgMemberUpdated, fu.Standup.PublicWebPath(), nil
}

func standupMemberRemove(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	if !fu.Admin() {
		return nil, "", "", errors.New("you do not have permission to remove this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	if *userID == fu.Self.UserID {
		return nil, "", "", errors.New("you can't remove yourself")
	}
	curr := fu.Members.Get(fu.Standup.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this standup", userID.String())
	}
	err := p.Svc.um.Delete(p.Ctx, nil, curr.StandupID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActMemberRemove, userID, &fu.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgMemberRemoved, fu.Standup.PublicWebPath(), nil
}

func standupUpdateSelf(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fu.Self.Name && picture == fu.Self.Picture && choice != KeyGlobal {
		return fu, MsgNoChangesNeeded, fu.Standup.PublicWebPath(), nil
	}

	fu.Self.Picture = picture
	fu.Self.Name = name
	err := p.Svc.um.Update(p.Ctx, nil, fu.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.SetName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fu.Self.UserID, "name": name, "role": fu.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActMemberUpdate, arg, &fu.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgProfileEdited, fu.Standup.PublicWebPath(), nil
}
