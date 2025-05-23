// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsprint/Edit.html:1
package vsprint

//line views/vsprint/Edit.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/Edit.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/Edit.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/Edit.html:12
type Edit struct {
	layout.Basic
	Model *sprint.Sprint
	Paths []string
	IsNew bool
}

//line views/vsprint/Edit.html:19
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Edit.html:19
	qw422016.N().S(`
  <div class="card">
`)
//line views/vsprint/Edit.html:21
	if p.IsNew {
//line views/vsprint/Edit.html:21
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vsprint/Edit.html:23
		components.StreamSVGIcon(qw422016, `sprint`, ps)
//line views/vsprint/Edit.html:23
		qw422016.N().S(` New Sprint</h3>
`)
//line views/vsprint/Edit.html:24
	} else {
//line views/vsprint/Edit.html:24
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vsprint/Edit.html:25
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vsprint/Edit.html:25
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete sprint [`)
//line views/vsprint/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vsprint/Edit.html:25
		qw422016.N().S(`]?"><button>`)
//line views/vsprint/Edit.html:25
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vsprint/Edit.html:25
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vsprint/Edit.html:26
		components.StreamSVGIcon(qw422016, `sprint`, ps)
//line views/vsprint/Edit.html:26
		qw422016.N().S(` Edit Sprint [`)
//line views/vsprint/Edit.html:26
		qw422016.E().S(p.Model.String())
//line views/vsprint/Edit.html:26
		qw422016.N().S(`]</h3>
`)
//line views/vsprint/Edit.html:27
	}
//line views/vsprint/Edit.html:27
	qw422016.N().S(`    <form action="`)
//line views/vsprint/Edit.html:28
	qw422016.E().S(util.Choose(p.IsNew, sprint.Route(p.Paths...)+`/_new`, p.Model.WebPath(p.Paths...)+`/edit`))
//line views/vsprint/Edit.html:28
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vsprint/Edit.html:31
	if p.IsNew {
//line views/vsprint/Edit.html:31
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vsprint/Edit.html:31
	}
//line views/vsprint/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:32
	edit.StreamStringTable(qw422016, "slug", "", "Slug", p.Model.Slug, 5, "String text")
//line views/vsprint/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:33
	edit.StreamStringTable(qw422016, "title", "", "Title", p.Model.Title, 5, "String text")
//line views/vsprint/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:34
	edit.StreamStringTable(qw422016, "icon", "", "Icon", p.Model.Icon, 5, "String text")
//line views/vsprint/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:35
	edit.StreamSelectTable(qw422016, "status", "", "Status", p.Model.Status.Key, enum.AllSessionStatuses.Keys(), enum.AllSessionStatuses.Strings(), 5, enum.AllSessionStatuses.Help())
//line views/vsprint/Edit.html:35
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:36
	edit.StreamUUIDTable(qw422016, "teamID", "input-teamID", "Team ID", p.Model.TeamID, 5, "UUID in format (00000000-0000-0000-0000-000000000000) (optional)")
//line views/vsprint/Edit.html:36
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:37
	edit.StreamTimestampDayTable(qw422016, "startDate", "", "Start Date", p.Model.StartDate, 5, "Calendar date (optional)")
//line views/vsprint/Edit.html:37
	qw422016.N().S(`
          `)
//line views/vsprint/Edit.html:38
	edit.StreamTimestampDayTable(qw422016, "endDate", "", "End Date", p.Model.EndDate, 5, "Calendar date (optional)")
//line views/vsprint/Edit.html:38
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-teamID"), "/admin/db/team?team.l=10", "q", (o) => (o["title"] || "[no title]") + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vsprint/Edit.html:49
}

//line views/vsprint/Edit.html:49
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Edit.html:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/Edit.html:49
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/Edit.html:49
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/Edit.html:49
}

//line views/vsprint/Edit.html:49
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/Edit.html:49
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/Edit.html:49
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/Edit.html:49
	qs422016 := string(qb422016.B)
//line views/vsprint/Edit.html:49
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/Edit.html:49
	return qs422016
//line views/vsprint/Edit.html:49
}
