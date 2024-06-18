// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vcomment/List.html:1
package vcomment

//line views/vcomment/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vcomment/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcomment/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcomment/List.html:11
type List struct {
	layout.Basic
	Models        comment.Comments
	UsersByUserID user.Users
	Params        filter.ParamSet
}

//line views/vcomment/List.html:18
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/List.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right mrs large-buttons">
`)
//line views/vcomment/List.html:21
	if len(p.Models) > 0 {
//line views/vcomment/List.html:21
		qw422016.N().S(`<a href="/admin/db/comment/_random"><button>`)
//line views/vcomment/List.html:21
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vcomment/List.html:21
		qw422016.N().S(` Random</button></a>`)
//line views/vcomment/List.html:21
	}
//line views/vcomment/List.html:21
	qw422016.N().S(`      <a href="/admin/db/comment/_new"><button>`)
//line views/vcomment/List.html:22
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vcomment/List.html:22
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vcomment/List.html:24
	components.StreamSVGIcon(qw422016, `comments`, ps)
//line views/vcomment/List.html:24
	qw422016.N().S(` `)
//line views/vcomment/List.html:24
	qw422016.E().S(ps.Title)
//line views/vcomment/List.html:24
	qw422016.N().S(`</h3>
`)
//line views/vcomment/List.html:25
	if len(p.Models) == 0 {
//line views/vcomment/List.html:25
		qw422016.N().S(`    <div class="mt"><em>No comments available</em></div>
`)
//line views/vcomment/List.html:27
	} else {
//line views/vcomment/List.html:27
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vcomment/List.html:29
		StreamTable(qw422016, p.Models, p.UsersByUserID, p.Params, as, ps)
//line views/vcomment/List.html:29
		qw422016.N().S(`
    </div>
`)
//line views/vcomment/List.html:31
	}
//line views/vcomment/List.html:31
	qw422016.N().S(`  </div>
`)
//line views/vcomment/List.html:33
}

//line views/vcomment/List.html:33
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/List.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcomment/List.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vcomment/List.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vcomment/List.html:33
}

//line views/vcomment/List.html:33
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcomment/List.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcomment/List.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vcomment/List.html:33
	qs422016 := string(qb422016.B)
//line views/vcomment/List.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcomment/List.html:33
	return qs422016
//line views/vcomment/List.html:33
}
