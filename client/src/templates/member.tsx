namespace member {
  function renderMember(member: Member): JSX.Element {
    const profile = system.cache.getProfile();
    return <div class="section" onclick={"events.openModal('member', '" + member.userID + "');"}>
      <div title="user is offline" class="right uk-article-meta online-indicator">offline</div>
      <div class={profile.linkColor + "-fg section-link"}>{member.name}</div>
    </div>;
  }

  export function renderMembers(members: Member[]): JSX.Element {
    if (members.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="events.openModal('invite');" type="button">Invite Members</button>
      </div>;
    } else {
      return <ul class="uk-list uk-list-divider">
        {members.map(m => <li id={"member-" + m.userID}>
          {renderMember(m)}
        </li>)}
      </ul>;
    }
  }
}
