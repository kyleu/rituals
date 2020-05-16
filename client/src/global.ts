namespace system {
  class Cache {
    profile?: rituals.Profile;
    session?: rituals.Session;

    activeMember?: string;

    currentService = "";
    currentID = "";
    connectTime = 0;

    members: member.Member[] = [];
    online: string[] = [];

    public getProfile(): rituals.Profile {
      if(this.profile === undefined) {
        throw "no active profile";
      }
      return this.profile;
    }
  }

  export function getMemberName(id: string) {
    const ret = cache.members.filter(m => m.userID === id);
    if(ret.length === 0) {
      return id;
    }
    return ret[0].name;
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
  export const sprint = {
    key: "sprint",
    title: "Sprint",
    plural: "sprints",
    icon: "git-fork"
  };
  export const estimate = {
    key: "estimate",
    title: "Estimate",
    plural: "estimates",
    icon: "settings"
  };
  export const standup = {
    key: "standup",
    title: "Standup",
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
    getActions: "get-actions",
    updateProfile: "update-profile",

    updateSession: "update-session",

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

    actions: "actions",
    sessionJoined: "session-joined",
    sessionUpdate: "session-update",

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
