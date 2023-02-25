import type {Message} from "./socket";
import {opt, req} from "./dom";
import {send} from "./app";
import {configFocus, setTeamSprint} from "./workspace";
import {tagsWire} from "./tags";
import {Feedback, feedbackAdd, feedbackRemove, initFeedbacks} from "./feedback";
import {loadPermsForm, Permission, permissionsSprintToggle, permissionsTeamToggle, permissionsUpdate} from "./permission";

export type Retro = {
  id: string;
  slug: string;
  title: string;
  categories: string[];
  icon: string;
  status: string;
  teamID: string;
  sprintID: string;
  owner: string;
}

export function initRetro() {
  configFocus("retro");
  const frm = opt<HTMLFormElement>("#modal-retro-config form");
  if (frm) {
    const teamEl = req<HTMLSelectElement>("select[name=\"team\"]", frm);
    teamEl.onchange = () => {
      permissionsTeamToggle(teamEl.value !== "");
    };
    const sprintEl = req<HTMLSelectElement>("select[name=\"sprint\"]", frm);
    sprintEl.onchange = () => {
      permissionsSprintToggle(sprintEl.value !== "");
    };
    permissionsTeamToggle(teamEl.value !== "");
    permissionsSprintToggle(sprintEl.value !== "");
    frm.onsubmit = () => {
      const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
      const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
      const categories = req<HTMLInputElement>("input[name=\"categories\"]", frm).value;
      send("update", {"title": title, "icon": icon, "categories": categories, "team": teamEl.value, "sprint": sprintEl.value, ...loadPermsForm(frm)});
      document.location.hash = "";
      return false;
    };
  }
  initFeedbacks();
}

function onUpdate(param: Retro) {
  req("#owner-id").innerText = param.owner;
  const panel = req<HTMLElement>("#modal-retro-config");
  const frm = opt<HTMLFormElement>("form", panel);
  if (frm) {
    const cat = req<HTMLInputElement>("input[name=\"categories\"]", frm);
    cat.value = param.categories.join(",");
    if (cat.parentElement) {
      tagsWire(cat.parentElement);
    }
  } else {
    req(".config-panel-categories", panel).innerText = param.categories.join(", ");
  }
  setTeamSprint("retro", panel, param.teamID, param.sprintID, param.title, param.icon);
  // const listEl = req("#category-list");
  // for (const catEl of els(" .category", listEl)) {
  //   TODO
  // }
}

export function handleRetro(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param as Retro);
    case "child-add":
      return feedbackAdd(m.param as Feedback);
    case "child-update":
      return console.log("TODO: child-update");
    // return feedbackUpdate(m.param as Feedback);
    case "child-remove":
      return feedbackRemove(m.param as string);
    case "permissions":
      return permissionsUpdate(m.param as Permission[]);
    default:
      throw new Error("invalid retro command [" + m.cmd + "]");
  }
}
