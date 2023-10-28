// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vcomment/Edit.html:2
package vcomment

//line views/vcomment/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vcomment/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcomment/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcomment/Edit.html:11
type Edit struct {
	layout.Basic
	Model *comment.Comment
	IsNew bool
}

//line views/vcomment/Edit.html:17
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Edit.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vcomment/Edit.html:19
	if p.IsNew {
//line views/vcomment/Edit.html:19
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vcomment/Edit.html:21
		components.StreamSVGRefIcon(qw422016, `comments`, ps)
//line views/vcomment/Edit.html:21
		qw422016.N().S(` New Comment</h3>
    <form action="/admin/db/comment/_new" class="mt" method="post">
`)
//line views/vcomment/Edit.html:23
	} else {
//line views/vcomment/Edit.html:23
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vcomment/Edit.html:24
		qw422016.E().S(p.Model.WebPath())
//line views/vcomment/Edit.html:24
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete comment [`)
//line views/vcomment/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vcomment/Edit.html:24
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vcomment/Edit.html:25
		components.StreamSVGRefIcon(qw422016, `comments`, ps)
//line views/vcomment/Edit.html:25
		qw422016.N().S(` Edit Comment [`)
//line views/vcomment/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vcomment/Edit.html:25
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vcomment/Edit.html:27
	}
//line views/vcomment/Edit.html:27
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vcomment/Edit.html:30
	if p.IsNew {
//line views/vcomment/Edit.html:30
		components.StreamTableInputUUID(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vcomment/Edit.html:30
	}
//line views/vcomment/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vcomment/Edit.html:31
	components.StreamTableSelect(qw422016, "svc", "", "Svc", p.Model.Svc.Key, enum.AllModelServices.Keys(), enum.AllModelServices.Strings(), 5, enum.AllModelServices.Help())
//line views/vcomment/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vcomment/Edit.html:32
	components.StreamTableInputUUID(qw422016, "modelID", "", "Model ID", &p.Model.ModelID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vcomment/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vcomment/Edit.html:33
	components.StreamTableInputUUID(qw422016, "userID", "input-userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vcomment/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vcomment/Edit.html:34
	components.StreamTableInput(qw422016, "content", "", "Content", p.Model.Content, 5, "String text")
//line views/vcomment/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vcomment/Edit.html:35
	components.StreamTableTextarea(qw422016, "html", "", "HTML", 8, p.Model.HTML, 5, "String text")
//line views/vcomment/Edit.html:35
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
//line views/vcomment/Edit.html:46
}

//line views/vcomment/Edit.html:46
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Edit.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcomment/Edit.html:46
	p.StreamBody(qw422016, as, ps)
//line views/vcomment/Edit.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vcomment/Edit.html:46
}

//line views/vcomment/Edit.html:46
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcomment/Edit.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcomment/Edit.html:46
	p.WriteBody(qb422016, as, ps)
//line views/vcomment/Edit.html:46
	qs422016 := string(qb422016.B)
//line views/vcomment/Edit.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcomment/Edit.html:46
	return qs422016
//line views/vcomment/Edit.html:46
}
