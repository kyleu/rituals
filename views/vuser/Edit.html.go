// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vuser/Edit.html:2
package vuser

//line views/vuser/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vuser/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vuser/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vuser/Edit.html:11
type Edit struct {
	layout.Basic
	Model *user.User
	IsNew bool
}

//line views/vuser/Edit.html:17
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/Edit.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vuser/Edit.html:19
	if p.IsNew {
//line views/vuser/Edit.html:19
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vuser/Edit.html:21
		components.StreamSVGIcon(qw422016, `profile`, ps)
//line views/vuser/Edit.html:21
		qw422016.N().S(` New User</h3>
    <form action="/admin/db/user/_new" class="mt" method="post">
`)
//line views/vuser/Edit.html:23
	} else {
//line views/vuser/Edit.html:23
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vuser/Edit.html:24
		qw422016.E().S(p.Model.WebPath())
//line views/vuser/Edit.html:24
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete user [`)
//line views/vuser/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vuser/Edit.html:24
		qw422016.N().S(`]?"><button>`)
//line views/vuser/Edit.html:24
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vuser/Edit.html:24
		qw422016.N().S(`Delete</button></a></div>
    <h3>`)
//line views/vuser/Edit.html:25
		components.StreamSVGIcon(qw422016, `profile`, ps)
//line views/vuser/Edit.html:25
		qw422016.N().S(` Edit User [`)
//line views/vuser/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vuser/Edit.html:25
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vuser/Edit.html:27
	}
//line views/vuser/Edit.html:27
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vuser/Edit.html:30
	if p.IsNew {
//line views/vuser/Edit.html:30
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vuser/Edit.html:30
	}
//line views/vuser/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vuser/Edit.html:31
	edit.StreamStringTable(qw422016, "name", "", "Name", p.Model.Name, 5, "String text")
//line views/vuser/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vuser/Edit.html:32
	edit.StreamStringTable(qw422016, "picture", "", "Picture", p.Model.Picture, 5, "URL in string form")
//line views/vuser/Edit.html:32
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/vuser/Edit.html:38
}

//line views/vuser/Edit.html:38
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/Edit.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vuser/Edit.html:38
	p.StreamBody(qw422016, as, ps)
//line views/vuser/Edit.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vuser/Edit.html:38
}

//line views/vuser/Edit.html:38
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vuser/Edit.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vuser/Edit.html:38
	p.WriteBody(qb422016, as, ps)
//line views/vuser/Edit.html:38
	qs422016 := string(qb422016.B)
//line views/vuser/Edit.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vuser/Edit.html:38
	return qs422016
//line views/vuser/Edit.html:38
}
