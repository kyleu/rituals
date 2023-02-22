import {Message} from "./socket";
import {req} from "./dom";
import {send} from "./app";
import {focusDelay} from "./util";
import {ChildAdd, ChildRemove, ChildUpdate, onChildAddModel, onChildRemoveModel, onChildUpdateModel, setTeamSprint} from "./workspace";
import {loadPermsForm, Permission, permissionsTeamToggle, permissionsUpdate} from "./permission";

export type Sprint = {
  id: string;
  slug: string;
  title: string;
  startDate: string;
  endDate: string;
  icon: string;
  status: string;
  teamID: string;
  owner: string;
}

export function initSprint() {
  req("#modal-sprint-config-link").onclick = function() {
    focusDelay(req("#modal-sprint-config form input[name=\"title\"]"));
  }
  const frm = req<HTMLFormElement>("#modal-sprint-config form");
  const teamEl = req<HTMLSelectElement>("select[name=\"team\"]", frm);
  teamEl.onchange = function() {
    permissionsTeamToggle(teamEl.value !== "");
  }
  permissionsTeamToggle(teamEl.value !== "");
  frm.onsubmit = function() {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    const startDate = req<HTMLInputElement>("input[name=\"startDate\"]", frm).value;
    const endDate = req<HTMLInputElement>("input[name=\"endDate\"]", frm).value;
    send("update", {"title": title, "icon": icon, "startDate": startDate, "endDate": endDate, "team": teamEl.value, ...loadPermsForm(frm)});
    document.location.hash = "";
    return false;
  };
}

export function handleSprint(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param as Sprint);
    case "child-add":
      return onChildAddModel(m.param as ChildAdd);
    case "child-update":
      return onChildUpdateModel(m.param as ChildUpdate);
    case "child-remove":
      return onChildRemoveModel(m.param as ChildRemove);
    case "permissions":
      return permissionsUpdate(m.param as Permission[]);
    default:
      throw "invalid sprint command [" + m.cmd + "]"
  }
}

function onUpdate(param: Sprint) {
  req("owner-id").innerText = param.owner;
  const frm = req<HTMLFormElement>("#modal-sprint-config form");
  req<HTMLInputElement>("input[name=\"startDate\"]", frm).value = ds(param.startDate);
  req<HTMLInputElement>("input[name=\"endDate\"]", frm).value = ds(param.endDate);
  req("#model-summary").innerText = summary(param);
  setTeamSprint("sprint", frm, param.teamID, null, param.title, param.icon);
}

function summary(param: Sprint) {
  let ret = "";
  if (param.startDate) {
    ret += "starts ";
    ret += ds(param.startDate);
    if (param.endDate) {
      ret += ", ";
    }
  }
  if (param.endDate) {
    ret += "ends ";
    ret += ds(param.endDate);
  }
  return ret;
}

function ds(s: string) {
  return `${s}`.split('T')[0];
}
