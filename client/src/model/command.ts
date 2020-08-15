namespace command {
  export const client = {
    ping: "ping",

    connect: "connect",
    setActive: "set-active",

    getActions: "get-actions",
    getTeams: "get-teams",
    getSprints: "get-sprints",

    updateSession: "update-session",

    addComment: "add-comment",
    updateComment: "update-comment",
    removeComment: "remove-comment",

    updateProfile: "update-profile",
    updateMember: "update-member",
    removeMember: "remove-member",

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
    sessionRemove: "session-remove",

    commentUpdate: "comment-update",
    commentRemove: "comment-remove",

    permissionsUpdate: "permissions-update",
    teamUpdate: "team-update",
    sprintUpdate: "sprint-update",
    contentUpdate: "content-update",

    actions: "actions",
    teams: "teams",
    sprints: "sprints",

    memberUpdate: "member-update",
    memberRemove: "member-remove",
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
