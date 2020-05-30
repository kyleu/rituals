// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/session/sprint.html
// DO NOT EDIT!
package templates

import (
	"bytes"

	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func ComponentSessionSprint(sprints sprint.Sessions, ctx web.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<div id="model-sprint-container" class="uk-margin">
  <label class="uk-form-label" for="model-sprint-select" name="sprint">Sprint</label>
  <div id="model-sprint-select">
    <select class="uk-select" name="sprint">
      <option value="">- no sprint -</option>
      `)
	for _, s := range sprints {
		buffer.WriteString(`
        <option value="`)
		hero.EscapeHTML(s.ID.String(), buffer)
		buffer.WriteString(`">`)
		hero.EscapeHTML(s.Title, buffer)
		buffer.WriteString(`</option>
      `)
	}
	buffer.WriteString(`
    </select>
  </div>
</div>
`)

}
