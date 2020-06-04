package act

import (
	"net/http"
	"strconv"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/database/query"
)

func ParamSetFromRequest(r *http.Request) query.ParamSet {
	ret := make(query.ParamSet)

	for qk, qs := range r.URL.Query() {
		if strings.Contains(qk, ".") {
			for _, qv := range qs {
				ret = apply(ret, qk, qv)
			}
		}
	}

	return ret
}

func apply(ps query.ParamSet, qk string, qv string) query.ParamSet {
	switch {
	case strings.HasSuffix(qk, ".o"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".o"))
		asc := true
		if strings.HasSuffix(qv, ".d") {
			asc = false
			qv = qv[0 : len(qv)-2]
		}
		curr.Orderings = append(curr.Orderings, &query.Ordering{Column: qv, Asc: asc})
	case strings.HasSuffix(qk, ".l"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".l"))
		li, err := strconv.ParseInt(qv, 10, 64)
		if err == nil {
			curr.Limit = int(li)
			max := 10000
			if curr.Limit > max {
				curr.Limit = max
			}
		}
	case strings.HasSuffix(qk, ".x"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".x"))
		xi, err := strconv.ParseInt(qv, 10, 64)
		if err == nil {
			curr.Offset = int(xi)
		}
	}
	return ps
}

func getCurr(q query.ParamSet, key string) *query.Params {
	curr, ok := q[key]
	if !ok {
		curr = &query.Params{Key: key}
		q[key] = curr
	}
	return curr
}

func IDFromParams(key string, m map[string]string) (*uuid.UUID, error) {
	retOut, ok := m[util.KeyID]
	if !ok {
		return nil, errors.New("params do not contain \"id\"")
	}

	ret := util.GetUUIDFromString(retOut)
	if ret == nil {
		return nil, util.IDError(key, retOut)
	}

	return ret, nil
}
