import type {Message} from "./socket";
import {opt, req} from "./dom";
import {send} from "./app";
import {configFocus, setTeamSprint} from "./workspace";
import {initReports, Report, reportAdd, reportRemove} from "./report";
import {initPermissions, loadPermsForm, Permission, permissionsSprintToggle, permissionsTeamToggle, permissionsUpdate} from "./permission";

export type Standup = {
  id: string;
  slug: string;
  title: string;
  icon: string;
  status: string;
  teamID: string;
  sprintID: string;
  owner: string;
}

export function initStandup() {
  configFocus("standup");
  const frm = opt<HTMLFormElement>("#modal-standup-config form");
  if (frm) {
    const teamEl = req<HTMLSelectElement>("select[name=\"team\"]", frm);
    teamEl.onchange = () => {
      permissionsTeamToggle(teamEl.value !== "");
    };
    const sprintEl = req<HTMLSelectElement>("select[name=\"sprint\"]", frm);
    sprintEl.onchange = () => {
      permissionsSprintToggle(sprintEl.value !== "");
    };
    initPermissions(teamEl, sprintEl);
    frm.onsubmit = () => {
      const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
      const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
      send("update", {"title": title, "icon": icon, "team": teamEl.value, "sprint": sprintEl.value, ...loadPermsForm(frm)});
      document.location.hash = "";
      return false;
    };
  }
  initReports();
}

function onUpdate(param: Standup) {
  req("#owner-id").innerText = param.owner;
  const frm = req<HTMLFormElement>("#modal-standup-config");
  setTeamSprint("standup", frm, param.teamID, param.sprintID, param.title, param.icon);
}

export function handleStandup(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param as Standup);
    case "child-add":
      return reportAdd(m.param as Report);
    case "child-update":
      return console.log("TODO: child-update");
      // return reportUpdate(m.param as Report);
    case "child-remove":
      return reportRemove(m.param as string);
    case "permissions":
      return permissionsUpdate(m.param as Permission[]);
    default:
      throw new Error("invalid standup command [" + m.cmd + "]");
  }
}
