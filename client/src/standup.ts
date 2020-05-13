namespace standup {
  interface Detail extends rituals.Session {

  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    reports: report.Report[]
  }

  class Cache {
    detail?: Detail;
    reports: report.Report[] = [];

    activeReport?: string;
  }

  export const cache = new Cache();

  export function onStandupMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.standup, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        setStandupDetail(sj.session);
        report.setReports(sj.reports)
        break;
      case command.server.sessionUpdate:
        setStandupDetail(param as Detail);
        break;
      case command.server.reportUpdate:
        onReportUpdate(param as report.Report);
        break;
      case command.server.reportRemove:
        onReportRemoved(param as string);
        break;
      default:
        console.warn("unhandled command [" + cmd + "] for standup");
    }
  }

  function setStandupDetail(detail: Detail) {
    cache.detail = detail;
    rituals.setDetail(detail);
  }

  export function onSubmitStandupSession() {
    const title = util.req<HTMLInputElement>("#model-title-input").value;
    const msg = {
      svc: services.standup,
      cmd: command.client.updateSession,
      param: {
        title: title,
      },
    };
    socket.send(msg);
  }

  function onReportUpdate(r: report.Report) {
    const x = preUpdate(r.id)
    x.push(r);
    postUpdate(x, r.id)
  }

  function onReportRemoved(id: string) {
    const x = preUpdate(id)
    postUpdate(x, id)
    UIkit.notification("report has been deleted", {status: "success", pos: "top-right"});
  }

  function preUpdate(id: string) {
    return cache.reports.filter((p) => p.id !== id);
  }

  function postUpdate(x: report.Report[], id: string) {
    report.setReports(x);

    if(id === standup.cache.activeReport) {
      UIkit.modal("#modal-report").hide();
    }
  }
}
