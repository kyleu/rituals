import {els, req} from "./dom";
import {send} from "./app";
import {snippetFeedback} from "./feedbacks";

export type Feedback = {
  category: string;
  idx: number;
  userID: string;
  content: string;
  html?: string
}

export function initFeedbacks() {
  els<HTMLAnchorElement>(".add-feedback-link").forEach((x) => x.onclick = function() {
    const category = x.dataset["category"];
    setTimeout(() => req<HTMLInputElement>("#feedback-add-content-" + category).focus(), 100);
    return true;
  });
  for(const feedbackAddModal of els(".modal-feedback")) {
    const feedbackAddForm = req("form", feedbackAddModal);
    feedbackAddForm.onsubmit = function () {
      const category = req<HTMLInputElement>("select[name=\"category\"]", feedbackAddForm).value;
      const input = req<HTMLInputElement>("textarea[name=\"content\"]", feedbackAddForm);
      const content = input.value;
      input.value = "";
      send("child-add", {"category": category, "content": content});
      document.location.hash = "";
      return false;
    }
  }
}

export function feedbackAdd(f: Feedback) {
  const list = req("#category-" + f.category + " .feedback-list");
  let idx = -1;
  for (let i = 0; i < list.children.length; i++) {
    const n = list.children.item(i) as HTMLElement;
    const title = req(".feedback-content", n).innerText;
    const currIdxStr = n.dataset["index"];
    if (currIdxStr) {
      const currIdx = parseInt(currIdxStr, 10);
      if (currIdx >= f.idx) {
        idx = i;
        break;
      } else {
        if (title.localeCompare(f.content, undefined, { sensitivity: 'accent' }) >= 0) {
          idx = i;
          break;
        }
      }
    }
  }
  const tr = snippetFeedback(f);
  if (idx == -1) {
    list.appendChild(tr);
  } else {
    list.insertBefore(tr, list.children[idx]);
  }
}
