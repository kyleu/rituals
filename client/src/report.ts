import {els, opt, req} from "./dom";
import {send} from "./app";
import {snippetReport, snippetReportContainer, snippetReportModalEdit, snippetReportModalView} from "./reports";
import {getSelfID, username} from "./member";
import {initCommentsModal} from "./comment";
import {flashCreate} from "./flash";
import {focusDelay} from "./util";

export type Report = {
  id: string;
  day: string;
  userID: string;
  content: string;
  html?: string
}

function initEditModal(modal: HTMLElement) {
  const frm = req("form", modal);
  const reportID = req<HTMLInputElement>("input[name=\"reportID\"]", frm).value;
  req<HTMLElement>(".report-edit-delete", frm).onclick = () => {
    if (confirm("Are you sure you want to delete this report?")) {
      send("child-remove", {"reportID": reportID});
      document.location.hash = "";
    }
    return false;
  };
  frm.onsubmit = () => {
    const day = req<HTMLInputElement>("input[name=\"day\"]", frm).value;
    const input = req<HTMLTextAreaElement>("textarea[name=\"content\"]", frm);
    const content = input.value;
    send("child-update", {"reportID": reportID, "day": day, "content": content});
    document.location.hash = "";
    return false;
  };
}

export function initReports() {
  els<HTMLAnchorElement>(".add-report-link").forEach((x) => {
    x.onclick = () => {
      return focusDelay(req<HTMLInputElement>("#report-add-content"));
    };
  });
  els<HTMLAnchorElement>(".modal-report-edit-link").forEach((x) => {
    x.onclick = () => {
      return focusDelay(req<HTMLInputElement>("#input-content-" + x.dataset.id));
    };
  });

  const reportAddModal = req("#modal-report--add");
  const reportAddForm = req("form", reportAddModal);
  reportAddForm.onsubmit = () => {
    const day = req<HTMLInputElement>("input[name=\"day\"]", reportAddForm).value;
    const input = req<HTMLTextAreaElement>("textarea[name=\"content\"]", reportAddForm);
    const content = input.value;
    input.value = "";
    send("child-add", {"day": day, "content": content});
    document.location.hash = "";
    return false;
  };

  els(".modal-report-edit").forEach(initEditModal);
}

export function reportAdd(r: Report) {
  if (r.day.length > 10) {
    r.day = r.day.substring(0, 10);
  }
  let list = opt("#report-group-" + r.day + " .bd");

  if (!list) {
    const x = snippetReportContainer(r.day);
    req("#report-groups").appendChild(x);
    list = req(".bd", x);
  }

  let idx = -1;
  const u = username(r.userID);
  for (let i = 0; i < list.children.length; i++) {
    const n = list.children.item(i) as HTMLElement;
    const tgt = req(".username", n).innerText;
    if (tgt.localeCompare(u, undefined, {sensitivity: "accent"}) >= 0) {
      idx = i;
      break;
    }
  }
  const div = snippetReport(r);
  if (idx === -1) {
    list.appendChild(div);
  } else {
    list.insertBefore(div, list.children[idx]);
  }
  if (getSelfID() === r.userID) {
    const modal = snippetReportModalEdit(r);
    initEditModal(modal);
    req("#report-modals").appendChild(modal);
  } else {
    req("#report-modals").appendChild(snippetReportModalView(r));
  }
  initCommentsModal(req(".modal", div));
}

export function reportUpdate(r: Report) {
  if (r.day.length > 10) {
    r.day = r.day.substring(0, 10);
  }
  let div = opt("#report-" + r.id);
  if (!div) {
    return reportAdd(r);
  }
  const body = req(".pt", div);
  if (r.html) {
    body.innerHTML = r.html;
  } else {
    body.innerText = r.content;
  }
}

export function reportRemove(id: string) {
  const rpt = req("#report-" + id);
  const bd = rpt.parentElement!;
  rpt.remove();
  if (bd.children.length === 0) {
    bd.parentElement?.remove();
  }
  flashCreate(id + "-removed", "success", `report has been removed`);
  req("#modal-report-" + id).remove();
}
