package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func estimateMemberUpdate(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	if !fe.Admin() {
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
	curr := fe.Members.Get(fe.Estimate.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this estimate", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.em.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActMemberUpdate, curr, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgMemberUpdated, fe.Estimate.PublicWebPath(), nil
}

func estimateMemberRemove(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	if !fe.Admin() {
		return nil, "", "", errors.New("you do not have permission to remove this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	if *userID == fe.Self.UserID {
		return nil, "", "", errors.New("you can't remove yourself")
	}
	curr := fe.Members.Get(fe.Estimate.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this estimate", userID.String())
	}
	err := p.Svc.em.Delete(p.Ctx, nil, curr.EstimateID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActMemberRemove, userID, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgMemberRemoved, fe.Estimate.PublicWebPath(), nil
}

func estimateUpdateSelf(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fe.Self.Name && picture == fe.Self.Picture && choice != KeyGlobal {
		return fe, MsgNoChangesNeeded, fe.Estimate.PublicWebPath(), nil
	}

	fe.Self.Picture = picture
	fe.Self.Name = name
	err := p.Svc.em.Update(p.Ctx, nil, fe.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.SetName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fe.Self.UserID, "name": name, "role": fe.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActMemberUpdate, arg, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgProfileEdited, fe.Estimate.PublicWebPath(), nil
}
