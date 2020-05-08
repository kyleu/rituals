interface EstimateDetail extends Session {
  choices: string[];
  options: object;
}

interface StoryStatusUpdate extends Session {
  storyID: string;
  status: { key: string };
}

interface EstimateSessionJoined extends SessionJoined{
  session: EstimateDetail;
  stories: Story[];
  votes: Vote[];
}

class EstimateCache {
  activeStory?: string;

  detail?: EstimateDetail;

  stories: Story[] = [];
  votes: Vote[] = [];

  activeVotes(): Vote[] {
    if(this.activeStory === undefined) {
      console.log("!!!")
      return [];
    }
    return this.votes.filter(x => x.storyID == this.activeStory);
  }
}

const estimateCache = new EstimateCache();

function onEstimateMessage(cmd: string, param: any) {
  switch(cmd) {
    case serverCmd.sessionJoined:
      let sj = param as EstimateSessionJoined
      onSessionJoin(sj);
      setEstimateDetail(sj.session);
      setStories(sj.stories);
      setVotes(sj.votes);
      break;
    case serverCmd.sessionUpdate:
      setEstimateDetail(param as EstimateDetail);
      break;
    case serverCmd.storyUpdate:
      onStoryUpdate(param as Story);
      break;
    case serverCmd.storyStatusChange:
      onStoryStatusChange(param as StoryStatusUpdate);
      break;
    case serverCmd.voteUpdate:
      onVoteUpdate(param as Vote);
      break;
    default:
      console.warn("unhandled command [" + cmd + "] for estimate");
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
      choices: choices,
    },
  };
  send(msg);
}

function getStory(id: string) {
  return estimateCache.stories.filter((x) => x.id === id).pop();
}

function onStoryUpdate(story: Story) {
  let x = estimateCache.stories;

  x = x.filter((p) => p.id !== story.id);
  x.push(story);
  x = x.sort((l, r) => (l.idx > r.idx ? 1 : -1));

  setStories(x);
}


