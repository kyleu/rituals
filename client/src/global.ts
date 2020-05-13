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

  export const cache = new Cache();
}

namespace services {
  export const system = "system";
  export const estimate = "estimate";
  export const standup = "standup";
  export const retro = "retro";
}

namespace command {
  export const client = {
    error: "error",
    ping: "ping",

    connect: "connect",
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
