// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrpermission/Table.html:2
package vrpermission

//line views/vretro/vrpermission/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/views/components"
)

//line views/vretro/vrpermission/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrpermission/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrpermission/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models rpermission.RetroPermissions, retros retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrpermission/Table.html:11
	qw422016.N().S(`
`)
//line views/vretro/vrpermission/Table.html:12
	prms := params.Get("rpermission", nil, ps.Logger).Sanitize("rpermission")

//line views/vretro/vrpermission/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vretro/vrpermission/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "rpermission", "retro_id", "Retro ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vrpermission/Table.html:16
	qw422016.N().S(`
        `)
//line views/vretro/vrpermission/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "rpermission", "key", "Key", "String text", prms, ps.URI, ps)
//line views/vretro/vrpermission/Table.html:17
	qw422016.N().S(`
        `)
//line views/vretro/vrpermission/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "rpermission", "value", "Value", "String text", prms, ps.URI, ps)
//line views/vretro/vrpermission/Table.html:18
	qw422016.N().S(`
        `)
//line views/vretro/vrpermission/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "rpermission", "access", "Access", "String text", prms, ps.URI, ps)
//line views/vretro/vrpermission/Table.html:19
	qw422016.N().S(`
        `)
//line views/vretro/vrpermission/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "rpermission", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vretro/vrpermission/Table.html:20
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vretro/vrpermission/Table.html:24
	for _, model := range models {
//line views/vretro/vrpermission/Table.html:24
		qw422016.N().S(`      <tr>
        <td>
          <div class="icon"><a href="/admin/db/retro/permission/`)
//line views/vretro/vrpermission/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.RetroID)
//line views/vretro/vrpermission/Table.html:27
		qw422016.N().S(`/`)
//line views/vretro/vrpermission/Table.html:27
		qw422016.N().U(model.Key)
//line views/vretro/vrpermission/Table.html:27
		qw422016.N().S(`/`)
//line views/vretro/vrpermission/Table.html:27
		qw422016.N().U(model.Value)
//line views/vretro/vrpermission/Table.html:27
		qw422016.N().S(`">`)
//line views/vretro/vrpermission/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.RetroID)
//line views/vretro/vrpermission/Table.html:27
		if x := retros.Get(model.RetroID); x != nil {
//line views/vretro/vrpermission/Table.html:27
			qw422016.N().S(` (`)
//line views/vretro/vrpermission/Table.html:27
			qw422016.E().S(x.TitleString())
//line views/vretro/vrpermission/Table.html:27
			qw422016.N().S(`)`)
//line views/vretro/vrpermission/Table.html:27
		}
//line views/vretro/vrpermission/Table.html:27
		qw422016.N().S(`</a></div>
          <a title="Retro" href="`)
//line views/vretro/vrpermission/Table.html:28
		qw422016.E().S(`/retro` + `/` + model.RetroID.String())
//line views/vretro/vrpermission/Table.html:28
		qw422016.N().S(`">`)
//line views/vretro/vrpermission/Table.html:28
		components.StreamSVGRefIcon(qw422016, "retro", ps)
//line views/vretro/vrpermission/Table.html:28
		qw422016.N().S(`</a>
        </td>
        <td><a href="/admin/db/retro/permission/`)
//line views/vretro/vrpermission/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.RetroID)
//line views/vretro/vrpermission/Table.html:30
		qw422016.N().S(`/`)
//line views/vretro/vrpermission/Table.html:30
		qw422016.N().U(model.Key)
//line views/vretro/vrpermission/Table.html:30
		qw422016.N().S(`/`)
//line views/vretro/vrpermission/Table.html:30
		qw422016.N().U(model.Value)
//line views/vretro/vrpermission/Table.html:30
		qw422016.N().S(`">`)
//line views/vretro/vrpermission/Table.html:30
		qw422016.E().S(model.Key)
//line views/vretro/vrpermission/Table.html:30
		qw422016.N().S(`</a></td>
        <td><a href="/admin/db/retro/permission/`)
//line views/vretro/vrpermission/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.RetroID)
//line views/vretro/vrpermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vretro/vrpermission/Table.html:31
		qw422016.N().U(model.Key)
//line views/vretro/vrpermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vretro/vrpermission/Table.html:31
		qw422016.N().U(model.Value)
//line views/vretro/vrpermission/Table.html:31
		qw422016.N().S(`">`)
//line views/vretro/vrpermission/Table.html:31
		qw422016.E().S(model.Value)
//line views/vretro/vrpermission/Table.html:31
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vretro/vrpermission/Table.html:32
		qw422016.E().S(model.Access)
//line views/vretro/vrpermission/Table.html:32
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vrpermission/Table.html:33
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vretro/vrpermission/Table.html:33
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrpermission/Table.html:35
	}
//line views/vretro/vrpermission/Table.html:36
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vretro/vrpermission/Table.html:36
		qw422016.N().S(`      <tr>
        <td colspan="5">`)
//line views/vretro/vrpermission/Table.html:38
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vretro/vrpermission/Table.html:38
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrpermission/Table.html:40
	}
//line views/vretro/vrpermission/Table.html:40
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vretro/vrpermission/Table.html:43
}

//line views/vretro/vrpermission/Table.html:43
func WriteTable(qq422016 qtio422016.Writer, models rpermission.RetroPermissions, retros retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrpermission/Table.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrpermission/Table.html:43
	StreamTable(qw422016, models, retros, params, as, ps)
//line views/vretro/vrpermission/Table.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrpermission/Table.html:43
}

//line views/vretro/vrpermission/Table.html:43
func Table(models rpermission.RetroPermissions, retros retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrpermission/Table.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrpermission/Table.html:43
	WriteTable(qb422016, models, retros, params, as, ps)
//line views/vretro/vrpermission/Table.html:43
	qs422016 := string(qb422016.B)
//line views/vretro/vrpermission/Table.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrpermission/Table.html:43
	return qs422016
//line views/vretro/vrpermission/Table.html:43
}
