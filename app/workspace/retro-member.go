package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func retroMemberUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if !fr.Admin() {
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
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	curr.Role = enum.AllMemberStatuses.Get(role, nil)
	err := p.Svc.rm.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, curr, &fr.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgMemberUpdated, fr.Retro.PublicWebPath(), nil
}

func retroMemberRemove(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if !fr.Admin() {
		return nil, "", "", errors.New("you do not have permission to remove this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	if *userID == fr.Self.UserID {
		return nil, "", "", errors.New("you can't remove yourself")
	}
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	err := p.Svc.rm.Delete(p.Ctx, nil, curr.RetroID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberRemove, userID, &fr.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgMemberRemoved, fr.Retro.PublicWebPath(), nil
}

func retroUpdateSelf(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fr.Self.Name && picture == fr.Self.Picture && choice != KeyGlobal {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}

	fr.Self.Picture = picture
	fr.Self.Name = name
	err := p.Svc.rm.Update(p.Ctx, nil, fr.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.SetName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fr.Self.UserID, "name": name, "role": fr.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, arg, &fr.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgProfileEdited, fr.Retro.PublicWebPath(), nil
}
