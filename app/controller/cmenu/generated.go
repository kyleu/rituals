// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/rituals/app/lib/menu"

//nolint:lll
func generatedMenu() menu.Items {
	return menu.Items{
		&menu.Item{Key: "action", Title: "Actions", Description: "An action within a workspace", Icon: "action", Route: "/admin/db/action"},
		&menu.Item{Key: "comment", Title: "Comments", Description: "A comment on a model within the system", Icon: "comments", Route: "/admin/db/comment"},
		&menu.Item{Key: "email", Title: "Emails", Description: "An email message that has been sent", Icon: "email", Route: "/admin/db/email"},
		&menu.Item{Key: "estimate", Title: "Estimates", Description: "Planning poker for any stories you need to work on", Icon: "estimate", Route: "/admin/db/estimate", Children: menu.Items{
			&menu.Item{Key: "ehistory", Title: "Histories", Description: "Historical names and URLs for estimate sessions", Icon: "history", Route: "/admin/db/estimate/history"},
			&menu.Item{Key: "emember", Title: "Members", Description: "Membership roster for estimate sessions", Icon: "users", Route: "/admin/db/estimate/member"},
			&menu.Item{Key: "epermission", Title: "Permissions", Description: "Security controls for estimate sessions", Icon: "permission", Route: "/admin/db/estimate/permission"},
			&menu.Item{Key: "story", Title: "Stories", Description: "The detailed use cases for an estimate session", Icon: "story", Route: "/admin/db/estimate/story", Children: menu.Items{
				&menu.Item{Key: "vote", Title: "Votes", Description: "An estimate for a story, from a user", Icon: "vote-yea", Route: "/admin/db/estimate/story/vote"},
			}},
		}},
		&menu.Item{Key: "retro", Title: "Retros", Description: "Discover improvements and praise for your work", Icon: "retro", Route: "/admin/db/retro", Children: menu.Items{
			&menu.Item{Key: "feedback", Title: "Feedbacks", Description: "User feedback for a retrospective", Icon: "comment", Route: "/admin/db/retro/feedback"},
			&menu.Item{Key: "rhistory", Title: "Histories", Description: "Historical names and URLs for retrospectives", Icon: "history", Route: "/admin/db/retro/history"},
			&menu.Item{Key: "rmember", Title: "Members", Description: "Membership roster for retrospectives", Icon: "users", Route: "/admin/db/retro/member"},
			&menu.Item{Key: "rpermission", Title: "Permissions", Description: "Security controls for retrospectives", Icon: "permission", Route: "/admin/db/retro/permission"},
		}},
		&menu.Item{Key: "sprint", Title: "Sprints", Description: "Plan your time and direct your efforts", Icon: "sprint", Route: "/admin/db/sprint", Children: menu.Items{
			&menu.Item{Key: "shistory", Title: "Histories", Description: "Historical names and URLs for sprints", Icon: "history", Route: "/admin/db/sprint/history"},
			&menu.Item{Key: "smember", Title: "Members", Description: "Membership roster for sprints", Icon: "users", Route: "/admin/db/sprint/member"},
			&menu.Item{Key: "spermission", Title: "Permissions", Description: "Security controls for sprints", Icon: "permission", Route: "/admin/db/sprint/permission"},
		}},
		&menu.Item{Key: "standup", Title: "Standups", Description: "Share your progress with your team", Icon: "standup", Route: "/admin/db/standup", Children: menu.Items{
			&menu.Item{Key: "report", Title: "Reports", Description: "Daily status reports for a standup", Icon: "file-alt", Route: "/admin/db/standup/report"},
			&menu.Item{Key: "uhistory", Title: "Histories", Description: "Historical names and URLs for standups", Icon: "history", Route: "/admin/db/standup/history"},
			&menu.Item{Key: "umember", Title: "Members", Description: "Membership roster for standups", Icon: "users", Route: "/admin/db/standup/member"},
			&menu.Item{Key: "upermission", Title: "Permissions", Description: "Security controls for standups", Icon: "permission", Route: "/admin/db/standup/permission"},
		}},
		&menu.Item{Key: "team", Title: "Teams", Description: "Join your friends and work towards a common goal", Icon: "team", Route: "/admin/db/team", Children: menu.Items{
			&menu.Item{Key: "thistory", Title: "Histories", Description: "Historical names and URLs for teams", Icon: "history", Route: "/admin/db/team/history"},
			&menu.Item{Key: "tmember", Title: "Members", Description: "Membership roster for teams", Icon: "users", Route: "/admin/db/team/member"},
			&menu.Item{Key: "tpermission", Title: "Permissions", Description: "Security controls for teams", Icon: "permission", Route: "/admin/db/team/permission"},
		}},
		&menu.Item{Key: "user", Title: "Users", Description: "A user of the system", Icon: "profile", Route: "/admin/db/user"},
	}
}
