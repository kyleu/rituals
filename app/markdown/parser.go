package markdown

import (
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

var policy = bluemonday.UGCPolicy()

func ToHTML(s string) string {
	html := string(blackfriday.Run([]byte(s)))
	ret := policy.Sanitize(html)
	return ret
}
