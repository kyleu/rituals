import {Message} from "./socket";
import {commentAdd, Comment} from "./comments";
import {memberAdd, memberRemove, memberUpdate, onlineUpdate} from "./members";
import {flashCreate} from "./flash";
import {storyAdd} from "./estimate";

export function handle(svc: string, m: Message) {
  switch (m.cmd) {
    case "error":
      return onError(m.param.message);
    case "comment":
      return commentAdd(m.param as Comment);
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
        case "story-add":
          return storyAdd();
      }
  }
  return;
}

function onError(message: string) {
  flashCreate("error", "error", message);
}

