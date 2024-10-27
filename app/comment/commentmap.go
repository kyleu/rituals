package comment

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (c *Comment) ToMap() util.ValueMap {
	return util.ValueMap{"id": c.ID, "svc": c.Svc, "modelID": c.ModelID, "userID": c.UserID, "content": c.Content, "html": c.HTML, "created": c.Created}
}

//nolint:gocognit
func CommentFromMap(m util.ValueMap, setPK bool) (*Comment, util.ValueMap, error) {
	ret := &Comment{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				retID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retID != nil {
					ret.ID = *retID
				}
			}
		case "svc":
			retSvc, err := m.ParseString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
			ret.Svc = enum.AllModelServices.Get(retSvc, nil)
		case "modelID":
			retModelID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retModelID != nil {
				ret.ModelID = *retModelID
			}
		case "userID":
			retUserID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retUserID != nil {
				ret.UserID = *retUserID
			}
		case "content":
			ret.Content, err = m.ParseString(k, true, true)
		case "html":
			ret.HTML, err = m.ParseString(k, true, true)
		default:
			extra[k] = v
		}
		if err != nil {
			return nil, nil, err
		}
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, extra, nil
}
