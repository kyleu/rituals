package query

import (
	"fmt"
	"logur.dev/logur"
	"net/url"
	"strings"
)

type Ordering struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc"`
}

type Params struct {
	Key       string
	Orderings []*Ordering
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
	var ret []string
	for _, o := range p.Orderings {
		dir := ""
		if !o.Asc {
			dir = " desc"
		}
		ret = append(ret, o.Column+dir)
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
		allowed := make([]*Ordering, 0)
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
				logger.Warn("no column [" + o.Column + "] available in allowed columns for [" + p.Key + "]")
			}
		}
		return &Params{Key: p.Key, Orderings: allowed, Limit: p.Limit, Offset: p.Offset}
	} else {
		return p
	}
}

type ParamSet map[string]*Params

func (s ParamSet) Get(key string, logger logur.Logger) *Params {
	x, ok := s[key]
	if !ok {
		return nil
	}
	return x.Filtered(logger)
}
