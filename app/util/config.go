package util

const AppName = "rituals.dev"

var AuthEnabled = false

type key int

const (
	ContextKey key = iota
	RoutesKey  key = iota
	InfoKey    key = iota
)
