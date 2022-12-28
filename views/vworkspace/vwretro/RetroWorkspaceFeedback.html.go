// Code generated by qtc from "RetroWorkspaceFeedback.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:1
package vwretro

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:1
import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
func StreamRetroWorkspaceFeedbacks(qw422016 *qt422016.Writer, userID uuid.UUID, username string, w *workspace.FullRetro, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
	qw422016.N().S(`
  <div class="categories">
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:16
	for _, grp := range w.Feedbacks.Grouped(w.Retro.Categories) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:16
		qw422016.N().S(`    <div class="category">
      <div class="right"><a href="#modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:18
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:18
		qw422016.N().S(`"><button>New</button></a></div>
      <h4><a href="#modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:19
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:19
		qw422016.N().S(`">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:19
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:19
		qw422016.N().S(`</a></h4>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:20
		for _, f := range grp.Feedbacks {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:22
			uname := "Former Guest"
			curr := w.Members.Get(w.Retro.ID, f.UserID)
			if curr != nil {
				uname = curr.Name
			}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:27
			qw422016.N().S(`      <div class="clear"></div>
      <div class="feedback mt">
        <div class="right">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:30
			vwutil.StreamComments(qw422016, enum.ModelServiceFeedback, f.ID, f.TitleString(), w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:30
			qw422016.N().S(`</div>
        <a href="#modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:31
			qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:31
			qw422016.N().S(`" class="clean">
          `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:32
			qw422016.E().S(uname)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:32
			qw422016.N().S(`
          <div class="pt">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:33
			qw422016.N().S(f.HTML)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:33
			qw422016.N().S(`</div>
        </a>
      </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:36
		}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:36
		qw422016.N().S(`      `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:37
		StreamRetroWorkspaceFeedbackModalAdd(qw422016, w.Retro.Categories, grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:37
		qw422016.N().S(`
    </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:39
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:39
	qw422016.N().S(`  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	for _, f := range w.Feedbacks {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:42
		if userID == f.UserID {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:42
			qw422016.N().S(`  `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:43
			StreamRetroWorkspaceFeedbackModalEdit(qw422016, f, w.Retro.Categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:43
			qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:44
		} else {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:44
			qw422016.N().S(`  `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:45
			StreamRetroWorkspaceFeedbackModalView(qw422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:45
			qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:46
		}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
func WriteRetroWorkspaceFeedbacks(qq422016 qtio422016.Writer, userID uuid.UUID, username string, w *workspace.FullRetro, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	StreamRetroWorkspaceFeedbacks(qw422016, userID, username, w, ps)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
func RetroWorkspaceFeedbacks(userID uuid.UUID, username string, w *workspace.FullRetro, ps *cutil.PageState) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	WriteRetroWorkspaceFeedbacks(qb422016, userID, username, w, ps)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:48
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:50
func StreamRetroWorkspaceFeedbackModalAdd(qw422016 *qt422016.Writer, categories []string, category string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:50
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:51
	if !slices.Contains(categories, category) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:52
		categories = append(categories, category)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:53
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:53
	qw422016.N().S(`  <div id="modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:54
	qw422016.E().S(category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:54
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>New Feedback</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:63
	qw422016.E().S(string(action.ActFeedbackAdd))
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:63
	qw422016.N().S(`" />
          <table class="mt expanded">
            <tbody>
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:66
	components.StreamFormVerticalSelect(qw422016, "category", "Category", category, categories, categories, 5)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:66
	qw422016.N().S(`
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:67
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, "", 5, "HTML and Markdown supported")
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:67
	qw422016.N().S(`
              <tr><td colspan="2"><button type="submit">Add Feedback</button></td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
func WriteRetroWorkspaceFeedbackModalAdd(qq422016 qtio422016.Writer, categories []string, category string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	StreamRetroWorkspaceFeedbackModalAdd(qw422016, categories, category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
func RetroWorkspaceFeedbackModalAdd(categories []string, category string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	WriteRetroWorkspaceFeedbackModalAdd(qb422016, categories, category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:77
func StreamRetroWorkspaceFeedbackModalEdit(qw422016 *qt422016.Writer, f *feedback.Feedback, categories []string, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:77
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:78
	if !slices.Contains(categories, f.Category) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:79
		categories = append(categories, f.Category)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:80
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:80
	qw422016.N().S(`  <div id="modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:86
	qw422016.E().S(f.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:86
	qw422016.N().S(` :: `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:86
	qw422016.E().S(username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:86
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="feedbackID" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:90
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:90
	qw422016.N().S(`" />
          <table class="mt expanded">
            <tbody>
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:93
	components.StreamFormVerticalSelect(qw422016, "category", "Category", string(f.Category), categories, categories, 5)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:93
	qw422016.N().S(`
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:94
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, f.Content, 5, "HTML and Markdown supported")
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:94
	qw422016.N().S(`
              <tr><td colspan="2">
                <div class="right"><button type="submit" name="action" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	qw422016.E().S(string(action.ActFeedbackRemove))
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	qw422016.N().S(`" onclick="return confirm('Are you sure you want to delete this feedback?');">Delete</button></div>
                <button type="submit" name="action" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:97
	qw422016.E().S(string(action.ActFeedbackUpdate))
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:97
	qw422016.N().S(`">Save Changes</button>
              </td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
func WriteRetroWorkspaceFeedbackModalEdit(qq422016 qtio422016.Writer, f *feedback.Feedback, categories []string, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	StreamRetroWorkspaceFeedbackModalEdit(qw422016, f, categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
func RetroWorkspaceFeedbackModalEdit(f *feedback.Feedback, categories []string, username string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	WriteRetroWorkspaceFeedbackModalEdit(qb422016, f, categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:105
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:107
func StreamRetroWorkspaceFeedbackModalView(qw422016 *qt422016.Writer, f *feedback.Feedback, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:107
	qw422016.N().S(`
<div id="modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.N().S(`" class="modal" style="display: none;">
  <a class="backdrop" href="#"></a>
  <div class="modal-content">
    <div class="modal-header">
      <a href="#" class="modal-close">×</a>
      <h2>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:113
	qw422016.E().S(f.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:113
	qw422016.N().S(` :: `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:113
	qw422016.E().S(username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:113
	qw422016.N().S(`</h2>
    </div>
    <div class="modal-body">
      <div>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:116
	qw422016.N().S(f.HTML)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:116
	qw422016.N().S(`</div>
    </div>
  </div>
</div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
func WriteRetroWorkspaceFeedbackModalView(qq422016 qtio422016.Writer, f *feedback.Feedback, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	StreamRetroWorkspaceFeedbackModalView(qw422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
func RetroWorkspaceFeedbackModalView(f *feedback.Feedback, username string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	WriteRetroWorkspaceFeedbackModalView(qb422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:120
}
