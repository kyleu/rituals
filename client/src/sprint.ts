namespace sprint {
  interface Detail extends rituals.Session {

  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
  }

  class Cache {
    detail?: Detail;
  }

  export const cache = new Cache();

  export function onSprintMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.sprint, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        setSprintDetail(sj.session);
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

  export function onSubmitSprintSession() {
    const title = util.req<HTMLInputElement>("#model-title-input").value;
    const msg = {
      svc: services.sprint,
      cmd: command.client.updateSession,
      param: {
        title: title,
      },
    };
    socket.send(msg);
  }
}
