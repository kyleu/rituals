namespace report {
  export interface Report {
    id: string;
    d: string;
    authorID: string;
    content: string;
    html: string;
    created: string;
  }

  export interface DayReports {
    d: string;
    reports: Report[];
  }

  export function onSubmitReport() {
    const d = dom.req<HTMLInputElement>("#standup-report-date").value;
    const content = dom.req<HTMLInputElement>("#standup-report-content").value;
    const msg = {svc: services.standup.key, cmd: command.client.addReport, param: {d: d, content: content}};
    socket.send(msg);
    return false;
  }

  export function onEditReport() {
    const d = dom.req<HTMLInputElement>("#standup-report-edit-date").value;
    const content = dom.req<HTMLInputElement>("#standup-report-edit-content").value;
    const msg = {svc: services.standup.key, cmd: command.client.updateReport, param: {id: standup.cache.activeReport, d: d, content: content}};
    socket.send(msg);
    return false;
  }

  export function onRemoveReport() {
    const id = standup.cache.activeReport;
    if (id) {
      UIkit.modal.confirm("Delete this report?").then(function () {
        const msg = {svc: services.standup.key, cmd: command.client.removeReport, param: id};
        socket.send(msg);
        UIkit.modal("#modal-report").hide();
      });
    }
    return false;
  }

  function getActiveReport() {
    if (standup.cache.activeReport === undefined) {
      console.warn("no active report");
      return undefined;
    }
    const curr = standup.cache.reports.filter(x => x.id === standup.cache.activeReport).shift();
    if (!curr) {
      console.warn(`cannot load active report [${standup.cache.activeReport}]`);
    }
    return curr;
  }

  export function viewActiveReport() {
    const profile = system.cache.getProfile();
    const report = getActiveReport();
    if (report === undefined) {
      console.warn("no active report");
      return;
    }

    dom.setText("#report-title", `${report.d} / ${system.getMemberName(report.authorID)}`);

    setFor(report, profile.userID);
  }

  export function setReports(reports: Report[]) {
    standup.cache.reports = reports;
    dom.setContent("#report-detail", renderReports(reports));
    UIkit.modal("#modal-add-report").hide();
  }

  export function getReportDates(reports: Report[]): DayReports[] {
    function distinct(v: string, i: number, s: any[]) {
      return s.indexOf(v) === i;
    }

    function toCollection(d: string): DayReports {
      const sorted = reports.filter(r => r.d === d).sort((l, r) => (l.created > r.created ? -1 : 1));
      return {"d": d, "reports": sorted};
    }

    return reports.map(r => r.d).filter(distinct).sort().reverse().map(toCollection);
  }

  function setFor(report: report.Report, userID: string) {
    const same = report.authorID === userID;

    dom.req("#modal-report .content-edit").style.display = same ? "block" : "none";
    dom.setValue(dom.req<HTMLInputElement>("#standup-report-edit-date"), same ? report.d : "");

    const contentEditTextarea = dom.req<HTMLTextAreaElement>("#standup-report-edit-content");
    dom.setValue(contentEditTextarea, same ? report.content : "");
    if (same) {
      dom.wireTextarea(contentEditTextarea);
    }

    const contentView = dom.req("#modal-report .content-view");
    contentView.style.display = same ? "none" : "block";
    dom.setHTML(contentView, same ? "" : report.html);

    dom.req("#modal-report .buttons-edit").style.display = same ? "block" : "none";
    dom.req("#modal-report .buttons-view").style.display = same ? "none" : "block";
  }
}
