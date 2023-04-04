// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/Detail.html:2
package vsprint

//line views/vsprint/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate"
	"github.com/kyleu/rituals/views/vretro"
	"github.com/kyleu/rituals/views/vsprint/vshistory"
	"github.com/kyleu/rituals/views/vsprint/vsmember"
	"github.com/kyleu/rituals/views/vsprint/vspermission"
	"github.com/kyleu/rituals/views/vstandup"
)

//line views/vsprint/Detail.html:24
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/Detail.html:24
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/Detail.html:24
type Detail struct {
	layout.Basic
	Model                          *sprint.Sprint
	TeamsByTeamID                  team.Teams
	Params                         filter.ParamSet
	RelEstimatesBySprintID         estimate.Estimates
	RelRetrosBySprintID            retro.Retros
	RelSprintHistoriesBySprintID   shistory.SprintHistories
	RelSprintMembersBySprintID     smember.SprintMembers
	RelSprintPermissionsBySprintID spermission.SprintPermissions
	RelStandupsBySprintID          standup.Standups
}

//line views/vsprint/Detail.html:37
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Detail.html:37
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-sprint"><button type="button">JSON</button></a>
      <a href="`)
//line views/vsprint/Detail.html:41
	qw422016.E().S(p.Model.WebPath())
//line views/vsprint/Detail.html:41
	qw422016.N().S(`/edit"><button>`)
//line views/vsprint/Detail.html:41
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vsprint/Detail.html:41
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vsprint/Detail.html:43
	components.StreamSVGRefIcon(qw422016, `sprint`, ps)
//line views/vsprint/Detail.html:43
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:43
	qw422016.E().S(p.Model.TitleString())
//line views/vsprint/Detail.html:43
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/sprint"><em>Sprint</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vsprint/Detail.html:49
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vsprint/Detail.html:49
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vsprint/Detail.html:53
	qw422016.E().S(p.Model.Slug)
//line views/vsprint/Detail.html:53
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vsprint/Detail.html:57
	qw422016.E().S(p.Model.Title)
//line views/vsprint/Detail.html:57
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Icon</th>
          <td>`)
//line views/vsprint/Detail.html:61
	qw422016.E().S(p.Model.Icon)
//line views/vsprint/Detail.html:61
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [new, active, complete]">Status</th>
          <td>`)
//line views/vsprint/Detail.html:65
	qw422016.E().V(p.Model.Status)
//line views/vsprint/Detail.html:65
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td class="nowrap">
            `)
//line views/vsprint/Detail.html:70
	components.StreamDisplayUUID(qw422016, p.Model.TeamID)
//line views/vsprint/Detail.html:70
	if p.Model.TeamID != nil {
//line views/vsprint/Detail.html:70
		if x := p.TeamsByTeamID.Get(*p.Model.TeamID); x != nil {
//line views/vsprint/Detail.html:70
			qw422016.N().S(` (`)
//line views/vsprint/Detail.html:70
			qw422016.E().S(x.TitleString())
//line views/vsprint/Detail.html:70
			qw422016.N().S(`)`)
//line views/vsprint/Detail.html:70
		}
//line views/vsprint/Detail.html:70
	}
//line views/vsprint/Detail.html:70
	qw422016.N().S(`
            `)
//line views/vsprint/Detail.html:71
	if p.Model.TeamID != nil {
//line views/vsprint/Detail.html:71
		qw422016.N().S(`<a title="Team" href="`)
//line views/vsprint/Detail.html:71
		qw422016.E().S(`/admin/db/team` + `/` + p.Model.TeamID.String())
//line views/vsprint/Detail.html:71
		qw422016.N().S(`">`)
//line views/vsprint/Detail.html:71
		components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vsprint/Detail.html:71
		qw422016.N().S(`</a>`)
//line views/vsprint/Detail.html:71
	}
//line views/vsprint/Detail.html:71
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Calendar date (optional)">Start Date</th>
          <td>`)
//line views/vsprint/Detail.html:76
	components.StreamDisplayTimestampDay(qw422016, p.Model.StartDate)
//line views/vsprint/Detail.html:76
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Calendar date (optional)">End Date</th>
          <td>`)
