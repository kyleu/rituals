interface Poll {
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
  userID:  string;
  choice:  string;
  updated: string;
  created: string;
}

function setPolls(polls: Poll[]) {
  estimateCache.polls = polls;
  const detail = $id("poll-detail");
  detail.innerHTML = "";
  detail.appendChild(renderPolls(polls));

  UIkit.modal("#modal-add-poll").hide();
}

function setVotes(votes: Vote[]) {
  console.log("todo: votes")
}

function onSubmitPoll() {
  const title = $req<HTMLInputElement>("#poll-title-input").value;
  const msg = {
    svc: services.estimate,
    cmd: clientCmd.addPoll,
    param: {
      title: title
    }
  }
  send(msg);
}

function getActivePoll() {
  if (estimateCache.activePoll === null) {
    console.warn("no active poll")
    return null;
  }
  const curr = estimateCache.polls.filter(x => x.id === estimateCache.activePoll);
  if (curr.length !== 1) {
    console.log("cannot load active poll [" + estimateCache.activePoll + "]");
    return null;
  }
  return curr[0];
}

function viewActivePoll() {
  const poll = getActivePoll();
  if (poll === null) {
    console.log("no active poll");
    return;
  }
  $req("#poll-title").innerText = poll.title;
}
