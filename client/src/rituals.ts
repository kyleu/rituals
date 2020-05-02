declare var UIkit: any;

let socket: WebSocket;
let debug = true;

interface Message {
  svc: string;
  cmd: string;
  param: any;
}

interface Profile {
  userID:    string;
  name:      string;
  role:      string;
  theme:     string;
  navColor:  string;
  linkColor: string;
  locale:    string;
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
    case "system":
      onSystemMessage(msg.cmd, msg.param);
      break;
    case "estimate":
      onEstimateMessage(msg.cmd, msg.param);
      break;
    default:
      console.warn("Unhandled message for service [" + msg.svc + "]");
  }
}

function setDetail(param: Detail) {
  $id("model-title").innerText = param.title;
}

let activeProfile: Profile | null = null;

function setProfile(profile: Profile) {
  activeProfile = profile
}

function onSystemMessage(cmd: string, param: any) {
  switch(cmd) {
    case "profile":
      setProfile(param);
      break;
    default:
      console.warn("Unhandled system message for command [" + cmd + "]");
  }
}
