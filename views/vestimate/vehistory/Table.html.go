// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vehistory/Table.html:2
package vehistory

//line views/vestimate/vehistory/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/views/components"
)

//line views/vestimate/vehistory/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vehistory/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vehistory/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models ehistory.EstimateHistories, estimates estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vehistory/Table.html:11
	qw422016.N().S(`
`)
//line views/vestimate/vehistory/Table.html:12
	prms := params.Get("ehistory", nil, ps.Logger).Sanitize("ehistory")

//line views/vestimate/vehistory/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vestimate/vehistory/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "ehistory", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vestimate/vehistory/Table.html:16
	qw422016.N().S(`
        `)
//line views/vestimate/vehistory/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "ehistory", "estimate_id", "Estimate ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vehistory/Table.html:17
	qw422016.N().S(`
        `)
//line views/vestimate/vehistory/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "ehistory", "estimate_name", "Estimate Name", "String text", prms, ps.URI, ps)
//line views/vestimate/vehistory/Table.html:18
	qw422016.N().S(`
        `)
//line views/vestimate/vehistory/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "ehistory", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/vehistory/Table.html:19
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vestimate/vehistory/Table.html:23
	for _, model := range models {
//line views/vestimate/vehistory/Table.html:23
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/estimate/history/`)
//line views/vestimate/vehistory/Table.html:25
		qw422016.N().U(model.Slug)
//line views/vestimate/vehistory/Table.html:25
		qw422016.N().S(`">`)
//line views/vestimate/vehistory/Table.html:25
		qw422016.E().S(model.Slug)
//line views/vestimate/vehistory/Table.html:25
		qw422016.N().S(`</a></td>
        <td>
          <div class="icon">`)
//line views/vestimate/vehistory/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vehistory/Table.html:27
		if x := estimates.Get(model.EstimateID); x != nil {
//line views/vestimate/vehistory/Table.html:27
			qw422016.N().S(` (`)
//line views/vestimate/vehistory/Table.html:27
			qw422016.E().S(x.TitleString())
//line views/vestimate/vehistory/Table.html:27
			qw422016.N().S(`)`)
//line views/vestimate/vehistory/Table.html:27
		}
//line views/vestimate/vehistory/Table.html:27
		qw422016.N().S(`</div>
          <a title="Estimate" href="`)
//line views/vestimate/vehistory/Table.html:28
		qw422016.E().S(`/estimate` + `/` + model.EstimateID.String())
//line views/vestimate/vehistory/Table.html:28
		qw422016.N().S(`">`)
//line views/vestimate/vehistory/Table.html:28
		components.StreamSVGRefIcon(qw422016, "estimate", ps)
//line views/vestimate/vehistory/Table.html:28
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vestimate/vehistory/Table.html:30
		qw422016.E().S(model.EstimateName)
//line views/vestimate/vehistory/Table.html:30
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/vehistory/Table.html:31
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vestimate/vehistory/Table.html:31
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vehistory/Table.html:33
	}
//line views/vestimate/vehistory/Table.html:34
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/vehistory/Table.html:34
		qw422016.N().S(`      <tr>
        <td colspan="4">`)
//line views/vestimate/vehistory/Table.html:36
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/vehistory/Table.html:36
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vehistory/Table.html:38
	}
//line views/vestimate/vehistory/Table.html:38
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vestimate/vehistory/Table.html:41
}

//line views/vestimate/vehistory/Table.html:41
func WriteTable(qq422016 qtio422016.Writer, models ehistory.EstimateHistories, estimates estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vehistory/Table.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vehistory/Table.html:41
	StreamTable(qw422016, models, estimates, params, as, ps)
//line views/vestimate/vehistory/Table.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vehistory/Table.html:41
}

//line views/vestimate/vehistory/Table.html:41
func Table(models ehistory.EstimateHistories, estimates estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vehistory/Table.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vehistory/Table.html:41
	WriteTable(qb422016, models, estimates, params, as, ps)
//line views/vestimate/vehistory/Table.html:41
	qs422016 := string(qb422016.B)
//line views/vestimate/vehistory/Table.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vehistory/Table.html:41
	return qs422016
//line views/vestimate/vehistory/Table.html:41
}