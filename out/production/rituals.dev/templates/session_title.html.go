// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/session/title.html
// DO NOT EDIT!
package templates

import (
	"bytes"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func ComponentSessionTitle(v string, ctx web.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<div class="uk-margin">
  <label class="uk-form-label" for="model-title-input">Title</label>
  <input class="uk-input" id="model-title-input" type="text" name="`)
	hero.EscapeHTML(util.KeyTitle, buffer)
	buffer.WriteString(`" placeholder="Name your session" value="`)
	hero.EscapeHTML(v, buffer)
	buffer.WriteString(`" />
</div>
`)

}
