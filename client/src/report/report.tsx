namespace report {
  function renderReport(model: Report) {
    const profile = system.cache.getProfile();
    const ret = (
      <div id={`report-${model.id}`} class="report-detail section" onclick={`modal.open('report', '${model.id}');`}>
        <div class="report-comments right">{comment.renderCount("report", model.id)}</div>
        <div class="left">
          <a class={`${profile.linkColor}-fg section-link`}>{member.renderTitle(member.getMember(model.userID))}</a>
        </div>
        <div class="clear" />
        <div class="report-content">loading...</div>
      </div>
    );

    if (model.html.length > 0) {
      dom.setHTML(dom.req(".report-content", ret), model.html).style.display = "block";
    }

    return ret;
  }

  export function renderReports(reports: readonly Report[]) {
    if (reports.length === 0) {
      return (
        <div>
          <button class="uk-button uk-button-default" onclick="modal.open('add-report');" type="button">
            Add Report
          </button>
        </div>
      );
    } else {
      const dates = getReportDates(reports);
      return (
        <ul class="uk-list">
          {dates.map(day => (
            <li id={`report-date-${day.d}`}>
              <h5>
                <div class="right uk-article-meta">{date.dow(date.dateFromYMD(day.d).getDay())}</div>
                {date.toDateString(date.dateFromYMD(day.d))}
              </h5>
              <ul>
                {day.reports.map(r => (
                  <li>{renderReport(r)}</li>
                ))}
              </ul>
            </li>
          ))}
        </ul>
      );
    }
  }
}
