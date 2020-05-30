// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/admin/retro/card.html
// DO NOT EDIT!
package templates

import (
	"bytes"

	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

func AdminRetroCard(retros retro.Sessions, params query.ParamSet, ctx web.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<div class="uk-card uk-card-body uk-margin-top">
  <h3 class="uk-card-title">Retros</h3>
  `)
	if len(retros) == 0 {
		buffer.WriteString(`
    <p>No retros available</p>
  `)
	} else {
		AdminRetroTable(retros, params.Get(util.SvcRetro.Key, ctx.Logger), ctx, buffer)
	}
	buffer.WriteString(`
</div>
`)

}
