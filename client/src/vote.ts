namespace vote {
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

  export function setVotes(votes: Vote[]) {
    estimate.cache.votes = votes;
    viewVotes();
  }

  export function onVoteUpdate(v: Vote) {
    let x = estimate.cache.votes;
    x = x.filter(v => v.userID != v.userID || v.storyID != v.storyID);
    x.push(v);
    estimate.cache.votes = x;
    if (v.storyID === estimate.cache.activeStory) {
      viewVotes();
    }
  }

  export function viewVotes() {
    const s = story.getActiveStory();
    if (s === undefined) {
      return;
    }
    const votes = estimate.cache.activeVotes();
    const activeVote = votes.filter(v => v.userID === system.cache.getProfile().userID).pop();

    if (s.status.key == "active") {
      viewActiveVotes(votes, activeVote);
    }
    if (s.status.key == "complete") {
      viewVoteResults(votes);
    }
  }

  function viewActiveVotes(votes: Vote[], activeVote: Vote | undefined) {
    util.setContent("#story-vote-members", renderVoteMembers(system.cache.members, votes));
    util.setContent("#story-vote-choices", renderVoteChoices(estimate.cache.detail!.choices, activeVote?.choice));
  }

  function viewVoteResults(votes: Vote[]) {
    util.setContent("#story-vote-results", renderVoteResults(system.cache.members, votes));
    util.setContent("#story-vote-summary", renderVoteSummary(votes));
  }

  // noinspection JSUnusedGlobalSymbols
  export function onSubmitVote(choice: string) {
    const msg = {
      svc: services.estimate,
      cmd: command.client.submitVote,
      param: {storyID: estimate.cache.activeStory, choice: choice}
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
