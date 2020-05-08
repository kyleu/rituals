namespace story {
  export interface Story {
    id: string;
    idx: number;
    author: string;
    title: string;
    status: {
      key: string;
    };
    finalVote: string;
    created: string;
  }

  export interface Vote {
    storyID: string;
    userID: string;
    choice: string;
    updated: string;
    created: string;
  }

  export function setStories(stories: Story[]) {
    estimate.cache.stories = stories;
    const detail = util.req("#story-detail");
    detail.innerHTML = "";
    detail.appendChild(renderStories(stories));

    UIkit.modal("#modal-add-story").hide();
  }

  export function setVotes(votes: Vote[]) {
    estimate.cache.votes = votes;
    if (estimate.cache.activeStory) {
      viewVotes();
    }
  }

  export function onSubmitStory() {
    const title = util.req<HTMLInputElement>("#story-title-input").value;
    const msg = {
      svc: services.estimate,
      cmd: command.client.addStory,
      param: {title: title},
    };
    socket.send(msg);
    return false;
  }

  function getActiveStory() {
    if (estimate.cache.activeStory === undefined) {
      console.warn("no active story");
      return undefined;
    }
    const curr = estimate.cache.stories.filter(x => x.id === estimate.cache.activeStory);
    if (curr.length !== 1) {
      console.warn("cannot load active story [" + estimate.cache.activeStory + "]");
      return undefined;
    }
    return curr[0];
  }

  export function viewActiveStory() {
    const story = getActiveStory();
    if (story === undefined) {
      console.log("no active story");
      return;
    }
    util.req("#story-title").innerText = story.title;
    viewStoryStatus(story.status.key);
  }

  function viewStoryStatus(status: string) {
    switch (status) {
      case "pending":
        break;
      case "active":
        viewVotes();
        break;
      case "complete":
        viewVotes();
        break;
    }
    for (let el of util.els(".story-status-section")) {
      const s = el.id.substr(el.id.lastIndexOf("-") + 1);
      if (s === status) {
        el.classList.add("active");
      } else {
        el.classList.remove("active");
      }
    }
  }

  export function onVoteUpdate(vote: Vote) {
    let x = estimate.cache.votes;
    x = x.filter(v => v.userID != vote.userID || v.storyID != vote.storyID);
    x.push(vote);
    estimate.cache.votes = x;
    if (vote.storyID === estimate.cache.activeStory) {
      viewVotes();
    }
  }

  export function viewVotes() {
    const votes = estimate.cache.activeVotes();
    const activeVote = votes.filter(v => v.userID === system.cache.profile!.userID).pop();

    viewActiveVotes(votes, activeVote);
    viewVoteResults(votes);
  }

  function viewActiveVotes(votes: story.Vote[], activeVote: story.Vote | undefined) {
    const m = util.req("#story-vote-members");
    m.innerHTML = "";
    m.appendChild(renderVoteMembers(system.cache.members, votes));

    const c = util.req("#story-vote-choices");
    c.innerHTML = "";
    c.appendChild(renderVoteChoices(estimate.cache.detail!.choices, activeVote?.choice));
  }

  function viewVoteResults(votes: story.Vote[]) {
    const c = util.req("#story-vote-results");
    c.innerHTML = "";
    c.appendChild(renderVoteResults(system.cache.members, votes));
  }

  export function requestStoryStatus(s: string) {
    const story = getActiveStory();
    if (story === undefined) {
      console.log("no active story");
      return;
    }
    const msg = {
      svc: services.estimate,
      cmd: command.client.setStoryStatus,
      param: {storyID: story.id, status: s},
    };
    socket.send(msg);
  }

  export function onStoryStatusChange(u: estimate.StoryStatusUpdate) {
    util.req("#story-" + u.storyID + " .story-status").innerText = u.status.key;
    viewStoryStatus(u.status.key);
  }

  // noinspection JSUnusedGlobalSymbols
  export function onSubmitVote(choice: string) {
    const msg = {
      svc: services.estimate,
      cmd: command.client.submitVote,
      param: {storyID: estimate.cache.activeStory, choice: choice},
    };
    socket.send(msg);
  }
}
