// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vepermission/Table.html:2
package vepermission

//line views/vestimate/vepermission/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/views/components"
)

//line views/vestimate/vepermission/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vepermission/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vepermission/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models epermission.EstimatePermissions, estimates estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Table.html:11
	qw422016.N().S(`
`)
//line views/vestimate/vepermission/Table.html:12
	prms := params.Get("epermission", nil, ps.Logger).Sanitize("epermission")

//line views/vestimate/vepermission/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vestimate/vepermission/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "epermission", "estimate_id", "Estimate ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:16
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "epermission", "key", "Key", "String text", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:17
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "epermission", "value", "Value", "String text", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:18
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "epermission", "access", "Access", "String text", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:19
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "epermission", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:20
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vestimate/vepermission/Table.html:24
	for _, model := range models {
//line views/vestimate/vepermission/Table.html:24
		qw422016.N().S(`      <tr>
        <td class="nowrap">
          <a href="/admin/db/estimate/permission/`)
//line views/vestimate/vepermission/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:27
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:27
		qw422016.N().U(model.Key)
//line views/vestimate/vepermission/Table.html:27
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:27
		qw422016.N().U(model.Value)
//line views/vestimate/vepermission/Table.html:27
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:27
		if x := estimates.Get(model.EstimateID); x != nil {
//line views/vestimate/vepermission/Table.html:27
			qw422016.N().S(` (`)
//line views/vestimate/vepermission/Table.html:27
			qw422016.E().S(x.TitleString())
//line views/vestimate/vepermission/Table.html:27
			qw422016.N().S(`)`)
//line views/vestimate/vepermission/Table.html:27
		}
//line views/vestimate/vepermission/Table.html:27
		qw422016.N().S(`</a>
          <a title="Estimate" href="`)
//line views/vestimate/vepermission/Table.html:28
		qw422016.E().S(`/estimate` + `/` + model.EstimateID.String())
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:28
		components.StreamSVGRef(qw422016, "estimate", 18, 18, "", ps)
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().S(`</a>
        </td>
        <td><a href="/admin/db/estimate/permission/`)
//line views/vestimate/vepermission/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:30
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:30
		qw422016.N().U(model.Key)
//line views/vestimate/vepermission/Table.html:30
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:30
		qw422016.N().U(model.Value)
//line views/vestimate/vepermission/Table.html:30
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:30
		qw422016.E().S(model.Key)
//line views/vestimate/vepermission/Table.html:30
		qw422016.N().S(`</a></td>
        <td><a href="/admin/db/estimate/permission/`)
//line views/vestimate/vepermission/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().U(model.Key)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().U(model.Value)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:31
		qw422016.E().S(model.Value)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vestimate/vepermission/Table.html:32
		qw422016.E().S(model.Access)
//line views/vestimate/vepermission/Table.html:32
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/vepermission/Table.html:33
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vestimate/vepermission/Table.html:33
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vepermission/Table.html:35
	}
//line views/vestimate/vepermission/Table.html:36
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/vepermission/Table.html:36
		qw422016.N().S(`      <tr>
        <td colspan="5">`)
//line views/vestimate/vepermission/Table.html:38
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/vepermission/Table.html:38
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vepermission/Table.html:40
	}
//line views/vestimate/vepermission/Table.html:40
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vestimate/vepermission/Table.html:43
}

//line views/vestimate/vepermission/Table.html:43
func WriteTable(qq422016 qtio422016.Writer, models epermission.EstimatePermissions, estimates estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Table.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vepermission/Table.html:43
	StreamTable(qw422016, models, estimates, params, as, ps)
//line views/vestimate/vepermission/Table.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vepermission/Table.html:43
}

//line views/vestimate/vepermission/Table.html:43
func Table(models epermission.EstimatePermissions, estimates estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vepermission/Table.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vepermission/Table.html:43
	WriteTable(qb422016, models, estimates, params, as, ps)
//line views/vestimate/vepermission/Table.html:43
	qs422016 := string(qb422016.B)
//line views/vestimate/vepermission/Table.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vepermission/Table.html:43
	return qs422016
//line views/vestimate/vepermission/Table.html:43
}
