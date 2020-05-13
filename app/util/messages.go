package util

// Client Messages
const ClientCmdError = "error"
const ClientCmdPing = "ping"

const ClientCmdConnect = "connect"
const ClientCmdUpdateProfile = "update-profile"
const ClientCmdUpdateSession = "update-session"

const ClientCmdAddStory = "add-story"
const ClientCmdUpdateStory = "update-story"
const ClientCmdRemoveStory = "remove-story"
const ClientCmdSetStoryStatus = "set-story-status"
const ClientCmdSubmitVote = "submit-vote"

const ClientCmdAddReport = "add-report"
const ClientCmdUpdateReport = "update-report"
const ClientCmdRemoveReport = "remove-report"

const ClientCmdAddFeedback = "add-feedback"
const ClientCmdUpdateFeedback = "update-feedback"
const ClientCmdRemoveFeedback = "remove-feedback"

// Server Messages
const ServerCmdError = "error"
const ServerCmdPong = "pong"

const ServerCmdSessionJoined = "session-joined"
const ServerCmdSessionUpdate = "session-update"

const ServerCmdMemberUpdate = "member-update"
const ServerCmdOnlineUpdate = "online-update"

const ServerCmdStoryUpdate = "story-update"
const ServerCmdStoryRemove = "story-remove"
const ServerCmdStoryStatusChange = "story-status-change"
const ServerCmdVoteUpdate = "vote-update"

const ServerCmdReportUpdate = "report-update"
const ServerCmdReportRemove = "report-remove"

const ServerCmdFeedbackUpdate = "feedback-update"
const ServerCmdFeedbackRemove = "feedback-remove"
