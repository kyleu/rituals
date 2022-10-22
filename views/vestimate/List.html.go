// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/List.html:2
package vestimate

//line views/vestimate/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/List.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/List.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/List.html:14
type List struct {
	layout.Basic
	Models  estimate.Estimates
	Users   user.Users
	Teams   team.Teams
	Sprints sprint.Sprints
	Params  filter.ParamSet
}

//line views/vestimate/List.html:23
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/List.html:23
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/estimate/new"><button>New</button></a></div>
    <h3>`)
//line views/vestimate/List.html:26
	components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vestimate/List.html:26
	qw422016.E().S(ps.Title)
//line views/vestimate/List.html:26
	qw422016.N().S(`</h3>
`)
//line views/vestimate/List.html:27
	if len(p.Models) == 0 {
//line views/vestimate/List.html:27
		qw422016.N().S(`    <div class="mt"><em>No estimates available</em></div>
`)
//line views/vestimate/List.html:29
	} else {
//line views/vestimate/List.html:29
		qw422016.N().S(`    <div class="overflow clear">
      `)
//line views/vestimate/List.html:31
		StreamTable(qw422016, p.Models, p.Users, p.Teams, p.Sprints, p.Params, as, ps)
//line views/vestimate/List.html:31
		qw422016.N().S(`
    </div>
`)
//line views/vestimate/List.html:33
	}
//line views/vestimate/List.html:33
	qw422016.N().S(`  </div>
`)
//line views/vestimate/List.html:35
}

//line views/vestimate/List.html:35
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/List.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/List.html:35
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/List.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/List.html:35
}

//line views/vestimate/List.html:35
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/List.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/List.html:35
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/List.html:35
	qs422016 := string(qb422016.B)
//line views/vestimate/List.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/List.html:35
	return qs422016
//line views/vestimate/List.html:35
}
