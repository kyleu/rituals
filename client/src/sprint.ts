import type {Message} from "./socket";
import {opt, req} from "./dom";
import {send} from "./app";
import {ChildAdd, ChildRemove, ChildUpdate, configFocus, onChildAddModel, onChildRemoveModel, onChildUpdateModel, setTeamSprint} from "./workspace";
import {initPermissions, loadPermsForm, Permission, permissionsTeamToggle, permissionsUpdate} from "./permission";

export type Sprint = {
  id: string;
  slug: string;
  title: string;
  startDate: string;
  endDate: string;
  icon: string;
  status: string;
  teamID: string;
}

export function initSprint() {
  configFocus("sprint");
  const frm = opt<HTMLFormElement>("#modal-sprint-config form");
  if (frm) {
    const teamEl = req<HTMLSelectElement>("select[name=\"team\"]", frm);
    teamEl.onchange = () => {
      permissionsTeamToggle(teamEl.value !== "");
    };
    initPermissions(teamEl);
    frm.onsubmit = () => {
      const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
      const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
      const startDate = req<HTMLInputElement>("input[name=\"startDate\"]", frm).value;
      const endDate = req<HTMLInputElement>("input[name=\"endDate\"]", frm).value;
      send("update", {"title": title, "icon": icon, "startDate": startDate, "endDate": endDate, "team": teamEl.value, ...loadPermsForm(frm)});
      document.location.hash = "";
      return false;
    };
  }
}

function ds(s: string) {
  return `${s}`.split("T")[0];
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

function onUpdate(param: Sprint) {
  const panel = req<HTMLElement>("#modal-sprint-config");
  const frm = opt<HTMLFormElement>("form", panel);
  if (frm) {
    req<HTMLInputElement>("input[name=\"startDate\"]", frm).value = ds(param.startDate);
    req<HTMLInputElement>("input[name=\"endDate\"]", frm).value = ds(param.endDate);
  } else {
    req(".config-panel-startDate", panel).innerText = ds(param.startDate);
    req(".config-panel-endDate", panel).innerText = ds(param.endDate);
  }
  req("#model-summary").innerText = summary(param);
  setTeamSprint("sprint", panel, param.teamID, null, param.title, param.icon);
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
      throw new Error("invalid sprint command [" + m.cmd + "]");
  }
}
