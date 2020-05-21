package markdown

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var policy = bluemonday.UGCPolicy()

func ToHTML(s string) string {
	html := string(blackfriday.Run([]byte(s)))
	ret := policy.Sanitize(html)
	ret = strings.TrimSuffix(ret, "\n")

	return ret
}
