namespace standup {
  interface Detail extends rituals.Session {
    status: { key: string };
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    sprint?: sprint.Detail;
    reports: report.Report[]
  }

  class Cache {
    activeReport?: string;

    detail?: Detail;
    sprint?: sprint.Detail;

    reports: report.Report[] = [];
  }

  export const cache = new Cache();

  export function onStandupMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.standup.key, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        rituals.setSprint(sj.sprint)
        setStandupDetail(sj.session);
        report.setReports(sj.reports)
        break;
      case command.server.sessionUpdate:
        setStandupDetail(param as Detail);
        break;
      case command.server.sprintUpdate:
        const x = param as sprint.Detail | undefined
        if (standup.cache.detail) {
          standup.cache.detail.sprintID = x?.id;
        }
        rituals.setSprint(x)
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
    const sprintID = util.req<HTMLSelectElement>("#model-sprint-select select").value;
    const msg = {svc: services.standup.key, cmd: command.client.updateSession, param: {title: title, sprintID: sprintID}};
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
