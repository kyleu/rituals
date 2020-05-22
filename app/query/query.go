package query

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kyleu/rituals.dev/app/util"

	"logur.dev/logur"
)

type Ordering struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc"`
}

type Orderings []*Ordering

var DefaultCreatedOrdering = Orderings{{Column: util.KeyCreated, Asc: false}}
var DefaultMCreatedOrdering = Orderings{{Column: "m." + util.KeyCreated, Asc: false}}

type Params struct {
	Key       string
	Orderings Orderings
	Limit     int
	Offset    int
}

func ParamsWithDefaultOrdering(key string, params *Params, orderings ...*Ordering) *Params {
	if params == nil {
		params = &Params{Key: key}
	}

	if len(params.Orderings) == 0 {
		params.Orderings = orderings
	}

	return params
}

func (p *Params) Clone(orderings ...*Ordering) *Params {
	return &Params{Key: p.Key, Orderings: orderings, Limit: p.Limit, Offset: p.Offset}
}

func (p *Params) GetOrdering(col string) *Ordering {
	var ret *Ordering

	for _, o := range p.Orderings {
		if o.Column == col {
			ret = o
		}
	}

	return ret
}

func (p *Params) OrderByString() string {
	var ret = make([]string, len(p.Orderings))

	for i, o := range p.Orderings {
		dir := ""
		if !o.Asc {
			dir = " desc"
		}
		snake := util.ToSnakeCase(o.Column)
		ret[i] = snake + dir
	}

	return strings.Join(ret, ", ")
}

func (p *Params) ToQueryString(u *url.URL) string {
	if p == nil {
		return ""
	}

	if u == nil {
		return ""
	}

	var ret = u.Query()

	delete(ret, p.Key+".o")
	delete(ret, p.Key+".l")
	delete(ret, p.Key+".x")

	for _, o := range p.Orderings {
		s := o.Column

		if !o.Asc {
			s += ".d"
		}

		ret.Add(p.Key+".o", s)
	}

	if p.Limit > 0 {
		ret.Add(p.Key+".l", fmt.Sprintf("%v", p.Limit))
	}

	if p.Offset > 0 {
		ret.Add(p.Key+".x", fmt.Sprintf("%v", p.Offset))
	}

	return ret.Encode()
}

func (p *Params) Filtered(logger logur.Logger) *Params {
	if len(p.Orderings) > 0 {
		allowed := make(Orderings, 0)

		for _, o := range p.Orderings {
			containsCol := false
			available, ok := allowedColumns[p.Key]

			if !ok {
				logger.Warn("no columns available for [" + p.Key + "]")
			}

			for _, c := range available {
				if c == o.Column {
					containsCol = true
				}
			}

			if containsCol {
				allowed = append(allowed, o)
			} else {
				msg := "no column [%v] for [%v] available in allowed columns [%v]"
				logger.Warn(fmt.Sprintf(msg, o.Column, p.Key, strings.Join(available, ", ")))
			}
		}

		return &Params{Key: p.Key, Orderings: allowed, Limit: p.Limit, Offset: p.Offset}
	}

	return p
}

type ParamSet map[string]*Params

func (s ParamSet) Get(key string, logger logur.Logger) *Params {
	x, ok := s[key]
	if !ok {
		return nil
	}

	return x.Filtered(logger)
}
