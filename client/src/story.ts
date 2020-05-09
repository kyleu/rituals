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

  export interface VoteResults {
    count: number;
    min: number;
    max: number;
    sum: number;
    mean: number;
    median: number;
    mode: number;
  }

  export function setStories(stories: Story[]) {
    estimate.cache.stories = stories;
    util.setContent("#story-detail", renderStories(stories));
    UIkit.modal("#modal-add-story").hide();
  }

  export function setVotes(votes: Vote[]) {
    estimate.cache.votes = votes;
    viewVotes();
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
    function setActive(el: HTMLElement, status: string) {
      const s = el.id.substr(el.id.lastIndexOf("-") + 1);
      if (s === status) {
        el.classList.add("active");
      } else {
        el.classList.remove("active");
      }
    }

    for (let el of util.els(".story-status-body")) {
      setActive(el, status);
    }
    for (let el of util.els(".story-status-actions")) {
      setActive(el, status);
    }

    let txt = "TODO";
    switch(status) {
      case "pending":
        txt = "Story";
        break;
      case "active":
        txt = "Voting";
        break;
      case "complete":
        txt = "Results";
        break;
    }
    util.req("#story-status").innerText = txt;

    viewVotes();
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
    const story = getActiveStory();
    if (story === undefined) {
      return;
    }
    const votes = estimate.cache.activeVotes();
    const activeVote = votes.filter(v => v.userID === system.cache.profile!.userID).pop();

    if (story.status.key == "active") {
      viewActiveVotes(votes, activeVote);
    }
    if (story.status.key == "complete") {
      viewVoteResults(votes);
    }
  }

  function viewActiveVotes(votes: story.Vote[], activeVote: story.Vote | undefined) {
    util.setContent("#story-vote-members", renderVoteMembers(system.cache.members, votes));
    util.setContent("#story-vote-choices", renderVoteChoices(estimate.cache.detail!.choices, activeVote?.choice));
  }

  function viewVoteResults(votes: story.Vote[]) {
    util.setContent("#story-vote-results", renderVoteResults(system.cache.members, votes));
    util.setContent("#story-vote-summary", renderVoteSummary(votes));
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
    estimate.cache.stories.forEach(s => {
      if (s.id == u.storyID) {
        s.status = u.status;
      }
    });

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

  export function getVoteResults(votes: Vote[]): VoteResults {
    let floats = votes.map(v => {
      let n = parseFloat(v.choice);
      if (isNaN(n)) {
        return -1;
      }
      return n;
    }).filter(x => x !== -1).sort();

    let count = floats.length;

    let min = Math.min(...floats);
    let max = Math.max(...floats);

    let sum = floats.reduce((x, y) => x + y, 0);

    var mode = floats.reduce(function(current: any, item) {
      var val = current.numMapping[item] = (current.numMapping[item] || 0) + 1;
      if (val > current.greatestFreq) {
        current.greatestFreq = val;
        current.mode = item;
      }
      return current;
    }, {mode: null, greatestFreq: -Infinity, numMapping: {}}).mode;

    return {
      count: count,
      min: min,
      max: max,
      sum: sum,
      mean: count == 0 ? 0 : sum / count,
      median: count == 0 ? 0 : floats[Math.floor(floats.length / 2)],
      mode: count == 0 ? 0 : mode
    };
  }
}
