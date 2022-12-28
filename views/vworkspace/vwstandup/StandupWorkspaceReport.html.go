// Code generated by qtc from "StandupWorkspaceReport.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:1
package vwstandup

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:1
import (
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:12
func StreamStandupWorkspaceReports(qw422016 *qt422016.Writer, w *workspace.FullStandup, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:12
	qw422016.N().S(`
  <ul class="accordion">
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
	for _, g := range w.Reports.Grouped() {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:14
		qw422016.N().S(`    <li>
      <input id="accordion-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:16
		qw422016.E().S(util.TimeToYMD(&g.Day))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:16
		qw422016.N().S(`" type="checkbox" hidden checked="checked" />
      <label for="accordion-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
		qw422016.E().S(util.TimeToYMD(&g.Day))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
		qw422016.N().S(`">`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
		qw422016.N().S(` `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
		qw422016.E().S(util.TimeToYMD(&g.Day))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:17
		qw422016.N().S(`</label>
      <div class="bd">
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:19
		for _, r := range g.Reports {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:19
			qw422016.N().S(`        <div class="report">
          <div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:23
			uname := "Former Guest"
			curr := w.Members.Get(w.Standup.ID, r.UserID)
			if curr != nil {
				uname = curr.Name
			}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:28
			qw422016.N().S(`            <div class="right">`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:29
			vwutil.StreamComments(qw422016, enum.ModelServiceReport, r.ID, r.TitleString(), w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:29
			qw422016.N().S(`</div>
            <a href="#modal-report-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:30
			qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:30
			qw422016.N().S(`" class="clean">
              <h4>`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:31
			qw422016.E().S(uname)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:31
			qw422016.N().S(`</h4>
              <div class="pt">`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:32
			qw422016.N().S(r.HTML)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:32
			qw422016.N().S(`</div>
            </a>
          </div>
        </div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:36
		}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:36
		qw422016.N().S(`      </div>
    </li>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:39
	}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:39
	qw422016.N().S(`  </ul>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:41
	for _, r := range w.Reports {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
		if ps.Profile.ID == r.UserID {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:42
			qw422016.N().S(`  `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:43
			StreamStandupWorkspaceReportModalEdit(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:43
			qw422016.N().S(`
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:44
		} else {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:44
			qw422016.N().S(`  `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:45
			StreamStandupWorkspaceReportModalView(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:45
			qw422016.N().S(`
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:46
		}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:47
	}
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
func WriteStandupWorkspaceReports(qq422016 qtio422016.Writer, w *workspace.FullStandup, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	StreamStandupWorkspaceReports(qw422016, w, ps)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
func StandupWorkspaceReports(w *workspace.FullStandup, ps *cutil.PageState) string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	WriteStandupWorkspaceReports(qb422016, w, ps)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:48
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:50
func StreamStandupWorkspaceReportModalAdd(qw422016 *qt422016.Writer) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:50
	qw422016.N().S(`
  <div id="modal-report--add" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>New Report</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:60
	qw422016.E().S(string(action.ActReportAdd))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:60
	qw422016.N().S(`" />
          <table class="mt expanded">
            <tbody>
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:63
	components.StreamFormVerticalInputTimestampDay(qw422016, "day", "Day", util.TimeToday(), 5)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:63
	qw422016.N().S(`
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:64
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, "", 5, "HTML and Markdown supported")
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:64
	qw422016.N().S(`
            <tr><td colspan="2"><button type="submit">Add Report</button></td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
func WriteStandupWorkspaceReportModalAdd(qq422016 qtio422016.Writer) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	StreamStandupWorkspaceReportModalAdd(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
func StandupWorkspaceReportModalAdd() string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	WriteStandupWorkspaceReportModalAdd(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:72
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:74
func StreamStandupWorkspaceReportModalEdit(qw422016 *qt422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:74
	qw422016.N().S(`
  <div id="modal-report-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:75
	qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:75
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:80
	qw422016.E().S(r.TitleString())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:80
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="reportID" value="`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:84
	qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:84
	qw422016.N().S(`" />
          <table class="mt expanded">
            <tbody>
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:87
	components.StreamFormVerticalInputTimestampDay(qw422016, "day", "Day", &r.Day, 5)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:87
	qw422016.N().S(`
            `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:88
	components.StreamFormVerticalTextarea(qw422016, "content", "Content", 8, r.Content, 5, "HTML and Markdown supported")
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:88
	qw422016.N().S(`
            <tr><td colspan="2">
              <div class="right"><button type="submit" name="action" value="`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:90
	qw422016.E().S(string(action.ActReportRemove))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:90
	qw422016.N().S(`" onclick="return confirm('Are you sure you want to delete this report?');">Delete</button></div>
              <button type="submit" name="action" value="`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	qw422016.E().S(string(action.ActReportUpdate))
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:91
	qw422016.N().S(`">Save Changes</button>
            </td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
func WriteStandupWorkspaceReportModalEdit(qq422016 qtio422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	StreamStandupWorkspaceReportModalEdit(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
func StandupWorkspaceReportModalEdit(r *report.Report) string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	WriteStandupWorkspaceReportModalEdit(qb422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:99
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:101
func StreamStandupWorkspaceReportModalView(qw422016 *qt422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:101
	qw422016.N().S(`
  <div id="modal-report-`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:102
	qw422016.E().S(r.ID.String())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:102
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:107
	qw422016.E().S(r.TitleString())
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:107
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        `)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:110
	qw422016.N().S(r.HTML)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:110
	qw422016.N().S(`
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
func WriteStandupWorkspaceReportModalView(qq422016 qtio422016.Writer, r *report.Report) {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	StreamStandupWorkspaceReportModalView(qw422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
}

//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
func StandupWorkspaceReportModalView(r *report.Report) string {
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	WriteStandupWorkspaceReportModalView(qb422016, r)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
	return qs422016
//line views/vworkspace/vwstandup/StandupWorkspaceReport.html:114
}
