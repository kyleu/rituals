import {els, opt, req} from "./dom";
import {snippetCommentsModal, snippetCommentsModalLink} from "./comments";
import {initCommentsModal} from "./comment";
import {svgRef} from "./util";

export type ChildAdd = {
  "type": string;
  "id": string;
  "title": string;
  "path": string;
  "icon": string;
}

export type ChildUpdate = {
  "type": string;
  "model": { id: string; title: string; icon: string };
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
  ret += (key === "retro" ? "retrospective" : key);
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
  tr.id = `${param.type}-list-${param.id}`;
  const td = document.createElement("td");

  const commentsDiv = document.createElement("div");
  commentsDiv.classList.add("right");
  commentsDiv.appendChild(snippetCommentsModalLink(param.type, param.id));
  commentsDiv.appendChild(snippetCommentsModal(param.type, param.id, param.id));
  td.appendChild(commentsDiv);

  const a = document.createElement("a");
  a.href = param.path;

  const iconSpan = document.createElement("span");
  iconSpan.classList.add("model-span-icon")
  iconSpan.innerHTML = svgRef(param.icon, 16, "icon");
  a.appendChild(iconSpan);

  const titleSpan = document.createElement("span");
  titleSpan.classList.add("model-span-title")
  titleSpan.innerText = param.title;
  a.appendChild(titleSpan);

  td.appendChild(a);
  tr.appendChild(td);
  tbody.appendChild(tr);
  initCommentsModal(req(".modal", commentsDiv));
}

export function onChildUpdateModel(param: ChildUpdate) {
  const tr = req(`#${param.type}-list-${param.model.id}`);
  req(".model-span-icon", tr).innerHTML = svgRef(param.model.icon);
  req(".model-span-title", tr).innerText = param.model.title;
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

export function setTeamSprint(key: string, frm: HTMLFormElement, teamID: string | null, sprintID: string | null, title: string, icon: string) {
  req<HTMLInputElement>("input[name=\"title\"]", frm).value = title;
  for (const r of els<HTMLInputElement>("input[name=\"icon\"]", frm)) {
    r.checked = icon === r.value;
  }
  const t = opt<HTMLSelectElement>("select[name=\"team\"]", frm)
  if (t !== null && t !== undefined) {
    t.value = teamID ? teamID : "";
  }
  const s = opt<HTMLSelectElement>("select[name=\"sprint\"]", frm)
  if (s !== null && s !== undefined) {
    s.value = sprintID ? sprintID : "";
  }
  req("#model-title").innerText = title;
  req("#model-icon").innerHTML = svgRef(icon, 20);
  req("#model-banner").innerHTML = modelBanner(key, frm, teamID ? teamID : "", sprintID ? sprintID : "");
  // flashCreate(key, "success", key + " updated");
}
