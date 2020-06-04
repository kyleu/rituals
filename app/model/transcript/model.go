package transcript

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/config"
)

type Resolver func(app *config.AppInfo, userID uuid.UUID, param interface{}, format string) (interface{}, error)

type Transcript struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Resolve     Resolver `json:"-"`
}

type Transcripts = []*Transcript

var AllTranscripts = Transcripts{&Email, &Team, &Sprint, &Estimate, &Standup, &Retro}

func FromString(s string) *Transcript {
	for _, t := range AllTranscripts {
		if t.Key == s {
			return t
		}
	}
	return nil
}
