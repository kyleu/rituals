import type {Message} from "./socket";
import {opt, req} from "./dom";
import {send} from "./app";
import {ChildAdd, ChildRemove, ChildUpdate, configFocus, onChildAddModel, onChildRemoveModel, onChildUpdateModel, setTeamSprint} from "./workspace";
import {loadPermsForm, Permission, permissionsUpdate} from "./permission";

export type Team = {
  id: string;
  slug: string;
  title: string;
  icon: string;
  status: string;
}

export function initTeam() {
  configFocus("team");
  const frm = opt<HTMLFormElement>("#modal-team-config form");
  if (frm) {
    frm.onsubmit = () => {
      const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
      const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
      send("update", {"title": title, "icon": icon, ...loadPermsForm(frm)});
      document.location.hash = "";
      return false;
    };
  }
}

function onUpdate(param: Team) {
  const frm = req<HTMLFormElement>("#modal-team-config");
  setTeamSprint("team", frm, null, null, param.title, param.icon);
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
      throw new Error("invalid team command [" + m.cmd + "]");
  }
}
