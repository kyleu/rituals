import {Message} from "./socket";
import {req} from "./dom";
import {send} from "./app";
import {focusDelay} from "./util";
import {ChildAdd, ChildRemove, ChildUpdate, onChildAddModel, onChildRemoveModel, onChildUpdateModel, setTeamSprint} from "./workspace";
import {loadPermsForm, Permission, permissionsUpdate} from "./permission";

export type Team = {
  id: string;
  slug: string;
  title: string;
  icon: string;
  status: string;
  owner: string;
}

export function initTeam() {
  req("#modal-team-config-link").onclick = function() {
    focusDelay(req("#modal-team-config form input[name=\"title\"]"));
  }
  const frm = req<HTMLFormElement>("#modal-team-config form");
  frm.onsubmit = function () {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    send("update", {"title": title, "icon": icon, ...loadPermsForm(frm)});
    document.location.hash = "";
    return false;
  };
}

export function handleTeam(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param as Team);
    case "child-add":
      return onChildAddModel(m.param as ChildAdd);
    case "child-update":
      return onChildUpdateModel(m.param as ChildUpdate);
    case "child-remove":
      return onChildRemoveModel(m.param as ChildRemove);
    case "permissions":
      return permissionsUpdate(m.param as Permission[]);
    default:
      throw "invalid team command [" + m.cmd + "]"
  }
}

function onUpdate(param: Team) {
  req("owner-id").innerText = param.owner;
  const frm = req<HTMLFormElement>("#modal-team-config form");
  setTeamSprint("team", frm, null, null, param.title, param.icon);
}
