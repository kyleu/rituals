import {els, req} from "./dom";
import {send} from "./app";
import {focusDelay} from "./util";

export function wireStoryModal(el: HTMLElement) {
  const id = el.dataset["id"];
  if (!id) {
    console.warn("no id in dataset", el);
    return;
  }
  els(".vote-submit-button", el).forEach((e) => e.style.display = "none");
  els<HTMLAnchorElement>(".link-edit", el).forEach((e) => wireStoryModalFormEdit(id, e));
  els<HTMLFormElement>(".status-new-form-delete", el).forEach((e) => wireStoryModalFormDelete(id, e));
  els<HTMLFormElement>(".status-new-form-next", el).forEach((e) => wireStoryModalFormToActive(id, e));
  els<HTMLFormElement>(".status-active-form-prev", el).forEach((e) => wireStoryModalFormToNew(id, e));
  els<HTMLFormElement>(".status-active-form-next", el).forEach((e) => wireStoryModalFormToComplete(id, e));
  els(".story-vote-options", el).forEach((e) => wireStoryModalFormVote(id, e));
  els<HTMLFormElement>(".status-complete-form-prev", el).forEach((e) => wireStoryModalFormToActive(id, e));
}

function wireStoryModalFormEdit(id: string, a: HTMLAnchorElement) {
  a.onclick = function () {
    focusDelay(req("#modal-story-" + id + "-edit form input[name=\"title\"]"));
  }
}

function wireStoryModalFormToNew(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function () {
    send("child-status", {"storyID": id, "status": "new"});
    console.log("new");
    return false;
  }
}

function wireStoryModalFormToActive(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function () {
    send("child-status", {"storyID": id, "status": "active"});
    console.log("active");
    return false;
  }
}

function wireStoryModalFormToComplete(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function () {
    send("child-status", {"storyID": id, "status": "complete"});
    console.log("complete");
    return false;
  }
}

function wireStoryModalFormVote(id: string, e: HTMLElement) {
  els(".vote-option", e).forEach((opt) => {
    opt.onclick = function (evt) {
      req<HTMLInputElement>("input[name=\"vote\"]", opt).checked = true;
      send("vote", {"storyID": id, "vote": opt.dataset["choice"]});
      return false;
    };
  });
}

function wireStoryModalFormDelete(id: string, frm: HTMLFormElement) {
  frm.onsubmit = function () {
    if (!confirm("Are you sure you want to delete this story?")) {
      return false;
    }
    send("child-remove", {"storyID": id});
    document.location.hash = "";
    return false;
  }
}
