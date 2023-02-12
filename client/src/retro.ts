import {Message} from "./socket";
import {els, req} from "./dom";
import {send} from "./app";
import {setTeamSprint} from "./workspace";
import {tagsWire} from "./tags";
import {Feedback, feedbackAdd, feedbackRemove, initFeedbacks} from "./feedback";
import {focusDelay} from "./util";

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
  req("#modal-retro-config-link").onclick = function() {
    focusDelay(req("#modal-retro-config form input[name=\"title\"]"));
  }
  const frm = req<HTMLFormElement>("#modal-retro-config form");
  frm.onsubmit = function() {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    const categories = req<HTMLInputElement>("input[name=\"categories\"]", frm).value;
    const team = req<HTMLInputElement>("select[name=\"team\"]", frm).value;
    const sprint = req<HTMLInputElement>("select[name=\"sprint\"]", frm).value;
    send("update", {"title": title, "icon": icon, "categories": categories, "team": team, "sprint": sprint});
    document.location.hash = "";
    return false;
  };
  initFeedbacks();
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
    default:
      throw "invalid retro command [" + m.cmd + "]"
  }
}

function onUpdate(param: Retro) {
  const frm = req<HTMLFormElement>("#modal-retro-config form");
  const cat = req<HTMLInputElement>("input[name=\"categories\"]", frm)
  cat.value = param.categories.join(",");
  if(cat.parentElement) {
    tagsWire(cat.parentElement);
  }
  const listEl = req("#category-list");
  for (const catEl of els(" .category", listEl)) {

  }

  setTeamSprint("retro", frm, param.teamID, param.sprintID, param.title, param.icon);
}
