package query

import (
	"strings"

	"logur.dev/logur"
)

type ParamSet map[string]*Params

func (s ParamSet) Get(key string, logger logur.Logger) *Params {
	x, ok := s[key]
	if !ok {
		return nil
	}

	return x.Filtered(logger)
}

func (s ParamSet) String() string {
	ret := make([]string, 0, len(s))
	for _, p := range s {
		ret = append(ret, p.String())
	}

	return strings.Join(ret, ", ")
}
