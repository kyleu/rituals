function renderMember(member: Member): any {
  return <pre>{ JSON.stringify(member, null, 2) }</pre>
}

function renderPoll(poll: Poll): any {
  return <pre>{ JSON.stringify(poll, null, 2) }</pre>
}

function renderVote(vote: Vote): any {
  return <pre>{ JSON.stringify(vote, null, 2) }</pre>
}
