// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vcomment/Table.html:2
package vcomment

//line views/vcomment/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vcomment/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcomment/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcomment/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models comment.Comments, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Table.html:11
	qw422016.N().S(`
`)
//line views/vcomment/Table.html:12
	prms := params.Get("comment", nil, ps.Logger).Sanitize("comment")

//line views/vcomment/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vcomment/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "comment", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vcomment/Table.html:16
	qw422016.N().S(`
        `)
//line views/vcomment/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "comment", "svc", "Svc", "Available options: [team, sprint, estimate, standup, retro, story, feedback, report]", prms, ps.URI, ps)
//line views/vcomment/Table.html:17
	qw422016.N().S(`
        `)
//line views/vcomment/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "comment", "model_id", "Model ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vcomment/Table.html:18
	qw422016.N().S(`
        `)
//line views/vcomment/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "comment", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vcomment/Table.html:19
	qw422016.N().S(`
        `)
//line views/vcomment/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "comment", "content", "Content", "String text", prms, ps.URI, ps)
//line views/vcomment/Table.html:20
	qw422016.N().S(`
        `)
//line views/vcomment/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "comment", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vcomment/Table.html:21
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vcomment/Table.html:25
	for _, model := range models {
//line views/vcomment/Table.html:25
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/comment/`)
//line views/vcomment/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vcomment/Table.html:27
		qw422016.N().S(`">`)
//line views/vcomment/Table.html:27
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vcomment/Table.html:27
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vcomment/Table.html:28
		qw422016.E().V(model.Svc)
//line views/vcomment/Table.html:28
		qw422016.N().S(`</td>
        <td>`)
//line views/vcomment/Table.html:29
		components.StreamDisplayUUID(qw422016, &model.ModelID)
//line views/vcomment/Table.html:29
		qw422016.N().S(`</td>
        <td class="nowrap">
          `)
//line views/vcomment/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vcomment/Table.html:31
		if x := users.Get(model.UserID); x != nil {
//line views/vcomment/Table.html:31
			qw422016.N().S(` (`)
//line views/vcomment/Table.html:31
			qw422016.E().S(x.TitleString())
//line views/vcomment/Table.html:31
			qw422016.N().S(`)`)
//line views/vcomment/Table.html:31
		}
//line views/vcomment/Table.html:31
		qw422016.N().S(`
          <a title="User" href="`)
//line views/vcomment/Table.html:32
		qw422016.E().S(`/user` + `/` + model.UserID.String())
//line views/vcomment/Table.html:32
		qw422016.N().S(`">`)
//line views/vcomment/Table.html:32
		components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vcomment/Table.html:32
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vcomment/Table.html:34
		qw422016.E().S(model.Content)
//line views/vcomment/Table.html:34
		qw422016.N().S(`</td>
        <td>`)
//line views/vcomment/Table.html:35
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vcomment/Table.html:35
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vcomment/Table.html:37
	}
//line views/vcomment/Table.html:38
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vcomment/Table.html:38
		qw422016.N().S(`      <tr>
        <td colspan="6">`)
//line views/vcomment/Table.html:40
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vcomment/Table.html:40
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vcomment/Table.html:42
	}
//line views/vcomment/Table.html:42
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vcomment/Table.html:45
}

//line views/vcomment/Table.html:45
func WriteTable(qq422016 qtio422016.Writer, models comment.Comments, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Table.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcomment/Table.html:45
	StreamTable(qw422016, models, users, params, as, ps)
//line views/vcomment/Table.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vcomment/Table.html:45
}

//line views/vcomment/Table.html:45
func Table(models comment.Comments, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vcomment/Table.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcomment/Table.html:45
	WriteTable(qb422016, models, users, params, as, ps)
//line views/vcomment/Table.html:45
	qs422016 := string(qb422016.B)
//line views/vcomment/Table.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcomment/Table.html:45
	return qs422016
//line views/vcomment/Table.html:45
}
