// Code generated by hero.
// source: github.com/kyleu/rituals.dev/web/templates/member/self.html
// DO NOT EDIT!
package templates

import (
	"bytes"

	"github.com/kyleu/rituals.dev/app/web"
	"github.com/shiyanhui/hero"
)

func MemberSelf(ctx web.RequestContext, buffer *bytes.Buffer) {
	buffer.WriteString(`
<div class="uk-card uk-card-body uk-transition-toggle" id="member-self">
  <h3 class="uk-card-title">
    <span class="right">
      <a class="`)
	hero.EscapeHTML(ctx.Profile.LinkClass(), buffer)
	buffer.WriteString(` uk-icon-button uk-transition-fade" data-uk-icon="pencil" onclick="events.openModal('self');return false;" title="edit name"></a>
    </span>
    <span class="model-icon h3-icon" onclick="events.openModal('self');" data-uk-icon="icon: user;"></span>
    <span class="member-name" onclick="events.openModal('self');">`)
	hero.EscapeHTML(ctx.Profile.Name, buffer)
	buffer.WriteString(`</span>
  </h3>
  <p class="member-role"></p>
</div>
<div id="modal-self" data-uk-modal>
  <div class="uk-modal-dialog uk-margin-auto-vertical">
    <button class="uk-button uk-modal-close-default" type="button" data-uk-close></button>
    <div class="uk-modal-header">
      <h2 class="uk-modal-title">Edit Self</h2>
    </div>
    <form class="uk-form-stacked" id="form-self-name" onsubmit="member.onSubmitSelf();return false;">
      <div class="uk-modal-body" data-uk-overflow-auto>
          <fieldset class="uk-fieldset">
            <div class="uk-margin">
              <label class="uk-form-label" for="self-name-input">Name</label>
              <input class="uk-input" id="self-name-input" type="text" />
            </div>

            <div class="uk-margin">
              <div>
                <label for="self-name-choice-local">
                  <input class="uk-radio" id="self-name-choice-local" checked="checked" type="radio" name="self-name-choice" value="local" />
                  Change for this session only
                </label>
              </div>
              <div>
                <label for="self-name-choice-global">
                  <input class="uk-radio" id="self-name-choice-global" type="radio" name="self-name-choice" value="global" />
                  Change global default
                </label>
              </div>
            </div>
          </fieldset>
      </div>
      <div class="uk-modal-footer uk-text-right">
        <button class="left uk-button uk-button-default" type="button" onclick="member.removeMember('self');">Leave</button>
        <button class="uk-button uk-button-default" type="submit">Save</button>
      </div>
    </form>
  </div>
</div>
`)

}
