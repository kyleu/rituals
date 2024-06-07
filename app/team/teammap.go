// Package team - Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func FromMap(m util.ValueMap, setPK bool) (*Team, util.ValueMap, error) {
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
