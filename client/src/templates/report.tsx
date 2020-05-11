namespace report {
  function renderReport(model: report.Report): JSX.Element {
    const profile = system.cache.getProfile();
    const ret = <div id={"report-" + model.id} class="report-detail uk-border-rounded section" onclick={"events.openModal('report', '" + model.id + "');"}>
      <a class={profile.linkColor + "-fg section-link"}>{member.getMemberName(model.author)}</a>
      <div class="report-content">loading...</div>
    </div>;

    if(model.html.length > 0) {
      util.setHTML(util.req(".report-content", ret), model.html).style.display = "block";
    }

    return ret;
  }

  export function renderReports(reports: report.Report[]): JSX.Element {
    if (reports.length === 0) {
      return <div>
        <button class="uk-button uk-button-default" onclick="events.openModal('add-report');" type="button">Add Report</button>
      </div>;
    } else {
      const dates = getReportDates(reports);
      return <ul class="uk-list">
        {dates.map(day => <li id={"report-date-" + day.d}>
          <div>{day.d}</div>
          {day.reports.map(r => <li>{renderReport(r)}</li>)}
        </li>)}
      </ul>;
    }
  }

}
