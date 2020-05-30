// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/admin/connection/detail.html
// DO NOT EDIT!
package templates

import (
	"fmt"
	"io"

	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func AdminConnectionDetail(model *socket.Status, msg *socket.Message, ctx web.RequestContext, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
`)
	Head(ctx, _buffer)
	_buffer.WriteString(`
<body style="visibility: hidden;">
`)
	Navbar(ctx, _buffer)
	_buffer.WriteString(`
<div id="`)
	hero.EscapeHTML(util.KeyContent, _buffer)
	_buffer.WriteString(`" data-uk-height-viewport="expand: true">
`)
	for _, flash := range ctx.Flashes {
		cls, content := web.ParseFlash(flash)
		_buffer.WriteString(`
  <div class="alert-top `)
		hero.EscapeHTML(cls, _buffer)
		_buffer.WriteString(`" data-uk-alert>
    <a class="uk-alert-close" href="#" data-uk-close=""></a>
    <p>`)
		hero.EscapeHTML(content, _buffer)
		_buffer.WriteString(`</p>
  </div>
`)
	}
	_buffer.WriteString(`
<div class="uk-section uk-section-small">
  <div class="uk-container">
    <div class="uk-card uk-card-body">
      <h3 class="uk-card-title">Socket Connection</h3>
      <table class="uk-table uk-table-divider uk-text-left">
        <tbody>
        <tr>
          <th>ID</th>
          <td>`)
	hero.EscapeHTML(model.ID.String(), _buffer)
	_buffer.WriteString(`</td>
        </tr>
        <tr>
          <th>User ID</th>
          <td><a class="`)
	hero.EscapeHTML(ctx.Profile.LinkClass(), _buffer)
	_buffer.WriteString(`" href="`)
	hero.EscapeHTML(ctx.Route(util.AdminLink(util.KeyUser, `detail`), `id`, model.UserID.String()), _buffer)
	_buffer.WriteString(`">`)
	hero.EscapeHTML(model.UserID.String(), _buffer)
	_buffer.WriteString(`</a></td>
        </tr>
        </tbody>
      </table>
    </div>
    <div class="uk-card uk-card-body">
      <form class="uk-form-stacked" action="`)
	hero.EscapeHTML(ctx.Route(util.AdminLink(util.KeyConnection, `post`)), _buffer)
	_buffer.WriteString(`" method="post">
        <input name=util.KeyID type="hidden" value="`)
	hero.EscapeHTML(fmt.Sprintf("%v", model.ID), _buffer)
	_buffer.WriteString(`" />
        <fieldset class="uk-fieldset">
          <div class="uk-margin-small">
            <label class="uk-form-label" for="`)
	hero.EscapeHTML(util.KeySvc, _buffer)
	_buffer.WriteString(`">Service</label>
            `)
	ComponentSelect(util.KeySvc, []string{util.SvcSystem.Key, util.SvcEstimate.Key, util.SvcStandup.Key, util.SvcRetro.Key}, msg.Svc, _buffer)
	_buffer.WriteString(`
          </div>

          <div class="uk-margin">
            <label class="uk-form-label" for="cmd">Command</label>
            <input class="uk-input" name="cmd" type="text" value="`)
	hero.EscapeHTML(msg.Cmd, _buffer)
	_buffer.WriteString(`" />
          </div>

          <div class="uk-margin">
            <label class="uk-form-label" for="param">Param</label>
            <textarea class="uk-textarea" name="param">`)
	hero.EscapeHTML(util.ToJSON(msg.Param), _buffer)
	_buffer.WriteString(`</textarea>
          </div>

          <div class="uk-margin-top">
            <button class="uk-button uk-button-default" type="submit">Send</button>
          </div>
        </fieldset>
      </form>
    </div>
  </div>
</div>
`)

	_buffer.WriteString(`
</div>
<script>window.addEventListener("load", function() { dom.initDom('`)
	hero.EscapeHTML(ctx.Profile.Theme.String(), _buffer)
	_buffer.WriteString(`'); }, false)</script>
</body>
</html>
`)
	return w.Write(_buffer.Bytes())

}
