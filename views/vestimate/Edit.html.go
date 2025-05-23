// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/Edit.html:1
package vestimate

//line views/vestimate/Edit.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/Edit.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/Edit.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/Edit.html:12
type Edit struct {
	layout.Basic
	Model *estimate.Estimate
	Paths []string
	IsNew bool
}

//line views/vestimate/Edit.html:19
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Edit.html:19
	qw422016.N().S(`
  <div class="card">
`)
//line views/vestimate/Edit.html:21
	if p.IsNew {
//line views/vestimate/Edit.html:21
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vestimate/Edit.html:23
		components.StreamSVGIcon(qw422016, `estimate`, ps)
//line views/vestimate/Edit.html:23
		qw422016.N().S(` New Estimate</h3>
`)
//line views/vestimate/Edit.html:24
	} else {
//line views/vestimate/Edit.html:24
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/vestimate/Edit.html:25
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/Edit.html:25
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete estimate [`)
//line views/vestimate/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vestimate/Edit.html:25
		qw422016.N().S(`]?"><button>`)
//line views/vestimate/Edit.html:25
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vestimate/Edit.html:25
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vestimate/Edit.html:26
		components.StreamSVGIcon(qw422016, `estimate`, ps)
//line views/vestimate/Edit.html:26
		qw422016.N().S(` Edit Estimate [`)
//line views/vestimate/Edit.html:26
		qw422016.E().S(p.Model.String())
//line views/vestimate/Edit.html:26
		qw422016.N().S(`]</h3>
`)
//line views/vestimate/Edit.html:27
	}
//line views/vestimate/Edit.html:27
	qw422016.N().S(`    <form action="`)
//line views/vestimate/Edit.html:28
	qw422016.E().S(util.Choose(p.IsNew, estimate.Route(p.Paths...)+`/_new`, p.Model.WebPath(p.Paths...)+`/edit`))
//line views/vestimate/Edit.html:28
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/vestimate/Edit.html:31
	if p.IsNew {
//line views/vestimate/Edit.html:31
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/Edit.html:31
	}
//line views/vestimate/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vestimate/Edit.html:32
	edit.StreamStringTable(qw422016, "slug", "", "Slug", p.Model.Slug, 5, "String text")
//line views/vestimate/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vestimate/Edit.html:33
	edit.StreamStringTable(qw422016, "title", "", "Title", p.Model.Title, 5, "String text")
//line views/vestimate/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vestimate/Edit.html:34
	edit.StreamStringTable(qw422016, "icon", "", "Icon", p.Model.Icon, 5, "String text")
//line views/vestimate/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vestimate/Edit.html:35
	edit.StreamSelectTable(qw422016, "status", "", "Status", p.Model.Status.Key, enum.AllSessionStatuses.Keys(), enum.AllSessionStatuses.Strings(), 5, enum.AllSessionStatuses.Help())
//line views/vestimate/Edit.html:35
	qw422016.N().S(`
          `)
//line views/vestimate/Edit.html:36
	edit.StreamUUIDTable(qw422016, "teamID", "input-teamID", "Team ID", p.Model.TeamID, 5, "UUID in format (00000000-0000-0000-0000-000000000000) (optional)")
//line views/vestimate/Edit.html:36
	qw422016.N().S(`
          `)
//line views/vestimate/Edit.html:37
	edit.StreamUUIDTable(qw422016, "sprintID", "input-sprintID", "Sprint ID", p.Model.SprintID, 5, "UUID in format (00000000-0000-0000-0000-000000000000) (optional)")
//line views/vestimate/Edit.html:37
	qw422016.N().S(`
          `)
//line views/vestimate/Edit.html:38
	edit.StreamTextareaTable(qw422016, "choices", "", "Choices", 8, util.ToJSON(p.Model.Choices), 5, "Comma-separated list of values")
//line views/vestimate/Edit.html:38
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-teamID"), "/admin/db/team?team.l=10", "q", (o) => (o["title"] || "[no title]") + " (" + o["id"] + ")", (o) => o["id"]);
      rituals.autocomplete(document.getElementById("input-sprintID"), "/admin/db/sprint?sprint.l=10", "q", (o) => (o["title"] || "[no title]") + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vestimate/Edit.html:50
}

//line views/vestimate/Edit.html:50
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Edit.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/Edit.html:50
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/Edit.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/Edit.html:50
}

//line views/vestimate/Edit.html:50
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/Edit.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/Edit.html:50
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/Edit.html:50
	qs422016 := string(qb422016.B)
//line views/vestimate/Edit.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/Edit.html:50
	return qs422016
//line views/vestimate/Edit.html:50
}
