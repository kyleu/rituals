package session

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/history"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
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
	Members     *member.Service
	Comments    *comment.Service
	Permissions *permission.Service
	History     *history.Service
	Actions     *action.Service
}

func (svcs *DataServices) GetData(id uuid.UUID, params query.ParamSet, logger logur.Logger) *Data {
	return &Data{
		Members:  svcs.Members.GetByModelID(id, params.Get(util.KeyMember, logger)),
		Comments: svcs.Comments.GetByModelID(id, params.Get(util.KeyComment, logger)),
		Perms:    svcs.Permissions.GetByModelID(id, params.Get(util.KeyPermission, logger)),
		History:  svcs.History.GetByModelID(id, params.Get(util.KeyHistory, logger)),
		Actions:  svcs.Actions.GetBySvcModel(util.SvcTeam, id, params.Get(util.KeyAction, logger)),
	}
}
