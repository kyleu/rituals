package retro

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

//nolint:lll
func (r *Retro) ToMap() util.ValueMap {
	return util.ValueMap{"id": r.ID, "slug": r.Slug, "title": r.Title, "icon": r.Icon, "status": r.Status, "teamID": r.TeamID, "sprintID": r.SprintID, "categories": r.Categories, "created": r.Created, "updated": r.Updated}
}

func RetroFromMap(m util.ValueMap, setPK bool) (*Retro, util.ValueMap, error) {
	ret := &Retro{}
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
		case "slug":
			ret.Slug, err = m.ParseString(k, true, true)
		case "title":
			ret.Title, err = m.ParseString(k, true, true)
		case "icon":
			ret.Icon, err = m.ParseString(k, true, true)
		case "status":
			retStatus, err := m.ParseString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
			ret.Status = enum.AllSessionStatuses.Get(retStatus, nil)
		case "teamID":
			ret.TeamID, err = m.ParseUUID(k, true, true)
		case "sprintID":
			ret.SprintID, err = m.ParseUUID(k, true, true)
		case "categories":
			ret.Categories, err = m.ParseArrayString(k, true, true)
			if err != nil {
				return nil, nil, err
			}
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

//nolint:lll
func (r *Retro) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: r.ID}, {K: "slug", V: r.Slug}, {K: "title", V: r.Title}, {K: "icon", V: r.Icon}, {K: "status", V: r.Status}, {K: "teamID", V: r.TeamID}, {K: "sprintID", V: r.SprintID}, {K: "categories", V: r.Categories}, {K: "created", V: r.Created}, {K: "updated", V: r.Updated}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
