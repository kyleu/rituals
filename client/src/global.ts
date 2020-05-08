declare var UIkit: any;

const debug = true;

let appUnloading = false;

class SystemCache {
  profile?: Profile;
  session?: Session;

  activeMember?: string;

  currentService = "";
  currentID = "";
  connectTime = 0;

  detail?: EstimateDetail;

  members: Member[] = [];
  online: string[] = [];
}

const systemCache = new SystemCache();

const services = {
  system: "system",
  estimate: "estimate",
  standup: "standup",
  retro: "retro",
};

const clientCmd = {
  error: "error",
  ping: "ping",

  connect: "connect",
  updateProfile: "update-profile",

  updateSession: "update-session",

  addStory: "add-story",
  updateStory: "update-story",
  setStoryStatus: "set-story-status",
  submitVote: "submit-vote",
};

const serverCmd = {
  error: "error",
  pong: "pong",

  sessionJoined: "session-joined",
  sessionUpdate: "session-update",

  memberUpdate: "member-update",
  onlineUpdate: "online-update",

  storyUpdate: "story-update",
  storyStatusChange: "story-status-change",
  voteUpdate: "vote-update",
};
