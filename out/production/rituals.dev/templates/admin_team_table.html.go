// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/admin/team/table.html
// DO NOT EDIT!
package templates

import (
	"bytes"

	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func AdminTeamTable(teams team.Sessions, params *query.Params, ctx web.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`

<table class="uk-table uk-table-divider uk-text-left">
  <thead>
    <tr>
      `)
	ComponentTableHeader(util.SvcTeam.Key, util.KeyID, util.Title(util.KeyID), params, ctx, buffer)
	ComponentTableHeader(util.SvcTeam.Key, util.KeySlug, util.Title(util.KeySlug), params, ctx, buffer)
	ComponentTableHeader(util.SvcTeam.Key, util.KeyTitle, util.Title(util.KeyTitle), params, ctx, buffer)
	ComponentTableHeader(util.SvcTeam.Key, util.KeyOwner, util.Title(util.KeyOwner), params, ctx, buffer)
	ComponentTableHeader(util.SvcTeam.Key, util.KeyCreated, util.Title(util.KeyCreated), params, ctx, buffer)
	buffer.WriteString(`
    </tr>
  </thead>
  <tbody>
  `)
	for _, model := range teams {
		buffer.WriteString(`
    <tr>
      <td class="uk-table-shrink uk-text-nowrap"><a class="`)
		hero.EscapeHTML(ctx.Profile.LinkClass(), buffer)
		buffer.WriteString(`" href="`)
		hero.EscapeHTML(ctx.Route(util.AdminLink(util.SvcTeam.Key, `detail`), `id`, model.ID.String()), buffer)
		buffer.WriteString(`">`)
		ComponentUUID(&model.ID, buffer)
		buffer.WriteString(`</a></td>
      <td><a class="`)
		hero.EscapeHTML(ctx.Profile.LinkClass(), buffer)
		buffer.WriteString(`" href="`)
		hero.EscapeHTML(ctx.Route(util.SvcTeam.Key, util.KeyKey, model.Slug), buffer)
		buffer.WriteString(`">`)
		hero.EscapeHTML(model.Slug, buffer)
		buffer.WriteString(`</a></td>
      <td>`)
		hero.EscapeHTML(model.Title, buffer)
		buffer.WriteString(`</td>
      <td class="uk-table-shrink uk-text-nowrap"><a class="`)
		hero.EscapeHTML(ctx.Profile.LinkClass(), buffer)
		buffer.WriteString(`" href="`)
		hero.EscapeHTML(ctx.Route(util.AdminLink(util.KeyUser, `detail`), `id`, model.Owner.String()), buffer)
		buffer.WriteString(`">`)
		ComponentUUID(&model.Owner, buffer)
		buffer.WriteString(`</a></td>
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
