namespace sprint {
  export interface Detail extends rituals.Session {
    readonly startDate: string;
    readonly endDate: string;
  }

  interface SessionJoined extends rituals.SessionJoined {
    readonly session: Detail;
    readonly team?: team.Detail;
    readonly estimates: rituals.Session[];
    readonly standups: rituals.Session[];
    readonly retros: rituals.Session[];
  }

  class Cache {
    detail?: Detail;
  }

  export const cache = new Cache();

  export function onSprintMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.sprint.key, param as string);
        break;
      case command.server.sessionJoined:
        const sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        setSprintDetail(sj.session);
        rituals.setTeam(sj.team)
        setSprintContents(sj);
        rituals.showWelcomeMessage(sj.members.length);
        break;
      case command.server.teamUpdate:
        const tm = param as team.Detail | undefined
        if (sprint.cache.detail) {
          sprint.cache.detail.teamID = tm?.id;
        }
        rituals.setTeam(tm);
        break;
      case command.server.sessionUpdate:
        setSprintDetail(param as Detail);
        break;
      case command.server.permissionsUpdate:
        system.setPermissions(param as permission.Permission[]);
        break;
      case command.server.contentUpdate:
        socket.socketConnect(system.cache.currentService, system.cache.currentID);
        break;
      default:
        console.warn(`unhandled command [${cmd}] for sprint`);
    }
  }

  function setSprintDetail(detail: Detail) {
    cache.detail = detail;
    const s = detail.startDate?.length === 0 ? undefined : new Date(detail.startDate);
    const e = detail.endDate?.length === 0 ? undefined : new Date(detail.endDate);
    dom.setContent("#sprint-date-display", renderSprintDates(s, e));
    dom.setValue("#sprint-start-date-input", s ? date.dateToYMD(s) : "");
    dom.setValue("#sprint-end-date-input", e ? date.dateToYMD(e) : "");
    rituals.setDetail(detail);
  }

  function setSprintContents(sj: SessionJoined) {
    dom.setContent("#sprint-estimate-list", contents.renderContents(services.sprint, services.estimate, sj.estimates));
    dom.setContent("#sprint-standup-list", contents.renderContents(services.sprint, services.standup, sj.standups));
    dom.setContent("#sprint-retro-list", contents.renderContents(services.sprint, services.retro, sj.retros));
  }

  export function onSubmitSprintSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const teamID = dom.req<HTMLSelectElement>("#model-team-select select").value;
    const startDate = dom.opt<HTMLInputElement>("#model-start-date-input")?.value;
    const endDate = dom.opt<HTMLInputElement>("#model-end-date-input")?.value;
    const permissions = permission.readPermissions();

    const msg = {svc: services.sprint.key, cmd: command.client.updateSession, param: {title, startDate, endDate, teamID, permissions}};
    socket.send(msg);
  }

  export function refreshSprints() {
    const sprintSelect = dom.opt("#model-sprint-select");
    if (sprintSelect) {
      socket.send({svc: services.system.key, cmd: command.client.getSprints, param: null});
    }
  }

  export function viewSprints(sprints: sprint.Detail[]) {
    const c = dom.opt("#model-sprint-container");
    if(c) {
      dom.setDisplay(c, sprints.length > 0)
      dom.setContent("#model-sprint-select", renderSprintSelect(sprints, system.cache.session?.sprintID));
      permission.setModelPerms("sprint");
    }
  }
}
