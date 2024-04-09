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

var (
	MemberStatusOwner    = MemberStatus{Key: "owner"}
	MemberStatusMember   = MemberStatus{Key: "member"}
	MemberStatusObserver = MemberStatus{Key: "observer"}

	AllMemberStatuses = MemberStatuses{MemberStatusOwner, MemberStatusMember, MemberStatusObserver}
)

type MemberStatus struct {
	Key         string
	Name        string
	Description string
	Icon        string
}

func (m MemberStatus) String() string {
	if m.Name != "" {
		return m.Name
	}
	return m.Key
}

func (m MemberStatus) Matches(xx MemberStatus) bool {
	return m.Key == xx.Key
}

func (m MemberStatus) MarshalJSON() ([]byte, error) {
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

func (m MemberStatus) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(m.Key, start)
}

func (m *MemberStatus) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
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
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, ok := converted.(string); ok {
			*m = AllMemberStatuses.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan MemberStatus enum from value [%v]", value)
}

func MemberStatusParse(logger util.Logger, keys ...string) MemberStatuses {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) MemberStatus {
		return AllMemberStatuses.Get(x, logger)
	})
}

type MemberStatuses []MemberStatus

func (m MemberStatuses) Keys() []string {
	return lo.Map(m, func(x MemberStatus, _ int) string {
		return x.Key
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
	msg := fmt.Sprintf("unable to find [MemberStatus] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return MemberStatus{Key: "_error", Name: "error: " + msg}
}

func (m MemberStatuses) GetByName(name string, logger util.Logger) MemberStatus {
	for _, value := range m {
		if value.Name == name {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [MemberStatus] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return MemberStatus{Key: "_error", Name: "error: " + msg}
}

func (m MemberStatuses) Random() MemberStatus {
	return m[util.RandomInt(len(m))]
}
