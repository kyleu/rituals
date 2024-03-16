// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vemail/Edit.html:2
package vemail

//line views/vemail/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vemail/Edit.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vemail/Edit.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vemail/Edit.html:12
type Edit struct {
	layout.Basic
	Model *email.Email
	IsNew bool
}

//line views/vemail/Edit.html:18
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/Edit.html:18
	qw422016.N().S(`
  <div class="card">
`)
//line views/vemail/Edit.html:20
	if p.IsNew {
//line views/vemail/Edit.html:20
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vemail/Edit.html:22
		components.StreamSVGRefIcon(qw422016, `email`, ps)
//line views/vemail/Edit.html:22
		qw422016.N().S(` New Email</h3>
    <form action="/admin/db/email/_new" class="mt" method="post">
`)
//line views/vemail/Edit.html:24
	} else {
//line views/vemail/Edit.html:24
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vemail/Edit.html:25
		qw422016.E().S(p.Model.WebPath())
//line views/vemail/Edit.html:25
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete email [`)
//line views/vemail/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vemail/Edit.html:25
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vemail/Edit.html:26
		components.StreamSVGRefIcon(qw422016, `email`, ps)
//line views/vemail/Edit.html:26
		qw422016.N().S(` Edit Email [`)
//line views/vemail/Edit.html:26
		qw422016.E().S(p.Model.String())
//line views/vemail/Edit.html:26
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vemail/Edit.html:28
	}
//line views/vemail/Edit.html:28
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vemail/Edit.html:31
	if p.IsNew {
//line views/vemail/Edit.html:31
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vemail/Edit.html:31
	}
//line views/vemail/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vemail/Edit.html:32
	edit.StreamTextareaTable(qw422016, "recipients", "", "Recipients", 8, util.ToJSON(p.Model.Recipients), 5, "Comma-separated list of values")
//line views/vemail/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vemail/Edit.html:33
	edit.StreamStringTable(qw422016, "subject", "", "Subject", p.Model.Subject, 5, "String text")
//line views/vemail/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vemail/Edit.html:34
	edit.StreamTextareaTable(qw422016, "data", "", "Data", 8, util.ToJSON(p.Model.Data.String()), 5, "JSON object")
//line views/vemail/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vemail/Edit.html:35
	edit.StreamStringTable(qw422016, "plain", "", "Plain", p.Model.Plain, 5, "String text")
//line views/vemail/Edit.html:35
	qw422016.N().S(`
          `)
//line views/vemail/Edit.html:36
	edit.StreamTextareaTable(qw422016, "html", "", "HTML", 8, p.Model.HTML, 5, "String text")
//line views/vemail/Edit.html:36
	qw422016.N().S(`
          `)
//line views/vemail/Edit.html:37
	edit.StreamUUIDTable(qw422016, "userID", "input-userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vemail/Edit.html:37
	qw422016.N().S(`
          `)
//line views/vemail/Edit.html:38
	edit.StreamStringTable(qw422016, "status", "", "Status", p.Model.Status, 5, "String text")
//line views/vemail/Edit.html:38
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-userID"), "/admin/db/user?user.l=10", "q", (o) => o["name"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vemail/Edit.html:49
}

//line views/vemail/Edit.html:49
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/Edit.html:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vemail/Edit.html:49
	p.StreamBody(qw422016, as, ps)
//line views/vemail/Edit.html:49
	qt422016.ReleaseWriter(qw422016)
//line views/vemail/Edit.html:49
}

//line views/vemail/Edit.html:49
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vemail/Edit.html:49
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vemail/Edit.html:49
	p.WriteBody(qb422016, as, ps)
//line views/vemail/Edit.html:49
	qs422016 := string(qb422016.B)
//line views/vemail/Edit.html:49
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vemail/Edit.html:49
	return qs422016
//line views/vemail/Edit.html:49
}