//line views/vsprint/Detail.html:80
	components.StreamDisplayTimestampDay(qw422016, p.Model.EndDate)
//line views/vsprint/Detail.html:80
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vsprint/Detail.html:84
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vsprint/Detail.html:84
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vsprint/Detail.html:88
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vsprint/Detail.html:88
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vsprint/Detail.html:95
	if len(p.RelEstimatesBySprintID) > 0 {
//line views/vsprint/Detail.html:95
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsprint/Detail.html:97
		components.StreamSVGRefIcon(qw422016, `estimate`, ps)
//line views/vsprint/Detail.html:97
		qw422016.N().S(` Related estimates by [sprint id]</h3>
    <div class="overflow clear">
      `)
//line views/vsprint/Detail.html:99
		vestimate.StreamTable(qw422016, p.RelEstimatesBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:99
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vsprint/Detail.html:102
	}
//line views/vsprint/Detail.html:103
	if len(p.RelRetrosBySprintID) > 0 {
//line views/vsprint/Detail.html:103
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsprint/Detail.html:105
		components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vsprint/Detail.html:105
		qw422016.N().S(` Related retros by [sprint id]</h3>
    <div class="overflow clear">
      `)
//line views/vsprint/Detail.html:107
		vretro.StreamTable(qw422016, p.RelRetrosBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:107
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vsprint/Detail.html:110
	}
//line views/vsprint/Detail.html:111
	if len(p.RelSprintHistoriesBySprintID) > 0 {
//line views/vsprint/Detail.html:111
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsprint/Detail.html:113
		components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vsprint/Detail.html:113
		qw422016.N().S(` Related histories by [sprint id]</h3>
    <div class="overflow clear">
      `)
//line views/vsprint/Detail.html:115
		vshistory.StreamTable(qw422016, p.RelSprintHistoriesBySprintID, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:115
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vsprint/Detail.html:118
	}
//line views/vsprint/Detail.html:119
	if len(p.RelSprintMembersBySprintID) > 0 {
//line views/vsprint/Detail.html:119
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsprint/Detail.html:121
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vsprint/Detail.html:121
		qw422016.N().S(` Related members by [sprint id]</h3>
    <div class="overflow clear">
      `)
//line views/vsprint/Detail.html:123
		vsmember.StreamTable(qw422016, p.RelSprintMembersBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:123
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vsprint/Detail.html:126
	}
//line views/vsprint/Detail.html:127
	if len(p.RelSprintPermissionsBySprintID) > 0 {
//line views/vsprint/Detail.html:127
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsprint/Detail.html:129
		components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vsprint/Detail.html:129
		qw422016.N().S(` Related permissions by [sprint id]</h3>
    <div class="overflow clear">
      `)
//line views/vsprint/Detail.html:131
		vspermission.StreamTable(qw422016, p.RelSprintPermissionsBySprintID, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:131
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vsprint/Detail.html:134
	}
//line views/vsprint/Detail.html:135
	if len(p.RelStandupsBySprintID) > 0 {
//line views/vsprint/Detail.html:135
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vsprint/Detail.html:137
		components.StreamSVGRefIcon(qw422016, `standup`, ps)
//line views/vsprint/Detail.html:137
		qw422016.N().S(` Related standups by [sprint id]</h3>
    <div class="overflow clear">
      `)
//line views/vsprint/Detail.html:139
		vstandup.StreamTable(qw422016, p.RelStandupsBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:139
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vsprint/Detail.html:142
	}
//line views/vsprint/Detail.html:142
	qw422016.N().S(`  `)
//line views/vsprint/Detail.html:143
	components.StreamJSONModal(qw422016, "sprint", "Sprint JSON", p.Model, 1)
//line views/vsprint/Detail.html:143
	qw422016.N().S(`
`)
//line views/vsprint/Detail.html:144
}

//line views/vsprint/Detail.html:144
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Detail.html:144
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/Detail.html:144
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/Detail.html:144
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/Detail.html:144
}

//line views/vsprint/Detail.html:144
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/Detail.html:144
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/Detail.html:144
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/Detail.html:144
	qs422016 := string(qb422016.B)
//line views/vsprint/Detail.html:144
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/Detail.html:144
	return qs422016
//line views/vsprint/Detail.html:144
}
