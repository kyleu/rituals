namespace member {
  function renderMember(member: Member) {
    const profile = system.cache.getProfile();
    return <div class="section" onclick={`modal.open('member', '${member.userID}');`}>
      <div title="user is offline" class="right uk-article-meta online-indicator">offline</div>
      <div class={`${profile.linkColor}-fg section-link`}>{renderTitle(member)}</div>
    </div>;
  }

  export function renderMembers(members: ReadonlyArray<Member>) {
    if (members.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="modal.open('invitation');" type="button">Invite Members</button>
      </div>;
    } else {
      return <ul class="uk-list uk-list-divider">
        {members.map(m => <li id={`member-${m.userID}`}>
          {renderMember(m)}
        </li>)}
      </ul>;
    }
  }

  export function renderTitle(member?: Member) {
    if (!member) {
      return <span>{"{former member}"}</span>
    }
    if (member.picture && member.picture.length > 0) {
      return <div>
        <div class="profile-image"><img class="uk-border-circle" src={member.picture} alt={member.name} /></div>
        <div class="left">{member.name}</div>
        <div class="clear"/>
      </div>;
    }
    return <div>
      <div class="profile-image"><span class="profile-icon uk-icon" data-uk-icon="user" /></div>
      <div class="left">{member.name}</div>
    </div>;
  }

  export function renderHeader(m: member.Member | undefined, t: string, close?: JSX.Element) {
    return <header class="uk-comment-header uk-position-relative">
      {close ? close : <span/>}
      <div class="uk-grid-collapse uk-flex-middle" uk-grid>
        <div class="uk-width-auto">
          {(m && m.picture && m.picture.length > 0) ? <div>
            <div class="profile-image"><img class="uk-border-circle" src={m.picture} alt={m.name} /></div>
          </div> : <div>
            <div class="profile-image"><span class="profile-icon uk-icon" data-uk-icon="user" /></div>
          </div>}
        </div>
        <div class="uk-width-expand">
          <h4 class="uk-comment-title uk-margin-remove">{ m?.name }</h4>
          <p class="uk-comment-meta uk-margin-remove-top">{ date.toDateTimeString(new Date(t)) }</p>
        </div>
      </div>
    </header>
  }

  export function viewSelf() {
    const selfInput = dom.setValue("#self-name-input", dom.req("#member-self .member-name").innerText);
    setTimeout(() => selfInput.focus(), 250);
  }

  export function setPicture(url?: string) {
    if (url && url.length > 0) {
      return <div class="model-icon profile-image"><img class="uk-border-circle" src={url} alt="your picture" /></div>
    }
    return <span class="model-icon h3-icon" onclick="modal.open('self');" data-uk-icon="icon: user;"/>
  }
}
