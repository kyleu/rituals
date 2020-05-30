// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/retro/list.html
// DO NOT EDIT!
package templates

import (
	"io"

	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func RetroList(sessions retro.Sessions, teams team.Sessions, sprints sprint.Sessions, auths auth.Records, params *query.Params, ctx web.RequestContext, w io.Writer) (int, error) {
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
    <div class="uk-card uk-card-body uk-margin-top">
      <h3 class="uk-card-title"><span class="h3-icon" data-uk-icon="icon: `)
	hero.EscapeHTML(util.SvcRetro.Icon, _buffer)
	_buffer.WriteString(`"></span> `)
	hero.EscapeHTML(util.SvcRetro.PluralTitle, _buffer)
	_buffer.WriteString(`</h3>
      <p>Discover improvements and praise for your work</p>
      `)
	if len(sessions) > 0 {
		RetroTable(sessions, params, ctx, _buffer)
	}
	_buffer.WriteString(`
    </div>

    <div class="uk-card uk-card-body uk-margin-top">
      <h3 class="uk-card-title">New `)
	hero.EscapeHTML(util.SvcRetro.Title, _buffer)
	_buffer.WriteString(`</h3>
      <form class="uk-form-stacked" action="`)
	hero.EscapeHTML(ctx.Route(util.SvcRetro.Key+`.new`), _buffer)
	_buffer.WriteString(`" method="post">
        `)
	RetroForm("", retro.DefaultCategories, teams, sprints, auths, ctx, _buffer)
	_buffer.WriteString(`
        <div class="uk-margin-top">
          <button class="uk-button uk-button-default" type="submit">Start</button>
        </div>
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
