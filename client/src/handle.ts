import {Message} from "./socket";
import {addComment, Comment} from "./comments";
import {memberAdd, memberRemove, memberUpdate, onlineUpdate} from "./member";
import {flashCreate} from "./flash";

export function handle(svc: string, m: Message) {
  switch (m.cmd) {
    case "error":
      return onError(m.param.message);
    case "comment":
      return addComment(m.param as Comment);
    case "online-update":
      return onlineUpdate(m.param.userID, m.param.connected);
    case "member-add":
      return memberAdd(m.param.userID, m.param.name, m.param.role);
    case "member-update":
      return memberUpdate(m.param.userID, m.param.name, m.param.role);
    case "member-remove":
      return memberRemove(m.param);
  }
  switch (svc) {
    case "estimate":
      switch (m.cmd) {
        case "fooo":
          return addComment(m.param as Comment);
      }
  }
  return;
}

function onError(message: string) {
  flashCreate("error", "error", message);
}

