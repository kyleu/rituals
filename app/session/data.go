package session

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/history"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Data struct {
	Members  member.Entries
	Comments comment.Comments
	Perms    permission.Permissions
	History  history.Entries
	Actions  action.Actions
}

type DataServices struct {
	Svc         util.Service
	Members     *member.Service
	Comments    *comment.Service
	Permissions *permission.Service
	History     *history.Service
	Actions     *action.Service
}

func NewDataServices(svc util.Service, actions *action.Service, users *user.Service, comments *comment.Service, db *npndatabase.Service, logger logur.Logger) *DataServices {
	return &DataServices{
		Svc:         svc,
		Members:     member.NewService(actions, users, db, logger, svc),
		Comments:    comments,
		Permissions: permission.NewService(actions, db, logger, svc),
		History:     history.NewService(actions, db, logger, svc),
		Actions:     actions,
	}
}

func (svcs *DataServices) GetData(id uuid.UUID, params npncore.ParamSet, logger logur.Logger) *Data {
	return &Data{
		Members:  svcs.Members.GetByModelID(id, params.Get(npncore.KeyMember, logger)),
		Comments: svcs.GetComments(id, params.Get(npncore.KeyComment, logger)),
		Perms:    svcs.Permissions.GetByModelID(id, params.Get(npncore.KeyPermission, logger)),
		History:  svcs.History.GetByModelID(id, params.Get(npncore.KeyHistory, logger)),
		Actions:  svcs.Actions.GetBySvcModel(util.SvcTeam.Key, id, params.Get(npncore.KeyAction, logger)),
	}
}

func (svcs *DataServices) GetComments(id uuid.UUID, params *npncore.Params) comment.Comments {
	return svcs.Comments.GetByModelID(svcs.Svc, id, params)
}
