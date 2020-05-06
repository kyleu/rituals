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
    case "polls":
      setPolls(param)
      break;
    case "votes":
      setVotes(param)
      break;
    default:
      console.warn("unhandled command [" + cmd + "] for estimate")
  }
}

function setEstimateDetail(param: EstimateDetail) {
  $id<HTMLInputElement>("model-choices-input").value = param.choices.join(", ");
  setDetail(param);
}

function setPolls(polls: Poll[]) {
  const detail = $id("poll-detail");
  detail.innerHTML = "";
  detail.appendChild(renderPolls(polls));

  UIkit.modal("#modal-poll").hide();
}

function setVotes(votes: Vote[]) {
  const detail = $id("vote-detail");
  detail.innerHTML = "";
  detail.appendChild(renderVotes(votes));
}

function onSubmitEstimateSession() {
  let title = $req<HTMLInputElement>("#model-title-input").value;
  let choices = $req<HTMLInputElement>("#model-choices-input").value;
  let msg = {
    svc: "estimate",
    cmd: "session-save",
    param: {
      title: title,
      choices: choices
    }
  }
  send(msg);
}

function onSubmitPoll() {
  let title = $req<HTMLInputElement>("#poll-title-input").value;
  let msg = {
    svc: "estimate",
    cmd: "new-poll-save",
    param: {
      title: title
    }
  }
  send(msg);
}
