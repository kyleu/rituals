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
  let profile = activeProfile
  if(profile == null) {
    return <div class="uk-margin-bottom">error</div>
  } else {
    let b = Math.random() >= 0.5;
    return <li>
      <div title="user is offline" id={"online-status-" + member.userID} class="right uk-article-meta online-indicator">offline</div>
      <a class={profile.linkColor + "-fg"} href="" onclick={"activeMember = '" + member.userID + "';"} data-uk-toggle="target: #modal-member">{member.name}</a>
    </li>
  }
}

function renderMembers(members: Member[]): any {
  return <ul class="uk-list uk-list-divider">
    {members.map(m => renderMember(m))}
  </ul>;
}
