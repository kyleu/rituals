import {Message} from "./socket";
import {els, req} from "./dom";
import {send} from "./app";
import {svgRef} from "./util";
import {flashCreate} from "./flash";
import {modelBanner} from "./workspace";

export function initSprint() {
  const frm = req<HTMLFormElement>("#modal-sprint-config form");
  frm.onsubmit = function() {
    const title = req<HTMLInputElement>("input[name=\"title\"]", frm).value;
    const icon = req<HTMLInputElement>("input[name=\"icon\"]:checked", frm).value;
    const startDate = req<HTMLInputElement>("input[name=\"startDate\"]", frm).value;
    const endDate = req<HTMLInputElement>("input[name=\"endDate\"]", frm).value;
    const team = req<HTMLInputElement>("select[name=\"team\"]", frm).value;
    send("update", {"title": title, "icon": icon, "startDate": startDate, "endDate": endDate, "team": team});
    document.location.hash = "";
    return false;
  };
}

export function handleSprint(m: Message) {
  switch (m.cmd) {
    case "update":
      return onUpdate(m.param);
    default:
      throw "invalid sprint command [" + m.cmd + "]"
  }
}

function onUpdate(param: any) {
  const frm = req<HTMLFormElement>("#modal-sprint-config form");
  req<HTMLInputElement>("input[name=\"title\"]", frm).value = param.title;
  for (const r of els<HTMLInputElement>("input[name=\"icon\"]", frm)) {
    r.checked = param.icon === r.value;
  }
  req<HTMLInputElement>("input[name=\"startDate\"]", frm).value = ds(param.startDate);
  req<HTMLInputElement>("input[name=\"endDate\"]", frm).value = ds(param.endDate);
  req<HTMLInputElement>("select[name=\"team\"]", frm).value = param.teamID ? param.teamID : "";

  req("#model-title").innerText = param.title;
  req("#model-icon").innerHTML = svgRef(param.icon, 20);
  req("#model-summary").innerText = summary(param);
  req("#model-banner").innerHTML = modelBanner("sprint", frm, param.teamID, param.sprintID);

  flashCreate("sprint", "success", "sprint updated");
}

function summary(param: any) {
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

function ds(s: string) {
  return `${s}`.split('T')[0];
}
