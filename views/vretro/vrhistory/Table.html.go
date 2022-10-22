// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrhistory/Table.html:2
package vrhistory

//line views/vretro/vrhistory/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/views/components"
)

//line views/vretro/vrhistory/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrhistory/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrhistory/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models rhistory.RetroHistories, retros retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Table.html:11
	qw422016.N().S(`
`)
//line views/vretro/vrhistory/Table.html:12
	prms := params.Get("rhistory", nil, ps.Logger).Sanitize("rhistory")

//line views/vretro/vrhistory/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vretro/vrhistory/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "rhistory", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:16
	qw422016.N().S(`
        `)
//line views/vretro/vrhistory/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "rhistory", "retro_id", "Retro ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:17
	qw422016.N().S(`
        `)
//line views/vretro/vrhistory/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "rhistory", "retro_name", "Retro Name", "String text", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:18
	qw422016.N().S(`
        `)
//line views/vretro/vrhistory/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "rhistory", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:19
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vretro/vrhistory/Table.html:23
	for _, model := range models {
//line views/vretro/vrhistory/Table.html:23
		qw422016.N().S(`      <tr>
        <td><a href="/retro/rhistory/`)
//line views/vretro/vrhistory/Table.html:25
		qw422016.N().U(model.Slug)
//line views/vretro/vrhistory/Table.html:25
		qw422016.N().S(`">`)
//line views/vretro/vrhistory/Table.html:25
		qw422016.E().S(model.Slug)
//line views/vretro/vrhistory/Table.html:25
		qw422016.N().S(`</a></td>
        <td>
          <div class="icon">`)
//line views/vretro/vrhistory/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.RetroID)
//line views/vretro/vrhistory/Table.html:27
		if x := retros.Get(model.RetroID); x != nil {
//line views/vretro/vrhistory/Table.html:27
			qw422016.N().S(` (`)
//line views/vretro/vrhistory/Table.html:27
			qw422016.E().S(x.TitleString())
//line views/vretro/vrhistory/Table.html:27
			qw422016.N().S(`)`)
//line views/vretro/vrhistory/Table.html:27
		}
//line views/vretro/vrhistory/Table.html:27
		qw422016.N().S(`</div>
          <a title="Retro" href="`)
//line views/vretro/vrhistory/Table.html:28
		qw422016.E().S(`/retro` + `/` + model.RetroID.String())
//line views/vretro/vrhistory/Table.html:28
		qw422016.N().S(`">`)
//line views/vretro/vrhistory/Table.html:28
		components.StreamSVGRefIcon(qw422016, "star", ps)
//line views/vretro/vrhistory/Table.html:28
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vretro/vrhistory/Table.html:30
		qw422016.E().S(model.RetroName)
//line views/vretro/vrhistory/Table.html:30
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vrhistory/Table.html:31
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vretro/vrhistory/Table.html:31
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrhistory/Table.html:33
	}
//line views/vretro/vrhistory/Table.html:34
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vretro/vrhistory/Table.html:34
		qw422016.N().S(`      <tr>
        <td colspan="4">`)
//line views/vretro/vrhistory/Table.html:36
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vretro/vrhistory/Table.html:36
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrhistory/Table.html:38
	}
//line views/vretro/vrhistory/Table.html:38
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vretro/vrhistory/Table.html:41
}

//line views/vretro/vrhistory/Table.html:41
func WriteTable(qq422016 qtio422016.Writer, models rhistory.RetroHistories, retros retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Table.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrhistory/Table.html:41
	StreamTable(qw422016, models, retros, params, as, ps)
//line views/vretro/vrhistory/Table.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrhistory/Table.html:41
}

//line views/vretro/vrhistory/Table.html:41
func Table(models rhistory.RetroHistories, retros retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrhistory/Table.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrhistory/Table.html:41
	WriteTable(qb422016, models, retros, params, as, ps)
//line views/vretro/vrhistory/Table.html:41
	qs422016 := string(qb422016.B)
//line views/vretro/vrhistory/Table.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrhistory/Table.html:41
	return qs422016
//line views/vretro/vrhistory/Table.html:41
}
