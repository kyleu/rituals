package standup

import (
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

	q := `insert into report (id, standup_id, d, author_id, content, html) values (
    $1, $2, $3, $4, $5, $6
	)`
	_, err := s.db.Exec(q, id, standupID, d, authorID, content, html)

	if err != nil {
		return nil, err
	}

	actionContent := map[string]interface{}{"reportID": id}
	s.actions.Post(util.SvcStandup.Key, standupID, authorID, action.ActReportAdd, actionContent, "")

	return s.GetReportByID(id)
}

var defaultReportOrdering = query.Orderings{{Column: "d", Asc: false}, {Column: util.KeyCreated, Asc: false}}

func (s *Service) GetReports(standupID uuid.UUID, params *query.Params) ([]*Report, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyReport, params, defaultReportOrdering...)
	var dtos []reportDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.KeyReport, "standup_id = $1", params.OrderByString(), params.Limit, params.Offset), standupID)

	if err != nil {
		return nil, err
	}

	ret := make([]*Report, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToReport())
	}
	return ret, nil
}

func (s *Service) GetReportByID(reportID uuid.UUID) (*Report, error) {
	dto := &reportDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.KeyReport, "id = $1", "", 0, 0), reportID)

	if err != nil {
		return nil, err
	}

	return dto.ToReport(), nil
}

func (s *Service) GetReportStandupID(reportID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := query.SQLSelect("standup_id", util.KeyReport, "id = $1", "", 0, 0)
	err := s.db.Get(&ret, q, reportID)

	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *Service) UpdateReport(reportID uuid.UUID, d time.Time, content string, authorID uuid.UUID) (*Report, error) {
	html := markdown.ToHTML(content)

	q := `update report set d = $1, author_id = $2, content = $3, html = $4 where id = $5`
	_, err := s.db.Exec(q, d, authorID, content, html, reportID)

	if err != nil {
		return nil, err
	}

	report, err := s.GetReportByID(reportID)
	if report == nil {
		return nil, errors.New("cannot load newly-updated report")
	}

	actionContent := map[string]interface{}{"reportID": reportID}
	s.actions.Post(util.SvcStandup.Key, report.StandupID, authorID, action.ActReportUpdate, actionContent, "")

	return report, err
}

func (s *Service) RemoveReport(reportID uuid.UUID, userID uuid.UUID) error {
	report, err := s.GetReportByID(reportID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot load report ["+reportID.String()+"] for removal"))
	}
	if report == nil {
		return errors.New("cannot load report [" + reportID.String() + "] for removal")
	}

	q := "delete from report where id = $1"
	_, err = s.db.Exec(q, reportID)

	actionContent := map[string]interface{}{"reportID": reportID}
	s.actions.Post(util.SvcStandup.Key, report.StandupID, userID, action.ActReportRemove, actionContent, "")

	return err
}
