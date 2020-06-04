package admin

import (
	"time"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/model/sandbox"
	"github.com/kyleu/rituals.dev/app/socket"
	"github.com/kyleu/rituals.dev/app/util"
)

func SectionCounts(sections []string, routes util.RouteDescriptions, db *database.Service, socket *socket.Service) (map[string]int64, map[string]*time.Time, error) {
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

func sectionCount(routes util.RouteDescriptions, db *database.Service, sck *socket.Service, section string) (int64, *time.Time, error) {
	switch section {
	case util.KeyGraphQL, util.KeyTranscript:
		return -1, nil, nil
	case util.KeyConnection:
		return int64(sck.Count()), nil, nil
	case util.KeyModules:
		return int64(len(util.ExtractModules().Deps)), nil, nil
	case util.KeyRoutes:
		return int64(len(routes)), nil, nil
	case util.KeySandbox:
		return int64(len(sandbox.AllSandboxes)), nil, nil
	case util.KeyUser:
		return databaseWork(db, util.KeySystemUser)
	default:
		return databaseWork(db, section)
	}
}

type recentResult struct {
	M *time.Time `db:"m"`
}

func databaseWork(db *database.Service, section string) (int64, *time.Time, error) {
	count, err := db.Count("select count(*) as c from "+section, nil)
	if err != nil {
		return -1, nil, errors.Wrap(err, "cannot get count from "+section)
	}

	rr := &recentResult{}
	err = db.Get(rr, "select max("+util.KeyCreated+") as m from "+section, nil)
	if err != nil {
		return count, nil, errors.Wrap(err, "cannot get recent records from "+section)
	}

	return count, rr.M, nil
}
