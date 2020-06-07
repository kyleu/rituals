namespace socket {
  export interface Message {
    readonly svc: string;
    readonly cmd: string;
    readonly param: any;
  }

  let socket: WebSocket;
  let appUnloading = false;

  function socketUrl() {
    const l = document.location;
    let protocol = "ws";
    if (l.protocol === "https:") {
      protocol = "wss";
    }
    return protocol + `://${l.host}/s`;
  }

  export function setAppUnloading() {
    appUnloading = true;
  }

  export function socketConnect(svc: services.Service, id: string) {
    system.cache.currentService = svc;
    system.cache.currentID = id;
    system.cache.connectTime = Date.now();

    socket = new WebSocket(socketUrl());
    socket.onopen = () => {
      send({ svc: svc.key, cmd: command.client.connect, param: id });
    };
    socket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      onSocketMessage(msg);
    };
    socket.onerror = (event) => {
      rituals.onError(services.system, event.type);
    };
    socket.onclose = () => {
      onSocketClose();
    };
  }

  export function send(msg: Message) {
    if (debug) {
      console.debug("out", msg);
    }
    socket.send(JSON.stringify(msg));
  }

  export function onSocketMessage(msg: Message) {
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

  function onSocketClose() {
    function disconnect(seconds: number) {
      if (debug) {
        console.info(`socket closed, reconnecting in ${seconds} seconds`);
      }
      setTimeout(() => {
        socketConnect(system.cache.currentService!, system.cache.currentID);
      }, seconds * 1000);
    }

    if (!appUnloading) {
      disconnect(10);
    }
  }
}
