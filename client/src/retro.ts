namespace retro {
  interface Detail extends rituals.Session {
    options: object;
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
  }

  class Cache {
    detail?: Detail;
  }

  const cache = new Cache();

  export function onRetroMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.retro, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined
        rituals.onSessionJoin(sj);
        setRetroDetail(sj.session);
        break;
      case command.server.sessionUpdate:
        setRetroDetail(param as Detail);
        break;
      default:
        console.warn("unhandled command [" + cmd + "] for retro");
    }
  }

  function setRetroDetail(detail: Detail) {
    cache.detail = detail;
    rituals.setDetail(detail);
  }

  export function onSubmitRetroSession() {
    const title = util.req<HTMLInputElement>("#model-title-input").value;
    const msg = {
      svc: services.retro,
      cmd: command.client.updateSession,
      param: {
        title: title,
      },
    };
    socket.send(msg);
  }
}
