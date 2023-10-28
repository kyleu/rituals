// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/Edit.html:2
package vretro

//line views/vretro/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/Edit.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/Edit.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/Edit.html:12
type Edit struct {
	layout.Basic
	Model *retro.Retro
	IsNew bool
}

//line views/vretro/Edit.html:18
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/Edit.html:18
	qw422016.N().S(`
  <div class="card">
`)
//line views/vretro/Edit.html:20
	if p.IsNew {
//line views/vretro/Edit.html:20
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vretro/Edit.html:22
		components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vretro/Edit.html:22
		qw422016.N().S(` New Retro</h3>
    <form action="/admin/db/retro/_new" class="mt" method="post">
`)
//line views/vretro/Edit.html:24
	} else {
//line views/vretro/Edit.html:24
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vretro/Edit.html:25
		qw422016.E().S(p.Model.WebPath())
//line views/vretro/Edit.html:25
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete retro [`)
//line views/vretro/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vretro/Edit.html:25
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vretro/Edit.html:26
		components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vretro/Edit.html:26
		qw422016.N().S(` Edit Retro [`)
//line views/vretro/Edit.html:26
		qw422016.E().S(p.Model.String())
//line views/vretro/Edit.html:26
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vretro/Edit.html:28
	}
//line views/vretro/Edit.html:28
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vretro/Edit.html:31
	if p.IsNew {
//line views/vretro/Edit.html:31
		components.StreamTableInputUUID(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/Edit.html:31
	}
//line views/vretro/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vretro/Edit.html:32
	components.StreamTableInput(qw422016, "slug", "", "Slug", p.Model.Slug, 5, "String text")
//line views/vretro/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vretro/Edit.html:33
	components.StreamTableInput(qw422016, "title", "", "Title", p.Model.Title, 5, "String text")
//line views/vretro/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vretro/Edit.html:34
	components.StreamTableInput(qw422016, "icon", "", "Icon", p.Model.Icon, 5, "String text")
//line views/vretro/Edit.html:34
	qw422016.N().S(`
          `)
//line views/vretro/Edit.html:35
	components.StreamTableSelect(qw422016, "status", "", "Status", p.Model.Status.Key, enum.AllSessionStatuses.Keys(), enum.AllSessionStatuses.Strings(), 5, enum.AllSessionStatuses.Help())
//line views/vretro/Edit.html:35
	qw422016.N().S(`
          `)
//line views/vretro/Edit.html:36
	components.StreamTableInputUUID(qw422016, "teamID", "input-teamID", "Team ID", p.Model.TeamID, 5, "UUID in format (00000000-0000-0000-0000-000000000000) (optional)")
//line views/vretro/Edit.html:36
	qw422016.N().S(`
          `)
//line views/vretro/Edit.html:37
	components.StreamTableInputUUID(qw422016, "sprintID", "input-sprintID", "Sprint ID", p.Model.SprintID, 5, "UUID in format (00000000-0000-0000-0000-000000000000) (optional)")
//line views/vretro/Edit.html:37
	qw422016.N().S(`
          `)
//line views/vretro/Edit.html:38
	components.StreamTableTextarea(qw422016, "categories", "", "Categories", 8, util.ToJSON(p.Model.Categories), 5, "Comma-separated list of values")
//line views/vretro/Edit.html:38
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-teamID"), "/admin/db/team?team.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " (" + o["id"] + ")", (o) => o["id"]);
      rituals.autocomplete(document.getElementById("input-sprintID"), "/admin/db/sprint?sprint.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vretro/Edit.html:50
}

//line views/vretro/Edit.html:50
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/Edit.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/Edit.html:50
	p.StreamBody(qw422016, as, ps)
//line views/vretro/Edit.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/Edit.html:50
}

//line views/vretro/Edit.html:50
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/Edit.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/Edit.html:50
	p.WriteBody(qb422016, as, ps)
//line views/vretro/Edit.html:50
	qs422016 := string(qb422016.B)
//line views/vretro/Edit.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/Edit.html:50
	return qs422016
//line views/vretro/Edit.html:50
}
