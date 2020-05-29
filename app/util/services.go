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
	Key:         KeySystem,
	Title:       Title(KeySystem),
	Plural:      Plural(KeySystem),
	PluralTitle: PluralTitle(KeySystem),
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
	Title:       "Estimate",
	Plural:      "estimates",
	PluralTitle: "Estimates",
	Icon:        "settings",
}
var SvcStandup = Service{
	Key:         "standup",
	Title:       "Standup",
	Plural:      "standups",
	PluralTitle: "Standups",
	Icon:        "future",
}
var SvcRetro = Service{
	Key:         "retro",
	Title:       "Retrospective",
	Plural:      "retros",
	PluralTitle: "Retrospectives",
	Icon:        "history",
}

var allServices = []Service{SvcSystem, SvcTeam, SvcSprint, SvcEstimate, SvcStandup, SvcRetro}

func ServiceFromString(str string) Service {
	for _, s := range allServices {
		if s.Key == str {
			return s
		}
	}
	return SvcSystem
}

func ServiceTitle(svc Service, title string) string {
	title = strings.TrimSpace(title)
	if title == "" {
		title = "Untitled " + svc.Title
	}

	return title
}
