function renderMember(member: Member): JSX.Element {
  let profile = systemCache.profile;
  if (profile === undefined) {
    return <div class="uk-margin-bottom">error</div>;
  } else {
    let b = Math.random() >= 0.5;
    return <div>
      <div title="user is offline" class="right uk-article-meta online-indicator">offline</div>
      <a class={profile.linkColor + "-fg"} href="" onclick={"systemCache.activeMember = '" + member.userID + "';openModal('member');return false;"}>{member.name}</a>
    </div>;
  }
}

function renderMembers(members: Member[]): JSX.Element {
  if (members.length === 0) {
    return <div>
      <button class="uk-button uk-button-default" onclick="openModal('invite');" type="button">Invite Members</button>
    </div>;
  } else {
    return <ul class="uk-list uk-list-divider">
      {members.map(m => <li id={"member-" + m.userID}>
        {renderMember(m)}
      </li>)}
    </ul>;
  }
}
