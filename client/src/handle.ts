import {Message} from "./socket";
import {addComment, Comment} from "./comments";
import {memberAdd, memberUpdate, onlineUpdate} from "./member";

export function handle(svc: string, m: Message) {
  switch (m.cmd) {
    case "comment":
      return addComment(m.param as Comment);
    case "online-update":
      return onlineUpdate(m.param.userID, m.param.connected);
    case "member-add":
      return memberAdd(m.param.userID, m.param.name, m.param.role);
    case "member-update":
      return memberUpdate(m.param.userID, m.param.name, m.param.role);
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
