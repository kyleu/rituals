function renderPoll(poll: Poll): any {
  let profile = systemCache.profile
  if(profile == null) {
    return <li>profile error</li>
  } else {
    return <li>
      <div id={"poll-status-" + poll.id} class="right uk-article-meta poll-status">{poll.status.key}</div>
      <a class={profile.linkColor + "-fg"} href="" onclick={"estimateCache.activePoll = '" + poll.id + "';"} data-uk-toggle="target: #modal-poll">{poll.title}</a>
    </li>
  }
}

function renderPolls(polls: Poll[]): any {
  return <ul class="uk-list uk-list-divider">
    {polls.map(p => renderPoll(p))}
  </ul>;
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
