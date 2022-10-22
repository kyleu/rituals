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
	Models action.Actions
	Users  user.Users
	Params filter.ParamSet
}

//line views/vaction/List.html:19
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/List.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/db/action/new"><button>New</button></a></div>
    <h3>`)
//line views/vaction/List.html:22
	components.StreamSVGRefIcon(qw422016, `gamepad`, ps)
//line views/vaction/List.html:22
	qw422016.E().S(ps.Title)
//line views/vaction/List.html:22
	qw422016.N().S(`</h3>
`)
//line views/vaction/List.html:23
	if len(p.Models) == 0 {
//line views/vaction/List.html:23
		qw422016.N().S(`    <div class="mt"><em>No actions available</em></div>
`)
//line views/vaction/List.html:25
	} else {
//line views/vaction/List.html:25
		qw422016.N().S(`    <div class="overflow clear">
      `)
//line views/vaction/List.html:27
		StreamTable(qw422016, p.Models, p.Users, p.Params, as, ps)
//line views/vaction/List.html:27
		qw422016.N().S(`
    </div>
`)
//line views/vaction/List.html:29
	}
//line views/vaction/List.html:29
	qw422016.N().S(`  </div>
`)
//line views/vaction/List.html:31
}

//line views/vaction/List.html:31
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/List.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/List.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vaction/List.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/List.html:31
}

//line views/vaction/List.html:31
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/List.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/List.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vaction/List.html:31
	qs422016 := string(qb422016.B)
//line views/vaction/List.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/List.html:31
	return qs422016
//line views/vaction/List.html:31
}
