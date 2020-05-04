interface EstimateDetail extends Detail{
  choices: string[];
  options: object;
}

interface Poll {
  id: string;
  idx: number;
  author: string;
  title: string;
  status: { key: string; };
  finalVote: string;
  created: string;
}

interface Vote {
  userID:  string;
  choice:  string;
  updated: string;
  created: string;
}

function onEstimateMessage(cmd: string, param: any) {
  switch(cmd) {
    case "detail":
      setEstimateDetail(param);
      break;
    case "members":
      setMembers(param);
      break;
    case "polls":
      setPolls(param)
      break;
    case "votes":
      setVotes(param)
      break;
    default:
      console.warn("Unhandled command [" + cmd + "] for estimate")
  }
}

function setEstimateDetail(param: EstimateDetail) {
  setDetail(param);
}

function setPolls(polls: Poll[]) {
  const detail = $id("poll-detail");
  detail.innerHTML = "";
  detail.appendChild(renderPolls(polls));
}

function setVotes(votes: Vote[]) {
  const detail = $id("vote-detail");
  detail.innerHTML = "";
  detail.appendChild(renderVotes(votes));
}
