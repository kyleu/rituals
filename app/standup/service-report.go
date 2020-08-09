package standup

import (
	"fmt"
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"github.com/kyleu/rituals.dev/app/action"

	"emperror.dev/errors"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) NewReport(standupID uuid.UUID, d time.Time, content string, userID uuid.UUID) (*Report, error) {
	id := npncore.UUID()
	html := util.ToHTML(content)

	q := npndatabase.SQLInsert(npncore.KeyReport, []string{npncore.KeyID, npncore.WithDBID(s.svc.Key), "d", npncore.WithDBID(npncore.KeyUser), npncore.KeyContent, npncore.KeyHTML}, 1)
	err := s.db.Insert(q, nil, id, standupID, d, userID, content, html)
	if err != nil {
		return nil, err
	}

	actionContent := map[string]interface{}{"reportID": id}
	s.Data.Actions.Post(s.svc, standupID, userID, action.ActReportAdd, actionContent)

	return s.GetReportByID(id)
}

var defaultReportOrdering = npncore.Orderings{{Column: "d", Asc: false}, {Column: npncore.KeyCreated, Asc: false}}

func (s *Service) GetReports(standupID uuid.UUID, params *npncore.Params) Reports {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyReport, params, defaultReportOrdering...)
	var dtos []reportDTO
	q := npndatabase.SQLSelect("*", npncore.KeyReport, "standup_id = $1", params.OrderByString(), params.Limit, params.Offset)
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
	q := npndatabase.SQLSelectSimple("*", npncore.KeyReport, npncore.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, reportID)

	if err != nil {
		return nil, err
	}

	return dto.toReport(), nil
}

func (s *Service) GetReportStandupID(reportID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := npndatabase.SQLSelectSimple(npncore.WithDBID(s.svc.Key), npncore.KeyReport, npncore.KeyID+" = $1")
	err := s.db.Get(&ret, q, nil, reportID)

	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *Service) UpdateReport(reportID uuid.UUID, d time.Time, content string, userID uuid.UUID) (*Report, error) {
	html := util.ToHTML(content)

	q := npndatabase.SQLUpdate(npncore.KeyReport, []string{"d", npncore.WithDBID(npncore.KeyUser), npncore.KeyContent, npncore.KeyHTML}, npncore.KeyID+" = $5")
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

	err = s.db.DeleteOne(npndatabase.SQLDelete(npncore.KeyReport, npncore.KeyID+" = $1"), nil, reportID)

	actionContent := map[string]interface{}{"reportID": reportID}
	s.Data.Actions.Post(s.svc, report.StandupID, userID, action.ActReportRemove, actionContent)

	return err
}
