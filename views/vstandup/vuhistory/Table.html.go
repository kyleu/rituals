// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vuhistory/Table.html:2
package vuhistory

//line views/vstandup/vuhistory/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vstandup/vuhistory/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vuhistory/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vuhistory/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models uhistory.StandupHistories, standupsByStandupID standup.Standups, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vuhistory/Table.html:12
	qw422016.N().S(`
`)
//line views/vstandup/vuhistory/Table.html:13
	prms := params.Get("uhistory", nil, ps.Logger).Sanitize("uhistory")

//line views/vstandup/vuhistory/Table.html:13
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vstandup/vuhistory/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "uhistory", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vstandup/vuhistory/Table.html:17
	qw422016.N().S(`
        `)
//line views/vstandup/vuhistory/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "uhistory", "standup_id", "Standup ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/vuhistory/Table.html:18
	qw422016.N().S(`
        `)
//line views/vstandup/vuhistory/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "uhistory", "standup_name", "Standup Name", "String text", prms, ps.URI, ps)
//line views/vstandup/vuhistory/Table.html:19
	qw422016.N().S(`
        `)
//line views/vstandup/vuhistory/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "uhistory", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vstandup/vuhistory/Table.html:20
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vstandup/vuhistory/Table.html:24
	for _, model := range models {
//line views/vstandup/vuhistory/Table.html:24
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/standup/history/`)
//line views/vstandup/vuhistory/Table.html:26
		qw422016.N().U(model.Slug)
//line views/vstandup/vuhistory/Table.html:26
		qw422016.N().S(`">`)
//line views/vstandup/vuhistory/Table.html:26
		view.StreamString(qw422016, model.Slug)
//line views/vstandup/vuhistory/Table.html:26
		qw422016.N().S(`</a></td>
        <td class="nowrap">
          `)
//line views/vstandup/vuhistory/Table.html:28
		view.StreamUUID(qw422016, &model.StandupID)
//line views/vstandup/vuhistory/Table.html:28
		if x := standupsByStandupID.Get(model.StandupID); x != nil {
//line views/vstandup/vuhistory/Table.html:28
			qw422016.N().S(` (`)
//line views/vstandup/vuhistory/Table.html:28
			qw422016.E().S(x.TitleString())
//line views/vstandup/vuhistory/Table.html:28
			qw422016.N().S(`)`)
//line views/vstandup/vuhistory/Table.html:28
		}
//line views/vstandup/vuhistory/Table.html:28
		qw422016.N().S(`
          <a title="Standup" href="`)
//line views/vstandup/vuhistory/Table.html:29
		qw422016.E().S(`/admin/db/standup` + `/` + model.StandupID.String())
//line views/vstandup/vuhistory/Table.html:29
		qw422016.N().S(`">`)
//line views/vstandup/vuhistory/Table.html:29
		components.StreamSVGRef(qw422016, "standup", 18, 18, "", ps)
//line views/vstandup/vuhistory/Table.html:29
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vstandup/vuhistory/Table.html:31
		view.StreamString(qw422016, model.StandupName)
//line views/vstandup/vuhistory/Table.html:31
		qw422016.N().S(`</td>
        <td>`)
//line views/vstandup/vuhistory/Table.html:32
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vstandup/vuhistory/Table.html:32
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/vuhistory/Table.html:34
	}
//line views/vstandup/vuhistory/Table.html:35
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vstandup/vuhistory/Table.html:35
		qw422016.N().S(`      <tr>
        <td colspan="4">`)
//line views/vstandup/vuhistory/Table.html:37
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vstandup/vuhistory/Table.html:37
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/vuhistory/Table.html:39
	}
//line views/vstandup/vuhistory/Table.html:39
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vstandup/vuhistory/Table.html:42
}

//line views/vstandup/vuhistory/Table.html:42
func WriteTable(qq422016 qtio422016.Writer, models uhistory.StandupHistories, standupsByStandupID standup.Standups, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vuhistory/Table.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vuhistory/Table.html:42
	StreamTable(qw422016, models, standupsByStandupID, params, as, ps)
//line views/vstandup/vuhistory/Table.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vuhistory/Table.html:42
}

//line views/vstandup/vuhistory/Table.html:42
func Table(models uhistory.StandupHistories, standupsByStandupID standup.Standups, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vuhistory/Table.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vuhistory/Table.html:42
	WriteTable(qb422016, models, standupsByStandupID, params, as, ps)
//line views/vstandup/vuhistory/Table.html:42
	qs422016 := string(qb422016.B)
//line views/vstandup/vuhistory/Table.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vuhistory/Table.html:42
	return qs422016
//line views/vstandup/vuhistory/Table.html:42
}
