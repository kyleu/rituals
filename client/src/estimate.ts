import {initStories, Story, storyAdd, storyRemove, storyStatus} from "./story";
import type {Message} from "./socket";
import {opt, req} from "./dom";
import {send} from "./app";
import {configFocus, setTeamSprint} from "./workspace";
import {tagsWire} from "./tags";
import {initPermissions, loadPermsForm, Permission, permissionsSprintToggle, permissionsTeamToggle, permissionsUpdate} from "./permission";

export type Estimate = {
  id: string;
  slug: string;
  title: string;
  choices: string[];
  icon: string;
  status: string;
  teamID: string;
  sprintID: string;
  owner: string;
}

export function initEstimate() {
  configFocus("estimate");
  const frm = opt<HTMLFormElement>("#modal-estimate-config form");
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
      const choices = req<HTMLInputElement>("input[name=\"choices\"]", frm).value;
      send("update", {"title": title, "icon": icon, "choices": choices, "team": teamEl.value, "sprint": sprintEl.value, ...loadPermsForm(frm)});
      document.location.hash = "";
      return false;
    };
  }
  initStories();
}

function onUpdate(param: Estimate) {
  req("#owner-id").innerText = param.owner;
  const panel = req<HTMLElement>("#modal-estimate-config");
  const frm = opt<HTMLFormElement>("form", panel);
  if (frm) {
    const ch = req<HTMLInputElement>("input[name=\"choices\"]", frm);
    ch.value = param.choices.join(",");
    if (ch.parentElement) {
      tagsWire(ch.parentElement);
    }
  } else {
    req(".config-panel-choices", panel).innerText = param.choices.join(", ");
  }
  setTeamSprint("estimate", panel, param.teamID, param.sprintID, param.title, param.icon);
}

export function handleEstimate(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param as Estimate);
    case "child-add":
      return storyAdd(m.param as Story);
    case "child-status":
      return storyStatus(m.param as Story);
    case "child-remove":
      return storyRemove(m.param as string);
    case "permissions":
      return permissionsUpdate(m.param as Permission[]);
    default:
      throw new Error("invalid estimate command [" + m.cmd + "]");
  }
}
