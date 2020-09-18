namespace socket {
  export interface Message {
    readonly svc: string;
    readonly cmd: string;
    readonly param: any;
  }

  let sock: WebSocket;
  let connected = false;
  let appUnloading = false;
  let pendingMessages: Message[] = [];

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

    sock = new WebSocket(socketUrl());
    sock.onopen = () => onSocketOpen;
    sock.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      onSocketMessage(msg);
    };
    sock.onerror = (event) => {
      // rituals.onError(services.system, event.type);
    };
    sock.onclose = onSocketClose;
  }

  export function send(msg: Message) {
    if (connected) {
      if (debug) {
        console.debug("out", msg);
      }
      const m = JSON.stringify(msg, null, 2);
      sock.send(m);
    } else {
      pendingMessages.push(msg);
    }
  }

  function onSocketOpen(svc: string, id: string) {
    connected = true;
    pendingMessages.forEach(send);
    pendingMessages = [];
    send({ svc: system.cache.currentService!.key, cmd: command.client.connect, param: system.cache.currentID });
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
