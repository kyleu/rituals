package transcript

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type EstimateResponse struct {
	Svc         util.Service           `json:"-"`
	Session     *estimate.Session      `json:"session"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Stories     estimate.Stories       `json:"stories"`
	Votes       estimate.Votes         `json:"votes"`
	Permissions permission.Permissions `json:"permissions"`
}

var Estimate = Transcript{
	Key:         util.SvcEstimate.Key,
	Title:       util.SvcEstimate.Title,
	Description: util.SvcEstimate.Description,
	Resolve: func(ai npnweb.AppInfo, userID uuid.UUID, slug string) (interface{}, error) {
		svc := app.Svc(ai)
		if len(slug) == 0 {
			return svc.Estimate.List(nil), nil
		}
		sess := svc.Estimate.GetBySlug(slug)
		if sess == nil {
			return nil, errors.New("no session available matching [" + slug + "]")
		}
		dataSvc := svc.Estimate.Data
		return EstimateResponse{
			Svc:         util.SvcEstimate,
			Session:     sess,
			Team:        svc.Team.GetByIDPointer(sess.TeamID),
			Sprint:      svc.Sprint.GetByIDPointer(sess.SprintID),
			Comments:    dataSvc.GetComments(sess.ID, nil),
			Members:     dataSvc.Members.GetByModelID(sess.ID, nil),
			Stories:     svc.Estimate.GetStories(sess.ID, nil),
			Votes:       svc.Estimate.GetEstimateVotes(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sess.ID, nil),
		}, nil
	},
}
