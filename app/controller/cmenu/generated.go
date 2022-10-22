// Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/rituals/app/lib/menu"

func generatedMenu() menu.Items {
	return menu.Items{
		&menu.Item{Key: "action", Title: "Actions", Description: "TODO", Icon: "star", Route: "/action"},
		&menu.Item{Key: "comment", Title: "Comments", Description: "TODO", Icon: "star", Route: "/comment"},
		&menu.Item{Key: "email", Title: "Emails", Description: "TODO", Icon: "star", Route: "/email"},
		&menu.Item{Key: "estimate", Title: "Estimates", Description: "TODO", Icon: "star", Route: "/estimate", Children: menu.Items{
			&menu.Item{Key: "ehistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/estimate/ehistory"},
			&menu.Item{Key: "emember", Title: "Members", Description: "TODO", Icon: "users", Route: "/estimate/emember"},
			&menu.Item{Key: "epermission", Title: "Permissions", Description: "TODO", Icon: "users", Route: "/estimate/epermission"},
			&menu.Item{Key: "story", Title: "Stories", Description: "TODO", Icon: "star", Route: "/estimate/story", Children: menu.Items{
				&menu.Item{Key: "vote", Title: "Votes", Description: "TODO", Icon: "star", Route: "/estimate/story/vote"},
			}},
		}},
		&menu.Item{Key: "retro", Title: "Retros", Description: "TODO", Icon: "star", Route: "/retro", Children: menu.Items{
			&menu.Item{Key: "feedback", Title: "Feedbacks", Description: "TODO", Icon: "star", Route: "/retro/feedback"},
			&menu.Item{Key: "rhistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/retro/rhistory"},
			&menu.Item{Key: "rmember", Title: "Members", Description: "TODO", Icon: "users", Route: "/retro/rmember"},
			&menu.Item{Key: "rpermission", Title: "Permissions", Description: "TODO", Icon: "users", Route: "/retro/rpermission"},
		}},
		&menu.Item{Key: "sprint", Title: "Sprints", Description: "TODO", Icon: "star", Route: "/sprint", Children: menu.Items{
			&menu.Item{Key: "shistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/sprint/shistory"},
			&menu.Item{Key: "smember", Title: "Members", Description: "TODO", Icon: "users", Route: "/sprint/smember"},
			&menu.Item{Key: "spermission", Title: "Permissions", Description: "TODO", Icon: "users", Route: "/sprint/spermission"},
		}},
		&menu.Item{Key: "standup", Title: "Standups", Description: "TODO", Icon: "star", Route: "/standup", Children: menu.Items{
			&menu.Item{Key: "report", Title: "Reports", Description: "TODO", Icon: "clock", Route: "/standup/report"},
			&menu.Item{Key: "uhistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/standup/uhistory"},
			&menu.Item{Key: "umember", Title: "Members", Description: "TODO", Icon: "users", Route: "/standup/umember"},
			&menu.Item{Key: "upermission", Title: "Permissions", Description: "TODO", Icon: "users", Route: "/standup/upermission"},
		}},
		&menu.Item{Key: "team", Title: "Teams", Description: "TODO", Icon: "users", Route: "/team", Children: menu.Items{
			&menu.Item{Key: "thistory", Title: "Histories", Description: "TODO", Icon: "clock", Route: "/team/thistory"},
			&menu.Item{Key: "tmember", Title: "Members", Description: "TODO", Icon: "users", Route: "/team/tmember"},
			&menu.Item{Key: "tpermission", Title: "Permissions", Description: "TODO", Icon: "users", Route: "/team/tpermission"},
		}},
		&menu.Item{Key: "user", Title: "Users", Description: "TODO", Icon: "profile", Route: "/user"},
	}
}
