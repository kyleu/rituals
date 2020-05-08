let socket: WebSocket;

function socketUrl() {
  const l = document.location;
  let protocol = "ws";
  if (l.protocol === "https:") {
    protocol = "wss";
  }
  return protocol + "://" + l.host + "/s";
}

function socketConnect(svc: string, id: string) {
  systemCache.currentService = svc;
  systemCache.currentID = id;
  systemCache.connectTime = Date.now();

  socket = new WebSocket(socketUrl());
  socket.onopen = function () {
    console.debug("socket connected");
    const msg = { svc: svc, cmd: clientCmd.connect, param: id };
    send(msg);
  };
  socket.onmessage = function (event) {
    const msg = JSON.parse(event.data);
    onSocketMessage(msg);
  };
  socket.onerror = function (event) {
    onError("socket error: " + event.type);
  };
  socket.onclose = function () {
    onSocketClose();
  };
}

function send(msg: Message) {
  console.log("sending message");
  console.log(msg);
  socket.send(JSON.stringify(msg));
}

function onSocketClose() {
  function disconnect(seconds: number) {
    if (seconds === 1) {
      console.warn("socket closed, reconnecting in a second");
    } else {
      console.warn("socket closed, reconnecting in " + seconds + " seconds");
    }
    setTimeout(() => {
      socketConnect(systemCache.currentService, systemCache.currentID);
    }, 4000);
  }
  if (!appUnloading) {
    const delta = Date.now() - systemCache.connectTime;
    if (delta < 2000) {
      disconnect(5);
    } else {
      disconnect(1);
    }
  }
}
