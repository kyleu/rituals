import {els, opt, req} from "./dom";
import {snippetCommentsModal, snippetCommentsModalLink} from "./comments";
import {initCommentsModal} from "./comment";
import {focusDelay, svgRef} from "./util";

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

export function getSelectName(key: string, panel: HTMLElement, id: string | null) {
  if (id) {
    const el = opt<HTMLSelectElement>(`select[name="${key}"] option[value="${id}"]`, panel);
    if (el) {
      const s = el.innerText;
      return `<a href="/${key}/${id}">${s}</a> `;
    }
    return `<a href="/${key}/${id}">${id}</a> `;
  }
  return "";
}

export function modelBanner(key: string, panel: HTMLElement, teamID: string, sprintID: string) {
  let ret = "";
  ret += getSelectName("sprint", panel, sprintID);
  ret += key === "retro" ? "retrospective" : key;
  if (teamID) {
    ret += " in " + getSelectName("team", panel, teamID);
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
  iconSpan.classList.add("model-span-icon");
  iconSpan.innerHTML = svgRef(param.icon, 16, "icon");
  a.appendChild(iconSpan);

  const titleSpan = document.createElement("span");
  titleSpan.classList.add("model-span-title");
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

export function setTeamSprint(key: string, panel: HTMLElement, teamID: string | null, sprintID: string | null, title: string, icon: string) {
  const tEl = opt<HTMLInputElement>("input[name=\"title\"]", panel);
  if (tEl) {
    tEl.value = title;
  }
  for (const pt of els(".config-panel-team", panel)) {
    pt.innerHTML = getSelectName("team", panel, teamID);
  }
  for (const ps of els(".config-panel-sprint", panel)) {
    ps.innerHTML = getSelectName("sprint", panel, sprintID);
  }
  for (const vi of els(".view-icon", panel)) {
    vi.innerHTML = svgRef(icon, 24, "icon");
  }
  for (const vt of els(".view-title", panel)) {
    vt.innerText = title;
  }
  for (const r of els<HTMLInputElement>("input[name=\"icon\"]", panel)) {
    r.checked = icon === r.value;
  }
  const t = opt<HTMLSelectElement>("select[name=\"team\"]", panel);
  if (t !== null && t !== undefined) {
    t.value = teamID ? teamID : "";
  }
  const s = opt<HTMLSelectElement>("select[name=\"sprint\"]", panel);
  if (s !== null && s !== undefined) {
    s.value = sprintID ? sprintID : "";
  }
  req("#model-title").innerText = title;
  req("#model-icon").innerHTML = svgRef(icon, 20);
  req("#model-banner").innerHTML = modelBanner(key, panel, teamID ? teamID : "", sprintID ? sprintID : "");
  // flashCreate(key, "success", key + " updated");
}

export function configFocus(k: string) {
  req(`#modal-${k}-config-link`).onclick = () => {
    const i = opt<HTMLInputElement>(`#modal-${k}-config form input[name="title"]`);
    if (i) {
      focusDelay(i);
    }
  };
}
