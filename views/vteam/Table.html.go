// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vteam/Table.html:2
package vteam

//line views/vteam/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vteam/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models team.Teams, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vteam/Table.html:11
	qw422016.N().S(`
`)
//line views/vteam/Table.html:12
	prms := params.Get("team", nil, ps.Logger).Sanitize("team")

//line views/vteam/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vteam/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "team", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vteam/Table.html:16
	qw422016.N().S(`
        `)
//line views/vteam/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "team", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vteam/Table.html:17
	qw422016.N().S(`
        `)
//line views/vteam/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "team", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vteam/Table.html:18
	qw422016.N().S(`
        `)
//line views/vteam/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "team", "status", "Status", "Available options: [new, active, complete, deleted]", prms, ps.URI, ps)
//line views/vteam/Table.html:19
	qw422016.N().S(`
        `)
//line views/vteam/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "team", "owner", "Owner", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vteam/Table.html:20
	qw422016.N().S(`
        `)
//line views/vteam/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "team", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vteam/Table.html:21
	qw422016.N().S(`
        `)
//line views/vteam/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "team", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vteam/Table.html:22
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vteam/Table.html:26
	for _, model := range models {
//line views/vteam/Table.html:26
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/team/`)
//line views/vteam/Table.html:28
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vteam/Table.html:28
		qw422016.N().S(`">`)
//line views/vteam/Table.html:28
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vteam/Table.html:28
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vteam/Table.html:29
		qw422016.E().S(model.Slug)
//line views/vteam/Table.html:29
		qw422016.N().S(`</td>
        <td><strong>`)
//line views/vteam/Table.html:30
		qw422016.E().S(model.Title)
//line views/vteam/Table.html:30
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vteam/Table.html:31
		qw422016.E().V(model.Status)
//line views/vteam/Table.html:31
		qw422016.N().S(`</td>
        <td>
          <div class="icon">`)
//line views/vteam/Table.html:33
		components.StreamDisplayUUID(qw422016, &model.Owner)
//line views/vteam/Table.html:33
		if x := users.Get(model.Owner); x != nil {
//line views/vteam/Table.html:33
			qw422016.N().S(` (`)
//line views/vteam/Table.html:33
			qw422016.E().S(x.TitleString())
//line views/vteam/Table.html:33
			qw422016.N().S(`)`)
//line views/vteam/Table.html:33
		}
//line views/vteam/Table.html:33
		qw422016.N().S(`</div>
          <a title="User" href="`)
//line views/vteam/Table.html:34
		qw422016.E().S(`/user` + `/` + model.Owner.String())
//line views/vteam/Table.html:34
		qw422016.N().S(`">`)
//line views/vteam/Table.html:34
		components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vteam/Table.html:34
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vteam/Table.html:36
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vteam/Table.html:36
		qw422016.N().S(`</td>
        <td>`)
//line views/vteam/Table.html:37
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vteam/Table.html:37
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vteam/Table.html:39
	}
//line views/vteam/Table.html:40
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vteam/Table.html:40
		qw422016.N().S(`      <tr>
        <td colspan="7">`)
//line views/vteam/Table.html:42
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vteam/Table.html:42
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vteam/Table.html:44
	}
//line views/vteam/Table.html:44
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vteam/Table.html:47
}

//line views/vteam/Table.html:47
func WriteTable(qq422016 qtio422016.Writer, models team.Teams, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vteam/Table.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/Table.html:47
	StreamTable(qw422016, models, users, params, as, ps)
//line views/vteam/Table.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/Table.html:47
}

//line views/vteam/Table.html:47
func Table(models team.Teams, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vteam/Table.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/Table.html:47
	WriteTable(qb422016, models, users, params, as, ps)
//line views/vteam/Table.html:47
	qs422016 := string(qb422016.B)
//line views/vteam/Table.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/Table.html:47
	return qs422016
//line views/vteam/Table.html:47
}