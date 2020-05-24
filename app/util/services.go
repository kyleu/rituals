package util

import "strings"

type Service struct {
	Key         string
	Title       string
	Plural      string
	PluralTitle string
	Icon        string
}

// Services

var SvcSystem = Service{
	Key:         "system",
	Title:       "System",
	Plural:      "systems",
	PluralTitle: "Systems",
	Icon:        "close",
}
var SvcTeam = Service{
	Key:         "team",
	Title:       "Team",
	Plural:      "teams",
	PluralTitle: "Teams",
	Icon:        "users",
}
var SvcSprint = Service{
	Key:         "sprint",
	Title:       "Sprint",
	Plural:      "sprints",
	PluralTitle: "Sprints",
	Icon:        "social",
}
var SvcEstimate = Service{
	Key:         "estimate",
	Title:       "Estimation Session",
	Plural:      "estimates",
	PluralTitle: "Estimation Sessions",
	Icon:        "settings",
}
var SvcStandup = Service{
	Key:         "standup",
	Title:       "Daily Standup",
	Plural:      "standups",
	PluralTitle: "Daily Standups",
	Icon:        "future",
}
var SvcRetro = Service{
	Key:         "retro",
	Title:       "Retrospective",
	Plural:      "retros",
	PluralTitle: "Retrospectives",
	Icon:        "history",
}

func ServiceTitle(svc string, title string) string {
	title = strings.TrimSpace(title)
	if title == "" {
		title = "Untitled " + svc
	}

	return title
}
