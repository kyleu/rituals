let socket: WebSocket;
let debug = true;

interface Message {
  t: String;
}

function onMessage(msg: Message) {
  console.log("message received", msg);
}

function socketUrl() {
  let l = document.location;
  let protocol = "ws";
  if(l.protocol == "https") {
    protocol = "wss";
  }
  return protocol + "://" + l.host + "/s";
}

function connect(k: String, v: String) {
  socket = new WebSocket(socketUrl());
  socket.onopen = function () {
    const msg = {"t": "connect", "k": k, "v": v};
    send(msg);
  };
  socket.onmessage = function (event) {
    const msg = JSON.parse(event.data)
    onMessage(msg)
  }
}

function send(msg: Message) {
  console.log("sending message", msg);
  socket.send(JSON.stringify(msg));
}

function sandbox() {
  send({t: "sandbox"});
}
