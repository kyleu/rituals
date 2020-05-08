interface Story {
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

interface Vote {
  storyID: string;
  userID: string;
  choice: string;
  updated: string;
  created: string;
}

function setStories(stories: Story[]) {
  estimateCache.stories = stories;
  const detail = $id("story-detail");
  detail.innerHTML = "";
  detail.appendChild(renderStories(stories));

  UIkit.modal("#modal-add-story").hide();
}

function setVotes(votes: Vote[]) {
  estimateCache.votes = votes;
  if (estimateCache.activeStory) {
    viewActiveVotes();
  }
}

function onSubmitStory() {
  const title = $req<HTMLInputElement>("#story-title-input").value;
  const msg = {
    svc: services.estimate,
    cmd: clientCmd.addStory,
    param: { title: title },
  };
  send(msg);
}

function getActiveStory() {
  if (estimateCache.activeStory === undefined) {
    console.warn("no active story");
    return undefined;
  }
  const curr = estimateCache.stories.filter(x => x.id === estimateCache.activeStory);
  if (curr.length !== 1) {
    console.warn("cannot load active story [" + estimateCache.activeStory + "]");
    return undefined;
  }
  return curr[0];
}

function viewActiveStory() {
  const story = getActiveStory();
  if (story === undefined) {
    console.log("no active story");
    return;
  }
  $req("#story-title").innerText = story.title;
  viewStoryStatus(story.status.key);
}

function viewStoryStatus(status: string) {
  switch (status) {
    case "pending":
      break;
    case "active":
      viewActiveVotes();
      break;
    case "complete":
      break;
  }
  for (let el of $(".story-status-section")) {
    const s = el.id.substr(el.id.lastIndexOf("-") + 1);
    if (s === status) {
      el.classList.add("active");
    } else {
      el.classList.remove("active");
    }
  }
}

function onVoteUpdate(vote: Vote) {
  let x = estimateCache.votes;
  x = x.filter(v => v.userID != vote.userID || v.storyID != vote.storyID);
  x.push(vote);
  estimateCache.votes = x;
  if(vote.storyID === estimateCache.activeStory) {
    viewActiveVotes();
  }
}

function viewActiveVotes() {
  const votes = estimateCache.activeVotes();
  const activeVote = votes.filter(v => v.userID === systemCache.profile!.userID).pop();
  const m = $id("story-vote-members");
  m.innerHTML = "";
  m.appendChild(renderVoteMembers(systemCache.members, votes));

  const c = $id("story-vote-choices");
  c.innerHTML = "";
  c.appendChild(renderVoteChoices(estimateCache.detail!.choices, activeVote?.choice));
}

function requestStoryStatus(s: string) {
  const story = getActiveStory();
  if (story === undefined) {
    console.log("no active story");
    return;
  }
  const msg = {
    svc: services.estimate,
    cmd: clientCmd.setStoryStatus,
    param: { storyID: story.id, status: s },
  };
  send(msg);
}

function onStoryStatusChange(u: StoryStatusUpdate) {
  $req("#story-" + u.storyID + " .story-status").innerText = u.status.key;
  viewStoryStatus(u.status.key);
}

function onSubmitVote(choice: string) {
  const msg = {
    svc: services.estimate,
    cmd: clientCmd.submitVote,
    param: { storyID: estimateCache.activeStory, choice: choice },
  };
  send(msg);
}

