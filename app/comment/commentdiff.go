// Package comment - Content managed by Project Forge, see [projectforge.md] for details.
package comment

import "github.com/kyleu/rituals/app/util"

func (c *Comment) Diff(cx *Comment) util.Diffs {
	var diffs util.Diffs
	if c.ID != cx.ID {
		diffs = append(diffs, util.NewDiff("id", c.ID.String(), cx.ID.String()))
	}
	if c.Svc != cx.Svc {
		diffs = append(diffs, util.NewDiff("svc", c.Svc.Key, cx.Svc.Key))
	}
	if c.ModelID != cx.ModelID {
		diffs = append(diffs, util.NewDiff("modelID", c.ModelID.String(), cx.ModelID.String()))
	}
	if c.UserID != cx.UserID {
		diffs = append(diffs, util.NewDiff("userID", c.UserID.String(), cx.UserID.String()))
	}
	if c.Content != cx.Content {
		diffs = append(diffs, util.NewDiff("content", c.Content, cx.Content))
	}
	if c.HTML != cx.HTML {
		diffs = append(diffs, util.NewDiff("html", c.HTML, cx.HTML))
	}
	if c.Created != cx.Created {
		diffs = append(diffs, util.NewDiff("created", c.Created.String(), cx.Created.String()))
	}
	return diffs
}
