package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/kyleu/rituals.dev/app/query"
)

func paramSetFromRequest(r *http.Request) query.ParamSet {
	ret := make(query.ParamSet)

	getCurr := func(key string) *query.Params {
		curr, ok := ret[key]
		if !ok {
			curr = &query.Params{Key: key}
			ret[key] = curr
		}
		return curr
	}

	for qk, qs := range r.URL.Query() {
		if strings.Contains(qk, ".") {
			for _, qv := range qs {
				if strings.HasSuffix(qk, ".o") {
					key := strings.TrimSuffix(qk, ".o")
					curr := getCurr(key)
					asc := true
					if strings.HasSuffix(qv, ".d") {
						asc = false
						qv = qv[0 : len(qv)-2]
					}
					curr.Orderings = append(curr.Orderings, &query.Ordering{Column: qv, Asc: asc})
				}
				if strings.HasSuffix(qk, ".l") {
					key := strings.TrimSuffix(qk, ".l")
					curr := getCurr(key)
					li, err := strconv.ParseInt(qv, 10, 64)
					if err == nil {
						curr.Limit = int(li)
						if curr.Limit > 10000 {
							curr.Limit = 10000
						}
					}
				}
				if strings.HasSuffix(qk, ".x") {
					key := strings.TrimSuffix(qk, ".x")
					curr := getCurr(key)
					xi, err := strconv.ParseInt(qv, 10, 64)
					if err == nil {
						curr.Offset = int(xi)
					}
				}
			}
		}
	}

	return ret
}
