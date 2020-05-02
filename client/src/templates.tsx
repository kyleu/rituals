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

function renderMembers(members: [Member]): any {
  return <div>
    {members.map(m => renderMember(m))}
  </div>;
}

function renderPoll(poll: Poll): any {
  return <div>
    <hr />
    <div>id: { poll.id }</div>
    <div>idx: { poll.idx }</div>
    <div>author: { poll.author }</div>
    <div>title: { poll.title }</div>
    <div>status: { poll.status.key }</div>
    <div>finalVote: { poll.finalVote }</div>
    <pre>{ JSON.stringify(poll, null, 2) }</pre>
  </div>
}

function renderPolls(polls: [Poll]): any {
  return <div>
    {polls.map(p => renderPoll(p))}
  </div>;
}

function renderVote(vote: Vote): any {
  return <pre>{ JSON.stringify(vote, null, 2) }</pre>
}

function renderVotes(votes: [Vote]): any {
  return <div>
    {votes.map(v => renderVote(v))}
  </div>;
}
