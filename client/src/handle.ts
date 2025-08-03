import type { Message } from "./socket";
import { commentAdd, Comment } from "./comment";
import { memberAdd, MemberMessage, memberRemove, memberUpdate, onlineUpdate } from "./member";
import { flashCreate } from "./flash";
import { handleTeam } from "./team";
import { handleSprint } from "./sprint";
import { handleEstimate } from "./estimate";
import { handleStandup } from "./standup";
import { handleRetro } from "./retro";

function onMessage(x: { level: "success" | "error"; message: string }) {
  flashCreate("adhoc-message", x.level, x.message);
}

function handleService(svc: string, m: Message) {
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
      throw new Error("invalid service [" + svc + "]");
  }
}

export function handle(svc: string, m: Message) {
  switch (m.cmd) {
    case "reset":
      document.location.href = "/";
      return;
    case "message":
      return onMessage(m.param as { level: "success" | "error"; message: string });
    case "comment":
      return commentAdd(m.param as Comment);
    case "online-update":
      return onlineUpdate(m.param as { userID: string; connected: boolean });
    case "member-add":
      if (svc === "estimate") {
        handleEstimate(m);
      }
      return memberAdd(m.param as MemberMessage);
    case "member-update":
      return memberUpdate(m.param as MemberMessage);
    case "member-remove":
      return memberRemove(m.param as string);
    default:
      return handleService(svc, m);
  }
}
