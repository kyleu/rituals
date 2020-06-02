package socket

import "github.com/kyleu/rituals.dev/app/util"

// Client Messages
const (
	ClientCmdPing = "ping"

	ClientCmdConnect = "connect"

	ClientCmdGetActions = "get-actions"
	ClientCmdGetTeams   = "get-teams"
	ClientCmdGetSprints = "get-sprints"

	ClientCmdUpdateSession = "update-session"

	ClientCmdAddComment    = "add-comment"
	ClientCmdUpdateComment = "update-comment"
	ClientCmdRemoveComment = "remove-comment"

	ClientCmdUpdateProfile = "update-profile"
	ClientCmdRemoveMember  = "remove-member"

	ClientCmdAddStory       = "add-story"
	ClientCmdUpdateStory    = "update-story"
	ClientCmdRemoveStory    = "remove-story"
	ClientCmdSetStoryStatus = "set-story-status"
	ClientCmdSubmitVote     = "submit-vote"

	ClientCmdAddReport    = "add-report"
	ClientCmdUpdateReport = "update-report"
	ClientCmdRemoveReport = "remove-report"

	ClientCmdAddFeedback    = "add-feedback"
	ClientCmdUpdateFeedback = "update-feedback"
	ClientCmdRemoveFeedback = "remove-feedback"
)

// Server Messages
const (
	ServerCmdError = util.KeyError
	ServerCmdPong  = "pong"

	ServerCmdSessionJoined = "session-joined"
	ServerCmdSessionUpdate = "session-update"

	ServerCmdCommentUpdate     = "comment-update"
	ServerCmdCommentRemove     = "comment-remove"

	ServerCmdPermissionsUpdate = "permissions-update"
	ServerCmdTeamUpdate        = "team-update"
	ServerCmdSprintUpdate      = "sprint-update"
	ServerCmdContentUpdate     = "content-update"

	ServerCmdActions = "actions"
	ServerCmdTeams   = "teams"
	ServerCmdSprints = "sprints"

	ServerCmdMemberUpdate = "member-update"
	ServerCmdOnlineUpdate = "online-update"

	ServerCmdStoryUpdate       = "story-update"
	ServerCmdStoryRemove       = "story-remove"
	ServerCmdStoryStatusChange = "story-status-change"
	ServerCmdVoteUpdate        = "vote-update"

	ServerCmdReportUpdate = "report-update"
	ServerCmdReportRemove = "report-remove"

	ServerCmdFeedbackUpdate = "feedback-update"
	ServerCmdFeedbackRemove = "feedback-remove"
)
