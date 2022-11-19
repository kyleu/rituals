// Code generated by qtc from "RetroWorkspaceFeedback.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:1
package vwretro

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:1
import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:10
func StreamRetroWorkspaceFeedbacks(qw422016 *qt422016.Writer, userID uuid.UUID, username string, w *workspace.FullRetro) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:10
	qw422016.N().S(`
  <div class="categories">
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:12
	for _, grp := range w.Feedbacks.Grouped(w.Retro.Categories) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:12
		qw422016.N().S(`    <div class="category">
      <div class="right"><a href="#modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
		qw422016.N().S(`"><button>New</button></a></div>
      <h4>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:15
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:15
		qw422016.N().S(`</h4>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:16
		for _, f := range grp.Feedbacks {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:18
			uname := "Former Guest"
			curr := w.Members.Get(w.Retro.ID, f.UserID)
			if curr != nil {
				uname = curr.Name
			}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:23
			qw422016.N().S(`      <div class="clear"></div>
      <div class="feedback mt">
        <a href="#modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:26
			qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:26
			qw422016.N().S(`">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:26
			qw422016.E().S(uname)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:26
			qw422016.N().S(`</a>
        <div class="mt">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:27
			qw422016.N().S(f.HTML)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:27
			qw422016.N().S(`</div>
      </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:29
		}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:29
		qw422016.N().S(`      `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:30
		StreamRetroWorkspaceFeedbackModalAdd(qw422016, w.Retro.Categories, grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:30
		qw422016.N().S(`
    </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:32
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:32
	qw422016.N().S(`  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:34
	for _, f := range w.Feedbacks {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:35
		if userID == f.UserID {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:35
			qw422016.N().S(`  `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:36
			StreamRetroWorkspaceFeedbackModalEdit(qw422016, f, w.Retro.Categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:36
			qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:37
		} else {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:37
			qw422016.N().S(`  `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:38
			StreamRetroWorkspaceFeedbackModalView(qw422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:38
			qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:39
		}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:40
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
func WriteRetroWorkspaceFeedbacks(qq422016 qtio422016.Writer, userID uuid.UUID, username string, w *workspace.FullRetro) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	StreamRetroWorkspaceFeedbacks(qw422016, userID, username, w)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
func RetroWorkspaceFeedbacks(userID uuid.UUID, username string, w *workspace.FullRetro) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	WriteRetroWorkspaceFeedbacks(qb422016, userID, username, w)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:43
func StreamRetroWorkspaceFeedbackModalAdd(qw422016 *qt422016.Writer, categories []string, category string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:43
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:44
	if !slices.Contains(categories, category) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:45
		categories = append(categories, category)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:46
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:46
	qw422016.N().S(`  <div id="modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	qw422016.E().S(category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>New Feedback</h2>
      </div>
      <div class="modal-body">
        <form action="" method="post">
          <input type="hidden" name="action" value="feedback-add" />
          <table class="mt expanded">
            <tbody>
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:59
	components.StreamFormVerticalSelect(qw422016, "category", "Category", category, categories, categories, 5)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:59
	qw422016.N().S(`
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:60
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, "", 5, "HTML and Markdown supported")
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:60
	qw422016.N().S(`
              <tr><td colspan="2"><button type="submit">Add Feedback</button></td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
func WriteRetroWorkspaceFeedbackModalAdd(qq422016 qtio422016.Writer, categories []string, category string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	StreamRetroWorkspaceFeedbackModalAdd(qw422016, categories, category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
func RetroWorkspaceFeedbackModalAdd(categories []string, category string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	WriteRetroWorkspaceFeedbackModalAdd(qb422016, categories, category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:68
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
func StreamRetroWorkspaceFeedbackModalEdit(qw422016 *qt422016.Writer, f *feedback.Feedback, categories []string, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:71
	if !slices.Contains(categories, f.Category) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:72
		categories = append(categories, f.Category)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:73
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:73
	qw422016.N().S(`  <div id="modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:74
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:74
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:79
	qw422016.E().S(f.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:79
	qw422016.N().S(` :: `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:79
	qw422016.E().S(username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:79
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        <form action="" method="post">
          <input type="hidden" name="action" value="feedback-edit" />
          <input type="hidden" name="feedbackID" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:84
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:84
	qw422016.N().S(`" />
          <table class="mt expanded">
            <tbody>
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:87
	components.StreamFormVerticalSelect(qw422016, "category", "Category", string(f.Category), categories, categories, 5)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:87
	qw422016.N().S(`
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:88
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, f.Content, 5, "HTML and Markdown supported")
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:88
	qw422016.N().S(`
              <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
func WriteRetroWorkspaceFeedbackModalEdit(qq422016 qtio422016.Writer, f *feedback.Feedback, categories []string, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	StreamRetroWorkspaceFeedbackModalEdit(qw422016, f, categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
func RetroWorkspaceFeedbackModalEdit(f *feedback.Feedback, categories []string, username string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	WriteRetroWorkspaceFeedbackModalEdit(qb422016, f, categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:96
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:98
func StreamRetroWorkspaceFeedbackModalView(qw422016 *qt422016.Writer, f *feedback.Feedback, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:98
	qw422016.N().S(`
<div id="modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:99
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:99
	qw422016.N().S(`" class="modal" style="display: none;">
  <a class="backdrop" href="#"></a>
  <div class="modal-content">
    <div class="modal-header">
      <a href="#" class="modal-close">×</a>
      <h2>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:104
	qw422016.E().S(f.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:104
	qw422016.N().S(` :: `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:104
	qw422016.E().S(username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:104
	qw422016.N().S(`</h2>
    </div>
    <div class="modal-body">
      <div>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:107
	qw422016.N().S(f.HTML)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:107
	qw422016.N().S(`</div>
    </div>
  </div>
</div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
func WriteRetroWorkspaceFeedbackModalView(qq422016 qtio422016.Writer, f *feedback.Feedback, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	StreamRetroWorkspaceFeedbackModalView(qw422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
func RetroWorkspaceFeedbackModalView(f *feedback.Feedback, username string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	WriteRetroWorkspaceFeedbackModalView(qb422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
}
