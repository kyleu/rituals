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
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:12
func StreamEstimateWorkspaceStories(qw422016 *qt422016.Writer, w *workspace.FullEstimate, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:12
	qw422016.N().S(`
  <table class="mt expanded">
    <thead>
      <tr>
        <th>Story</th>
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
		qw422016.N().S(`      <tr>
        <td><a href="#modal-story-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.N().S(`">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.E().S(s.TitleString())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:25
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.E().S(string(s.Status))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:26
		qw422016.N().S(`</td>
        <td>`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:27
		qw422016.E().S(s.FinalVoteSafe())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:27
		qw422016.N().S(`</td>
        <td>`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:28
		vwutil.StreamComments(qw422016, enum.ModelServiceStory, s.ID, s.TitleString(), w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:28
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:30
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:30
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:33
	for _, s := range w.Stories {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:34
		if ps.Profile.ID == s.UserID {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:34
			qw422016.N().S(`  `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:35
			StreamEstimateWorkspaceStoryModalEdit(qw422016, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:35
			qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:36
		}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:36
		qw422016.N().S(`  `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:37
		StreamEstimateWorkspaceStoryModal(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:37
		qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:38
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
func WriteEstimateWorkspaceStories(qq422016 qtio422016.Writer, w *workspace.FullEstimate, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	StreamEstimateWorkspaceStories(qw422016, w, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
func EstimateWorkspaceStories(w *workspace.FullEstimate, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	WriteEstimateWorkspaceStories(qb422016, w, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:39
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:41
func StreamEstimateWorkspaceStoryModalAdd(qw422016 *qt422016.Writer) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:41
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
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:51
	qw422016.E().S(string(action.ActStoryAdd))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:51
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:52
	components.StreamFormVerticalInput(qw422016, "title", "Title", "", 5, "Story title")
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:52
	qw422016.N().S(`
          <div class="mt">
            <a href="#"><button type="button">Cancel</button></a>
            <button type="submit">Add Story</button>
          </div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
func WriteEstimateWorkspaceStoryModalAdd(qq422016 qtio422016.Writer) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	StreamEstimateWorkspaceStoryModalAdd(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
func EstimateWorkspaceStoryModalAdd() string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	WriteEstimateWorkspaceStoryModalAdd(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:61
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:63
func StreamEstimateWorkspaceStoryModal(qw422016 *qt422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:63
	qw422016.N().S(`
  <div id="modal-story-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:64
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Story</h2>
      </div>
      <div class="modal-body">
        <h2 class="billboard">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:72
	qw422016.E().S(s.TitleString())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:72
	qw422016.N().S(`</h2>
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:73
	StreamEstimateWorkspaceStoryPanelNew(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:73
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:74
	StreamEstimateWorkspaceStoryPanelActive(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:74
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:75
	StreamEstimateWorkspaceStoryPanelComplete(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:75
	qw422016.N().S(`
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
func WriteEstimateWorkspaceStoryModal(qq422016 qtio422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	StreamEstimateWorkspaceStoryModal(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
func EstimateWorkspaceStoryModal(w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	WriteEstimateWorkspaceStoryModal(qb422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:79
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:81
func StreamEstimateWorkspaceStoryPanelNew(qw422016 *qt422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:81
	qw422016.N().S(`
  <div class="mt" `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	if s.Status != enum.SessionStatusNew {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
		qw422016.N().S(` style="display: none;"`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:82
	qw422016.N().S(`>
    <form action="" method="post">
      <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:84
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:84
	qw422016.N().S(`" />
      <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.E().S(string(action.ActStoryStatus))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:85
	qw422016.N().S(`" />
      <input type="hidden" name="status" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:86
	qw422016.E().S(string(enum.SessionStatusActive))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:86
	qw422016.N().S(`" />
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:87
	if ps.Profile.ID == s.UserID {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:87
		qw422016.N().S(`      <div>Your story is available to <a href="#modal-story-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:88
		qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:88
		qw422016.N().S(`-edit">edit</a> and ready to <button type="submit" class="button-link">start voting</button></div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:89
	} else {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:90
		if mem := w.Members.Get(s.EstimateID, s.UserID); mem == nil {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:90
			qw422016.N().S(`      <div>This story is ready to start voting</div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:92
		} else {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:92
			qw422016.N().S(`      <div>This story was created by <a href="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:93
			qw422016.E().S(mem.PublicWebPath(w.Estimate.Slug))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:93
			qw422016.N().S(`">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:93
			qw422016.E().S(mem.Name)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:93
			qw422016.N().S(`</a> and is ready to <button type="submit" class="button-link">start voting</button></div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:94
		}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:95
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:95
	qw422016.N().S(`    </form>
    <hr />
    <div class="mt">
      <form action="" method="post">
        <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:100
	qw422016.N().S(`" />
        <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:101
	qw422016.E().S(string(action.ActStoryStatus))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:101
	qw422016.N().S(`" />
        <input type="hidden" name="status" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:102
	qw422016.E().S(string(enum.SessionStatusActive))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:102
	qw422016.N().S(`" />
        <div class="right"><button type="submit">Start Voting</button></div>
      </form>
      <form action="" method="post" onsubmit="return confirm('Are you sure you want to delete this story?')">
        <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:106
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:106
	qw422016.N().S(`" />
        <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	qw422016.E().S(string(action.ActStoryRemove))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:107
	qw422016.N().S(`" />
        <div><button type="submit">Delete Story</button></div>
      </form>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
func WriteEstimateWorkspaceStoryPanelNew(qq422016 qtio422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	StreamEstimateWorkspaceStoryPanelNew(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
func EstimateWorkspaceStoryPanelNew(w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	WriteEstimateWorkspaceStoryPanelNew(qb422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:112
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:114
func StreamEstimateWorkspaceStoryPanelActive(qw422016 *qt422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:114
	qw422016.N().S(`
  <div class="mt" `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:115
	if s.Status != enum.SessionStatusActive {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:115
		qw422016.N().S(` style="display: none;"`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:115
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:115
	qw422016.N().S(`>
    <h4>Members</h4>
    <div class="story-members">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:118
	for _, m := range w.Members {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:119
		v := w.Votes.Get(s.ID, m.UserID)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:119
		qw422016.N().S(`      <div class="member">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:121
		if v == nil {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:121
			qw422016.N().S(`        <div>-</div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:123
		} else {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:123
			qw422016.N().S(`        <div>`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:124
			qw422016.E().S(v.Choice)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:124
			qw422016.N().S(`</div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:125
		}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:125
		qw422016.N().S(`        `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
		qw422016.E().S(m.Name)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:126
		qw422016.N().S(`
      </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:128
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:128
	qw422016.N().S(`    </div>
    <hr />
    <h4>Choices</h4>
    <form action="" method="post">
      <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:133
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:133
	qw422016.N().S(`" />
      <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:134
	qw422016.E().S(string(action.ActVote))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:134
	qw422016.N().S(`" />
      <div class="story-votes">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:136
	selfVote := w.Votes.Get(s.ID, ps.Profile.ID)

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:137
	for _, c := range w.Estimate.Choices {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:137
		qw422016.N().S(`        <div class="vote">
          <label>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:140
		if selfVote != nil && selfVote.Choice == c {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:140
			qw422016.N().S(`            <input type="radio" name="vote" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:141
			qw422016.E().S(c)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:141
			qw422016.N().S(`" checked="checked" />
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:142
		} else {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:142
			qw422016.N().S(`            <input type="radio" name="vote" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:143
			qw422016.E().S(c)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:143
			qw422016.N().S(`" />
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:144
		}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:144
		qw422016.N().S(`            <div class="vote-choice">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:145
		qw422016.E().S(c)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:145
		qw422016.N().S(`</div>
          </label>
        </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:148
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:148
	qw422016.N().S(`      </div>
      <div class="mt">
        <div class="right"><button type="submit">Submit vote</button></div>
      </div>
      <div class="clear"></div>
    </form>
    <hr />
    <div class="mt">
      <form action="" method="post">
        <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:158
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:158
	qw422016.N().S(`" />
        <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:159
	qw422016.E().S(string(action.ActStoryStatus))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:159
	qw422016.N().S(`" />
        <input type="hidden" name="status" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:160
	qw422016.E().S(string(enum.SessionStatusComplete))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:160
	qw422016.N().S(`" />
        <div class="right"><button type="submit">Finish Voting</button></div>
      </form>
      <form action="" method="post">
        <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:164
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:164
	qw422016.N().S(`" />
        <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:165
	qw422016.E().S(string(action.ActStoryStatus))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:165
	qw422016.N().S(`" />
        <input type="hidden" name="status" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:166
	qw422016.E().S(string(enum.SessionStatusNew))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:166
	qw422016.N().S(`" />
        <div><button type="submit">Restart</button></div>
      </form>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
func WriteEstimateWorkspaceStoryPanelActive(qq422016 qtio422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	StreamEstimateWorkspaceStoryPanelActive(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
func EstimateWorkspaceStoryPanelActive(w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	WriteEstimateWorkspaceStoryPanelActive(qb422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:171
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:173
func StreamEstimateWorkspaceStoryPanelComplete(qw422016 *qt422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:173
	qw422016.N().S(`
  <div class="mt" `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:174
	if s.Status != enum.SessionStatusComplete {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:174
		qw422016.N().S(` style="display: none;"`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:174
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:174
	qw422016.N().S(`>
    <div class="mt story-results">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:176
	if s.Status == enum.SessionStatusComplete {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:176
		qw422016.N().S(`      <div class="vote-results">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:178
		for _, m := range w.Members {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:178
			qw422016.N().S(`        <div class="vote-result">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:180
			if v := w.Votes.Get(s.ID, m.UserID); v == nil {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:180
				qw422016.N().S(`          <div class="number" title="user did not vote">-</div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:182
			} else {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:182
				qw422016.N().S(`          <div class="number">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:183
				qw422016.E().S(v.Choice)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:183
				qw422016.N().S(`</div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:184
			}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:184
			qw422016.N().S(`          `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:185
			qw422016.E().V(m.Name)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:185
			qw422016.N().S(`
        </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:187
		}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:187
		qw422016.N().S(`      </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:189
	} else {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:189
		qw422016.N().S(`      <em>Voting in progress...</em>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:191
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:191
	qw422016.N().S(`    </div>
    <hr />
    <div class="mt">
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:195
	if s.Status == enum.SessionStatusComplete {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:195
		qw422016.N().S(`      `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:196
		StreamEstimateWorkspaceStoryResults(qw422016, w.Votes.GetByStoryIDs(s.ID), ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:196
		qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:197
	} else {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:197
		qw422016.N().S(`      <em>Voting is still active</em>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:199
	}
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:199
	qw422016.N().S(`    </div>
    <hr />
    <div class="mt">
      <div class="right">
        <a href="#"><button type="button">Close</button></a>
      </div>
      <form action="" method="post">
        <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:207
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:207
	qw422016.N().S(`" />
        <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:208
	qw422016.E().S(string(action.ActStoryStatus))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:208
	qw422016.N().S(`" />
        <input type="hidden" name="status" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:209
	qw422016.E().S(string(enum.SessionStatusActive))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:209
	qw422016.N().S(`" />
        <div><button type="submit">Reopen</button></div>
      </form>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
func WriteEstimateWorkspaceStoryPanelComplete(qq422016 qtio422016.Writer, w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	StreamEstimateWorkspaceStoryPanelComplete(qw422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
func EstimateWorkspaceStoryPanelComplete(w *workspace.FullEstimate, s *story.Story, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	WriteEstimateWorkspaceStoryPanelComplete(qb422016, w, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:214
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:216
func StreamEstimateWorkspaceStoryResults(qw422016 *qt422016.Writer, v vote.Votes, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:216
	qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:217
	r := v.Results()

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:217
	qw422016.N().S(`  <div class="vote-calculations">
    <div class="vote-calculation" title="portion of votes that were able to be parsed as numbers">
      <div class="value">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:220
	qw422016.N().D(r.Count)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:220
	qw422016.N().S(`/`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:220
	qw422016.N().D(len(v))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:220
	qw422016.N().S(`</div>
      Counted
    </div>
    <div class="vote-calculation" title="the minimum and maximum vote recorded">
      <div class="value">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:224
	qw422016.N().F(r.Min)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:224
	qw422016.N().S(`-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:224
	qw422016.N().F(r.Max)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:224
	qw422016.N().S(`</div>
      Range
    </div>
    <div class="vote-calculation" title="mean average of all votes">
      <div class="value">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:228
	qw422016.N().F(r.Mean)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:228
	qw422016.N().S(`</div>
      Average
    </div>
    <div class="vote-calculation" title="median value from collected votes">
      <div class="value">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:232
	qw422016.N().F(r.Median)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:232
	qw422016.N().S(`</div>
      Median
    </div>
    <div class="vote-calculation" title="mode value(s) from collected votes">
      <div class="value">`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:236
	qw422016.E().S(r.ModeString())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:236
	qw422016.N().S(`</div>
      Mode
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
func WriteEstimateWorkspaceStoryResults(qq422016 qtio422016.Writer, v vote.Votes, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	StreamEstimateWorkspaceStoryResults(qw422016, v, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
func EstimateWorkspaceStoryResults(v vote.Votes, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	WriteEstimateWorkspaceStoryResults(qb422016, v, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:240
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:242
func StreamEstimateWorkspaceStoryModalEdit(qw422016 *qt422016.Writer, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:242
	qw422016.N().S(`
  <div id="modal-story-`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:243
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:243
	qw422016.N().S(`-edit" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Edit Story</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="storyID" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:252
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:252
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:253
	components.StreamFormVerticalInput(qw422016, "title", "Title", s.Title, 5, "Story title")
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:253
	qw422016.N().S(`
          <div class="mt">
            <div class="right"><button type="submit" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:255
	qw422016.E().S(string(action.ActStoryRemove))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:255
	qw422016.N().S(`" onclick="return confirm('Are you sure you want to delete this story?');">Delete</button></div>
            <button type="submit" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:256
	qw422016.E().S(string(action.ActStoryUpdate))
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:256
	qw422016.N().S(`">Save Changes</button>
          </div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
func WriteEstimateWorkspaceStoryModalEdit(qq422016 qtio422016.Writer, s *story.Story, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	StreamEstimateWorkspaceStoryModalEdit(qw422016, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
}

//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
func EstimateWorkspaceStoryModalEdit(s *story.Story, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	WriteEstimateWorkspaceStoryModalEdit(qb422016, s, ps)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspaceStory.html:262
}
