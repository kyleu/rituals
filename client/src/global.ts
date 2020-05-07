declare var UIkit: any;

const debug = true;

let appUnloading = false;

class SystemCache {
  profile: Profile | null = null;

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
  retro: "retro"
}

const clientCmd = {
  error: "error",
  ping: "ping",

  connect: "connect",
  updateProfile: "update-profile",

  updateSession: "update-session",

  addPoll: "add-poll",
  updatePoll: "update-poll",
  setPollStatus: "set-poll-status",
  submitVote: "submit-vote"
}

const serverCmd = {
  error: "error",
  pong: "pong",

  sessionJoined: "session-joined",
  sessionUpdate: "session-update",

  memberUpdate: "member-update",
  onlineUpdate: "online-update",

  pollUpdate: "poll-update",
  voteUpdate: "vote-update"
}
