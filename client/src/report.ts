namespace report {
  export interface Report {
    id: string;
    d: string;
    author: string;
    content: string;
    html: string;
    created: string;
  }

  export interface DayReports {
    d: string;
    reports: Report[];
  }

  export function onSubmitReport() {
    const d = util.req<HTMLInputElement>("#standup-report-date").value;
    const content = util.req<HTMLInputElement>("#standup-report-content").value;
    const msg = {
      svc: services.standup,
      cmd: command.client.addReport,
      param: {d: d, content: content},
    };
    socket.send(msg);
    return false;
  }

  function getActiveReport() {
    if (standup.cache.activeReport === undefined) {
      console.warn("no active report");
      return undefined;
    }
    const curr = standup.cache.reports.filter(x => x.id === standup.cache.activeReport);
    if (curr.length !== 1) {
      console.log("cannot load active report [" + standup.cache.activeReport + "]");
      return undefined;
    }
    return curr[0];
  }

  export function viewActiveReport() {
    const profile = system.cache.getProfile();
    const report = getActiveReport();
    if (report === undefined) {
      return;
    }

    util.setText("#report-title", report.d + " / " + member.getMemberName(report.author));
    const contentEdit = util.req("#modal-report .content-edit");
    const contentEditTextarea = util.req<HTMLTextAreaElement>(".uk-textarea", contentEdit);
    const contentView = util.req("#modal-report .content-view");
    const buttonsEdit = util.req("#modal-report .buttons-edit");
    const buttonsView = util.req("#modal-report .buttons-view");

    if(report.author === profile.userID) {
      contentEdit.style.display = "block";
      util.setValue(contentEditTextarea, report.content);
      util.wireTextarea(contentEditTextarea);
      contentView.style.display = "none";
      util.setHTML(contentView, "");
      buttonsEdit.style.display = "block";
      buttonsView.style.display = "none";
    } else {
      contentEdit.style.display = "none";
      util.setValue(contentEditTextarea, "");
      contentView.style.display = "block";
      util.setHTML(contentView, report.html);
      buttonsEdit.style.display = "none";
      buttonsView.style.display = "block";
    }

    console.log("viewActiveReport: " + report?.id);
  }

  export function setReports(reports: Report[]) {
    standup.cache.reports = reports;
    util.setContent("#report-detail", renderReports(reports));
    UIkit.modal("#modal-add-report").hide();
  }

  export function getReportDates(reports: Report[]): DayReports[] {
    function distinct(v: string, i: number, s: any[]) {
      return s.indexOf(v) === i;
    }

    function toCollection(d: string): DayReports {
      return {
        "d": d,
        "reports": reports.filter(r => r.d === d)
      }
    }

    return reports.map(r => r.d).filter(distinct).sort().reverse().map(toCollection);
  }
}
