import {Message} from "./socket";
import {commentAdd, Comment} from "./comment";
import {memberAdd, memberRemove, memberUpdate, onlineUpdate} from "./member";
import {flashCreate} from "./flash";
import {handleTeam} from "./team";
import {handleSprint} from "./sprint";
import {handleEstimate} from "./estimate";
import {handleStandup} from "./standup";
import {handleRetro} from "./retro";

export function handle(svc: string, m: Message) {
  switch (m.cmd) {
    case "error":
      return onError((m.param as { message: string }).message);
    case "comment":
      return commentAdd(m.param as Comment);
    case "online-update":
      const ou = m.param as { userID: string; connected: boolean; }
      return onlineUpdate(ou.userID, ou.connected);
    case "member-add":
      const ma = m.param as { userID: string; name: string; role: string; }
      return memberAdd(ma.userID, ma.name, ma.role);
    case "member-update":
      const mu = m.param as { userID: string; name: string; role: string; }
      return memberUpdate(mu.userID, mu.name, mu.role);
    case "member-remove":
      return memberRemove(m.param as string);
  }
  switch (svc) {
    case "team":
      return handleTeam(m);
    case "sprint":
      return handleSprint(m);
    case "estimate":
      return handleEstimate(m);
    case "standup":
      return handleStandup(m);
    case "retro":
      return handleRetro(m);
    default:
      throw "invalid service [" + svc + "]";
  }
}

function onError(log: string) {
  flashCreate("error", "error", log);
}
