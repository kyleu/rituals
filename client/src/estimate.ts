interface EstimateDetail extends Session{
  choices: string[];
  options: object;
}

class EstimateCache {
  activePoll?: string;

  detail?: EstimateDetail;

  polls: Poll[] = [];
  votes: Vote[] = [];
}

const estimateCache = new EstimateCache();

function onEstimateMessage(cmd: string, param: any) {
  switch(cmd) {
    case serverCmd.sessionJoined:
      onSessionJoin(param);
      setEstimateDetail(param.session);
      setPolls(param.polls);
      setVotes(param.votes);
      break;
    case serverCmd.sessionUpdate:
      setEstimateDetail(param.session);
      break;
    case serverCmd.pollUpdate:
      onPollUpdate(param);
      break;
    default:
      console.warn("unhandled command [" + cmd + "] for estimate")
  }
}

function setEstimateDetail(detail: EstimateDetail) {
  estimateCache.detail = detail;
  $id<HTMLInputElement>("model-choices-input").value = detail.choices.join(", ");
  setDetail(detail);
}

function onSubmitEstimateSession() {
  const title = $req<HTMLInputElement>("#model-title-input").value;
  const choices = $req<HTMLInputElement>("#model-choices-input").value;
  const msg = {
    svc: services.estimate,
    cmd: clientCmd.updateSession,
    param: {
      title: title,
      choices: choices
    }
  }
  send(msg);
}

function onPollUpdate(poll: Poll) {
  let x = estimateCache.polls;

  x = x.filter(p => p.id != poll.id);
  x.push(poll);
  x = x.sort((l, r) => (l.idx > r.idx) ? 1 : -1);

  setPolls(x);
}

