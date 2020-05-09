namespace story {
  function renderStory(story: story.Story): JSX.Element {
    let profile = system.cache.profile;
    if (profile === undefined) {
      return <li>profile error</li>;
    } else {
      return <li id={"story-" + story.id}>
        <div class="right uk-article-meta story-status">{story.status.key}</div>
        <a class={profile.linkColor + "-fg"} href="" onclick={"return events.openModal('story', '" + story.id + "');"}>{story.title}</a>
      </li>;
    }
  }

  export function renderStories(stories: story.Story[]): JSX.Element {
    if (stories.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="events.openModal('add-story');" type="button">Add Story</button>
      </div>;
    } else {
      return <ul class="uk-list uk-list-divider">
        {stories.map(s => renderStory(s))}
      </ul>;
    }
  }

  function renderVoteMember(member: member.Member, hasVote: boolean): JSX.Element {
    return <div class="vote-member" title={member.name + " has " + (hasVote ? "voted" : "not voted")}>
      <div>
        <span data-uk-icon={"icon: " + (hasVote ? "check" : "minus") + "; ratio: 1.6"}/>
      </div>
      {member.name}
    </div>;
  }

  export function renderVoteMembers(members: member.Member[], votes: story.Vote[]): JSX.Element {
    return <div class="uk-flex uk-flex-wrap uk-flex-around">
      {members.map(m => renderVoteMember(m, votes.filter(v => v.userID == m.userID).length > 0))}
    </div>;
  }

  export function renderVoteChoices(choices: string[], choice: string | undefined): JSX.Element {
    return <div class="uk-flex uk-flex-wrap uk-flex-center">
      {choices.map(c => <div
        class={"vote-choice uk-border-circle uk-box-shadow-hover-medium" + (c === choice ? " active " + system.cache.profile!.linkColor + "-border" : "")}
        onclick={"story.onSubmitVote('" + c + "');"}>{c}</div>)}
    </div>;
  }

  function renderVoteResult(member: member.Member, choice: string | undefined): JSX.Element {
    if (choice === undefined) {
      return <div class="vote-result"><div><span class="uk-border-circle"><span data-uk-icon="icon: minus; ratio: 1.6" /></span></div> {member.name}</div>;
    }
    return <div class="vote-result"><div><span class="uk-border-circle">{choice}</span></div> {member.name}</div>;
  }

  export function renderVoteResults(members: member.Member[], votes: story.Vote[]): JSX.Element {
    return <div class="uk-flex uk-flex-wrap uk-flex-around">
      {members.map(m => {
        let vote = votes.filter(v => v.userID == m.userID);
        return renderVoteResult(m, vote.length > 0 ? vote[0].choice : undefined);
      })}
    </div>;
  }

  export function renderVoteSummary(votes: story.Vote[]): JSX.Element {
    const results = getVoteResults(votes);
    function trim(n: number) { return n.toString().substr(0, 4) }
    return <div class="uk-flex uk-flex-wrap uk-flex-center result-container">
      <div class="result"><div class="secondary uk-border-circle">{trim(results.count)} / {trim(votes.length)}</div> <div>votes counted</div></div>
      <div class="result"><div class="secondary uk-border-circle">{trim(results.min)}-{trim(results.max)}</div> <div>vote range</div></div>
      <div class="result mean-result"><div class={ "mean uk-border-circle " + system.cache.profile?.linkColor + "-border" }>{trim(results.mean)}</div> <div>average</div></div>
      <div class="result"><div class="secondary uk-border-circle">{trim(results.median)}</div> <div>median</div></div>
      <div class="result"><div class="secondary uk-border-circle">{trim(results.mode)}</div> <div>mode</div></div>
    </div>;
  }
}
