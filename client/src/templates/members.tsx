function renderMember(member: Member): any {
  return <div>
    <hr />
    <div>user: { member.userID }</div>
    <div>name: { member.name }</div>
    <div>role: { member.role.key }</div>
    <div>created: { member.created }</div>
    <pre>{ JSON.stringify(member, null, 2) }</pre>
  </div>
}

function renderSelf(self: Member): any {
  return <div>
    {renderMember(self)}
  </div>;
}

function renderMembers(members: Member[]): any {
  return <div>
    {members.map(m => renderMember(m))}
  </div>;
}
