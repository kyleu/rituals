package invitation

import (
	"encoding/json"
	"time"

	"github.com/kyleu/rituals.dev/app/util"

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

type Type struct {
	Key string
}

var TypeSprint = Type{Key: util.SvcSprint.Key}
var TypeEstimate = Type{Key: util.SvcEstimate.Key}
var TypeStandup = Type{Key: util.SvcStandup.Key}
var TypeRetro = Type{Key: util.SvcRetro.Key}

var AllTypes = []Type{TypeSprint, TypeEstimate, TypeStandup, TypeRetro}

func typeFromString(s string) Type {
	for _, t := range AllTypes {
		if t.Key == s {
			return t
		}
	}
	return TypeSprint
}

func (t *Type) String() string {
	return t.Key
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*t = typeFromString(s)
	return nil
}

type Status struct {
	Key string
}

var StatusPending = Status{Key: "pending"}
var StatusRedeemed = Status{Key: "redeemed"}
var StatusDeleted = Status{Key: "deleted"}

var AllStatuses = []Status{StatusPending, StatusRedeemed, StatusDeleted}

func statusFromString(s string) Status {
	for _, t := range AllStatuses {
		if t.Key == s {
			return t
		}
	}
	return StatusPending
}

func (t *Status) String() string {
	return t.Key
}

func (t Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Status) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*t = statusFromString(s)
	return nil
}

type Invitation struct {
	Key      string
	K        Type
	V        string
	Src      *uuid.UUID
	Tgt      *uuid.UUID
	Note     string
	Status   Status
	Redeemed *time.Time
	Created  time.Time
}

type Invitations = []*Invitation

func (dto *invitationDTO) ToInvitation() *Invitation {
	return &Invitation{
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
