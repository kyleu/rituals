// Code generated by qtc from "MarkdownPage.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsite/MarkdownPage.html:2
package vsite

//line views/vsite/MarkdownPage.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsite/MarkdownPage.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/MarkdownPage.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/MarkdownPage.html:8
type MarkdownPage struct {
	layout.Basic
	Title string
	HTML  string
}

//line views/vsite/MarkdownPage.html:14
func (p *MarkdownPage) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/MarkdownPage.html:14
	qw422016.N().S(`
  <div class="card markdown">
    `)
//line views/vsite/MarkdownPage.html:16
	qw422016.N().S(p.HTML)
//line views/vsite/MarkdownPage.html:16
	qw422016.N().S(`
  </div>
`)
//line views/vsite/MarkdownPage.html:18
}

//line views/vsite/MarkdownPage.html:18
func (p *MarkdownPage) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/MarkdownPage.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/MarkdownPage.html:18
	p.StreamBody(qw422016, as, ps)
//line views/vsite/MarkdownPage.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/MarkdownPage.html:18
}

//line views/vsite/MarkdownPage.html:18
func (p *MarkdownPage) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/MarkdownPage.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/MarkdownPage.html:18
	p.WriteBody(qb422016, as, ps)
//line views/vsite/MarkdownPage.html:18
	qs422016 := string(qb422016.B)
//line views/vsite/MarkdownPage.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/MarkdownPage.html:18
	return qs422016
//line views/vsite/MarkdownPage.html:18
}
