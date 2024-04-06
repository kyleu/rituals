// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vteam/vthistory/Table.html:2
package vthistory

//line views/vteam/vthistory/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vteam/vthistory/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/vthistory/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/vthistory/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models thistory.TeamHistories, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vteam/vthistory/Table.html:12
	qw422016.N().S(`
`)
//line views/vteam/vthistory/Table.html:13
	prms := params.Sanitized("thistory", ps.Logger)

//line views/vteam/vthistory/Table.html:13
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vteam/vthistory/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "thistory", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vteam/vthistory/Table.html:18
	qw422016.N().S(`
          `)
//line views/vteam/vthistory/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "thistory", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vteam/vthistory/Table.html:19
	qw422016.N().S(`
          `)
//line views/vteam/vthistory/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "thistory", "team_name", "Team Name", "String text", prms, ps.URI, ps)
//line views/vteam/vthistory/Table.html:20
	qw422016.N().S(`
          `)
//line views/vteam/vthistory/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "thistory", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vteam/vthistory/Table.html:21
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vteam/vthistory/Table.html:25
	for _, model := range models {
//line views/vteam/vthistory/Table.html:25
		qw422016.N().S(`        <tr>
          <td><a href="/admin/db/team/history/`)
//line views/vteam/vthistory/Table.html:27
		qw422016.N().U(model.Slug)
//line views/vteam/vthistory/Table.html:27
		qw422016.N().S(`">`)
//line views/vteam/vthistory/Table.html:27
		view.StreamString(qw422016, model.Slug)
//line views/vteam/vthistory/Table.html:27
		qw422016.N().S(`</a></td>
          <td class="nowrap">
            `)
//line views/vteam/vthistory/Table.html:29
		view.StreamUUID(qw422016, &model.TeamID)
//line views/vteam/vthistory/Table.html:29
		if x := teamsByTeamID.Get(model.TeamID); x != nil {
//line views/vteam/vthistory/Table.html:29
			qw422016.N().S(` (`)
//line views/vteam/vthistory/Table.html:29
			qw422016.E().S(x.TitleString())
//line views/vteam/vthistory/Table.html:29
			qw422016.N().S(`)`)
//line views/vteam/vthistory/Table.html:29
		}
//line views/vteam/vthistory/Table.html:29
		qw422016.N().S(`
            <a title="Team" href="`)
//line views/vteam/vthistory/Table.html:30
		qw422016.E().S(`/admin/db/team` + `/` + model.TeamID.String())
//line views/vteam/vthistory/Table.html:30
		qw422016.N().S(`">`)
//line views/vteam/vthistory/Table.html:30
		components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vteam/vthistory/Table.html:30
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vteam/vthistory/Table.html:32
		view.StreamString(qw422016, model.TeamName)
//line views/vteam/vthistory/Table.html:32
		qw422016.N().S(`</td>
          <td>`)
//line views/vteam/vthistory/Table.html:33
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vteam/vthistory/Table.html:33
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vteam/vthistory/Table.html:35
	}
//line views/vteam/vthistory/Table.html:35
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vteam/vthistory/Table.html:39
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vteam/vthistory/Table.html:39
		qw422016.N().S(`  <hr />
  `)
//line views/vteam/vthistory/Table.html:41
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vteam/vthistory/Table.html:41
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vteam/vthistory/Table.html:43
	}
//line views/vteam/vthistory/Table.html:44
}

//line views/vteam/vthistory/Table.html:44
func WriteTable(qq422016 qtio422016.Writer, models thistory.TeamHistories, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vteam/vthistory/Table.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/vthistory/Table.html:44
	StreamTable(qw422016, models, teamsByTeamID, params, as, ps)
//line views/vteam/vthistory/Table.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/vthistory/Table.html:44
}

//line views/vteam/vthistory/Table.html:44
func Table(models thistory.TeamHistories, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vteam/vthistory/Table.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/vthistory/Table.html:44
	WriteTable(qb422016, models, teamsByTeamID, params, as, ps)
//line views/vteam/vthistory/Table.html:44
	qs422016 := string(qb422016.B)
//line views/vteam/vthistory/Table.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/vthistory/Table.html:44
	return qs422016
//line views/vteam/vthistory/Table.html:44
}
