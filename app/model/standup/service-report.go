package standup

import (
	"fmt"
	"time"

	"github.com/kyleu/rituals.dev/app/model/action"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/database/query"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) NewReport(standupID uuid.UUID, d time.Time, content string, userID uuid.UUID) (*Report, error) {
	id := util.UUID()
	html := util.ToHTML(content)

	q := query.SQLInsert(util.KeyReport, []string{util.KeyID, util.WithDBID(s.svc.Key), "d", util.WithDBID(util.KeyUser), util.KeyContent, util.KeyHTML}, 1)
	err := s.db.Insert(q, nil, id, standupID, d, userID, content, html)
	if err != nil {
		return nil, err
	}

	actionContent := map[string]interface{}{"reportID": id}
	s.Data.Actions.Post(s.svc, standupID, userID, action.ActReportAdd, actionContent)

	return s.GetReportByID(id)
}

var defaultReportOrdering = query.Orderings{{Column: "d", Asc: false}, {Column: util.KeyCreated, Asc: false}}

func (s *Service) GetReports(standupID uuid.UUID, params *query.Params) Reports {
	params = query.ParamsWithDefaultOrdering(util.KeyReport, params, defaultReportOrdering...)
	var dtos []reportDTO
	q := query.SQLSelect("*", util.KeyReport, "standup_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, standupID)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving reports for standup [%v]: %+v", standupID, err))
		return nil
	}

	ret := make(Reports, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toReport())
	}
	return ret
}

func (s *Service) GetReportByID(reportID uuid.UUID) (*Report, error) {
	dto := &reportDTO{}
	q := query.SQLSelectSimple("*", util.KeyReport, util.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, reportID)

	if err != nil {
		return nil, err
	}

	return dto.toReport(), nil
}

func (s *Service) GetReportStandupID(reportID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := query.SQLSelectSimple(util.WithDBID(s.svc.Key), util.KeyReport, util.KeyID+" = $1")
	err := s.db.Get(&ret, q, nil, reportID)

	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *Service) UpdateReport(reportID uuid.UUID, d time.Time, content string, userID uuid.UUID) (*Report, error) {
	html := util.ToHTML(content)

	q := query.SQLUpdate(util.KeyReport, []string{"d", util.WithDBID(util.KeyUser), util.KeyContent, util.KeyHTML}, util.KeyID+" = $5")
	err := s.db.UpdateOne(q, nil, d, userID, content, html, reportID)
	if err != nil {
		return nil, err
	}

	report, err := s.GetReportByID(reportID)
	if report == nil {
		return nil, errors.New("cannot load newly-updated report")
	}

	actionContent := map[string]interface{}{"reportID": reportID}
	s.Data.Actions.Post(s.svc, report.StandupID, userID, action.ActReportUpdate, actionContent)

	return report, err
}

func (s *Service) RemoveReport(reportID uuid.UUID, userID uuid.UUID) error {
	report, err := s.GetReportByID(reportID)
	if err != nil {
		return errors.Wrap(err, "cannot load report ["+reportID.String()+"] for removal")
	}
	if report == nil {
		return errors.New("cannot load report [" + reportID.String() + "] for removal")
	}

	err = s.db.DeleteOne(query.SQLDelete(util.KeyReport, util.KeyID+" = $1"), nil, reportID)

	actionContent := map[string]interface{}{"reportID": reportID}
	s.Data.Actions.Post(s.svc, report.StandupID, userID, action.ActReportRemove, actionContent)

	return err
}
