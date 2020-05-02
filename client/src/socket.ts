function socketUrl() {
  let l = document.location;
  let protocol = "ws";
  if(l.protocol === "https:") {
    protocol = "wss";
  }
  return protocol + "://" + l.host + "/s";
}

let currentService: string | null = null;
let currentId: string | null = null;

function connect(svc: string, id: any) {
  currentService = svc;
  currentId = id;
  socket = new WebSocket(socketUrl());
  socket.onopen = function () {
    const msg = {"svc": svc, "cmd": "connect", "param": id};
    send(msg);
  };
  socket.onmessage = function (event) {
    const msg = JSON.parse(event.data)
    onMessage(msg)
  }
}

function send(msg: Message) {
  console.log("sending message");
  console.log(msg);
  socket.send(JSON.stringify(msg));
}
