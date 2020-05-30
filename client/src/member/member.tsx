namespace member {
  function renderMember(member: Member): JSX.Element {
    const profile = system.cache.getProfile();
    return <div class="section" onclick={`modal.open('member', '${member.userID}');`}>
      <div title="user is offline" class="right uk-article-meta online-indicator">offline</div>
      <div class={`${profile.linkColor}-fg section-link`}>{member.name}</div>
    </div>;
  }

  export function renderMembers(members: Member[]): JSX.Element {
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

  export function viewSelf() {
    const selfInput = dom.setValue("#self-name-input", dom.req("#member-self .member-name").innerText);
    setTimeout(() => selfInput.focus(), 250);
  }
}
