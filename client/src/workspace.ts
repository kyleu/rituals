import {opt, req} from "./dom";

export type ChildAdd = {
  "type": string;
  "id": string;
  "title": string;
  "path": string;
}

export type ChildRemove = {
  "type": string;
  "id": string;
}

export function modelBanner(key: string, frm: HTMLFormElement, teamID: string, sprintID: string) {
  let ret = "";
  if (sprintID) {
    const el = opt<HTMLInputElement>(`select[name="sprint"] option[value="${sprintID}"]`, frm);
    if (el) {
      const s = el.innerText;
      ret += `<a href="/sprint/${sprintID}">${s}</a> `;
    }
  }
  ret += key;
  if (teamID) {
    const el = opt<HTMLInputElement>(`select[name="team"] option[value="${teamID}"]`, frm);
    if (el) {
      const t = el.innerText;
      ret += ` in <a href="/team/${teamID}">${t}</a>`;
    }
  }
  return ret;
}

export function onChildAddModel(param: ChildAdd) {
  const tbody = req(`#${param.type}-list tbody`);
  const empty = opt(".empty", tbody);
  if (empty) {
    empty.remove();
  }

  const tr = document.createElement("tr");
  tr.id = `${ param.type }-list-${ param.id }`;
  const td = document.createElement("td");
  const a = document.createElement("a");
  a.href = param.path;
  a.innerText = param.title;
  td.appendChild(a);
  tr.appendChild(td);
  tbody.appendChild(tr);
}

export function onChildRemoveModel(param: ChildRemove) {
  req(`#${param.type}-list-${param.id}`).remove();
  const tbody = req(`#${param.type}-list tbody`);
  if (tbody.children.length === 0) {
    const empty = document.createElement("tr");
    empty.classList.add("empty");
    const em = document.createElement("em");
    em.innerText = "no " + param.type + "s";
    empty.appendChild(em);
    tbody.appendChild(empty);
  }
}
