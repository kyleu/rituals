namespace standup {
  interface Detail extends rituals.Session {
    options: object;
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    updates: update.Update[]
  }

  class Cache {
    detail?: Detail;
    updates: update.Update[] = [];

    activeUpdate?: string;
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
        update.setUpdates(sj.updates)
        break;
      case command.server.sessionUpdate:
        setStandupDetail(param as Detail);
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
}
