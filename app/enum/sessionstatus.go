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

type SessionStatus struct {
	Key         string
	Title       string
	Description string
}

func (s SessionStatus) String() string {
	if s.Title != "" {
		return s.Title
	}
	return s.Key
}

func (s *SessionStatus) MarshalJSON() ([]byte, error) {
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

func (s SessionStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(s.Key, start)
}

func (s *SessionStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := d.DecodeElement(&key, &start); err != nil {
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
	if sv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := sv.(string); ok {
			*s = AllSessionStatuses.Get(v, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan SessionStatus enum from value [%v]", value)
}

type SessionStatuses []SessionStatus

func (s SessionStatuses) Keys() []string {
	return lo.Map(s, func(x SessionStatus, _ int) string {
		return x.Key
	})
}

func (s SessionStatuses) Titles() []string {
	return lo.Map(s, func(x SessionStatus, _ int) string {
		return x.Title
	})
}

func (s SessionStatuses) Strings() []string {
	return lo.Map(s, func(x SessionStatus, _ int) string {
		return x.String()
	})
}

func (s SessionStatuses) Help() string {
	return "Available options: [" + strings.Join(s.Strings(), ", ") + "]"
}

func (s SessionStatuses) Get(key string, logger util.Logger) SessionStatus {
	for _, value := range s {
		if value.Key == key {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [SessionStatus] enum with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return SessionStatus{Key: "_error", Title: "error: " + msg}
}

func (s SessionStatuses) Random() SessionStatus {
	return s[util.RandomInt(len(s))]
}

var (
	SessionStatusNew      = SessionStatus{Key: "new", Title: "New"}
	SessionStatusActive   = SessionStatus{Key: "active", Title: "Active"}
	SessionStatusComplete = SessionStatus{Key: "complete", Title: "Complete"}

	AllSessionStatuses = SessionStatuses{SessionStatusNew, SessionStatusActive, SessionStatusComplete}
)
