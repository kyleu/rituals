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
	SessionStatusNew      = SessionStatus{Key: "new"}
	SessionStatusActive   = SessionStatus{Key: "active"}
	SessionStatusComplete = SessionStatus{Key: "complete"}

	AllSessionStatuses = SessionStatuses{SessionStatusNew, SessionStatusActive, SessionStatusComplete}
)

type SessionStatus struct {
	Key         string
	Name        string
	Description string
	Icon        string
}

func (s SessionStatus) String() string {
	if s.Name != "" {
		return s.Name
	}
	return s.Key
}

func (s SessionStatus) Matches(xx SessionStatus) bool {
	return s.Key == xx.Key
}

func (s SessionStatus) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(s.Key, false), nil
}

func (s *SessionStatus) UnmarshalJSON(data []byte) error {
	var key string
	if err := util.FromJSON(data, &key); err != nil {
		return err
	}
	*s = AllSessionStatuses.Get(key, nil)
	return nil
}

func (s SessionStatus) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(s.Key, start)
}

func (s *SessionStatus) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
		return err
	}
	*s = AllSessionStatuses.Get(key, nil)
	return nil
}

func (s SessionStatus) Value() (driver.Value, error) {
	return s.Key, nil
}

func (s *SessionStatus) Scan(value any) error {
	if value == nil {
		return nil
	}
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, ok := converted.(string); ok {
			*s = AllSessionStatuses.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan SessionStatus enum from value [%v]", value)
}

func SessionStatusParse(logger util.Logger, keys ...string) SessionStatuses {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) SessionStatus {
		return AllSessionStatuses.Get(x, logger)
	})
}

type SessionStatuses []SessionStatus

func (s SessionStatuses) Keys() []string {
	return lo.Map(s, func(x SessionStatus, _ int) string {
		return x.Key
	})
}

func (s SessionStatuses) Strings() []string {
	return lo.Map(s, func(x SessionStatus, _ int) string {
		return x.String()
	})
}

func (s SessionStatuses) Help() string {
	return "Available session status options: [" + strings.Join(s.Strings(), ", ") + "]"
}

func (s SessionStatuses) Get(key string, logger util.Logger) SessionStatus {
	for _, value := range s {
		if value.Key == key {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [SessionStatus] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return SessionStatus{Key: "_error", Name: "error: " + msg}
}

func (s SessionStatuses) GetByName(name string, logger util.Logger) SessionStatus {
	for _, value := range s {
		if value.Name == name {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [SessionStatus] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return SessionStatus{Key: "_error", Name: "error: " + msg}
}

func (s SessionStatuses) Random() SessionStatus {
	return s[util.RandomInt(len(s))]
}
