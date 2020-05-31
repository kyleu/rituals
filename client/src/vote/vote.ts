namespace vote {
  export interface Vote {
    readonly storyID: string;
    readonly userID: string;
    readonly choice: string;
    readonly updated: string;
    readonly created: string;
  }

  export interface VoteResults {
    readonly count: number;
    readonly min: number;
    readonly max: number;
    readonly sum: number;
    readonly mean: number;
    readonly median: number;
    readonly mode: number;
  }

  export function setVotes(votes: Vote[]) {
    estimate.cache.votes = votes;
    viewVotes();
  }

  export function onVoteUpdate(v: Vote) {
    let x = estimate.cache.votes;
    x = x.filter(v => v.userID !== v.userID || v.storyID !== v.storyID);
    x.push(v);
    estimate.cache.votes = x;
    if (v.storyID === estimate.cache.activeStory) {
      viewVotes();
    }
  }

  export function viewVotes() {
    const s = story.getActiveStory();
    if (!s) {
      return;
    }
    const votes = estimate.cache.activeVotes();
    const activeVote = votes.filter(v => v.userID === system.cache.getProfile().userID).pop();

    switch (s.status) {
      case "pending":
        const same = system.cache.getProfile().userID === s.userID
        dom.setDisplay("#story-edit-section", same)
        dom.setDisplay("#story-view-section", !same)
        break;
      case "active":
        viewActiveVotes(votes, activeVote);
        break;
      case "complete":
        viewVoteResults(votes);
        break;
      default:
        console.warn(`invalid story status [${s.status}]`);
    }
  }

  function viewActiveVotes(votes: ReadonlyArray<Vote>, activeVote: Vote | undefined) {
    dom.setContent("#story-vote-members", renderVoteMembers(member.getMembers(), votes));
    dom.setContent("#story-vote-choices", renderVoteChoices(estimate.cache.detail!.choices, activeVote?.choice));
  }

  function viewVoteResults(votes: ReadonlyArray<Vote>) {
    dom.setContent("#story-vote-results", renderVoteResults(member.getMembers(), votes));
    dom.setContent("#story-vote-summary", renderVoteSummary(votes));
  }

  // noinspection JSUnusedGlobalSymbols
  export function onSubmitVote(choice: string) {
    const msg = {svc: services.estimate.key, cmd: command.client.submitVote, param: {storyID: estimate.cache.activeStory, choice}};
    socket.send(msg);
  }

  export function getVoteResults(votes: ReadonlyArray<Vote>): VoteResults {
    const floats = votes.map(v => {
      const n = parseFloat(v.choice);
      if (isNaN(n)) {
        return -1;
      }
      return n;
    }).filter(x => x !== -1).sort();

    const count = floats.length;

    const min = Math.min(...floats);
    const max = Math.max(...floats);

    const sum = floats.reduce((x, y) => x + y, 0);

    const mode = floats.reduce(function (current: any, item) {
      const val = current.numMapping[item] = (current.numMapping[item] || 0) + 1;
      if (val > current.greatestFreq) {
        current.greatestFreq = val;
        current.mode = item;
      }
      return current;
    }, {mode: null, greatestFreq: -Infinity, numMapping: {}}).mode;

    return {
      count, min, max, sum,
      mean: count === 0 ? 0 : sum / count,
      median: count === 0 ? 0 : floats[Math.floor(floats.length / 2)],
      mode: count === 0 ? 0 : mode
    };
  }
}
