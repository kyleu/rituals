package auth

import (
	"github.com/kyleu/rituals.dev/app/secrets"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/slack"
)

var googleConf = oauth2.Config{
	ClientID:     secrets.GoogleClientID,
	ClientSecret: secrets.GoogleClientSecret,
	Endpoint:     google.Endpoint,
	RedirectURL:  callbackUrl(host, "google"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	},
}

var githubConf = oauth2.Config{
	ClientID:     secrets.GithubClientID,
	ClientSecret: secrets.GithubClientSecret,
	Endpoint:     github.Endpoint,
	RedirectURL:  callbackUrl(host, "github"),
	Scopes: []string{"profile"},
}

var slackConf = oauth2.Config{
	ClientID:     secrets.SlackClientID,
	ClientSecret: secrets.SlackClientSecret,
	Endpoint:     slack.Endpoint,
	RedirectURL:  callbackUrl(host, "slack"),
	Scopes: []string{"users.profile:read"},
}

func callbackUrl(host, k string) string {
	return "http://" + host + "/auth/callback/" + k
}
