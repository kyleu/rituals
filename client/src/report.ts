namespace report {
  export interface Report {
    readonly id: string;
    readonly d: string;
    readonly authorID: string;
    readonly content: string;
    readonly html: string;
    readonly created: string;
  }

  export interface DayReports {
    readonly d: string;
    readonly reports: Report[];
  }

  export function onSubmitReport() {
    const d = dom.req<HTMLInputElement>("#report-date").value;
    const content = dom.req<HTMLInputElement>("#report-content").value;
    const msg = {svc: services.standup.key, cmd: command.client.addReport, param: {d, content}};
    socket.send(msg);
    return false;
  }

  export function onEditReport() {
    const d = dom.req<HTMLInputElement>("#report-edit-date").value;
    const content = dom.req<HTMLInputElement>("#report-edit-content").value;
    const msg = {svc: services.standup.key, cmd: command.client.updateReport, param: {id: standup.cache.activeReport, d, content}};
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
    if (!standup.cache.activeReport) {
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
    if (!report) {
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
    contents.onContentDisplay("report", same, report.content, report.html)
    dom.setValue(dom.req<HTMLInputElement>("#report-edit-date"), same ? report.d : "");
  }
}
