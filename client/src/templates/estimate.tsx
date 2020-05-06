function debugPoll(poll: Poll): any {
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

function renderPoll(poll: Poll): any {
  let profile = activeProfile
  if(profile == null) {
    return <li>error</li>
  } else {
    return <li>
      <a class={profile.linkColor + "-fg"} href="">{poll.title}</a>
    </li>
  }
}

function renderPolls(polls: Poll[]): any {
  return <ul class="uk-list uk-list-divider">
    {polls.map(p => renderPoll(p))}
  </ul>;
}

function debugVote(vote: Vote): any {
  return <pre>{ JSON.stringify(vote, null, 2) }</pre>
}

function renderVote(vote: Vote): any {
  return <li>
    { vote.userID }: { vote.choice }
  </li>
}

function renderVotes(votes: Vote[]): any {
  return <ul class="uk-list uk-list-divider">
    {votes.map(v => renderVote(v))}
  </ul>;
}
