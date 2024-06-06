// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/vsmember/List.html:2
package vsmember

//line views/vsprint/vsmember/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/vsmember/List.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vsmember/List.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vsmember/List.html:13
type List struct {
	layout.Basic
	Models            smember.SprintMembers
	SprintsBySprintID sprint.Sprints
	UsersByUserID     user.Users
	Params            filter.ParamSet
}

//line views/vsprint/vsmember/List.html:21
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/List.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/db/sprint/member/_new"><button>`)
//line views/vsprint/vsmember/List.html:23
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vsprint/vsmember/List.html:23
	qw422016.N().S(`New</button></a></div>
    <h3>`)
//line views/vsprint/vsmember/List.html:24
	components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vsprint/vsmember/List.html:24
	qw422016.E().S(ps.Title)
//line views/vsprint/vsmember/List.html:24
	qw422016.N().S(`</h3>
`)
//line views/vsprint/vsmember/List.html:25
	if len(p.Models) == 0 {
//line views/vsprint/vsmember/List.html:25
		qw422016.N().S(`    <div class="mt"><em>No members available</em></div>
`)
//line views/vsprint/vsmember/List.html:27
	} else {
//line views/vsprint/vsmember/List.html:27
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vsprint/vsmember/List.html:29
		StreamTable(qw422016, p.Models, p.SprintsBySprintID, p.UsersByUserID, p.Params, as, ps)
//line views/vsprint/vsmember/List.html:29
		qw422016.N().S(`
    </div>
`)
//line views/vsprint/vsmember/List.html:31
	}
//line views/vsprint/vsmember/List.html:31
	qw422016.N().S(`  </div>
`)
//line views/vsprint/vsmember/List.html:33
}

//line views/vsprint/vsmember/List.html:33
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/List.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vsmember/List.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/vsmember/List.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vsmember/List.html:33
}

//line views/vsprint/vsmember/List.html:33
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vsmember/List.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vsmember/List.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/vsmember/List.html:33
	qs422016 := string(qb422016.B)
//line views/vsprint/vsmember/List.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vsmember/List.html:33
	return qs422016
//line views/vsprint/vsmember/List.html:33
}
