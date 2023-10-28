// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrmember/Edit.html:2
package vrmember

//line views/vretro/vrmember/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vrmember/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrmember/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrmember/Edit.html:11
type Edit struct {
	layout.Basic
	Model *rmember.RetroMember
	IsNew bool
}

//line views/vretro/vrmember/Edit.html:17
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Edit.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vretro/vrmember/Edit.html:19
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:19
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/vretro/vrmember/Edit.html:21
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vretro/vrmember/Edit.html:21
		qw422016.N().S(` New Member</h3>
    <form action="/admin/db/retro/member/_new" class="mt" method="post">
`)
//line views/vretro/vrmember/Edit.html:23
	} else {
//line views/vretro/vrmember/Edit.html:23
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vretro/vrmember/Edit.html:24
		qw422016.E().S(p.Model.WebPath())
//line views/vretro/vrmember/Edit.html:24
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete member [`)
//line views/vretro/vrmember/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vretro/vrmember/Edit.html:24
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vretro/vrmember/Edit.html:25
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vretro/vrmember/Edit.html:25
		qw422016.N().S(` Edit Member [`)
//line views/vretro/vrmember/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vretro/vrmember/Edit.html:25
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vretro/vrmember/Edit.html:27
	}
//line views/vretro/vrmember/Edit.html:27
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vretro/vrmember/Edit.html:30
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:30
		components.StreamTableInputUUID(qw422016, "retroID", "input-retroID", "Retro ID", &p.Model.RetroID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vrmember/Edit.html:30
	}
//line views/vretro/vrmember/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:31
	if p.IsNew {
//line views/vretro/vrmember/Edit.html:31
		components.StreamTableInputUUID(qw422016, "userID", "input-userID", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vretro/vrmember/Edit.html:31
	}
//line views/vretro/vrmember/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:32
	components.StreamTableInput(qw422016, "name", "", "Name", p.Model.Name, 5, "String text")
//line views/vretro/vrmember/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:33
	components.StreamTableInput(qw422016, "picture", "", "Picture", p.Model.Picture, 5, "URL in string form")
//line views/vretro/vrmember/Edit.html:33
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Edit.html:34
	components.StreamTableSelect(qw422016, "role", "", "Role", p.Model.Role.Key, enum.AllMemberStatuses.Keys(), enum.AllMemberStatuses.Strings(), 5, enum.AllMemberStatuses.Help())
//line views/vretro/vrmember/Edit.html:34
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
//line views/vretro/vrmember/Edit.html:46
}

//line views/vretro/vrmember/Edit.html:46
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Edit.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrmember/Edit.html:46
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vrmember/Edit.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrmember/Edit.html:46
}

//line views/vretro/vrmember/Edit.html:46
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrmember/Edit.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrmember/Edit.html:46
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vrmember/Edit.html:46
	qs422016 := string(qb422016.B)
//line views/vretro/vrmember/Edit.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrmember/Edit.html:46
	return qs422016
//line views/vretro/vrmember/Edit.html:46
}
