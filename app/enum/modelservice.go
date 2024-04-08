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

//nolint:lll
var (
	ModelServiceTeam     = ModelService{Key: "team"}
	ModelServiceSprint   = ModelService{Key: "sprint"}
	ModelServiceEstimate = ModelService{Key: "estimate"}
	ModelServiceStandup  = ModelService{Key: "standup"}
	ModelServiceRetro    = ModelService{Key: "retro"}
	ModelServiceStory    = ModelService{Key: "story"}
	ModelServiceFeedback = ModelService{Key: "feedback"}
	ModelServiceReport   = ModelService{Key: "report"}

	AllModelServices = ModelServices{ModelServiceTeam, ModelServiceSprint, ModelServiceEstimate, ModelServiceStandup, ModelServiceRetro, ModelServiceStory, ModelServiceFeedback, ModelServiceReport}
)

type ModelService struct {
	Key         string
	Name        string
	Description string
	Icon        string
}

func (m ModelService) String() string {
	if m.Name != "" {
		return m.Name
	}
	return m.Key
}

func (m ModelService) MarshalJSON() ([]byte, error) {
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

func (m ModelService) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(m.Key, start)
}

func (m *ModelService) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
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
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, ok := converted.(string); ok {
			*m = AllModelServices.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan ModelService enum from value [%v]", value)
}

func ModelServiceParse(logger util.Logger, keys ...string) ModelServices {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) ModelService {
		return AllModelServices.Get(x, logger)
	})
}

type ModelServices []ModelService

func (m ModelServices) Keys() []string {
	return lo.Map(m, func(x ModelService, _ int) string {
		return x.Key
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
	msg := fmt.Sprintf("unable to find [ModelService] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return ModelService{Key: "_error", Name: "error: " + msg}
}

func (m ModelServices) GetByName(name string, logger util.Logger) ModelService {
	for _, value := range m {
		if value.Name == name {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [ModelService] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return ModelService{Key: "_error", Name: "error: " + msg}
}

func (m ModelServices) Random() ModelService {
	return m[util.RandomInt(len(m))]
}
