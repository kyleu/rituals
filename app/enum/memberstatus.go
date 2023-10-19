// Package enum - Content managed by Project Forge, see [projectforge.md] for details.
package enum

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type MemberStatus struct {
	Key         string
	Title       string
	Description string
}

func (m MemberStatus) String() string {
	if m.Title != "" {
		return m.Title
	}
	return m.Key
}

func (m *MemberStatus) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(m.Key, false), nil
}

func (m *MemberStatus) UnmarshalJSON(data []byte) error {
	var key string
	if err := util.FromJSON(data, &key); err != nil {
		return err
	}
	*m = AllMemberStatuses.Get(key, nil)
	return nil
}

func (m MemberStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.Key, start)
}

func (m *MemberStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := d.DecodeElement(&key, &start); err != nil {
		return err
	}
	*m = AllMemberStatuses.Get(key, nil)
	return nil
}

func (m MemberStatus) Value() (driver.Value, error) {
	return m.Key, nil
}

func (m *MemberStatus) Scan(value any) error {
	if value == nil {
		return nil
	}
	if sv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := sv.(string); ok {
			*m = AllMemberStatuses.Get(v, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan MemberStatus enum from value [%v]", value)
}

type MemberStatuses []MemberStatus

func (m MemberStatuses) Keys() []string {
	return lo.Map(m, func(x MemberStatus, _ int) string {
		return x.Key
	})
}

func (m MemberStatuses) Titles() []string {
	return lo.Map(m, func(x MemberStatus, _ int) string {
		return x.Title
	})
}

func (m MemberStatuses) Strings() []string {
	return lo.Map(m, func(x MemberStatus, _ int) string {
		return x.String()
	})
}

func (m MemberStatuses) Help() string {
	return "Available options: [" + strings.Join(m.Strings(), ", ") + "]"
}

func (m MemberStatuses) Get(key string, logger util.Logger) MemberStatus {
	for _, value := range m {
		if value.Key == key {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [MemberStatus] enum with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return MemberStatus{Key: "_error", Title: "error: " + msg}
}

func (m MemberStatuses) Random() MemberStatus {
	return m[util.RandomInt(len(m))]
}

var (
	MemberStatusOwner    = MemberStatus{Key: "owner", Title: "Owner"}
	MemberStatusMember   = MemberStatus{Key: "member", Title: "Member"}
	MemberStatusObserver = MemberStatus{Key: "observer", Title: "Observer"}

	AllMemberStatuses = MemberStatuses{MemberStatusOwner, MemberStatusMember, MemberStatusObserver}
)
