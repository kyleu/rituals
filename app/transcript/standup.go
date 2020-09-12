package transcript

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type StandupResponse struct {
	Svc         util.Service           `json:"-"`
	Session     *standup.Session       `json:"session"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Reports     standup.Reports        `json:"reports"`
	Permissions permission.Permissions `json:"permissions"`
}

var Standup = Transcript{
	Key:         util.SvcStandup.Key,
	Title:       util.SvcStandup.Title,
	Description: util.SvcStandup.Description,
	Resolve: func(ai npnweb.AppInfo, userID uuid.UUID, slug string) (interface{}, error) {
		svc := app.Svc(ai)
		if len(slug) == 0 {
			return svc.Standup.List(nil), nil
		}
		sess := svc.Standup.GetBySlug(slug)
		dataSvc := svc.Standup.Data
		if sess == nil {
			return nil, errors.New("no session available matching [" + slug + "]")
		}
		return StandupResponse{
			Svc:         util.SvcStandup,
			Session:     sess,
			Team:        svc.Team.GetByIDPointer(sess.TeamID),
			Sprint:      svc.Sprint.GetByIDPointer(sess.SprintID),
			Comments:    dataSvc.GetComments(sess.ID, nil),
			Members:     dataSvc.Members.GetByModelID(sess.ID, nil),
			Reports:     svc.Standup.GetReports(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sess.ID, nil),
		}, nil
	},
}
