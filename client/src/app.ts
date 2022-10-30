// $PF_IGNORE$
import {JSX} from "./jsx";

export function initWorkspace(t: string) {
  console.log("!!!!");
}

export function appInit(): void {
  (window as any).rituals.initWorkspace = initWorkspace;
}
