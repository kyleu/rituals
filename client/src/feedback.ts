import {els, opt, req} from "./dom";
import {send} from "./app";
import {snippetFeedback, snippetFeedbackContainer, snippetFeedbackModalEdit, snippetFeedbackModalView} from "./feedbacks";
import {initCommentsModal} from "./comment";
import {flashCreate} from "./flash";
import {focusDelay} from "./util";
import {getSelfID} from "./member";

export type Feedback = {
  id: string;
  category: string;
  idx: number;
  userID: string;
  content: string;
  html?: string
}

function initAddModal(feedbackAddModal: HTMLElement) {
  const feedbackAddForm = req("form", feedbackAddModal);
  feedbackAddForm.onsubmit = () => {
    const category = req<HTMLSelectElement>("select[name=\"category\"]", feedbackAddForm).value;
    const input = req<HTMLTextAreaElement>("textarea[name=\"content\"]", feedbackAddForm);
    const content = input.value;
    input.value = "";
    send("child-add", {"category": category, "content": content});
    document.location.hash = "";
    return false;
  };
}

function initEditModal(feedbackEditModal: HTMLElement) {
  const frm = req("form", feedbackEditModal);
  const feedbackID = req<HTMLInputElement>("input[name=\"feedbackID\"]", frm).value;
  req<HTMLElement>(".feedback-edit-delete", frm).onclick = () => {
    if (confirm("Are you sure you want to delete this feedback?")) {
      send("child-remove", {"feedbackID": feedbackID});
      document.location.hash = "";
    }
    return false;
  };
  frm.onsubmit = () => {
    const category = req<HTMLSelectElement>("select[name=\"category\"]", frm).value;
    const input = req<HTMLTextAreaElement>("textarea[name=\"content\"]", frm);
    const content = input.value;
    send("child-update", {"feedbackID": feedbackID, "category": category, "content": content});
    document.location.hash = "";
    return false;
  };
}

export function initFeedbacks() {
  els<HTMLAnchorElement>(".add-feedback-link").forEach((x) => {
    x.onclick = () => {
      return focusDelay(req<HTMLInputElement>("#feedback-add-content-" + x.dataset.category));
    };
  });
  els<HTMLAnchorElement>(".modal-feedback-edit-link").forEach((x) => {
    x.onclick = () => {
      return focusDelay(req<HTMLInputElement>("#input-content-" + x.dataset.id));
    };
  });
  for (const feedbackAddModal of els(".modal-feedback-add")) {
    initAddModal(feedbackAddModal);
  }
  els(".modal-feedback-edit").forEach(initEditModal);
}

export function feedbackAdd(f: Feedback) {
  let list = opt("#category-" + f.category + " .feedback-list");

  if (!list) {
    const x = snippetFeedbackContainer(f.category);
    req("#category-list").appendChild(x);
    list = req(".feedback-list", x);
  }

  let idx = -1;
  for (let i = 0; i < list.children.length; i++) {
    const n = list.children.item(i) as HTMLElement;
    const title = req(".feedback-content", n).innerText;
    const currIdxStr = n.dataset.index;
    if (currIdxStr) {
      const currIdx = parseInt(currIdxStr, 10);
      if (currIdx >= f.idx) {
        idx = i;
        break;
      } else if (title.localeCompare(f.content, undefined, {sensitivity: "accent"}) >= 0) {
        idx = i;
        break;
      }
    }
  }
  const tr = snippetFeedback(f);
  if (idx === -1) {
    list.appendChild(tr);
  } else {
    list.insertBefore(tr, list.children[idx]);
  }
  if (getSelfID() === f.userID) {
    const modal = snippetFeedbackModalEdit(f);
    initEditModal(modal);
    req("#feedback-modals").appendChild(modal);
  } else {
    req("#feedback-modals").appendChild(snippetFeedbackModalView(f));
  }

  initCommentsModal(req(".modal", tr));
}

export function feedbackRemove(id: string) {
  req("#feedback-" + id).remove();
  flashCreate(id + "-removed", "success", `feedback has been removed`);
  req("#modal-feedback-" + id).remove();
}
