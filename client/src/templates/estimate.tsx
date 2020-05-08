function renderStory(story: Story): JSX.Element {
  let profile = systemCache.profile;
  if (profile === undefined) {
    return <li>profile error</li>;
  } else {
    return <li id={"story-" + story.id}>
      <div class="right uk-article-meta story-status">{story.status.key}</div>
      <a class={profile.linkColor + "-fg"} href="" onclick={"estimateCache.activeStory = '" + story.id + "';openModal('story');return false;"}>{story.title}</a>
    </li>;
  }
}

function renderStories(stories: Story[]): JSX.Element {
  if (stories.length === 0) {
    return <div>
      <button class="uk-button uk-button-default" onclick="openModal('add-story');" type="button">Add Story</button>
    </div>;
  } else {
    return <ul class="uk-list uk-list-divider">
      {stories.map(s => renderStory(s))}
    </ul>;
  }
}

function renderVoteMember(member: Member, hasVote: boolean): JSX.Element {
  return <div class="vote-member" title={ member.name + " has " + (hasVote ? "voted" : "not voted") }>
    <div>
      <span data-uk-icon={"icon: " + (hasVote ? "check" : "minus") + "; ratio: 1.6"} />
    </div>
    {member.name}
  </div>;
}

function renderVoteMembers(members: Member[], votes: Vote[]): JSX.Element {
  return <div class="uk-flex uk-flex-wrap uk-flex-around">
    {members.map(m => renderVoteMember(m, votes.filter(v => v.userID == m.userID).length > 0))}
  </div>;
}

function renderVoteChoices(choices: string[], choice: string | undefined): JSX.Element {
  return <div class="uk-flex uk-flex-wrap uk-flex-center">
    {choices.map(c => <div class={ "vote-choice uk-border-circle uk-box-shadow-hover-medium" + (c === choice ? " active" : "") } onclick={ "onSubmitVote('" + c + "');" }>{c}</div>)}
  </div>;
}

