// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vupermission/Edit.html:2
package vupermission

//line views/vstandup/vupermission/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vupermission/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vupermission/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vupermission/Edit.html:10
type Edit struct {
	layout.Basic
	Model *upermission.StandupPermission
	IsNew bool
}

//line views/vstandup/vupermission/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Edit.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vstandup/vupermission/Edit.html:18
	if p.IsNew {
//line views/vstandup/vupermission/Edit.html:18
		qw422016.N().S(`    <div class="right"><a href="/standup/upermission/random"><button>Random</button></a></div>
    <h3>`)
//line views/vstandup/vupermission/Edit.html:20
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vstandup/vupermission/Edit.html:20
		qw422016.N().S(` New Permission</h3>
    <form action="/standup/upermission/new" class="mt" method="post">
`)
//line views/vstandup/vupermission/Edit.html:22
	} else {
//line views/vstandup/vupermission/Edit.html:22
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vstandup/vupermission/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vstandup/vupermission/Edit.html:23
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete permission [`)
//line views/vstandup/vupermission/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vstandup/vupermission/Edit.html:23
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vstandup/vupermission/Edit.html:24
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vstandup/vupermission/Edit.html:24
		qw422016.N().S(` Edit Permission [`)
//line views/vstandup/vupermission/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vstandup/vupermission/Edit.html:24
		qw422016.N().S(`]</h3>
`)
//line views/vstandup/vupermission/Edit.html:25
	}
//line views/vstandup/vupermission/Edit.html:25
	qw422016.N().S(`    <form action="" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vstandup/vupermission/Edit.html:29
	if p.IsNew {
//line views/vstandup/vupermission/Edit.html:29
		components.StreamTableInputUUID(qw422016, "standupID", "Standup ID", &p.Model.StandupID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vstandup/vupermission/Edit.html:29
	}
//line views/vstandup/vupermission/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vstandup/vupermission/Edit.html:30
	if p.IsNew {
//line views/vstandup/vupermission/Edit.html:30
		components.StreamTableInput(qw422016, "k", "K", p.Model.K, 5, "String text")
//line views/vstandup/vupermission/Edit.html:30
	}
//line views/vstandup/vupermission/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vstandup/vupermission/Edit.html:31
	if p.IsNew {
//line views/vstandup/vupermission/Edit.html:31
		components.StreamTableInput(qw422016, "v", "V", p.Model.V, 5, "String text")
//line views/vstandup/vupermission/Edit.html:31
	}
//line views/vstandup/vupermission/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vstandup/vupermission/Edit.html:32
	components.StreamTableInput(qw422016, "access", "Access", p.Model.Access, 5, "String text")
//line views/vstandup/vupermission/Edit.html:32
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vstandup/vupermission/Edit.html:38
}

//line views/vstandup/vupermission/Edit.html:38
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Edit.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vupermission/Edit.html:38
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vupermission/Edit.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vupermission/Edit.html:38
}

//line views/vstandup/vupermission/Edit.html:38
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vupermission/Edit.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vupermission/Edit.html:38
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vupermission/Edit.html:38
	qs422016 := string(qb422016.B)
//line views/vstandup/vupermission/Edit.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vupermission/Edit.html:38
	return qs422016
//line views/vstandup/vupermission/Edit.html:38
}
