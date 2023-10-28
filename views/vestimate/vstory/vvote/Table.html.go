// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vstory/vvote/Table.html:2
package vvote

//line views/vestimate/vstory/vvote/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
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
	prms := params.Get("vote", nil, ps.Logger).Sanitize("vote")

//line views/vestimate/vstory/vvote/Table.html:13
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vestimate/vstory/vvote/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "vote", "story_id", "Story ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:17
	qw422016.N().S(`
        `)
//line views/vestimate/vstory/vvote/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "vote", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:18
	qw422016.N().S(`
        `)
//line views/vestimate/vstory/vvote/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "vote", "choice", "Choice", "String text", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:19
	qw422016.N().S(`
        `)
//line views/vestimate/vstory/vvote/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "vote", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:20
	qw422016.N().S(`
        `)
//line views/vestimate/vstory/vvote/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "vote", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vestimate/vstory/vvote/Table.html:21
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vestimate/vstory/vvote/Table.html:25
	for _, model := range models {
//line views/vestimate/vstory/vvote/Table.html:25
		qw422016.N().S(`      <tr>
        <td class="nowrap">
          <a href="/admin/db/estimate/story/vote/`)
//line views/vestimate/vstory/vvote/Table.html:28
		components.StreamDisplayUUID(qw422016, &model.StoryID)
//line views/vestimate/vstory/vvote/Table.html:28
		qw422016.N().S(`/`)
//line views/vestimate/vstory/vvote/Table.html:28
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vestimate/vstory/vvote/Table.html:28
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:28
		components.StreamDisplayUUID(qw422016, &model.StoryID)
//line views/vestimate/vstory/vvote/Table.html:28
		if x := storiesByStoryID.Get(model.StoryID); x != nil {
//line views/vestimate/vstory/vvote/Table.html:28
			qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Table.html:28
			qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/vvote/Table.html:28
			qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Table.html:28
		}
//line views/vestimate/vstory/vvote/Table.html:28
		qw422016.N().S(`</a>
          <a title="Story" href="`)
//line views/vestimate/vstory/vvote/Table.html:29
		qw422016.E().S(`/admin/db/estimate/story` + `/` + model.StoryID.String())
//line views/vestimate/vstory/vvote/Table.html:29
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:29
		components.StreamSVGRef(qw422016, "story", 18, 18, "", ps)
//line views/vestimate/vstory/vvote/Table.html:29
		qw422016.N().S(`</a>
        </td>
        <td class="nowrap">
          <a href="/admin/db/estimate/story/vote/`)
//line views/vestimate/vstory/vvote/Table.html:32
		components.StreamDisplayUUID(qw422016, &model.StoryID)
//line views/vestimate/vstory/vvote/Table.html:32
		qw422016.N().S(`/`)
//line views/vestimate/vstory/vvote/Table.html:32
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vestimate/vstory/vvote/Table.html:32
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:32
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vestimate/vstory/vvote/Table.html:32
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vestimate/vstory/vvote/Table.html:32
			qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Table.html:32
			qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/vvote/Table.html:32
			qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Table.html:32
		}
//line views/vestimate/vstory/vvote/Table.html:32
		qw422016.N().S(`</a>
          <a title="User" href="`)
//line views/vestimate/vstory/vvote/Table.html:33
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vestimate/vstory/vvote/Table.html:33
		qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Table.html:33
		components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vestimate/vstory/vvote/Table.html:33
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vestimate/vstory/vvote/Table.html:35
		qw422016.E().S(model.Choice)
//line views/vestimate/vstory/vvote/Table.html:35
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/vstory/vvote/Table.html:36
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vestimate/vstory/vvote/Table.html:36
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/vstory/vvote/Table.html:37
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vestimate/vstory/vvote/Table.html:37
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vstory/vvote/Table.html:39
	}
//line views/vestimate/vstory/vvote/Table.html:40
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/vstory/vvote/Table.html:40
		qw422016.N().S(`      <tr>
        <td colspan="5">`)
//line views/vestimate/vstory/vvote/Table.html:42
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/vstory/vvote/Table.html:42
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vstory/vvote/Table.html:44
	}
//line views/vestimate/vstory/vvote/Table.html:44
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vestimate/vstory/vvote/Table.html:47
}

//line views/vestimate/vstory/vvote/Table.html:47
func WriteTable(qq422016 qtio422016.Writer, models vote.Votes, storiesByStoryID story.Stories, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Table.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/vvote/Table.html:47
	StreamTable(qw422016, models, storiesByStoryID, usersByUserID, params, as, ps)
//line views/vestimate/vstory/vvote/Table.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/vvote/Table.html:47
}

//line views/vestimate/vstory/vvote/Table.html:47
func Table(models vote.Votes, storiesByStoryID story.Stories, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/vvote/Table.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/vvote/Table.html:47
	WriteTable(qb422016, models, storiesByStoryID, usersByUserID, params, as, ps)
//line views/vestimate/vstory/vvote/Table.html:47
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/vvote/Table.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/vvote/Table.html:47
	return qs422016
//line views/vestimate/vstory/vvote/Table.html:47
}
