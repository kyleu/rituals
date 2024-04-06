// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vreport/List.html:2
package vreport

//line views/vstandup/vreport/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vreport/List.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vreport/List.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vreport/List.html:13
type List struct {
	layout.Basic
	Models              report.Reports
	StandupsByStandupID standup.Standups
	UsersByUserID       user.Users
	Params              filter.ParamSet
}

//line views/vstandup/vreport/List.html:21
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vreport/List.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/db/standup/report/_new"><button>New</button></a></div>
    <h3>`)
//line views/vstandup/vreport/List.html:24
	components.StreamSVGRefIcon(qw422016, `file-alt`, ps)
//line views/vstandup/vreport/List.html:24
	qw422016.E().S(ps.Title)
//line views/vstandup/vreport/List.html:24
	qw422016.N().S(`</h3>
`)
//line views/vstandup/vreport/List.html:25
	if len(p.Models) == 0 {
//line views/vstandup/vreport/List.html:25
		qw422016.N().S(`    <div class="mt"><em>No reports available</em></div>
`)
//line views/vstandup/vreport/List.html:27
	} else {
//line views/vstandup/vreport/List.html:27
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vstandup/vreport/List.html:29
		StreamTable(qw422016, p.Models, p.StandupsByStandupID, p.UsersByUserID, p.Params, as, ps)
//line views/vstandup/vreport/List.html:29
		qw422016.N().S(`
    </div>
`)
//line views/vstandup/vreport/List.html:31
	}
//line views/vstandup/vreport/List.html:31
	qw422016.N().S(`  </div>
`)
//line views/vstandup/vreport/List.html:33
}

//line views/vstandup/vreport/List.html:33
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vreport/List.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vreport/List.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vreport/List.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vreport/List.html:33
}

//line views/vstandup/vreport/List.html:33
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vreport/List.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vreport/List.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vreport/List.html:33
	qs422016 := string(qb422016.B)
//line views/vstandup/vreport/List.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vreport/List.html:33
	return qs422016
//line views/vstandup/vreport/List.html:33
}
