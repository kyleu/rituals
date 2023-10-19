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

type ModelService struct {
	Key         string
	Title       string
	Description string
}

func (m ModelService) String() string {
	if m.Title != "" {
		return m.Title
	}
	return m.Key
}

func (m *ModelService) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(m.Key, false), nil
}

func (m *ModelService) UnmarshalJSON(data []byte) error {
	var key string
	if err := util.FromJSON(data, &key); err != nil {
		return err
	}
	*m = AllModelServices.Get(key, nil)
	return nil
}

func (m ModelService) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.Key, start)
}

func (m *ModelService) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := d.DecodeElement(&key, &start); err != nil {
		return err
	}
	*m = AllModelServices.Get(key, nil)
	return nil
}

func (m ModelService) Value() (driver.Value, error) {
	return m.Key, nil
}

func (m *ModelService) Scan(value any) error {
	if value == nil {
		return nil
	}
	if sv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := sv.(string); ok {
			*m = AllModelServices.Get(v, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan ModelService enum from value [%v]", value)
}

type ModelServices []ModelService

func (m ModelServices) Keys() []string {
	return lo.Map(m, func(x ModelService, _ int) string {
		return x.Key
	})
}

func (m ModelServices) Titles() []string {
	return lo.Map(m, func(x ModelService, _ int) string {
		return x.Title
	})
}

func (m ModelServices) Strings() []string {
	return lo.Map(m, func(x ModelService, _ int) string {
		return x.String()
	})
}

func (m ModelServices) Help() string {
	return "Available options: [" + strings.Join(m.Strings(), ", ") + "]"
}

func (m ModelServices) Get(key string, logger util.Logger) ModelService {
	for _, value := range m {
		if value.Key == key {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [ModelService] enum with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return ModelService{Key: "_error", Title: "error: " + msg}
}

func (m ModelServices) Random() ModelService {
	return m[util.RandomInt(len(m))]
}

//nolint:lll
var (
	ModelServiceTeam     = ModelService{Key: "team", Title: "Team"}
	ModelServiceSprint   = ModelService{Key: "sprint", Title: "Sprint"}
	ModelServiceEstimate = ModelService{Key: "estimate", Title: "Estimate"}
	ModelServiceStandup  = ModelService{Key: "standup", Title: "Standup"}
	ModelServiceRetro    = ModelService{Key: "retro", Title: "Retro"}
	ModelServiceStory    = ModelService{Key: "story", Title: "Story"}
	ModelServiceFeedback = ModelService{Key: "feedback", Title: "Feedback"}
	ModelServiceReport   = ModelService{Key: "report", Title: "Report"}

	AllModelServices = ModelServices{ModelServiceTeam, ModelServiceSprint, ModelServiceEstimate, ModelServiceStandup, ModelServiceRetro, ModelServiceStory, ModelServiceFeedback, ModelServiceReport}
)
