package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func sprintMemberUpdate(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	if !fs.Admin() {
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
	curr := fs.Members.Get(fs.Sprint.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this sprint", userID.String())
	}
	curr.Role = enum.AllMemberStatuses.Get(role, nil)
	err := p.Svc.sm.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, curr, &fs.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, MsgMemberUpdated, fs.Sprint.PublicWebPath(), nil
}

func sprintMemberRemove(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	if !fs.Admin() {
		return nil, "", "", errors.New("you do not have permission to remove this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	if *userID == fs.Self.UserID {
		return nil, "", "", errors.New("you can't remove yourself")
	}
	curr := fs.Members.Get(fs.Sprint.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this sprint", userID.String())
	}
	err := p.Svc.sm.Delete(p.Ctx, nil, curr.SprintID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberRemove, userID, &fs.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, MsgMemberRemoved, fs.Sprint.PublicWebPath(), nil
}

func sprintUpdateSelf(p *Params, fs *FullSprint) (*FullSprint, string, string, error) {
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fs.Self.Name && picture == fs.Self.Picture && choice != KeyGlobal {
		return fs, MsgNoChangesNeeded, fs.Sprint.PublicWebPath(), nil
	}

	fs.Self.Picture = picture
	fs.Self.Name = name
	err := p.Svc.sm.Update(p.Ctx, nil, fs.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.SetName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fs.Self.UserID, "name": name, "role": fs.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, arg, &fs.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fs, MsgProfileEdited, fs.Sprint.PublicWebPath(), nil
}
