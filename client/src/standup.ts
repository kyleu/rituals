import {Message} from "./socket";
import {els, req} from "./dom";
import {send} from "./app";
import {svgRef} from "./util";
import {modelBanner} from "./workspace";
import {flashCreate} from "./flash";
import {initReports, Report, reportAdd} from "./report";

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
  const frm = req<HTMLFormElement>("#modal-standup-config form");
  frm.onsubmit = function () {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    const team = req<HTMLInputElement>("select[name=\"team\"]", frm).value;
    const sprint = req<HTMLInputElement>("select[name=\"sprint\"]", frm).value;
    send("update", {"title": title, "icon": icon, "team": team, "sprint": sprint});
    document.location.hash = "";
    return false;
  };

  initReports();
}

export function handleStandup(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param as Standup);
    case "child-add":
      return reportAdd(m.param as Report);
    default:
      throw "invalid standup command [" + m.cmd + "]"
  }
}

function onUpdate(param: Standup) {
  const frm = req<HTMLFormElement>("#modal-standup-config form");
  req<HTMLInputElement>("input[name=\"title\"]", frm).value = param.title;
  for (const r of els<HTMLInputElement>("input[name=\"icon\"]", frm)) {
    r.checked = param.icon === r.value;
  }
  req<HTMLInputElement>("select[name=\"team\"]", frm).value = param.teamID ? param.teamID : "";
  req<HTMLInputElement>("select[name=\"sprint\"]", frm).value = param.sprintID ? param.sprintID : "";

  req("#model-title").innerText = param.title;
  req("#model-icon").innerHTML = svgRef(param.icon, 20);
  req("#model-banner").innerHTML = modelBanner("standup", frm, param.teamID, param.sprintID);

  flashCreate("standup", "success", "standup updated");
}
