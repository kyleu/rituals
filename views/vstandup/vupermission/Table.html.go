// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vupermission/Table.html:2
package vupermission

//line views/vstandup/vupermission/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/views/components"
)

//line views/vstandup/vupermission/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vupermission/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vupermission/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models upermission.StandupPermissions, standups standup.Standups, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Table.html:11
	qw422016.N().S(`
`)
//line views/vstandup/vupermission/Table.html:12
	prms := params.Get("upermission", nil, ps.Logger).Sanitize("upermission")

//line views/vstandup/vupermission/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vstandup/vupermission/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "upermission", "standup_id", "Standup ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/vupermission/Table.html:16
	qw422016.N().S(`
        `)
//line views/vstandup/vupermission/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "upermission", "k", "K", "String text", prms, ps.URI, ps)
//line views/vstandup/vupermission/Table.html:17
	qw422016.N().S(`
        `)
//line views/vstandup/vupermission/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "upermission", "v", "V", "String text", prms, ps.URI, ps)
//line views/vstandup/vupermission/Table.html:18
	qw422016.N().S(`
        `)
//line views/vstandup/vupermission/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "upermission", "access", "Access", "String text", prms, ps.URI, ps)
//line views/vstandup/vupermission/Table.html:19
	qw422016.N().S(`
        `)
//line views/vstandup/vupermission/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "upermission", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vstandup/vupermission/Table.html:20
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vstandup/vupermission/Table.html:24
	for _, model := range models {
//line views/vstandup/vupermission/Table.html:24
		qw422016.N().S(`      <tr>
        <td>
          <div class="icon"><a href="/admin/db/standup/permission/`)
//line views/vstandup/vupermission/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.StandupID)
//line views/vstandup/vupermission/Table.html:27
		qw422016.N().S(`/`)
//line views/vstandup/vupermission/Table.html:27
		qw422016.N().U(model.K)
//line views/vstandup/vupermission/Table.html:27
		qw422016.N().S(`/`)
//line views/vstandup/vupermission/Table.html:27
		qw422016.N().U(model.V)
//line views/vstandup/vupermission/Table.html:27
		qw422016.N().S(`">`)
//line views/vstandup/vupermission/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.StandupID)
//line views/vstandup/vupermission/Table.html:27
		if x := standups.Get(model.StandupID); x != nil {
//line views/vstandup/vupermission/Table.html:27
			qw422016.N().S(` (`)
//line views/vstandup/vupermission/Table.html:27
			qw422016.E().S(x.TitleString())
//line views/vstandup/vupermission/Table.html:27
			qw422016.N().S(`)`)
//line views/vstandup/vupermission/Table.html:27
		}
//line views/vstandup/vupermission/Table.html:27
		qw422016.N().S(`</a></div>
          <a title="Standup" href="`)
//line views/vstandup/vupermission/Table.html:28
		qw422016.E().S(`/standup` + `/` + model.StandupID.String())
//line views/vstandup/vupermission/Table.html:28
		qw422016.N().S(`">`)
//line views/vstandup/vupermission/Table.html:28
		components.StreamSVGRefIcon(qw422016, "shoe-prints", ps)
//line views/vstandup/vupermission/Table.html:28
		qw422016.N().S(`</a>
        </td>
        <td><a href="/admin/db/standup/permission/`)
//line views/vstandup/vupermission/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.StandupID)
//line views/vstandup/vupermission/Table.html:30
		qw422016.N().S(`/`)
//line views/vstandup/vupermission/Table.html:30
		qw422016.N().U(model.K)
//line views/vstandup/vupermission/Table.html:30
		qw422016.N().S(`/`)
//line views/vstandup/vupermission/Table.html:30
		qw422016.N().U(model.V)
//line views/vstandup/vupermission/Table.html:30
		qw422016.N().S(`">`)
//line views/vstandup/vupermission/Table.html:30
		qw422016.E().S(model.K)
//line views/vstandup/vupermission/Table.html:30
		qw422016.N().S(`</a></td>
        <td><a href="/admin/db/standup/permission/`)
//line views/vstandup/vupermission/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.StandupID)
//line views/vstandup/vupermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vstandup/vupermission/Table.html:31
		qw422016.N().U(model.K)
//line views/vstandup/vupermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vstandup/vupermission/Table.html:31
		qw422016.N().U(model.V)
//line views/vstandup/vupermission/Table.html:31
		qw422016.N().S(`">`)
//line views/vstandup/vupermission/Table.html:31
		qw422016.E().S(model.V)
//line views/vstandup/vupermission/Table.html:31
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vstandup/vupermission/Table.html:32
		qw422016.E().S(model.Access)
//line views/vstandup/vupermission/Table.html:32
		qw422016.N().S(`</td>
        <td>`)
//line views/vstandup/vupermission/Table.html:33
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vstandup/vupermission/Table.html:33
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/vupermission/Table.html:35
	}
//line views/vstandup/vupermission/Table.html:36
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vstandup/vupermission/Table.html:36
		qw422016.N().S(`      <tr>
        <td colspan="5">`)
//line views/vstandup/vupermission/Table.html:38
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vstandup/vupermission/Table.html:38
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vstandup/vupermission/Table.html:40
	}
//line views/vstandup/vupermission/Table.html:40
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vstandup/vupermission/Table.html:43
}

//line views/vstandup/vupermission/Table.html:43
func WriteTable(qq422016 qtio422016.Writer, models upermission.StandupPermissions, standups standup.Standups, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Table.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vupermission/Table.html:43
	StreamTable(qw422016, models, standups, params, as, ps)
//line views/vstandup/vupermission/Table.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vupermission/Table.html:43
}

//line views/vstandup/vupermission/Table.html:43
func Table(models upermission.StandupPermissions, standups standup.Standups, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vupermission/Table.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vupermission/Table.html:43
	WriteTable(qb422016, models, standups, params, as, ps)
//line views/vstandup/vupermission/Table.html:43
	qs422016 := string(qb422016.B)
//line views/vstandup/vupermission/Table.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vupermission/Table.html:43
	return qs422016
//line views/vstandup/vupermission/Table.html:43
}
