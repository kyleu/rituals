package auth

import (
	"time"

	"github.com/gofrs/uuid"
)

type Provider struct {
	Key   string
	Title string
	Icon  string
}

var ProviderGitHub = Provider{
	Key:   "github",
	Title: "GitHub",
	Icon:  "github-alt",
}

var ProviderGoogle = Provider{
	Key:   "google",
	Title: "Google",
	Icon:  "google",
}

var ProviderSlack = Provider{
	Key:   "slack",
	Title: "Slack",
	Icon:  "hashtag",
}

var AllProviders = []*Provider{&ProviderGitHub, &ProviderGoogle, &ProviderSlack}

func ProviderFromString(s string) *Provider {
	for _, t := range AllProviders {
		if t.Key == s {
			return t
		}
	}
	return &ProviderGitHub
}

type Record struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Provider   *Provider
	ProviderID string
	Expires    *time.Time
	Name       string
	Email      string
	Picture    string
	Created    time.Time
}

type recordDTO struct {
	ID         uuid.UUID  `db:"id"`
	UserID     uuid.UUID  `db:"user_id"`
	Provider   string     `db:"provider"`
	ProviderID string     `db:"provider_id"`
	Expires    *time.Time `db:"expires"`
	Name       string     `db:"name"`
	Email      string     `db:"email"`
	Picture    string     `db:"picture"`
	Created    time.Time  `db:"created"`
}

func (dto *recordDTO) ToRecord() *Record {
	return &Record{
		ID:         dto.ID,
		UserID:     dto.UserID,
		Provider:   ProviderFromString(dto.Provider),
		ProviderID: dto.ProviderID,
		Expires:    dto.Expires,
		Name:       dto.Name,
		Email:      dto.Email,
		Picture:    dto.Picture,
		Created:    dto.Created,
	}
}
