// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/admin/team/list.html
// DO NOT EDIT!
package templates

import (
	"io"

	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func AdminTeamList(teams team.Sessions, params query.ParamSet, ctx web.RequestContext, w io.Writer) (int, error) {
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
      <h3 class="uk-card-title">Teams</h3>
      `)
	AdminTeamTable(teams, params.Get(util.SvcTeam.Key, ctx.Logger), ctx, _buffer)
	_buffer.WriteString(`
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
