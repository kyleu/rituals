// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/admin/action/table.html
// DO NOT EDIT!
package templates

import (
	"bytes"

	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func AdminActionTable(actions action.Actions, params *query.Params, ctx web.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`

<table class="uk-table uk-table-divider uk-text-left">
  <thead>
    <tr>
      `)
	ComponentTableHeader(util.KeyAction, util.KeyID, util.Title(util.KeyID), params, ctx, buffer)
	ComponentTableHeader(util.KeyAction, util.WithID(util.KeyModel), util.Title(util.KeyModel), params, ctx, buffer)
	ComponentTableHeader(util.KeyAction, util.WithID(util.KeyAuthor), util.WithID(util.KeyAuthor), params, ctx, buffer)
	ComponentTableHeader(util.KeyAction, util.KeyAct, util.Title(util.KeyAct), params, ctx, buffer)
	ComponentTableHeader(util.KeyAction, util.KeyNote, util.Title(util.KeyNote), params, ctx, buffer)
	ComponentTableHeader(util.KeyAction, util.KeyCreated, util.Title(util.KeyCreated), params, ctx, buffer)
	buffer.WriteString(`
    </tr>
  </thead>
  <tbody>
  `)
	for _, model := range actions {
		buffer.WriteString(`
    <tr>
      <td class="uk-table-shrink uk-text-nowrap"><a class="`)
		hero.EscapeHTML(ctx.Profile.LinkClass(), buffer)
		buffer.WriteString(`" href="`)
		hero.EscapeHTML(ctx.Route(util.AdminLink(util.KeyAction, `detail`), `id`, model.ID.String()), buffer)
		buffer.WriteString(`">`)
		ComponentUUID(&model.ID, buffer)
		buffer.WriteString(`</a></td>
      <td class="uk-table-shrink uk-text-nowrap">`)
		ComponentChannel(model.Svc, &model.ModelID, ctx, buffer)
		buffer.WriteString(`</td>
      <td class="uk-table-shrink uk-text-nowrap"><a class="`)
		hero.EscapeHTML(ctx.Profile.LinkClass(), buffer)
		buffer.WriteString(`" href="`)
		hero.EscapeHTML(ctx.Route(util.AdminLink(util.KeyUser, `detail`), `id`, model.ID.String()), buffer)
		buffer.WriteString(`">`)
		ComponentUUID(&model.AuthorID, buffer)
		buffer.WriteString(`</a></td>
      <td class="uk-table-shrink uk-text-nowrap">`)
		hero.EscapeHTML(model.Act, buffer)
		buffer.WriteString(`</td>
      <td>`)
		hero.EscapeHTML(model.Note, buffer)
		buffer.WriteString(`</td>
      <td class="uk-table-shrink uk-text-nowrap">`)
		ComponentDateTime(&model.Created, buffer)
		buffer.WriteString(`</td>
    </tr>
  `)
	}
	buffer.WriteString(`
  </tbody>
</table>
`)

}
