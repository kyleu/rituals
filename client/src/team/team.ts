namespace team {
  export interface Detail extends session.Session {
    readonly startDate: string;
    readonly endDate: string;
  }

  interface SessionJoined extends session.SessionJoined {
    readonly session: Detail;
    readonly sprints: session.Session[];
    readonly estimates: session.Session[];
    readonly standups: session.Session[];
    readonly retros: session.Session[];
  }

  class Cache {
    detail?: Detail;
  }

  export const cache = new Cache();

  export function onTeamMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.team, param as string);
        break;
      case command.server.sessionJoined:
        const sj = param as SessionJoined;
        session.onSessionJoin(sj);
        setTeamDetail(sj.session);
        setTeamHistory(sj);
        session.showWelcomeMessage(sj.members.length);
        break;
      case command.server.sessionUpdate:
        setTeamDetail(param as Detail);
        break;
      case command.server.sessionRemove:
        system.onSessionRemove(services.team);
        break;
      case command.server.permissionsUpdate:
        system.setPermissions(param as permission.Permission[]);
        break;
      case command.server.contentUpdate:
        socket.socketConnect(system.cache.currentService!, system.cache.currentID);
        break;
      default:
        console.warn(`unhandled command [${cmd}] for team`);
    }
  }

  function setTeamDetail(detail: Detail) {
    cache.detail = detail;
    session.setDetail(detail);
  }

  function setTeamHistory(sj: SessionJoined) {
    dom.setContent("#team-sprint-list", contents.renderContents(services.team, services.sprint, sj.sprints));
    dom.setContent("#team-estimate-list", contents.renderContents(services.team, services.estimate, sj.estimates));
    dom.setContent("#team-standup-list", contents.renderContents(services.team, services.standup, sj.standups));
    dom.setContent("#team-retro-list", contents.renderContents(services.team, services.retro, sj.retros));
  }

  export function onSubmitTeamSession() {
    const title = dom.req<HTMLInputElement>("#model-title-input").value;
    const permissions = permission.readPermissions();

    const msg = { svc: services.team.key, cmd: command.client.updateSession, param: { title, permissions } };
    socket.send(msg);
  }

  export function refreshTeams() {
    const teamSelect = dom.opt("#model-team-select");
    if (teamSelect) {
      socket.send({ svc: services.system.key, cmd: command.client.getTeams, param: null });
    }
  }

  export function viewTeams(teams: ReadonlyArray<team.Detail>) {
    const c = dom.opt("#model-team-container");
    if (c) {
      // dom.setDisplay(c, teams.length > 0)
      dom.setContent("#model-team-select", renderTeamSelect(teams, system.cache.session?.teamID));
      permission.setModelPerms("team");
    }
  }
}
