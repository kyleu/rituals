function debugMember(member: Member): any {
  return <div>
    <hr />
    <div>user: { member.userID }</div>
    <div>name: { member.name }</div>
    <div>role: { member.role.key }</div>
    <div>created: { member.created }</div>
    <pre>{ JSON.stringify(member, null, 2) }</pre>
  </div>
}

function renderMember(member: Member): any {
  let profile = systemCache.profile
  if(profile == null) {
    return <div class="uk-margin-bottom">error</div>
  } else {
    let b = Math.random() >= 0.5;
    return <div>
      <div title="user is offline" class="right uk-article-meta online-indicator">offline</div>
      <a class={profile.linkColor + "-fg"} href="" onclick={"systemCache.activeMember = '" + member.userID + "';"} data-uk-toggle="target: #modal-member">{member.name}</a>
    </div>
  }
}

function renderMembers(members: Member[]): any {
  return <ul class="uk-list uk-list-divider">
    { members.map(m => <li id={ "member-" + m.userID }>
      { renderMember(m) }
    </li>) }
  </ul>;
}
