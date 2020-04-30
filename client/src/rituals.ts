declare var UIkit: any;

let socket: WebSocket;
let debug = true;

let $: (selector: string, context?: any) => [HTMLElement] = UIkit.util.$$;

interface Message {
  svc: String;
  cmd: String;
  param: any;
}

function onMessage(msg: Message) {
  console.log("message received");
  console.log(msg);
  switch(msg.svc) {
    case "estimate":
      onEstimateMessage(msg.cmd, msg.param);
      break;
    default:
      console.warn("Unhandled message for service [" + msg.svc + "]")
  }
}

function socketUrl() {
  let l = document.location;
  let protocol = "ws";
  if(l.protocol === "https:") {
    protocol = "wss";
  }
  return protocol + "://" + l.host + "/s";
}

function connect(svc: String, value: any) {
  socket = new WebSocket(socketUrl());
  socket.onopen = function () {
    const msg = {"svc": svc, "cmd": "connect", "param": value};
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

function sandbox() {
  send({svc: "estimate", cmd: "sandbox", param: null});
}
