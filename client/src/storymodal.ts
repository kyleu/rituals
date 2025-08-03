import { els, req } from "./dom";
import { send } from "./app";
import { focusDelay, svgRef } from "./util";
import { getSelfID } from "./member";

function wireStoryModalFormEdit(id: string, a: HTMLAnchorElement) {
  a.onclick = () => {
    focusDelay(req("#modal-story-" + id + '-edit form input[name="title"]'));
  };
}

function wireStoryModalFormToNew(id: string, frm: HTMLFormElement) {
  frm.onsubmit = () => {
    send("child-status", { storyID: id, status: "new" });
    return false;
  };
}

function wireStoryModalFormToActive(id: string, frm: HTMLFormElement) {
  frm.onsubmit = () => {
    send("child-status", { storyID: id, status: "active" });
    return false;
  };
}

function wireStoryModalFormToComplete(id: string, frm: HTMLFormElement) {
  frm.onsubmit = () => {
    send("child-status", { storyID: id, status: "complete" });
    return false;
  };
}

function wireStoryModalFormVote(id: string, e: HTMLElement) {
  els(".vote-option", e).forEach((opt) => {
    opt.onclick = () => {
      req<HTMLInputElement>('input[name="vote"]', opt).checked = true;
      send("vote", { storyID: id, vote: opt.dataset.choice });
      els("#modal-story-" + id + " .story-members .member").forEach((m) => {
        if (m.dataset.member === getSelfID()) {
          req(".choice", m).innerHTML = svgRef("check", 18, "");
        }
      });
      return false;
    };
  });
}

export function wireStoryModalFormDelete(id: string, frm: HTMLFormElement) {
  frm.onsubmit = () => {
    if (!confirm("Are you sure you want to delete this story?")) {
      return false;
    }
    send("child-remove", { storyID: id });
    document.location.hash = "";
    return false;
  };
}

export function wireStoryModal(el: HTMLElement) {
  const id = el.dataset.id;
  if (!id) {
    console.warn("no id in dataset", el);
    return;
  }
  els(".vote-submit-button", el).forEach((e) => {
    e.style.display = "none";
  });
  els<HTMLAnchorElement>(".link-edit", el).forEach((e) => wireStoryModalFormEdit(id, e));
  els<HTMLFormElement>(".status-new-form-delete", el).forEach((e) => wireStoryModalFormDelete(id, e));
  els<HTMLFormElement>(".status-new-form-next", el).forEach((e) => wireStoryModalFormToActive(id, e));
  els<HTMLFormElement>(".status-active-form-prev", el).forEach((e) => wireStoryModalFormToNew(id, e));
  els<HTMLFormElement>(".status-active-form-next", el).forEach((e) => wireStoryModalFormToComplete(id, e));
  els(".story-vote-options", el).forEach((e) => wireStoryModalFormVote(id, e));
  els<HTMLFormElement>(".status-complete-form-prev", el).forEach((e) => wireStoryModalFormToActive(id, e));
}
