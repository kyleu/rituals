package admin

import (
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/rituals.dev/app/transcript"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/sandbox"
	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/util"
)

func SectionCounts(sections []string, routes npnweb.RouteDescriptions, db *npndatabase.Service, socket *socket.Service) (map[string]int64, map[string]*time.Time, error) {
	countMap := make(map[string]int64)
	recentMap := make(map[string]*time.Time)
	for _, section := range sections {
		count, recent, err := sectionCount(routes, db, socket, section)
		if err != nil {
			return nil, nil, err
		}
		countMap[section] = count
		recentMap[section] = recent
	}
	return countMap, recentMap, nil
}

func sectionCount(routes npnweb.RouteDescriptions, db *npndatabase.Service, sck *socket.Service, section string) (int64, *time.Time, error) {
	switch section {
	case npncore.KeyGraphQL:
		return -1, nil, nil
	case npncore.KeyConnection:
		return int64(sck.Count()), nil, nil
	case npncore.KeyModules:
		return int64(len(npnweb.ExtractModules().Deps)), nil, nil
	case npncore.KeyRoutes:
		return int64(len(routes)), nil, nil
	case npncore.KeySandbox:
		return int64(len(sandbox.AllSandboxes)), nil, nil
	case util.KeyTranscript:
		return int64(len(transcript.AllTranscripts)), nil, nil
	case npncore.KeyUser:
		return databaseWork(db, "system_user")
	default:
		return databaseWork(db, section)
	}
}

type recentResult struct {
	M *time.Time `db:"m"`
}

func databaseWork(db *npndatabase.Service, section string) (int64, *time.Time, error) {
	count, err := db.SingleInt("select count(*) as x from "+section, nil)
	if err != nil {
		return -1, nil, errors.Wrap(err, "cannot get count from "+section)
	}

	rr := &recentResult{}
	err = db.Get(rr, "select max("+npncore.KeyCreated+") as m from "+section, nil)
	if err != nil {
		return count, nil, errors.Wrap(err, "cannot get recent records from "+section)
	}

	return count, rr.M, nil
}
