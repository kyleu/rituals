// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vaction/List.html:2
package vaction

//line views/vaction/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vaction/List.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/List.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/List.html:12
type List struct {
	layout.Basic
	Models        action.Actions
	UsersByUserID user.Users
	Params        filter.ParamSet
}

//line views/vaction/List.html:19
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/List.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right mrs large-buttons">
`)
//line views/vaction/List.html:22
	if len(p.Models) > 0 {
//line views/vaction/List.html:22
		qw422016.N().S(`<a href="/admin/db/action/_random"><button>`)
//line views/vaction/List.html:22
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vaction/List.html:22
		qw422016.N().S(`Random</button></a>`)
//line views/vaction/List.html:22
	}
//line views/vaction/List.html:22
	qw422016.N().S(`      <a href="/admin/db/action/_new"><button>`)
//line views/vaction/List.html:23
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vaction/List.html:23
	qw422016.N().S(`New</button></a>
    </div>
    <h3>`)
//line views/vaction/List.html:25
	components.StreamSVGIcon(qw422016, `action`, ps)
//line views/vaction/List.html:25
	qw422016.E().S(ps.Title)
//line views/vaction/List.html:25
	qw422016.N().S(`</h3>
`)
//line views/vaction/List.html:26
	if len(p.Models) == 0 {
//line views/vaction/List.html:26
		qw422016.N().S(`    <div class="mt"><em>No actions available</em></div>
`)
//line views/vaction/List.html:28
	} else {
//line views/vaction/List.html:28
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vaction/List.html:30
		StreamTable(qw422016, p.Models, p.UsersByUserID, p.Params, as, ps)
//line views/vaction/List.html:30
		qw422016.N().S(`
    </div>
`)
//line views/vaction/List.html:32
	}
//line views/vaction/List.html:32
	qw422016.N().S(`  </div>
`)
//line views/vaction/List.html:34
}

//line views/vaction/List.html:34
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/List.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/List.html:34
	p.StreamBody(qw422016, as, ps)
//line views/vaction/List.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/List.html:34
}

//line views/vaction/List.html:34
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/List.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/List.html:34
	p.WriteBody(qb422016, as, ps)
//line views/vaction/List.html:34
	qs422016 := string(qb422016.B)
//line views/vaction/List.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/List.html:34
	return qs422016
//line views/vaction/List.html:34
}
