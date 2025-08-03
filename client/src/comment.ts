import { els, req } from "./dom";
import { send } from "./app";
import { getSelfID, username } from "./member";
import { snippetComment } from "./comments";
import { focusDelay, svgRef } from "./util";

export class Comment {
  svc?: string;
  modelID?: string;
  content?: string;
  userID?: string;
}

export function initCommentsLink(link: HTMLElement) {
  link.onclick = () => focusDelay(req("#modal-" + link.dataset.key + "-comments form textarea"));
}

export function commentAdd(c: Comment) {
  const ul = req("#comment-list-" + c.svc + "-" + c.modelID);
  const un = username(c.userID);
  const li = snippetComment(c, un);
  ul.appendChild(li);

  const count = ul.children.length;
  const link = req("#comment-link-" + c.svc + "-" + c.modelID);
  link.title = count + (count === 1 ? " comment" : " comments");
  if (link.innerHTML.indexOf("comment-dots") === -1) {
    link.innerHTML = svgRef("comment-dots", 18, "right");
  }
}

export function initCommentsModal(modal: HTMLElement) {
  const form = req<HTMLFormElement>("form", modal);
  form.onsubmit = () => {
    const c: Comment = {
      svc: req<HTMLInputElement>('input[name="svc"]', form).value,
      modelID: req<HTMLInputElement>('input[name="modelID"]', form).value
    };
    const ta = req<HTMLTextAreaElement>("textarea", form);
    c.content = ta.value;
    c.userID = getSelfID();
    send("comment", c);
    commentAdd(c);
    ta.value = "";
    return false;
  };
}

export function initComments() {
  for (const link of els(".comment-link")) {
    initCommentsLink(link);
  }
  for (const modal of els(".modal.comments")) {
    initCommentsModal(modal);
  }
}
