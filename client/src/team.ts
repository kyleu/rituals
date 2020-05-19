namespace team {
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
      default:
        console.warn("unhandled command [" + cmd + "] for team");
    }
  }

  function setTeamDetail(detail: Detail) {
    cache.detail = detail;
    rituals.setDetail(detail);
  }

  function setTeamHistory(sj: SessionJoined) {
  }

  export function onSubmitTeamSession() {
    const title = util.req<HTMLInputElement>("#model-title-input").value;
    const msg = {svc: services.team.key, cmd: command.client.updateSession, param: {title: title}};
    socket.send(msg);
  }
}
