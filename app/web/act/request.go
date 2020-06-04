package act

import (
	"net/http"
	"strings"
)

func GetContentType(r *http.Request) string {
	ret := r.Header.Get("content-type")
	idx := strings.Index(ret, ";")
	if idx > -1 {
		ret = ret[0:idx]
	}
	return strings.TrimSpace(ret)
}

func IsContentTypeJSON(c string) bool {
	return c == "application/json" || c == "text/json"
}
