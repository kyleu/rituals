import {Message} from "./socket";
import {req} from "./dom";
import {send} from "./app";
import {setTeamSprint} from "./workspace";
import {initReports, Report, reportAdd, reportRemove} from "./report";
import {focusDelay} from "./util";

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
  req("#modal-standup-config-link").onclick = function() {
    focusDelay(req("#modal-standup-config form input[name=\"title\"]"));
  }
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
    case "child-remove":
      return reportRemove(m.param as string);
    default:
      throw "invalid standup command [" + m.cmd + "]"
  }
}

function onUpdate(param: Standup) {
  const frm = req<HTMLFormElement>("#modal-standup-config form");
  setTeamSprint("standup", frm, param.teamID, param.sprintID, param.title, param.icon);
}
