import {initStories, Story, storyAdd, storyRemove, storyStatus} from "./story";
import {Message} from "./socket";
import {req} from "./dom";
import {send} from "./app";
import {setTeamSprint} from "./workspace";
import {tagsWire} from "./tags";
import {focusDelay} from "./util";
import {loadPermsForm, permissionsSprintToggle, permissionsTeamToggle} from "./permission";

export type Estimate = {
  id: string;
  slug: string;
  title: string;
  choices: string;
  icon: string;
  status: string;
  teamID: string;
  sprintID: string;
  owner: string;
}

export function initEstimate() {
  req("#modal-estimate-config-link").onclick = function() {
    focusDelay(req("#modal-estimate-config form input[name=\"title\"]"));
  }
  const frm = req<HTMLFormElement>("#modal-estimate-config form");
  const teamEl = req<HTMLSelectElement>("select[name=\"team\"]", frm);
  teamEl.onchange = function() {
    permissionsTeamToggle(teamEl.value !== "");
  }
  const sprintEl = req<HTMLSelectElement>("select[name=\"sprint\"]", frm);
  sprintEl.onchange = function() {
    permissionsSprintToggle(sprintEl.value !== "");
  }
  permissionsTeamToggle(teamEl.value !== "");
  permissionsSprintToggle(sprintEl.value !== "");
  frm.onsubmit = function() {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    const choices = req<HTMLInputElement>("input[name=\"choices\"]", frm).value;
    send("update", {"title": title, "icon": icon, "choices": choices, "team": teamEl.value, "sprint": sprintEl.value, ...loadPermsForm(frm)});
    document.location.hash = "";
    return false;
  };
  initStories();
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
    default:
      throw "invalid estimate command [" + m.cmd + "]"
  }
}

function onUpdate(param: Estimate) {
  const frm = req<HTMLFormElement>("#modal-estimate-config form");
  const ch = req<HTMLInputElement>("input[name=\"choices\"]", frm)
  ch.value = param.choices;
  if(ch.parentElement) {
    tagsWire(ch.parentElement);
  }
  setTeamSprint("estimate", frm, param.teamID, param.sprintID, param.title, param.icon);
}
