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

const ClientCmdAddPoll = "add-poll"
const ClientCmdUpdatePoll = "update-poll"
const ClientCmdSetPollStatus = "set-poll-status"
const ClientCmdSubmitVote = "submit-vote"

// Server Messages
const ServerCmdError = "error"
const ServerCmdPong = "pong"

const ServerCmdSessionJoined = "session-joined"
const ServerCmdSessionUpdate = "session-update"

const ServerCmdMemberUpdate = "member-update"
const ServerCmdOnlineUpdate = "online-update"

const ServerCmdPollUpdate = "poll-update"
const ServerCmdVoteUpdate = "vote-update"
