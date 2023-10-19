// Code generated by qtc from "EstimateWorkspaceStory.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:1
package vwestimate

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:1
import (
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:11
func StreamEstimateWorkspaceStories(qw422016 *qt422016.Writer, w *workspace.FullEstimate, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:11
	qw422016.N().S(`
  <table class="mt expanded">
    <thead>
      <tr>
        <th>Story</th>
        <th class="shrink">Author</th>
        <th class="shrink">Status</th>
        <th class="shrink">Score</th>
        <th class="shrink"></th>
      </tr>
    </thead>
    <tbody>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:23
	for _, s := range w.Stories {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:23
		qw422016.N().S(`      <tr class="story-row" id="story-row-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:24
		qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:24
		qw422016.N().S(`" data-idx="s.Idx">
        <td><a href="#modal-story-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.N().S(`"><div class="story-title">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.E().S(s.TitleString())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.N().S(`</div></a></td>
        <td class="story-author nowrap"><a href="#modal-member-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.E().S(s.UserID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.N().S(`"><em class="member-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.E().S(s.UserID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.N().S(`-name">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.E().S(w.UtilMembers.Name(&s.UserID))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.N().S(`</em></a></td>
        <td class="story-status">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:27
		qw422016.E().S(s.Status.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:27
		qw422016.N().S(`</td>
        <td class="story-final-vote">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:28
		qw422016.E().S(s.FinalVoteSafe())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:28
		qw422016.N().S(`</td>
        <td>`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:29
		vwutil.StreamComments(qw422016, enum.ModelServiceStory, s.ID, s.TitleString(), w.Comments, w.UtilMembers, "member-icon", ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:29
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:31
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:31
	qw422016.N().S(`    </tbody>
  </table>
  <div id="story-modals">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:35
	for _, s := range w.Stories {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:36
		if ps.Profile.ID == s.UserID {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:36
			qw422016.N().S(`    `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:37
			StreamEstimateWorkspaceStoryModalEdit(qw422016, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:37
			qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:38
		}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:38
		qw422016.N().S(`    `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
		StreamEstimateWorkspaceStoryModal(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
		qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:40
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:40
	qw422016.N().S(`  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
func WriteEstimateWorkspaceStories(qq422016 qtio422016.Writer, w *workspace.FullEstimate, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	StreamEstimateWorkspaceStories(qw422016, w, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
func EstimateWorkspaceStories(w *workspace.FullEstimate, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	WriteEstimateWorkspaceStories(qb422016, w, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:42
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:44
func StreamEstimateWorkspaceStoryModalAdd(qw422016 *qt422016.Writer) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:44
	qw422016.N().S(`
  <div id="modal-story--add" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Add Story</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:54
	qw422016.E().S(string(action.ActChildAdd))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:54
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:55
	components.StreamFormVerticalInput(qw422016, "title", "story-add-title", "Title", "", 5, "Story title")
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:55
	qw422016.N().S(`
          <div class="mt">
            <button class="right" type="submit">Add Story</button>
            <a href="#"><button type="button">Cancel</button></a>
          </div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
func WriteEstimateWorkspaceStoryModalAdd(qq422016 qtio422016.Writer) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	StreamEstimateWorkspaceStoryModalAdd(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
func EstimateWorkspaceStoryModalAdd() string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	WriteEstimateWorkspaceStoryModalAdd(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:66
func StreamEstimateWorkspaceStoryModalEmpty(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:66
	qw422016.N().S(`
  <div id="modal-story-new" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Story</h2>
      </div>
      <div class="modal-body">
        <h2 class="billboard">empty</h2>
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:76
	StreamEstimateWorkspaceStoryPanelNew(qw422016, nil, nil, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:76
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:77
	StreamEstimateWorkspaceStoryPanelActive(qw422016, nil, nil, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:77
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:78
	StreamEstimateWorkspaceStoryPanelComplete(qw422016, nil, nil, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:78
	qw422016.N().S(`
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
func WriteEstimateWorkspaceStoryModalEmpty(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	StreamEstimateWorkspaceStoryModalEmpty(qw422016, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
func EstimateWorkspaceStoryModalEmpty(ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	WriteEstimateWorkspaceStoryModalEmpty(qb422016, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:84
func StreamEstimateWorkspaceStoryModal(qw422016 *qt422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:84
	qw422016.N().S(`
  <div id="modal-story-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.N().S(`" class="modal modal-story" style="display: none;" data-id="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.N().S(`" data-status="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.N().S(`">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Story</h2>
      </div>
      <div class="modal-body">
        <h2 class="billboard">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:93
	qw422016.E().S(s.TitleString())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:93
	qw422016.N().S(`</h2>
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:94
	StreamEstimateWorkspaceStoryPanelNew(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:94
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:95
	StreamEstimateWorkspaceStoryPanelActive(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:95
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:96
	StreamEstimateWorkspaceStoryPanelComplete(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:96
	qw422016.N().S(`
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
func WriteEstimateWorkspaceStoryModal(qq422016 qtio422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	StreamEstimateWorkspaceStoryModal(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
func EstimateWorkspaceStoryModal(w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	WriteEstimateWorkspaceStoryModal(qb422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:102
func StreamEstimateWorkspaceStoryModalEdit(qw422016 *qt422016.Writer, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:102
	qw422016.N().S(`
  `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:103
	StreamEstimateWorkspaceStoryModalEditPanel(qw422016, s.ID.String(), s.Title, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:103
	qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
func WriteEstimateWorkspaceStoryModalEdit(qq422016 qtio422016.Writer, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	StreamEstimateWorkspaceStoryModalEdit(qw422016, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
func EstimateWorkspaceStoryModalEdit(s *story.Story, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	WriteEstimateWorkspaceStoryModalEdit(qb422016, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:104
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:106
func StreamEstimateWorkspaceStoryModalEditPanel(qw422016 *qt422016.Writer, id string, title string, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:106
	qw422016.N().S(`
  <div `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	if id != "" {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
		qw422016.N().S(`id="modal-story-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
		qw422016.E().S(id)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
		qw422016.N().S(`-edit" `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	qw422016.N().S(`class="modal modal-story-edit`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	if id == `` {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
		qw422016.N().S(` modal-story-edit-new`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	qw422016.N().S(`" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Edit Story</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:116
	qw422016.E().S(id)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:116
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:117
	components.StreamFormVerticalInput(qw422016, "title", "", "Title", title, 5, "Story title")
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:117
	qw422016.N().S(`
          <div class="mt">
            <div class="right"><button type="submit" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:119
	qw422016.E().S(string(action.ActChildUpdate))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:119
	qw422016.N().S(`">Save Changes</button></div>
            <button class="story-delete-button" data-id="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:120
	qw422016.E().S(id)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:120
	qw422016.N().S(`" type="submit" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:120
	qw422016.E().S(string(action.ActChildRemove))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:120
	qw422016.N().S(`" onclick="return confirm('Are you sure you want to delete this story?');">Delete</button>
          </div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
func WriteEstimateWorkspaceStoryModalEditPanel(qq422016 qtio422016.Writer, id string, title string, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	StreamEstimateWorkspaceStoryModalEditPanel(qw422016, id, title, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
func EstimateWorkspaceStoryModalEditPanel(id string, title string, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	WriteEstimateWorkspaceStoryModalEditPanel(qb422016, id, title, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
}
