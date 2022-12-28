package util

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var policy = bluemonday.UGCPolicy()

func ToHTML(s string, trimParagraph bool) string {
	html := string(blackfriday.MarkdownCommon([]byte(s)))
	ret := policy.Sanitize(html)
	ret = strings.TrimSuffix(ret, "\n")
	if trimParagraph {
		ret = strings.TrimSuffix(strings.TrimPrefix(ret, "<p>"), "</p>")
	}
	return ret
}
