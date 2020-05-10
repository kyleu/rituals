package util

const AppName = "rituals.dev"

// Services
const SvcSystem = "system"
const SvcEstimate = "estimate"
const SvcStandup = "standup"
const SvcRetro = "retro"

var AllServices = []string{SvcEstimate, SvcStandup, SvcRetro}

// Client Messages
const ClientCmdError = "error"
const ClientCmdPing = "ping"

const ClientCmdConnect = "connect"
const ClientCmdUpdateProfile = "update-profile"
const ClientCmdUpdateSession = "update-session"

const ClientCmdAddStory = "add-story"
const ClientCmdUpdateStory = "update-story"
const ClientCmdSetStoryStatus = "set-story-status"
const ClientCmdSubmitVote = "submit-vote"

const ClientCmdAddReport = "add-report"

// Server Messages
const ServerCmdError = "error"
const ServerCmdPong = "pong"

const ServerCmdSessionJoined = "session-joined"
const ServerCmdSessionUpdate = "session-update"

const ServerCmdMemberUpdate = "member-update"
const ServerCmdOnlineUpdate = "online-update"

const ServerCmdStoryUpdate = "story-update"
const ServerCmdStoryStatusChange = "story-status-change"
const ServerCmdVoteUpdate = "vote-update"

const ServerCmdReportUpdate = "report-update"
