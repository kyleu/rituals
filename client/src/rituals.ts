declare var UIkit: any;

let socket: WebSocket;
let debug = true;

interface Message {
  svc: string;
  cmd: string;
  param: any;
}

interface Detail {
  id: string;
  slug: string;
  password: string;
  title: string;
  owner: string;
  status: { key: string; };
  created: string;
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

function setDetail(param: Detail) {
  $id("model-title").innerText = param.title + "!!!!";
}

function sandbox() {
  send({svc: "estimate", cmd: "sandbox", param: null});
}
