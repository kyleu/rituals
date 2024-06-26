// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vretro/vfeedback/Edit.html:1
package vfeedback

//line views/vretro/vfeedback/Edit.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vfeedback/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vfeedback/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vfeedback/Edit.html:11
type Edit struct {
	layout.Basic
	Model *feedback.Feedback
	IsNew bool
}

//line views/vretro/vfeedback/Edit.html:17
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vfeedback/Edit.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vretro/vfeedback/Edit.html:19
	if p.IsNew {
//line views/vretro/vfeedback/Edit.html:19
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vretro/vfeedback/Edit.html:21
		components.StreamSVGIcon(qw422016, `comment`, ps)
//line views/vretro/vfeedback/Edit.html:21
		qw422016.N().S(` New Feedback</h3>
`)
//line views/vretro/vfeedback/Edit.html:22
	} else {
//line views/vretro/vfeedback/Edit.html:22
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vretro/vfeedback/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vretro/vfeedback/Edit.html:23
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete feedback [`)
//line views/vretro/vfeedback/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vretro/vfeedback/Edit.html:23
		qw422016.N().S(`]?"><button>`)
//line views/vretro/vfeedback/Edit.html:23
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vretro/vfeedback/Edit.html:23
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vretro/vfeedback/Edit.html:24
		components.StreamSVGIcon(qw422016, `comment`, ps)
//line views/vretro/vfeedback/Edit.html:24
		qw422016.N().S(` Edit Feedback [`)
//line views/vretro/vfeedback/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vretro/vfeedback/Edit.html:24
		qw422016.N().S(`]</h3>
`)
//line views/vretro/vfeedback/Edit.html:25
	}
//line views/vretro/vfeedback/Edit.html:25
	qw422016.N().S(`    <form action="`)
//line views/vretro/vfeedback/Edit.html:26
	qw422016.E().S(util.Choose(p.IsNew, `/admin/db/retro/feedback/_new`, ``))
//line views/vretro/vfeedback/Edit.html:26
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vretro/vfeedback/Edit.html:29
	if p.IsNew {
//line views/vretro/vfeedback/Edit.html:29
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vfeedback/Edit.html:29
	}
//line views/vretro/vfeedback/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vretro/vfeedback/Edit.html:30
	edit.StreamUUIDTable(qw422016, "retroID", "input-retroID", "Retro ID", &p.Model.RetroID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vfeedback/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vretro/vfeedback/Edit.html:31
	edit.StreamIntTable(qw422016, "idx", "", "Idx", p.Model.Idx, 5, "Integer")
//line views/vretro/vfeedback/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vretro/vfeedback/Edit.html:32
	edit.StreamUUIDTable(qw422016, "userID", "input-userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vfeedback/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vretro/vfeedback/Edit.html:33
	edit.StreamStringTable(qw422016, "category", "", "Category", p.Model.Category, 5, "String text")
//line views/vretro/vfeedback/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vretro/vfeedback/Edit.html:34
	edit.StreamStringTable(qw422016, "content", "", "Content", p.Model.Content, 5, "String text")
//line views/vretro/vfeedback/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vretro/vfeedback/Edit.html:35
	edit.StreamTextareaTable(qw422016, "html", "", "HTML", 8, p.Model.HTML, 5, "HTML code, in string form")
//line views/vretro/vfeedback/Edit.html:35
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-retroID"), "/admin/db/retro?retro.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " / " + o["categories"] + " (" + o["id"] + ")", (o) => o["id"]);
      rituals.autocomplete(document.getElementById("input-userID"), "/admin/db/user?user.l=10", "q", (o) => o["name"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vretro/vfeedback/Edit.html:47
}

//line views/vretro/vfeedback/Edit.html:47
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vfeedback/Edit.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vfeedback/Edit.html:47
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vfeedback/Edit.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vfeedback/Edit.html:47
}

//line views/vretro/vfeedback/Edit.html:47
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vfeedback/Edit.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vfeedback/Edit.html:47
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vfeedback/Edit.html:47
	qs422016 := string(qb422016.B)
//line views/vretro/vfeedback/Edit.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vfeedback/Edit.html:47
	return qs422016
//line views/vretro/vfeedback/Edit.html:47
}
