import {els, req} from "./dom";
import {send} from "./app";

export type Report = {
  day: string;
  userID: string;
  content: string;
  html?: string
}

export function initReports() {
  els<HTMLAnchorElement>(".add-report-link").forEach((x) => x.onclick = function() {
    setTimeout(() => req<HTMLInputElement>("#report-add-content").focus(), 100);
    return true;
  });

  const reportAddModal = req("#modal-report--add");
  const reportAddForm = req("form", reportAddModal);
  reportAddForm.onsubmit = function () {
    const day = req<HTMLInputElement>("input[name=\"day\"]", reportAddForm).value;
    const input = req<HTMLInputElement>("textarea[name=\"content\"]", reportAddForm);
    const content = input.value;
    input.value = "";
    send("child-add", {"day": day, "content": content});
    document.location.hash = "";
    return false;
  }
}

export function reportAdd(r: Report) {
  console.log("TODO: reportAdd");
}
