namespace sprint {
  export interface Detail extends rituals.Session {
    startDate: string;
    endDate: string;
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    estimates: rituals.Session[];
    standups: rituals.Session[];
    retros: rituals.Session[];
  }

  class Cache {
    detail?: Detail;
    estimates: rituals.Session[] = [];
    standups: rituals.Session[] = [];
    retros: rituals.Session[] = [];
  }

  export const cache = new Cache();

  export function onSprintMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.sprint.key, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        setSprintDetail(sj.session);
        setSprintContents(sj);
        break;
      case command.server.sessionUpdate:
        setSprintDetail(param as Detail);
        break;
      default:
        console.warn("unhandled command [" + cmd + "] for sprint");
    }
  }

  function setSprintDetail(detail: Detail) {
    cache.detail = detail;
    rituals.setDetail(detail);
  }

  function setSprintContents(sj: SessionJoined) {
    viewEstimates(sj.estimates);
    viewStandups(sj.standups);
    viewRetros(sj.retros);
  }

  function viewEstimates(estimates: rituals.Session[]) {
    cache.estimates = estimates;
    util.setContent("#sprint-estimate-list", renderContents(services.estimate, cache.estimates));
  }

  function viewStandups(standups: rituals.Session[]) {
    cache.standups = standups;
    util.setContent("#sprint-standup-list", renderContents(services.standup, cache.standups));
  }

  function viewRetros(retros: rituals.Session[]) {
    cache.retros = retros;
    util.setContent("#sprint-retro-list", renderContents(services.retro, cache.retros));
  }

  export function onSubmitSprintSession() {
    const title = util.req<HTMLInputElement>("#model-title-input").value;
    const msg = {svc: services.sprint.key, cmd: command.client.updateSession, param: {title: title}};
    socket.send(msg);
  }

  export function refreshSprints() {
    const sprintSelect = util.opt("#model-sprint-select");
    if (sprintSelect) {
      const msg = {svc: services.system.key, cmd: command.client.getSprints, param: null};
      socket.send(msg);
    }
  }

  export function viewSprints(sprints: sprint.Detail[]) {
    const c = util.opt("#model-sprint-container");
    if(c) {
      c.style.display = sprints.length > 0 ? "inline" : "none";
    }
    util.setContent("#model-sprint-select", renderSprintSelect(sprints, system.cache.session?.sprintID));
  }
}
