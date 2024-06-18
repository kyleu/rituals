// Code generated by qtc from "SocketTap.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/SocketTap.html:1
package vadmin

//line views/vadmin/SocketTap.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vadmin/SocketTap.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/SocketTap.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/SocketTap.html:7
type SocketTap struct {
	layout.Basic
}

//line views/vadmin/SocketTap.html:11
func (p *SocketTap) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/SocketTap.html:11
	qw422016.N().S(`
  <div class="card">
    <h3>Tap Activity</h3>
    <em>Shows all WebSocket traffic in real-time</em>
    <div class="overflow full-width">
      <table class="mt">
        <thead>
          <tr>
            <th>From</th>
            <th>Channel</th>
            <th>Command</th>
            <th>Param</th>
          </tr>
        </thead>
        <tbody id="tap-logs">
          <tr>
            <td colspan="4">Connecting...</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
  <script>
    function open() {
      document.getElementById("tap-logs").innerHTML = "";
      addMessage({"from": ":tap", "channel": ":tap", "cmd": "open", "param": null});
    }
    function recv(m) {
      addMessage(m);
    }
    function err(e) {
      addMessage({"from": ":tap", "channel": ":tap", "cmd": "error", "param": e});
    }
    function addMessage(m) {
      const t = document.getElementById("tap-logs");
      const row = document.createElement("tr");

      const tdFrom = document.createElement("td");
      tdFrom.innerText = m.from;
      row.appendChild(tdFrom);

      const tdChannel = document.createElement("td");
      tdChannel.innerText = m.channel;
      row.appendChild(tdChannel);

      const tdCmd = document.createElement("td");
      tdCmd.innerText = m.cmd;
      row.appendChild(tdCmd);

      const tdParam = document.createElement("td");
      const pre = document.createElement("pre");
      pre.innerText = JSON.stringify(m.param, null, 2);
      tdParam.appendChild(pre);
      row.appendChild(tdParam);

      t.appendChild(row);
    }
    document.addEventListener("DOMContentLoaded", function() {
      new rituals.Socket(true, open, recv, err, "/admin/sockets/tap-socket");
    });
  </script>
`)
//line views/vadmin/SocketTap.html:72
}

//line views/vadmin/SocketTap.html:72
func (p *SocketTap) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/SocketTap.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/SocketTap.html:72
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/SocketTap.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/SocketTap.html:72
}

//line views/vadmin/SocketTap.html:72
func (p *SocketTap) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/SocketTap.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/SocketTap.html:72
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/SocketTap.html:72
	qs422016 := string(qb422016.B)
//line views/vadmin/SocketTap.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/SocketTap.html:72
	return qs422016
//line views/vadmin/SocketTap.html:72
}
