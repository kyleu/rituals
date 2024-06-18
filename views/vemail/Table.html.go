// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vemail/Table.html:1
package vemail

//line views/vemail/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vemail/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vemail/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vemail/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models email.Emails, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vemail/Table.html:11
	qw422016.N().S(`
`)
//line views/vemail/Table.html:12
	prms := params.Sanitized("email", ps.Logger)

//line views/vemail/Table.html:12
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vemail/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "email", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vemail/Table.html:17
	qw422016.N().S(`
          `)
//line views/vemail/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "email", "recipients", "Recipients", "Comma-separated list of values", prms, ps.URI, ps)
//line views/vemail/Table.html:18
	qw422016.N().S(`
          `)
//line views/vemail/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "email", "subject", "Subject", "String text", prms, ps.URI, ps)
//line views/vemail/Table.html:19
	qw422016.N().S(`
          `)
//line views/vemail/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "email", "data", "Data", "JSON object", prms, ps.URI, ps)
//line views/vemail/Table.html:20
	qw422016.N().S(`
          `)
//line views/vemail/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "email", "plain", "Plain", "String text", prms, ps.URI, ps)
//line views/vemail/Table.html:21
	qw422016.N().S(`
          `)
//line views/vemail/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "email", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vemail/Table.html:22
	qw422016.N().S(`
          `)
//line views/vemail/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "email", "status", "Status", "String text", prms, ps.URI, ps)
//line views/vemail/Table.html:23
	qw422016.N().S(`
          `)
//line views/vemail/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "email", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vemail/Table.html:24
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vemail/Table.html:28
	for _, model := range models {
//line views/vemail/Table.html:28
		qw422016.N().S(`        <tr>
          <td><a href="/admin/db/email/`)
//line views/vemail/Table.html:30
		view.StreamUUID(qw422016, &model.ID)
//line views/vemail/Table.html:30
		qw422016.N().S(`">`)
//line views/vemail/Table.html:30
		view.StreamUUID(qw422016, &model.ID)
//line views/vemail/Table.html:30
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vemail/Table.html:31
		view.StreamStringArray(qw422016, model.Recipients)
//line views/vemail/Table.html:31
		qw422016.N().S(`</td>
          <td>`)
//line views/vemail/Table.html:32
		view.StreamString(qw422016, model.Subject)
//line views/vemail/Table.html:32
		qw422016.N().S(`</td>
          <td>`)
//line views/vemail/Table.html:33
		components.StreamJSON(qw422016, model.Data)
//line views/vemail/Table.html:33
		qw422016.N().S(`</td>
          <td>`)
//line views/vemail/Table.html:34
		view.StreamString(qw422016, model.Plain)
//line views/vemail/Table.html:34
		qw422016.N().S(`</td>
          <td class="nowrap">
            `)
//line views/vemail/Table.html:36
		view.StreamUUID(qw422016, &model.UserID)
//line views/vemail/Table.html:36
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vemail/Table.html:36
			qw422016.N().S(` (`)
//line views/vemail/Table.html:36
			qw422016.E().S(x.TitleString())
//line views/vemail/Table.html:36
			qw422016.N().S(`)`)
//line views/vemail/Table.html:36
		}
//line views/vemail/Table.html:36
		qw422016.N().S(`
            <a title="User" href="`)
//line views/vemail/Table.html:37
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vemail/Table.html:37
		qw422016.N().S(`">`)
//line views/vemail/Table.html:37
		components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vemail/Table.html:37
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vemail/Table.html:39
		view.StreamString(qw422016, model.Status)
//line views/vemail/Table.html:39
		qw422016.N().S(`</td>
          <td>`)
//line views/vemail/Table.html:40
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vemail/Table.html:40
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vemail/Table.html:42
	}
//line views/vemail/Table.html:42
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vemail/Table.html:46
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vemail/Table.html:46
		qw422016.N().S(`  <hr />
  `)
//line views/vemail/Table.html:48
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vemail/Table.html:48
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vemail/Table.html:50
	}
//line views/vemail/Table.html:51
}

//line views/vemail/Table.html:51
func WriteTable(qq422016 qtio422016.Writer, models email.Emails, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vemail/Table.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vemail/Table.html:51
	StreamTable(qw422016, models, usersByUserID, params, as, ps)
//line views/vemail/Table.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vemail/Table.html:51
}

//line views/vemail/Table.html:51
func Table(models email.Emails, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vemail/Table.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vemail/Table.html:51
	WriteTable(qb422016, models, usersByUserID, params, as, ps)
//line views/vemail/Table.html:51
	qs422016 := string(qb422016.B)
//line views/vemail/Table.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vemail/Table.html:51
	return qs422016
//line views/vemail/Table.html:51
}
