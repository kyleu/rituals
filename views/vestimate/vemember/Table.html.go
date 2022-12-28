// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vemember/Table.html:2
package vemember

//line views/vestimate/vemember/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vestimate/vemember/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vemember/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vemember/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models emember.EstimateMembers, estimates estimate.Estimates, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vemember/Table.html:12
	qw422016.N().S(`
`)
//line views/vestimate/vemember/Table.html:13
	prms := params.Get("emember", nil, ps.Logger).Sanitize("emember")

//line views/vestimate/vemember/Table.html:13
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vestimate/vemember/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "emember", "estimate_id", "Estimate ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vemember/Table.html:17
	qw422016.N().S(`
        `)
//line views/vestimate/vemember/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "emember", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vemember/Table.html:18
	qw422016.N().S(`
        `)
//line views/vestimate/vemember/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "emember", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vestimate/vemember/Table.html:19
	qw422016.N().S(`
        `)
//line views/vestimate/vemember/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "emember", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vestimate/vemember/Table.html:20
	qw422016.N().S(`
        `)
//line views/vestimate/vemember/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "emember", "role", "Role", "Available options: [owner, member, observer]", prms, ps.URI, ps)
//line views/vestimate/vemember/Table.html:21
	qw422016.N().S(`
        `)
//line views/vestimate/vemember/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "emember", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/vemember/Table.html:22
	qw422016.N().S(`
        `)
//line views/vestimate/vemember/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "emember", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vestimate/vemember/Table.html:23
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vestimate/vemember/Table.html:27
	for _, model := range models {
//line views/vestimate/vemember/Table.html:27
		qw422016.N().S(`      <tr>
        <td class="nowrap">
          <a href="/admin/db/estimate/member/`)
//line views/vestimate/vemember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vemember/Table.html:30
		qw422016.N().S(`/`)
//line views/vestimate/vemember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vestimate/vemember/Table.html:30
		qw422016.N().S(`">`)
//line views/vestimate/vemember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vemember/Table.html:30
		if x := estimates.Get(model.EstimateID); x != nil {
//line views/vestimate/vemember/Table.html:30
			qw422016.N().S(` (`)
//line views/vestimate/vemember/Table.html:30
			qw422016.E().S(x.TitleString())
//line views/vestimate/vemember/Table.html:30
			qw422016.N().S(`)`)
//line views/vestimate/vemember/Table.html:30
		}
//line views/vestimate/vemember/Table.html:30
		qw422016.N().S(`</a>
          <a title="Estimate" href="`)
//line views/vestimate/vemember/Table.html:31
		qw422016.E().S(`/estimate` + `/` + model.EstimateID.String())
//line views/vestimate/vemember/Table.html:31
		qw422016.N().S(`">`)
//line views/vestimate/vemember/Table.html:31
		components.StreamSVGRef(qw422016, "estimate", 18, 18, "", ps)
//line views/vestimate/vemember/Table.html:31
		qw422016.N().S(`</a>
        </td>
        <td class="nowrap">
          <a href="/admin/db/estimate/member/`)
//line views/vestimate/vemember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.EstimateID)
//line views/vestimate/vemember/Table.html:34
		qw422016.N().S(`/`)
//line views/vestimate/vemember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vestimate/vemember/Table.html:34
		qw422016.N().S(`">`)
//line views/vestimate/vemember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vestimate/vemember/Table.html:34
		if x := users.Get(model.UserID); x != nil {
//line views/vestimate/vemember/Table.html:34
			qw422016.N().S(` (`)
//line views/vestimate/vemember/Table.html:34
			qw422016.E().S(x.TitleString())
//line views/vestimate/vemember/Table.html:34
			qw422016.N().S(`)`)
//line views/vestimate/vemember/Table.html:34
		}
//line views/vestimate/vemember/Table.html:34
		qw422016.N().S(`</a>
          <a title="User" href="`)
//line views/vestimate/vemember/Table.html:35
		qw422016.E().S(`/user` + `/` + model.UserID.String())
//line views/vestimate/vemember/Table.html:35
		qw422016.N().S(`">`)
//line views/vestimate/vemember/Table.html:35
		components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vestimate/vemember/Table.html:35
		qw422016.N().S(`</a>
        </td>
        <td><strong>`)
//line views/vestimate/vemember/Table.html:37
		qw422016.E().S(model.Name)
//line views/vestimate/vemember/Table.html:37
		qw422016.N().S(`</strong></td>
        <td><a href="`)
//line views/vestimate/vemember/Table.html:38
		qw422016.E().S(model.Picture)
//line views/vestimate/vemember/Table.html:38
		qw422016.N().S(`" target="_blank">`)
//line views/vestimate/vemember/Table.html:38
		qw422016.E().S(model.Picture)
//line views/vestimate/vemember/Table.html:38
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vestimate/vemember/Table.html:39
		qw422016.E().V(model.Role)
//line views/vestimate/vemember/Table.html:39
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/vemember/Table.html:40
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vestimate/vemember/Table.html:40
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/vemember/Table.html:41
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vestimate/vemember/Table.html:41
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vemember/Table.html:43
	}
//line views/vestimate/vemember/Table.html:44
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/vemember/Table.html:44
		qw422016.N().S(`      <tr>
        <td colspan="7">`)
//line views/vestimate/vemember/Table.html:46
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/vemember/Table.html:46
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vemember/Table.html:48
	}
//line views/vestimate/vemember/Table.html:48
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vestimate/vemember/Table.html:51
}

//line views/vestimate/vemember/Table.html:51
func WriteTable(qq422016 qtio422016.Writer, models emember.EstimateMembers, estimates estimate.Estimates, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vemember/Table.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vemember/Table.html:51
	StreamTable(qw422016, models, estimates, users, params, as, ps)
//line views/vestimate/vemember/Table.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vemember/Table.html:51
}

//line views/vestimate/vemember/Table.html:51
func Table(models emember.EstimateMembers, estimates estimate.Estimates, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vemember/Table.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vemember/Table.html:51
	WriteTable(qb422016, models, estimates, users, params, as, ps)
//line views/vestimate/vemember/Table.html:51
	qs422016 := string(qb422016.B)
//line views/vestimate/vemember/Table.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vemember/Table.html:51
	return qs422016
//line views/vestimate/vemember/Table.html:51
}
