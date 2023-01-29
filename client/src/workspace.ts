import {req} from "./dom";

export function modelBanner(key: string, frm: HTMLFormElement, teamID: string, sprintID: string) {
  let ret = "";
  if (sprintID) {
    const s = req<HTMLInputElement>(`select[name="sprint"] option[value="${teamID}"]`, frm).innerText;
    ret += `<a href="/sprint/${sprintID}">${s}</a> `
  }
  ret += key;
  if (teamID) {
    const t = req<HTMLInputElement>(`select[name="team"] option[value="${teamID}"]`, frm).innerText;
    ret += ` in <a href="/team/${teamID}">${t}</a>`
  }
  return ret;
}
