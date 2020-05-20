namespace team {
  export interface Detail extends rituals.Session {
    startDate: string;
    endDate: string;
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    sprints: rituals.Session[];
    estimates: rituals.Session[];
    standups: rituals.Session[];
    retros: rituals.Session[];
  }

  class Cache {
    detail?: Detail;
  }

  export const cache = new Cache();

  export function onTeamMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.team.key, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        setTeamDetail(sj.session);
        setTeamHistory(sj);
        break;
      case command.server.sessionUpdate:
        setTeamDetail(param as Detail);
        break;
      case command.server.contentUpdate:
        socket.socketConnect(system.cache.currentService, system.cache.currentID);
        break;
      default:
        console.warn(`unhandled command [${cmd}] for team`);
    }
  }

  function setTeamDetail(detail: Detail) {
    cache.detail = detail;
    rituals.setDetail(detail);
  }


  function setTeamHistory(sj: SessionJoined) {
    dom.setContent("#team-sprint-list", contents.renderContents(services.sprint, sj.sprints));
    dom.setContent("#team-estimate-list", contents.renderContents(services.estimate, sj.estimates));
    dom.setContent("#team-standup-list", contents.renderContents(services.standup, sj.standups));
    dom.setContent("#team-retro-list", contents.renderContents(services.retro, sj.retros));
  }

  export function onSubmitTeamSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const msg = {svc: services.team.key, cmd: command.client.updateSession, param: {title: title}};
    socket.send(msg);
  }

  export function refreshTeams() {
    const teamSelect = dom.opt("#model-team-select");
    if (teamSelect) {
      socket.send({svc: services.system.key, cmd: command.client.getTeams, param: null});
    }
  }

  export function viewTeams(teams: team.Detail[]) {
    const c = dom.opt("#model-team-container");
    if(c) {
      c.style.display = teams.length > 0 ? "block" : "none";
      dom.setContent("#model-team-select", renderTeamSelect(teams, system.cache.session?.teamID));
    }
  }
}
