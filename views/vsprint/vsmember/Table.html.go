// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/vsmember/Table.html:2
package vsmember

//line views/vsprint/vsmember/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vsprint/vsmember/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vsmember/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vsmember/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models smember.SprintMembers, sprintsBySprintID sprint.Sprints, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Table.html:13
	qw422016.N().S(`
`)
//line views/vsprint/vsmember/Table.html:14
	prms := params.Get("smember", nil, ps.Logger).Sanitize("smember")

//line views/vsprint/vsmember/Table.html:14
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vsprint/vsmember/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "smember", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:18
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "smember", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:19
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "smember", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:20
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "smember", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:21
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "smember", "role", "Role", enum.AllMemberStatuses.Help(), prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:22
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "smember", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:23
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "smember", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:24
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vsprint/vsmember/Table.html:28
	for _, model := range models {
//line views/vsprint/vsmember/Table.html:28
		qw422016.N().S(`      <tr>
        <td class="nowrap">
          <a href="/admin/db/sprint/member/`)
//line views/vsprint/vsmember/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.SprintID)
//line views/vsprint/vsmember/Table.html:31
		qw422016.N().S(`/`)
//line views/vsprint/vsmember/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vsprint/vsmember/Table.html:31
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.SprintID)
//line views/vsprint/vsmember/Table.html:31
		if x := sprintsBySprintID.Get(model.SprintID); x != nil {
//line views/vsprint/vsmember/Table.html:31
			qw422016.N().S(` (`)
//line views/vsprint/vsmember/Table.html:31
			qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Table.html:31
			qw422016.N().S(`)`)
//line views/vsprint/vsmember/Table.html:31
		}
//line views/vsprint/vsmember/Table.html:31
		qw422016.N().S(`</a>
          <a title="Sprint" href="`)
//line views/vsprint/vsmember/Table.html:32
		qw422016.E().S(`/admin/db/sprint` + `/` + model.SprintID.String())
//line views/vsprint/vsmember/Table.html:32
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:32
		components.StreamSVGRef(qw422016, "sprint", 18, 18, "", ps)
//line views/vsprint/vsmember/Table.html:32
		qw422016.N().S(`</a>
        </td>
        <td class="nowrap">
          <a href="/admin/db/sprint/member/`)
//line views/vsprint/vsmember/Table.html:35
		components.StreamDisplayUUID(qw422016, &model.SprintID)
//line views/vsprint/vsmember/Table.html:35
		qw422016.N().S(`/`)
//line views/vsprint/vsmember/Table.html:35
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vsprint/vsmember/Table.html:35
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:35
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vsprint/vsmember/Table.html:35
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vsprint/vsmember/Table.html:35
			qw422016.N().S(` (`)
//line views/vsprint/vsmember/Table.html:35
			qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Table.html:35
			qw422016.N().S(`)`)
//line views/vsprint/vsmember/Table.html:35
		}
//line views/vsprint/vsmember/Table.html:35
		qw422016.N().S(`</a>
          <a title="User" href="`)
//line views/vsprint/vsmember/Table.html:36
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vsprint/vsmember/Table.html:36
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:36
		components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vsprint/vsmember/Table.html:36
		qw422016.N().S(`</a>
        </td>
        <td><strong>`)
//line views/vsprint/vsmember/Table.html:38
		qw422016.E().S(model.Name)
//line views/vsprint/vsmember/Table.html:38
		qw422016.N().S(`</strong></td>
        <td><a href="`)
//line views/vsprint/vsmember/Table.html:39
		qw422016.E().S(model.Picture)
//line views/vsprint/vsmember/Table.html:39
		qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vsprint/vsmember/Table.html:39
		qw422016.E().S(model.Picture)
//line views/vsprint/vsmember/Table.html:39
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vsprint/vsmember/Table.html:40
		qw422016.E().S(model.Role.String())
//line views/vsprint/vsmember/Table.html:40
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/vsmember/Table.html:41
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vsprint/vsmember/Table.html:41
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/vsmember/Table.html:42
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vsprint/vsmember/Table.html:42
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vsprint/vsmember/Table.html:44
	}
//line views/vsprint/vsmember/Table.html:45
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vsprint/vsmember/Table.html:45
		qw422016.N().S(`      <tr>
        <td colspan="7">`)
//line views/vsprint/vsmember/Table.html:47
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vsprint/vsmember/Table.html:47
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vsprint/vsmember/Table.html:49
	}
//line views/vsprint/vsmember/Table.html:49
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vsprint/vsmember/Table.html:52
}

//line views/vsprint/vsmember/Table.html:52
func WriteTable(qq422016 qtio422016.Writer, models smember.SprintMembers, sprintsBySprintID sprint.Sprints, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Table.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vsmember/Table.html:52
	StreamTable(qw422016, models, sprintsBySprintID, usersByUserID, params, as, ps)
//line views/vsprint/vsmember/Table.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vsmember/Table.html:52
}

//line views/vsprint/vsmember/Table.html:52
func Table(models smember.SprintMembers, sprintsBySprintID sprint.Sprints, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vsmember/Table.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vsmember/Table.html:52
	WriteTable(qb422016, models, sprintsBySprintID, usersByUserID, params, as, ps)
//line views/vsprint/vsmember/Table.html:52
	qs422016 := string(qb422016.B)
//line views/vsprint/vsmember/Table.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vsmember/Table.html:52
	return qs422016
//line views/vsprint/vsmember/Table.html:52
}
