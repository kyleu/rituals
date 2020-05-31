package act

import (
	"fmt"
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

func IsContentTypeJson(c string) bool {
	return c == "application/json" || c == "text/json"
}

func RequestToString(r *http.Request) string {
	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	return strings.Join(request, "\n")
}

