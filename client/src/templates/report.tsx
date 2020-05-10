namespace report {
  function renderReport(model: report.Report): JSX.Element {
    const profile = system.cache.getProfile();
    return <div>
      <a class={profile.linkColor + "-fg"} href="" onclick={"return events.openModal('report', '" + model.id + "');"}>{member.getMemberName(model.author)}</a>
    </div>;
  }

  export function renderReports(reports: report.Report[]): JSX.Element {
    if (reports.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="events.openModal('add-report');" type="button">Add Report</button>
      </div>;
    } else {
      const dates = getReportDates(reports);
      return <ul class="uk-list uk-list-divider">
        {dates.map(day => <li id={"report-date-" + day.d}>
          <div>{day.d}</div>
          <ul class="uk-list">
            {day.reports.map(r => <li>{renderReport(r)}</li>)}
          </ul>
        </li>)}
      </ul>;
    }
  }

}
