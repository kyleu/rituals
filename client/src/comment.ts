import {els, req} from "./dom";
import {send} from "./app";
import {getSelfID, username} from "./member";
import {snippetComment} from "./comments";

export class Comment {
  svc?: string;
  modelID?: string;
  content?: string;
  userID?: string;
}

export function initComments() {
  const modals = els(".modal.comments");
  for (const modal of modals) {
    initCommentsModal(modal);
  }
}

export function initCommentsModal(modal: HTMLElement) {
  const form = req<HTMLFormElement>("form", modal);
  form.onsubmit = function () {
    const c: Comment = {
      svc: req<HTMLInputElement>("input[name=\"svc\"]", form).value,
      modelID: req<HTMLInputElement>("input[name=\"modelID\"]", form).value
    };
    const ta = req<HTMLTextAreaElement>("textarea", form);
    c.content = ta.value;
    c.userID = getSelfID();
    send("comment", c);
    commentAdd(c);
    return false;
  }
}

export function commentAdd(c: Comment) {
  const ul = req("#comment-list-" + c.svc + "-" + c.modelID);
  const un = username(c.userID);
  const li = snippetComment(c, un);
  ul.appendChild(li);

  const count = ul.childNodes.length - 1;
  const link = req("#comment-link-" + c.svc + "-" + c.modelID);
  link.title = count + (count == 1 ? " comment" : " comments");
  if (link.innerHTML.indexOf("comment-dots") == -1) {
    link.innerHTML = `<svg style="width: 18px; height: 18px;" class="right"><use xlink:href="#svg-comment-dots"></use></svg>`;
  }
}
