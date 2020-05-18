package util

import "strings"

type Service struct {
	Key    string
	Title  string
	Plural string
	Icon   string
}

// Services
var SvcSystem = Service{
	Key:    "system",
	Title:  "System",
	Plural: "systems",
	Icon:   "close",
}
var SvcSprint = Service{
	Key:    "sprint",
	Title:  "Sprint",
	Plural: "sprints",
	Icon:   "git-fork",
}
var SvcEstimate = Service{
	Key:    "estimate",
	Title:  "Estimate",
	Plural: "estimates",
	Icon:   "settings",
}
var SvcStandup = Service{
	Key:    "standup",
	Title:  "Standup",
	Plural: "standups",
	Icon:   "future",
}
var SvcRetro = Service{
	Key:    "retro",
	Title:  "Retrospective",
	Plural: "retros",
	Icon:   "history",
}

var AllServiceKeys = []string{SvcEstimate.Key, SvcStandup.Key, SvcRetro.Key}

func ServiceTitle(title string) string {
	title = strings.TrimSpace(title)
	if title == "" {
		title = "Untitled"
	}
	return title
}
