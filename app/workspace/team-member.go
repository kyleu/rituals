package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func teamMemberUpdate(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	if !ft.Admin() {
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
	curr := ft.Members.Get(ft.Team.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this team", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.tm.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberUpdate, curr, &ft.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return ft, MsgMemberUpdated, ft.Team.PublicWebPath(), nil
}

func teamMemberRemove(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	if !ft.Admin() {
		return nil, "", "", errors.New("you do not have permission to remove this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	if *userID == ft.Self.UserID {
		return nil, "", "", errors.New("you can't remove yourself")
	}
	curr := ft.Members.Get(ft.Team.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this team", userID.String())
	}
	err := p.Svc.tm.Delete(p.Ctx, nil, curr.TeamID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberRemove, userID, &ft.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return ft, MsgMemberRemoved, ft.Team.PublicWebPath(), nil
}

func teamUpdateSelf(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == ft.Self.Name && picture == ft.Self.Picture && choice != KeyGlobal {
		return ft, MsgNoChangesNeeded, ft.Team.PublicWebPath(), nil
	}

	ft.Self.Picture = picture
	ft.Self.Name = name
	err := p.Svc.tm.Update(p.Ctx, nil, ft.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.SetName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": ft.Self.UserID, "name": name, "role": ft.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberUpdate, arg, &ft.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return ft, MsgProfileEdited, ft.Team.PublicWebPath(), nil
}
