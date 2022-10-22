// Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/rituals/app/lib/menu"

func generatedMenu() menu.Items {
	return menu.Items{
		&menu.Item{Key: "action", Title: "Actions", Description: "TODO", Icon: "gamepad", Route: "/admin/db/action"},
		&menu.Item{Key: "comment", Title: "Comments", Description: "TODO", Icon: "comments", Route: "/admin/db/comment"},
		&menu.Item{Key: "email", Title: "Emails", Description: "TODO", Icon: "envelope", Route: "/admin/db/email"},
		&menu.Item{Key: "estimate", Title: "Estimates", Description: "TODO", Icon: "ruler-horizontal", Route: "/admin/db/estimate", Children: menu.Items{
			&menu.Item{Key: "ehistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/admin/db/estimate/history"},
			&menu.Item{Key: "emember", Title: "Members", Description: "TODO", Icon: "users", Route: "/admin/db/estimate/member"},
			&menu.Item{Key: "epermission", Title: "Permissions", Description: "TODO", Icon: "lock", Route: "/admin/db/estimate/permission"},
			&menu.Item{Key: "story", Title: "Stories", Description: "TODO", Icon: "book-reader", Route: "/admin/db/estimate/story", Children: menu.Items{
				&menu.Item{Key: "vote", Title: "Votes", Description: "TODO", Icon: "vote-yea", Route: "/admin/db/estimate/story/vote"},
			}},
		}},
		&menu.Item{Key: "retro", Title: "Retros", Description: "TODO", Icon: "glasses", Route: "/admin/db/retro", Children: menu.Items{
			&menu.Item{Key: "feedback", Title: "Feedbacks", Description: "TODO", Icon: "comment", Route: "/admin/db/retro/feedback"},
			&menu.Item{Key: "rhistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/admin/db/retro/history"},
			&menu.Item{Key: "rmember", Title: "Members", Description: "TODO", Icon: "users", Route: "/admin/db/retro/member"},
			&menu.Item{Key: "rpermission", Title: "Permissions", Description: "TODO", Icon: "lock", Route: "/admin/db/retro/permission"},
		}},
		&menu.Item{Key: "sprint", Title: "Sprints", Description: "TODO", Icon: "running", Route: "/admin/db/sprint", Children: menu.Items{
			&menu.Item{Key: "shistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/admin/db/sprint/history"},
			&menu.Item{Key: "smember", Title: "Members", Description: "TODO", Icon: "users", Route: "/admin/db/sprint/member"},
			&menu.Item{Key: "spermission", Title: "Permissions", Description: "TODO", Icon: "lock", Route: "/admin/db/sprint/permission"},
		}},
		&menu.Item{Key: "standup", Title: "Standups", Description: "TODO", Icon: "shoe-prints", Route: "/admin/db/standup", Children: menu.Items{
			&menu.Item{Key: "report", Title: "Reports", Description: "TODO", Icon: "file-alt", Route: "/admin/db/standup/report"},
			&menu.Item{Key: "uhistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/admin/db/standup/history"},
			&menu.Item{Key: "umember", Title: "Members", Description: "TODO", Icon: "users", Route: "/admin/db/standup/member"},
			&menu.Item{Key: "upermission", Title: "Permissions", Description: "TODO", Icon: "lock", Route: "/admin/db/standup/permission"},
		}},
		&menu.Item{Key: "team", Title: "Teams", Description: "TODO", Icon: "users", Route: "/admin/db/team", Children: menu.Items{
			&menu.Item{Key: "thistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/admin/db/team/history"},
			&menu.Item{Key: "tmember", Title: "Members", Description: "TODO", Icon: "users", Route: "/admin/db/team/member"},
			&menu.Item{Key: "tpermission", Title: "Permissions", Description: "TODO", Icon: "lock", Route: "/admin/db/team/permission"},
		}},
		&menu.Item{Key: "user", Title: "Users", Description: "TODO", Icon: "profile", Route: "/admin/db/user"},
	}
}
