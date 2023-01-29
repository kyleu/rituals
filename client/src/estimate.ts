import {initStories, storyAdd} from "./story";
import {Message} from "./socket";
import {els, req} from "./dom";
import {send} from "./app";
import {svgRef} from "./util";
import {modelBanner} from "./workspace";
import {flashCreate} from "./flash";
import {tagsWire} from "./tags";

export function initEstimate() {
  const frm = req<HTMLFormElement>("#modal-estimate-config form");
  frm.onsubmit = function() {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    const choices = req<HTMLInputElement>("input[name=\"choices\"]", frm).value;
    const team = req<HTMLInputElement>("select[name=\"team\"]", frm).value;
    const sprint = req<HTMLInputElement>("select[name=\"sprint\"]", frm).value;
    send("update", {"title": title, "icon": icon, "choices": choices, "team": team, "sprint": sprint});
    document.location.hash = "";
    return false;
  };

  initStories();
}

export function handleEstimate(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param);
    case "story-add":
      return storyAdd(m.param);
    default:
      throw "invalid estimate command [" + m.cmd + "]"
  }
}

function onUpdate(param: any) {
  const frm = req<HTMLFormElement>("#modal-estimate-config form");
  req<HTMLInputElement>("input[name=\"title\"]", frm).value = param.title;
  for (const r of els<HTMLInputElement>("input[name=\"icon\"]", frm)) {
    r.checked = param.icon === r.value;
  }
  const ch = req<HTMLInputElement>("input[name=\"choices\"]", frm)
  ch.value = param.choices;
  if(ch.parentElement) {
    tagsWire(ch.parentElement);
  }
  req<HTMLInputElement>("select[name=\"team\"]", frm).value = param.teamID ? param.teamID : "";
  req<HTMLInputElement>("select[name=\"sprint\"]", frm).value = param.sprintID ? param.sprintID : "";

  req("#model-title").innerText = param.title;
  req("#model-icon").innerHTML = svgRef(param.icon, 20);
  req("#model-banner").innerHTML = modelBanner("estimate", frm, param.teamID, param.sprintID);

  flashCreate("estimate", "success", "estimate updated");
}
