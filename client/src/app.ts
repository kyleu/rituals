// $PF_IGNORE$
import {JSX} from "./jsx";


function open() {
  // console.log("[tap]: open");
}

function recv(m: any) {
  const pre = document.createElement("pre");
  pre.innerText = JSON.stringify(m, null, 2);
  document.getElementById("tap-list")?.append(pre);
}

function err(e: any) {
  console.log("[tap error]: " + e);
}

export function initWorkspace(t: string, x: any, members?: any[], permissions?: any[]) {
  new (window as any).rituals.Socket(true, open, recv, err, "/" + t + "/" + x.slug + "/connect");
  console.log("loaded [" + t + "] workspace with [" + members?.length + "] members and [" + permissions?.length + "] permissions");
}

export function appInit(): void {
  (window as any).rituals.initWorkspace = initWorkspace;
}
