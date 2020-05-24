namespace standup {
  interface Detail extends rituals.Session {
    readonly status: { readonly key: string };
  }

  interface SessionJoined extends rituals.SessionJoined {
    readonly session: Detail;
    readonly team?: team.Detail;
    readonly sprint?: sprint.Detail;
    readonly reports: report.Report[];
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
        const sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        rituals.setTeam(sj.team);
        rituals.setSprint(sj.sprint);
        setStandupDetail(sj.session);
        report.setReports(sj.reports);
        rituals.showWelcomeMessage(sj.members.length);
        break;
      case command.server.sessionUpdate:
        setStandupDetail(param as Detail);
        break;
      case command.server.permissionsUpdate:
        system.setPermissions(param as permission.Permission[]);
        break;
      case command.server.authUpdate:
        system.setAuth(param as permission.Auth[]);
        break;
      case command.server.teamUpdate:
        const tm = param as team.Detail | undefined;
        if (standup.cache.detail) {
          standup.cache.detail.teamID = tm?.id;
        }
        rituals.setTeam(tm);
        break;
      case command.server.sprintUpdate:
        const x = param as sprint.Detail | undefined;
        if (standup.cache.detail) {
          standup.cache.detail.sprintID = x?.id;
        }
        rituals.setSprint(x);
        break;
      case command.server.reportUpdate:
        onReportUpdate(param as report.Report);
        break;
      case command.server.reportRemove:
        onReportRemoved(param as string);
        break;
      default:
        console.warn(`unhandled command [${cmd}] for standup`);
    }
  }

  function setStandupDetail(detail: Detail) {
    cache.detail = detail;
    rituals.setDetail(detail);
  }

  export function onSubmitStandupSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const teamID = dom.req<HTMLSelectElement>("#model-team-select select").value;
    const sprintID = dom.req<HTMLSelectElement>("#model-sprint-select select").value;
    const permissions = permission.readPermissions();

    const msg = {svc: services.standup.key, cmd: command.client.updateSession, param: {title, teamID, sprintID, permissions}};
    socket.send(msg);
  }

  function onReportUpdate(r: report.Report) {
    const x = preUpdate(r.id);
    x.push(r);
    postUpdate(x, r.id);
  }

  function onReportRemoved(id: string) {
    const x = preUpdate(id);
    postUpdate(x, id);
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
