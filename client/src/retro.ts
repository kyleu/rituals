import {Message} from "./socket";
import {els, req} from "./dom";
import {send} from "./app";
import {svgRef} from "./util";
import {modelBanner} from "./workspace";
import {flashCreate} from "./flash";
import {tagsWire} from "./tags";

export function initRetro() {
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
}

export function handleRetro(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param);
    default:
      throw "invalid retro command [" + m.cmd + "]"
  }
}

function onUpdate(param: any) {
  const frm = req<HTMLFormElement>("#modal-retro-config form");
  req<HTMLInputElement>("input[name=\"title\"]", frm).value = param.title;
  for (const r of els<HTMLInputElement>("input[name=\"icon\"]", frm)) {
    r.checked = param.icon === r.value;
  }
  const cat = req<HTMLInputElement>("input[name=\"categories\"]", frm)
  cat.value = param.categories;
  if(cat.parentElement) {
    tagsWire(cat.parentElement);
  }
  req<HTMLInputElement>("select[name=\"team\"]", frm).value = param.teamID ? param.teamID : "";
  req<HTMLInputElement>("select[name=\"sprint\"]", frm).value = param.sprintID ? param.sprintID : "";

  req("#model-title").innerText = param.title;
  req("#model-icon").innerHTML = svgRef(param.icon, 20);
  req("#model-banner").innerHTML = modelBanner("retro", frm, param.teamID, param.sprintID);

  flashCreate("retro", "success", "retro updated");
}
