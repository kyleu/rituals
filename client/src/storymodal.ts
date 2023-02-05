import {els} from "./dom";

export function wireStoryModal(el: HTMLElement) {
  const id = el.dataset["id"];
  if (!id) {
    console.warn("no id in dataset", el);
    return;
  }
  els(".status-new-form-delete", el).forEach((e) => wireStoryModalFormDelete(id, e as HTMLFormElement));
  els(".status-new-form-next", el).forEach((e) => wireStoryModalFormToActive(id, e as HTMLFormElement));
  els(".status-active-form-prev", el).forEach((e) => wireStoryModalFormToNew(id, e as HTMLFormElement));
  els(".status-active-form-next", el).forEach((e) => wireStoryModalFormToComplete(id, e as HTMLFormElement));
  els(".status-active-form-vote", el).forEach((e) => wireStoryModalFormVote(id, e as HTMLFormElement));
  els(".status-complete-form-prev", el).forEach((e) => wireStoryModalFormToActive(id, e as HTMLFormElement));
}

function wireStoryModalFormToNew(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function() {
    console.log("new");
    return false;
  }
}

function wireStoryModalFormToActive(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function() {
    console.log("active");
    return false;
  }
}

function wireStoryModalFormToComplete(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function() {
    console.log("complete");
    return false;
  }
}

function wireStoryModalFormVote(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function() {
    console.log("vote");
    return false;
  }
}

function wireStoryModalFormDelete(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function() {
    console.log("delete");
    return false;
  }
}
