package team

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (t *Team) ToMap() util.ValueMap {
	return util.ValueMap{"id": t.ID, "slug": t.Slug, "title": t.Title, "icon": t.Icon, "status": t.Status, "created": t.Created, "updated": t.Updated}
}

func TeamFromMap(m util.ValueMap, setPK bool) (*Team, util.ValueMap, error) {
	ret := &Team{}
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
func (t *Team) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: t.ID}, {K: "slug", V: t.Slug}, {K: "title", V: t.Title}, {K: "icon", V: t.Icon}, {K: "status", V: t.Status}, {K: "created", V: t.Created}, {K: "updated", V: t.Updated}}
	return util.NewOrderedMap(false, 4, pairs...)
}
