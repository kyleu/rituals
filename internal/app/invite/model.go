package invite

import (
	"time"

	"github.com/kyleu/rituals.dev/internal/app/util"

	"github.com/gofrs/uuid"
)

type invitationDTO struct {
	Key      string     `db:"key"`
	K        string     `db:"k"`
	V        string     `db:"v"`
	Src      *uuid.UUID `db:"src"`
	Tgt      *uuid.UUID `db:"tgt"`
	Note     string     `db:"note"`
	Status   string     `db:"status"`
	Redeemed *time.Time `db:"redeemed"`
	Created  time.Time  `db:"created"`
}

type InvitationType struct {
	Key string
}

var InvitationTypeEstimate = InvitationType{Key: util.SvcEstimate}
var InvitationTypeStandup = InvitationType{Key: util.SvcStandup}
var InvitationTypeRetro = InvitationType{Key: util.SvcRetro}

var AllInvitationTypes = []InvitationType{InvitationTypeEstimate, InvitationTypeStandup, InvitationTypeRetro}

func typeFromString(s string) InvitationType {
	for _, t := range AllInvitationTypes {
		if t.String() == s {
			return t
		}
	}
	return InvitationTypeEstimate
}

func (t InvitationType) String() string {
	return t.Key
}

type InvitationStatus struct {
	Key string
}

var InvitationStatusPending = InvitationStatus{Key: "pending"}
var InvitationStatusRedeemed = InvitationStatus{Key: "redeemed"}
var InvitationStatusDeleted = InvitationStatus{Key: "deleted"}

var AllInvitationStatuses = []InvitationStatus{InvitationStatusPending, InvitationStatusRedeemed, InvitationStatusDeleted}

func statusFromString(s string) InvitationStatus {
	for _, t := range AllInvitationStatuses {
		if t.String() == s {
			return t
		}
	}
	return InvitationStatusPending
}

func (t InvitationStatus) String() string {
	return t.Key
}

type Invitation struct {
	Key      string
	K        InvitationType
	V        string
	Src      *uuid.UUID
	Tgt      *uuid.UUID
	Note     string
	Status   InvitationStatus
	Redeemed *time.Time
	Created  time.Time
}

func (dto invitationDTO) ToInvitation() Invitation {
	return Invitation{
		Key:      dto.Key,
		K:        typeFromString(dto.K),
		V:        dto.V,
		Src:      dto.Src,
		Tgt:      dto.Tgt,
		Note:     dto.Note,
		Status:   statusFromString(dto.Status),
		Redeemed: dto.Redeemed,
		Created:  dto.Created,
	}
}
