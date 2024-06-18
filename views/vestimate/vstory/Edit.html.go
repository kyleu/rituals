// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/vstory/Edit.html:1
package vstory

//line views/vestimate/vstory/Edit.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vstory/Edit.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/Edit.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/Edit.html:12
type Edit struct {
	layout.Basic
	Model *story.Story
	IsNew bool
}

//line views/vestimate/vstory/Edit.html:18
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Edit.html:18
	qw422016.N().S(`
  <div class="card">
`)
//line views/vestimate/vstory/Edit.html:20
	if p.IsNew {
//line views/vestimate/vstory/Edit.html:20
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vestimate/vstory/Edit.html:22
		components.StreamSVGIcon(qw422016, `story`, ps)
//line views/vestimate/vstory/Edit.html:22
		qw422016.N().S(` New Story</h3>
`)
//line views/vestimate/vstory/Edit.html:23
	} else {
//line views/vestimate/vstory/Edit.html:23
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vestimate/vstory/Edit.html:24
		qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vstory/Edit.html:24
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete story [`)
//line views/vestimate/vstory/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vestimate/vstory/Edit.html:24
		qw422016.N().S(`]?"><button>`)
//line views/vestimate/vstory/Edit.html:24
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vestimate/vstory/Edit.html:24
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vestimate/vstory/Edit.html:25
		components.StreamSVGIcon(qw422016, `story`, ps)
//line views/vestimate/vstory/Edit.html:25
		qw422016.N().S(` Edit Story [`)
//line views/vestimate/vstory/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vestimate/vstory/Edit.html:25
		qw422016.N().S(`]</h3>
`)
//line views/vestimate/vstory/Edit.html:26
	}
//line views/vestimate/vstory/Edit.html:26
	qw422016.N().S(`    <form action="`)
//line views/vestimate/vstory/Edit.html:27
	qw422016.E().S(util.Choose(p.IsNew, `/admin/db/estimate/story/_new`, ``))
//line views/vestimate/vstory/Edit.html:27
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vestimate/vstory/Edit.html:30
	if p.IsNew {
//line views/vestimate/vstory/Edit.html:30
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vstory/Edit.html:30
	}
//line views/vestimate/vstory/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:31
	edit.StreamUUIDTable(qw422016, "estimateID", "input-estimateID", "Estimate ID", &p.Model.EstimateID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vstory/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:32
	edit.StreamIntTable(qw422016, "idx", "", "Idx", p.Model.Idx, 5, "Integer")
//line views/vestimate/vstory/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:33
	edit.StreamUUIDTable(qw422016, "userID", "input-userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vstory/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:34
	edit.StreamStringTable(qw422016, "title", "", "Title", p.Model.Title, 5, "String text")
//line views/vestimate/vstory/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:35
	edit.StreamSelectTable(qw422016, "status", "", "Status", p.Model.Status.Key, enum.AllSessionStatuses.Keys(), enum.AllSessionStatuses.Strings(), 5, enum.AllSessionStatuses.Help())
//line views/vestimate/vstory/Edit.html:35
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Edit.html:36
	edit.StreamStringTable(qw422016, "finalVote", "", "Final Vote", p.Model.FinalVote, 5, "String text")
//line views/vestimate/vstory/Edit.html:36
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-estimateID"), "/admin/db/estimate?estimate.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " / " + o["choices"] + " (" + o["id"] + ")", (o) => o["id"]);
      rituals.autocomplete(document.getElementById("input-userID"), "/admin/db/user?user.l=10", "q", (o) => o["name"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vestimate/vstory/Edit.html:48
}

//line views/vestimate/vstory/Edit.html:48
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Edit.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/Edit.html:48
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/Edit.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/Edit.html:48
}

//line views/vestimate/vstory/Edit.html:48
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/Edit.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/Edit.html:48
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/Edit.html:48
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/Edit.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/Edit.html:48
	return qs422016
//line views/vestimate/vstory/Edit.html:48
}
