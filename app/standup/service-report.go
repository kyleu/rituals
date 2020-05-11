package standup

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/markdown"
	"github.com/kyleu/rituals.dev/app/util"
	"time"
)

func (s *Service) NewReport(standupID uuid.UUID, d time.Time, content string, authorID uuid.UUID) (*Report, error) {
	id := util.UUID()
	html := markdown.ToHTML(content)

	sql := `insert into report (id, standup_id, d, author_id, content, html) values (
    $1, $2, $3, $4, $5, $6
	)`
	_, err := s.db.Exec(sql, id, standupID, d, authorID, content, html)
	if err != nil {
		return nil, err
	}

	return s.GetReportByID(id)
}

func (s *Service) UpdateReport(reportID uuid.UUID, d time.Time, content string, authorID uuid.UUID) (*Report, error) {
	html := markdown.ToHTML(content)

	sql := `update report set d = $1, author_id = $2, content = $3, html = $4 where id = $5`
	_, err := s.db.Exec(sql, d, authorID, content, html, reportID)
	if err != nil {
		return nil, err
	}

	return s.GetReportByID(reportID)
}

func (s *Service) GetReports(standupID uuid.UUID) ([]Report, error) {
	var dtos []reportDTO
	err := s.db.Select(&dtos, "select * from report where standup_id = $1 order by d desc, created", standupID)
	if err != nil {
		return nil, err
	}
	ret := make([]Report, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToReport())
	}
	return ret, nil
}

func (s *Service) GetReportByID(reportID uuid.UUID) (*Report, error) {
	dto := &reportDTO{}
	err := s.db.Get(dto, "select * from report where id = $1", reportID)
	if err != nil {
		return nil, err
	}
	ret := dto.ToReport()
	return &ret, nil
}

func (s *Service) GetReportStandupID(reportID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	err := s.db.Get(&ret, "select standup_id from report where id = $1", reportID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
