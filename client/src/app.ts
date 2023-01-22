// $PF_IGNORE$
import {handle} from "./handle";
import {initComments} from "./comments";
import {Socket} from "./socket";

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
  handle(m);
}

function err(e: any) {
  console.log("[socket error]: " + e);
}

export function initWorkspace(t: string, idStr: string) {
  svc = t;
  id = idStr;
  initComments();
  sock = new Socket(true, open, recv, err, "/" + t + "/" + id + "/connect");
  console.log("loaded [" + t + "] workspace");
}

export function send(cmd: string, param: any) {
  sock.send({channel: svc + ":" + id, cmd: cmd, param: param})
}

export function appInit(): void {
  (window as any).rituals.initWorkspace = initWorkspace;
}
