var debug = true;

namespace system {
  class Cache {
    profile?: rituals.Profile;
    session?: rituals.Session;

    activeMember?: string;

    currentService = "";
    currentID = "";
    connectTime = 0;

    permissions: permission.Permission[] = [];
    auths: auth.Auth[] = [];
    members: member.Member[] = [];
    online: string[] = [];

    public getProfile(): rituals.Profile {
      if (this.profile === undefined) {
        throw "no active profile";
      }
      return this.profile;
    }
  }

  export function getMemberName(id: string) {
    const ret = cache.members.filter(m => m.userID === id).shift();
    if (ret) {
      return ret.name;
    }
    return id;
  }

  export const cache = new Cache();
}

namespace services {
  export interface Service {
    key: string,
    title: string,
    plural: string,
    icon: string,
  }

  export const system = {
    key: "system",
    title: "System",
    plural: "systems",
    icon: "close"
  };
  export const team = {
    key: "team",
    title: "Team",
    plural: "teams",
    icon: "users"
  };
  export const sprint = {
    key: "sprint",
    title: "Sprint",
    plural: "sprints",
    icon: "git-fork"
  };
  export const estimate = {
    key: "estimate",
    title: "Estimate Session",
    plural: "estimates",
    icon: "settings"
  };
  export const standup = {
    key: "standup",
    title: "Daily Standup",
    plural: "standups",
    icon: "future"
  };
  export const retro = {
    key: "retro",
    title: "Retrospective",
    plural: "retros",
    icon: "history"
  };
}

namespace command {
  export const client = {
    error: "error",
    ping: "ping",

    connect: "connect",
    updateSession: "update-session",

    getActions: "get-actions",
    getTeams: "get-teams",
    getSprints: "get-sprints",

    updateProfile: "update-profile",

    addStory: "add-story",
    updateStory: "update-story",
    removeStory: "remove-story",
    setStoryStatus: "set-story-status",
    submitVote: "submit-vote",

    addReport: "add-report",
    updateReport: "update-report",
    removeReport: "remove-report",

    addFeedback: "add-feedback",
    updateFeedback: "update-feedback",
    removeFeedback: "remove-feedback",
  };

  export const server = {
    error: "error",
    pong: "pong",

    sessionJoined: "session-joined",
    sessionUpdate: "session-update",
    teamUpdate: "team-update",
    sprintUpdate: "sprint-update",
    contentUpdate: "content-update",

    actions: "actions",
    teams: "teams",
    sprints: "sprints",

    memberUpdate: "member-update",
    onlineUpdate: "online-update",

    storyUpdate: "story-update",
    storyRemove: "story-remove",
    storyStatusChange: "story-status-change",
    voteUpdate: "vote-update",

    reportUpdate: "report-update",
    reportRemove: "report-remove",

    feedbackUpdate: "feedback-update",
    feedbackRemove: "feedback-remove",
  };
}
