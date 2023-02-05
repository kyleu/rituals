import {Message} from "./socket";
import {els, req} from "./dom";
import {send} from "./app";
import {svgRef} from "./util";
import {flashCreate} from "./flash";
import {ChildAdd, ChildRemove, modelBanner, onChildAddModel, onChildRemoveModel} from "./workspace";

export type Team = {
  id: string;
  slug: string;
  title: string;
  icon: string;
  status: string;
  owner: string;
}

export function initTeam() {
  const frm = req<HTMLFormElement>("#modal-team-config form");
  frm.onsubmit = function () {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    send("update", {"title": title, "icon": icon});
    document.location.hash = "";
    return false;
  };
}

export function handleTeam(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param as Team);
    case "child-add":
      return onChildAddModel(m.param as ChildAdd);
    case "child-remove":
      return onChildRemoveModel(m.param as ChildRemove);
    default:
      throw "invalid team command [" + m.cmd + "]"
  }
}

function onUpdate(param: Team) {
  const frm = req<HTMLFormElement>("#modal-team-config form");
  req<HTMLInputElement>("input[name=\"title\"]", frm).value = param.title;
  for (const r of els<HTMLInputElement>("input[name=\"icon\"]", frm)) {
    r.checked = param.icon === r.value;
  }

  req("#model-title").innerText = param.title;
  req("#model-icon").innerHTML = svgRef(param.icon, 20);
  req("#model-banner").innerHTML = modelBanner("team", frm, "", "");

  flashCreate("team", "success", "team updated");
}
