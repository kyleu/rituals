namespace socket {
  let socket: WebSocket;
  let appUnloading = false;

  function socketUrl() {
    const l = document.location;
    let protocol = "ws";
    if (l.protocol === "https:") {
      protocol = "wss";
    }
    return protocol + "://" + l.host + "/s";
  }

  export function setAppUnloading() {
    appUnloading = true;
  }

  export function socketConnect(svc: string, id: string) {
    system.cache.currentService = svc;
    system.cache.currentID = id;
    system.cache.connectTime = Date.now();

    socket = new WebSocket(socketUrl());
    socket.onopen = function () {
      if(debug) {
        console.debug("socket connected");
      }
      const msg = {svc: svc, cmd: command.client.connect, param: id};
      send(msg);
    };
    socket.onmessage = function (event) {
      const msg = JSON.parse(event.data);
      rituals.onSocketMessage(msg);
    };
    socket.onerror = function (event) {
      rituals.onError("socket", event.type);
    };
    socket.onclose = function () {
      onSocketClose();
    };
  }

  export function send(msg: rituals.Message) {
    if(debug) {
      console.debug("sending message");
      console.debug(msg);
    }
    socket.send(JSON.stringify(msg));
  }

  function onSocketClose() {
    function disconnect(seconds: number) {
      if(debug) {
        if (seconds === 1) {
          console.info("socket closed, reconnecting in a second");
        } else {
          console.info("socket closed, reconnecting in " + seconds + " seconds");
        }
      }
      setTimeout(() => {
        socketConnect(system.cache.currentService, system.cache.currentID);
      }, seconds * 1000);
    }

    if (!appUnloading) {
      const delta = Date.now() - system.cache.connectTime;
      if (delta < 2000) {
        disconnect(6);
      } else {
        disconnect(1);
      }
    }
  }
}
