package action

type Act string

const (
	ActCreate  Act = "create"
	ActConnect Act = "connect"
	ActUpdate  Act = "update"

	ActMemberAdd   Act = "member-add"
	ActPermissions Act = "permissions"

	ActContentAdd Act = "content-add"

	ActFeedbackAdd    Act = "feedback-add"
	ActFeedbackUpdate Act = "feedback-update"
	ActFeedbackRemove Act = "feedback-remove"

	ActReportAdd    Act = "report-add"
	ActReportUpdate Act = "report-update"
	ActReportRemove Act = "report-remove"

	ActStoryAdd    Act = "story-add"
	ActStoryUpdate Act = "story-update"
	ActStoryRemove Act = "story-remove"
	ActStoryStatus Act = "story-status"

	ActVoteAdd    Act = "vote-add"
	ActVoteUpdate Act = "vote-update"
)
