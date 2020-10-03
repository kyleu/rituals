namespace rituals {
  export function onError(svc: string, err: string) {
    console.error(`${svc}: ${err}`);
    const idx = err.lastIndexOf(":");
    if (idx > -1) {
      err = err.substr(idx + 1);
    }
    notify.notify(`${svc} error: ${err}`, false);
  }

  export function init(svc: string, id: string) {
    window.onbeforeunload = () => {
      socket.setAppUnloading();
    };

    log.init();

    socket.init(onSocketOpen, recv, onError);
    socket.socketConnect(svc, id);
  }

  function recv(msg: socket.Message) {
    if (debug) {
      console.debug("in", msg);
    }
    switch (msg.svc) {
      case services.system.key:
        system.onSystemMessage(msg.cmd, msg.param);
        break;
      case services.team.key:
        team.onTeamMessage(msg.cmd, msg.param);
        break;
      case services.sprint.key:
        sprint.onSprintMessage(msg.cmd, msg.param);
        break;
      case services.estimate.key:
        estimate.onEstimateMessage(msg.cmd, msg.param);
        break;
      case services.standup.key:
        standup.onStandupMessage(msg.cmd, msg.param);
        break;
      case services.retro.key:
        retro.onRetroMessage(msg.cmd, msg.param);
        break;
      default:
        console.warn(`unhandled message for service [${msg.svc}]`);
    }
  }

  function onSocketOpen(id: string) {
    system.cache.currentID = id;
    socket.send({ svc: socket.currentService!, cmd: command.client.connect, param: system.cache.currentID });
  }
}
