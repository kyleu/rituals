// $PF_IGNORE$
import {initComments} from "./comments";
import {handle} from "./handle";
import {initMembers} from "./members";
import {Socket} from "./socket";
import {initEstimate} from "./estimate";
import {initTeam} from "./team";
import {initSprint} from "./sprint";
import {initStandup} from "./standup";
import {initRetro} from "./retro";

let sock: Socket
let svc: string
let id: string

function open() {
  console.log("[socket]: open");
}

function recv(m: any) {
  const list = document.getElementById("socket-list");
  if (list) {
    const pre = document.createElement("pre");
    pre.innerText = JSON.stringify(m, null, 2);
    list.append(pre);
  }
  handle(svc, m);
}

function err(e: any) {
  console.log("[socket error]: " + e);
}

export function initWorkspace(t: string, idStr: string) {
  svc = t;
  id = idStr;
  initComments();
  initMembers();
  switch(svc) {
    case "team":
      initTeam();
      break;
    case "sprint":
      initSprint();
      break;
    case "estimate":
      initEstimate();
      break;
    case "standup":
      initStandup();
      break;
    case "retro":
      initRetro();
      break;
  }
  sock = new Socket(true, open, recv, err, "/" + svc + "/" + id + "/connect");
  console.log("loaded [" + svc + "] workspace [" + id + "]");
}

export function send(cmd: string, param: any) {
  sock.send({channel: svc + ":" + id, cmd: cmd, param: param})
}


declare global {
  interface Window {
    initWorkspace: (t: string, idStr: string) => void
  }
}

export function appInit(): void {
  window.initWorkspace = initWorkspace;
}
