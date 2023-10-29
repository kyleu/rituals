// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/rituals/app/lib/menu"

//nolint:lll
var (
	menuItemAction              = &menu.Item{Key: "action", Title: "Actions", Description: "An action within a workspace", Icon: "action", Route: "/admin/db/action"}
	menuItemComment             = &menu.Item{Key: "comment", Title: "Comments", Description: "A comment on a model within the system", Icon: "comments", Route: "/admin/db/comment"}
	menuItemEmail               = &menu.Item{Key: "email", Title: "Emails", Description: "An email message that has been sent", Icon: "email", Route: "/admin/db/email"}
	menuItemEstimate            = &menu.Item{Key: "estimate", Title: "Estimates", Description: "Planning poker for any stories you need to work on", Icon: "estimate", Route: "/admin/db/estimate", Children: menu.Items{menuItemEstimateEhistory, menuItemEstimateEmember, menuItemEstimateEpermission, menuItemEstimateStory}}
	menuItemEstimateEhistory    = &menu.Item{Key: "ehistory", Title: "Histories", Description: "Historical names and URLs for estimate sessions", Icon: "history", Route: "/admin/db/estimate/history"}
	menuItemEstimateEmember     = &menu.Item{Key: "emember", Title: "Members", Description: "Membership roster for estimate sessions", Icon: "users", Route: "/admin/db/estimate/member"}
	menuItemEstimateEpermission = &menu.Item{Key: "epermission", Title: "Permissions", Description: "Security controls for estimate sessions", Icon: "permission", Route: "/admin/db/estimate/permission"}
	menuItemEstimateStory       = &menu.Item{Key: "story", Title: "Stories", Description: "The detailed use cases for an estimate session", Icon: "story", Route: "/admin/db/estimate/story", Children: menu.Items{menuItemStoryVote}}
	menuItemRetro               = &menu.Item{Key: "retro", Title: "Retros", Description: "Discover improvements and praise for your work", Icon: "retro", Route: "/admin/db/retro", Children: menu.Items{menuItemRetroFeedback, menuItemRetroRhistory, menuItemRetroRmember, menuItemRetroRpermission}}
	menuItemRetroFeedback       = &menu.Item{Key: "feedback", Title: "Feedbacks", Description: "User feedback for a retrospective", Icon: "comment", Route: "/admin/db/retro/feedback"}
	menuItemRetroRhistory       = &menu.Item{Key: "rhistory", Title: "Histories", Description: "Historical names and URLs for retrospectives", Icon: "history", Route: "/admin/db/retro/history"}
	menuItemRetroRmember        = &menu.Item{Key: "rmember", Title: "Members", Description: "Membership roster for retrospectives", Icon: "users", Route: "/admin/db/retro/member"}
	menuItemRetroRpermission    = &menu.Item{Key: "rpermission", Title: "Permissions", Description: "Security controls for retrospectives", Icon: "permission", Route: "/admin/db/retro/permission"}
	menuItemSprint              = &menu.Item{Key: "sprint", Title: "Sprints", Description: "Plan your time and direct your efforts", Icon: "sprint", Route: "/admin/db/sprint", Children: menu.Items{menuItemSprintShistory, menuItemSprintSmember, menuItemSprintSpermission}}
	menuItemSprintShistory      = &menu.Item{Key: "shistory", Title: "Histories", Description: "Historical names and URLs for sprints", Icon: "history", Route: "/admin/db/sprint/history"}
	menuItemSprintSmember       = &menu.Item{Key: "smember", Title: "Members", Description: "Membership roster for sprints", Icon: "users", Route: "/admin/db/sprint/member"}
	menuItemSprintSpermission   = &menu.Item{Key: "spermission", Title: "Permissions", Description: "Security controls for sprints", Icon: "permission", Route: "/admin/db/sprint/permission"}
	menuItemStandup             = &menu.Item{Key: "standup", Title: "Standups", Description: "Share your progress with your team", Icon: "standup", Route: "/admin/db/standup", Children: menu.Items{menuItemStandupReport, menuItemStandupUhistory, menuItemStandupUmember, menuItemStandupUpermission}}
	menuItemStandupReport       = &menu.Item{Key: "report", Title: "Reports", Description: "Daily status reports for a standup", Icon: "file-alt", Route: "/admin/db/standup/report"}
	menuItemStandupUhistory     = &menu.Item{Key: "uhistory", Title: "Histories", Description: "Historical names and URLs for standups", Icon: "history", Route: "/admin/db/standup/history"}
	menuItemStandupUmember      = &menu.Item{Key: "umember", Title: "Members", Description: "Membership roster for standups", Icon: "users", Route: "/admin/db/standup/member"}
	menuItemStandupUpermission  = &menu.Item{Key: "upermission", Title: "Permissions", Description: "Security controls for standups", Icon: "permission", Route: "/admin/db/standup/permission"}
	menuItemStoryVote           = &menu.Item{Key: "vote", Title: "Votes", Description: "An estimate for a story, from a user", Icon: "vote-yea", Route: "/admin/db/estimate/story/vote"}
	menuItemTeam                = &menu.Item{Key: "team", Title: "Teams", Description: "Join your friends and work towards a common goal", Icon: "team", Route: "/admin/db/team", Children: menu.Items{menuItemTeamThistory, menuItemTeamTmember, menuItemTeamTpermission}}
	menuItemTeamThistory        = &menu.Item{Key: "thistory", Title: "Histories", Description: "Historical names and URLs for teams", Icon: "history", Route: "/admin/db/team/history"}
	menuItemTeamTmember         = &menu.Item{Key: "tmember", Title: "Members", Description: "Membership roster for teams", Icon: "users", Route: "/admin/db/team/member"}
	menuItemTeamTpermission     = &menu.Item{Key: "tpermission", Title: "Permissions", Description: "Security controls for teams", Icon: "permission", Route: "/admin/db/team/permission"}
	menuItemUser                = &menu.Item{Key: "user", Title: "Users", Description: "A user of the system", Icon: "profile", Route: "/admin/db/user"}
)

//nolint:unused
func generatedMenu() menu.Items {
	return menu.Items{
		menuItemAction,
		menuItemComment,
		menuItemEmail,
		menuItemEstimate,
		menuItemRetro,
		menuItemSprint,
		menuItemStandup,
		menuItemTeam,
		menuItemUser,
	}
}
