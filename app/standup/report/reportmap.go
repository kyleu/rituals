package report

import "github.com/kyleu/rituals/app/util"

//nolint:lll
func (r *Report) ToMap() util.ValueMap {
	return util.ValueMap{"id": r.ID, "standupID": r.StandupID, "day": r.Day, "userID": r.UserID, "content": r.Content, "html": r.HTML, "created": r.Created, "updated": r.Updated}
}

//nolint:gocognit
func ReportFromMap(m util.ValueMap, setPK bool) (*Report, util.ValueMap, error) {
	ret := &Report{}
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
		case "standupID":
			retStandupID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retStandupID != nil {
				ret.StandupID = *retStandupID
			}
		case "day":
			retDay, e := m.ParseTime(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retDay != nil {
				ret.Day = *retDay
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

//nolint:lll
func (r *Report) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: r.ID}, {K: "standupID", V: r.StandupID}, {K: "day", V: r.Day}, {K: "userID", V: r.UserID}, {K: "content", V: r.Content}, {K: "html", V: r.HTML}, {K: "created", V: r.Created}, {K: "updated", V: r.Updated}}
	return util.NewOrderedMap(false, 4, pairs...)
}
