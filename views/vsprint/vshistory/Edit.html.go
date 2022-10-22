// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/vshistory/Edit.html:2
package vshistory

//line views/vsprint/vshistory/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/vshistory/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vshistory/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vshistory/Edit.html:10
type Edit struct {
	layout.Basic
	Model *shistory.SprintHistory
	IsNew bool
}

//line views/vsprint/vshistory/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vshistory/Edit.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vsprint/vshistory/Edit.html:18
	if p.IsNew {
//line views/vsprint/vshistory/Edit.html:18
		qw422016.N().S(`    <div class="right"><a href="/sprint/shistory/random"><button>Random</button></a></div>
    <h3>`)
//line views/vsprint/vshistory/Edit.html:20
		components.StreamSVGRefIcon(qw422016, `clock`, ps)
//line views/vsprint/vshistory/Edit.html:20
		qw422016.N().S(` New History</h3>
    <form action="/sprint/shistory/new" class="mt" method="post">
`)
//line views/vsprint/vshistory/Edit.html:22
	} else {
//line views/vsprint/vshistory/Edit.html:22
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vsprint/vshistory/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vsprint/vshistory/Edit.html:23
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete history [`)
//line views/vsprint/vshistory/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vsprint/vshistory/Edit.html:23
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vsprint/vshistory/Edit.html:24
		components.StreamSVGRefIcon(qw422016, `clock`, ps)
//line views/vsprint/vshistory/Edit.html:24
		qw422016.N().S(` Edit History [`)
//line views/vsprint/vshistory/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vsprint/vshistory/Edit.html:24
		qw422016.N().S(`]</h3>
`)
//line views/vsprint/vshistory/Edit.html:25
	}
//line views/vsprint/vshistory/Edit.html:25
	qw422016.N().S(`    <form action="" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vsprint/vshistory/Edit.html:29
	if p.IsNew {
//line views/vsprint/vshistory/Edit.html:29
		components.StreamTableInput(qw422016, "slug", "Slug", p.Model.Slug, 5, "String text")
//line views/vsprint/vshistory/Edit.html:29
	}
//line views/vsprint/vshistory/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vsprint/vshistory/Edit.html:30
	components.StreamTableInputUUID(qw422016, "sprintID", "Sprint ID", &p.Model.SprintID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vsprint/vshistory/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vsprint/vshistory/Edit.html:31
	components.StreamTableInput(qw422016, "sprintName", "Sprint Name", p.Model.SprintName, 5, "String text")
//line views/vsprint/vshistory/Edit.html:31
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vsprint/vshistory/Edit.html:37
}

//line views/vsprint/vshistory/Edit.html:37
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vshistory/Edit.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vshistory/Edit.html:37
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/vshistory/Edit.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vshistory/Edit.html:37
}

//line views/vsprint/vshistory/Edit.html:37
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vshistory/Edit.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vshistory/Edit.html:37
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/vshistory/Edit.html:37
	qs422016 := string(qb422016.B)
//line views/vsprint/vshistory/Edit.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vshistory/Edit.html:37
	return qs422016
//line views/vsprint/vshistory/Edit.html:37
}
