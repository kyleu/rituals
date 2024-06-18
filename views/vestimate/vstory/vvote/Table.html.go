// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/vstory/vvote/Table.html:1
package vvote

//line views/vestimate/vstory/vvote/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vestimate/vstory/vvote/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/vvote/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/vvote/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models vote.Votes, storiesByStoryID story.Stories, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Table.html:12
	qw422016.N().S(`
`)
//line views/vestimate/vstory/vvote/Table.html:13
	prms := params.Sanitized("vote", ps.Logger)

//line views/vestimate/vstory/vvote/Table.html:13
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vestimate/vstory/vvote/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "vote", "story_id", "Story ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:18
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/vvote/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "vote", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:19
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/vvote/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "vote", "choice", "Choice", "String text", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:20
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/vvote/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "vote", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:21
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/vvote/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "vote", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:22
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vestimate/vstory/vvote/Table.html:26
	for _, model := range models {
//line views/vestimate/vstory/vvote/Table.html:26
		qw422016.N().S(`        <tr>
          <td class="nowrap">
            <a href="/admin/db/estimate/story/vote/`)
//line views/vestimate/vstory/vvote/Table.html:29
		view.StreamUUID(qw422016, &model.StoryID)
//line views/vestimate/vstory/vvote/Table.html:29
		qw422016.N().S(`/`)
//line views/vestimate/vstory/vvote/Table.html:29
		view.StreamUUID(qw422016, &model.UserID)
//line views/vestimate/vstory/vvote/Table.html:29
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:29
		view.StreamUUID(qw422016, &model.StoryID)
//line views/vestimate/vstory/vvote/Table.html:29
		if x := storiesByStoryID.Get(model.StoryID); x != nil {
//line views/vestimate/vstory/vvote/Table.html:29
			qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Table.html:29
			qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/vvote/Table.html:29
			qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Table.html:29
		}
//line views/vestimate/vstory/vvote/Table.html:29
		qw422016.N().S(`</a>
            <a title="Story" href="`)
//line views/vestimate/vstory/vvote/Table.html:30
		qw422016.E().S(`/admin/db/estimate/story` + `/` + model.StoryID.String())
//line views/vestimate/vstory/vvote/Table.html:30
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:30
		components.StreamSVGLink(qw422016, `story`, ps)
//line views/vestimate/vstory/vvote/Table.html:30
		qw422016.N().S(`</a>
          </td>
          <td class="nowrap">
            <a href="/admin/db/estimate/story/vote/`)
//line views/vestimate/vstory/vvote/Table.html:33
		view.StreamUUID(qw422016, &model.StoryID)
//line views/vestimate/vstory/vvote/Table.html:33
		qw422016.N().S(`/`)
//line views/vestimate/vstory/vvote/Table.html:33
		view.StreamUUID(qw422016, &model.UserID)
//line views/vestimate/vstory/vvote/Table.html:33
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:33
		view.StreamUUID(qw422016, &model.UserID)
//line views/vestimate/vstory/vvote/Table.html:33
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vestimate/vstory/vvote/Table.html:33
			qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Table.html:33
			qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/vvote/Table.html:33
			qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Table.html:33
		}
//line views/vestimate/vstory/vvote/Table.html:33
		qw422016.N().S(`</a>
            <a title="User" href="`)
//line views/vestimate/vstory/vvote/Table.html:34
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vestimate/vstory/vvote/Table.html:34
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:34
		components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vestimate/vstory/vvote/Table.html:34
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vestimate/vstory/vvote/Table.html:36
		view.StreamString(qw422016, model.Choice)
//line views/vestimate/vstory/vvote/Table.html:36
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/vstory/vvote/Table.html:37
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vestimate/vstory/vvote/Table.html:37
		qw422016.N().S(`</td>
          <td>`)
//line views/vestimate/vstory/vvote/Table.html:38
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vestimate/vstory/vvote/Table.html:38
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vestimate/vstory/vvote/Table.html:40
	}
//line views/vestimate/vstory/vvote/Table.html:40
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vestimate/vstory/vvote/Table.html:44
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/vstory/vvote/Table.html:44
		qw422016.N().S(`  <hr />
  `)
//line views/vestimate/vstory/vvote/Table.html:46
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/vstory/vvote/Table.html:46
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vestimate/vstory/vvote/Table.html:48
	}
//line views/vestimate/vstory/vvote/Table.html:49
}

//line views/vestimate/vstory/vvote/Table.html:49
func WriteTable(qq422016 qtio422016.Writer, models vote.Votes, storiesByStoryID story.Stories, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Table.html:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/vvote/Table.html:49
	StreamTable(qw422016, models, storiesByStoryID, usersByUserID, params, as, ps)
//line views/vestimate/vstory/vvote/Table.html:49
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/vvote/Table.html:49
}

//line views/vestimate/vstory/vvote/Table.html:49
func Table(models vote.Votes, storiesByStoryID story.Stories, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/vvote/Table.html:49
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/vvote/Table.html:49
	WriteTable(qb422016, models, storiesByStoryID, usersByUserID, params, as, ps)
//line views/vestimate/vstory/vvote/Table.html:49
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/vvote/Table.html:49
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/vvote/Table.html:49
	return qs422016
//line views/vestimate/vstory/vvote/Table.html:49
}
