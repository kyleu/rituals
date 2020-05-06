function socketUrl() {
  let l = document.location;
  let protocol = "ws";
  if(l.protocol === "https:") {
    protocol = "wss";
  }
  return protocol + "://" + l.host + "/s";
}

let currentService = "";
let currentId = "";

let connectTime = 0;

function socketConnect(svc: string, id: any) {
  currentService = svc;
  currentId = id;
  connectTime = Date.now()

  socket = new WebSocket(socketUrl());
  socket.onopen = function () {
    console.debug("socket connected")
    const msg = {"svc": svc, "cmd": "connect", "param": id};
    send(msg);
  };
  socket.onmessage = function (event) {
    const msg = JSON.parse(event.data);
    onSocketMessage(msg);
  }
  socket.onerror = function (event) {
    onError("socket error: " + event.type);
  }
  socket.onclose = function (event) {
    onSocketClose();
  }
}

function send(msg: Message) {
  console.log("sending message");
  console.log(msg);
  socket.send(JSON.stringify(msg));
}

function onSocketClose() {
  if(!appUnloading) {
    let delta = Date.now() - connectTime;
    if(delta < 2000) {
      console.warn("socket closed immediately, reconnecting in 4 seconds");
      setTimeout(() => { socketConnect(currentService, currentId) }, 4000)
    } else {
      console.warn("socket closed, reconnecting in a second");
      setTimeout(() => { socketConnect(currentService, currentId) }, 1000)
    }
  }
}
