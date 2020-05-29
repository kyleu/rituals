package standup

import (
	"fmt"
	"time"

	"github.com/kyleu/rituals.dev/app/action"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/query"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/markdown"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) NewReport(standupID uuid.UUID, d time.Time, content string, authorID uuid.UUID) (*Report, error) {
	id := util.UUID()
	html := markdown.ToHTML(content)

	q := query.SQLInsert(util.KeyReport, []string{util.KeyID, util.WithDBID(util.SvcStandup.Key), "d", util.WithDBID(util.KeyAuthor), util.KeyContent, util.KeyHTML}, 1)
	err := s.db.Insert(q, nil, id, standupID, d, authorID, content, html)
	if err != nil {
		return nil, err
	}

	actionContent := map[string]interface{}{"reportID": id}
	s.actions.Post(util.SvcStandup, standupID, authorID, action.ActReportAdd, actionContent, "")

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
		ret = append(ret, dto.ToReport())
	}
	return ret
}

func (s *Service) GetReportByID(reportID uuid.UUID) (*Report, error) {
	dto := &reportDTO{}
	q := query.SQLSelectSimple("*", util.KeyReport, util.KeyID + " = $1")
	err := s.db.Get(dto, q, nil, reportID)

	if err != nil {
		return nil, err
	}

	return dto.ToReport(), nil
}

func (s *Service) GetReportStandupID(reportID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := query.SQLSelectSimple(util.WithDBID(util.SvcStandup.Key), util.KeyReport, util.KeyID + " = $1")
	err := s.db.Get(&ret, q, nil, reportID)

	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *Service) UpdateReport(reportID uuid.UUID, d time.Time, content string, authorID uuid.UUID) (*Report, error) {
	html := markdown.ToHTML(content)

	q := query.SQLUpdate(util.KeyReport, []string{"d", util.WithDBID(util.KeyAuthor), util.KeyContent, util.KeyHTML}, util.KeyID + " = $5")
	err := s.db.UpdateOne(q, nil, d, authorID, content, html, reportID)
	if err != nil {
		return nil, err
	}

	report, err := s.GetReportByID(reportID)
	if report == nil {
		return nil, errors.New("cannot load newly-updated report")
	}

	actionContent := map[string]interface{}{"reportID": reportID}
	s.actions.Post(util.SvcStandup, report.StandupID, authorID, action.ActReportUpdate, actionContent, "")

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

	err = s.db.DeleteOne(query.SQLDelete(util.KeyReport, util.KeyID + " = $1"), nil, reportID)

	actionContent := map[string]interface{}{"reportID": reportID}
	s.actions.Post(util.SvcStandup, report.StandupID, userID, action.ActReportRemove, actionContent, "")

	return err
}
