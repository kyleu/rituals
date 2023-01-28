import {els, req} from "./dom";
import {send} from "./app";
import {getSelfID, username} from "./members";
import {snippetComment} from "./comment";

export class Comment {
  svc?: string;
  modelID?: string;
  content?: string;
  userID?: string;
}

export function initComments() {
  const modals = els(".modal.comments");
  for (const modal of modals) {
    const form = req<HTMLFormElement>("form", modal);
    form.onsubmit = function () {
      const inputs = els<HTMLInputElement>("input", form);
      const c: Comment = {};
      for (const input of inputs) {
        switch (input.name) {
          case "svc":
            c.svc = input.value;
            break;
          case "modelID":
            c.modelID = input.value;
            break;
        }
      }
      const ta = req<HTMLTextAreaElement>("textarea", form);
      c.content = ta.value;
      c.userID = getSelfID();
      send("comment", c);
      commentAdd(c);
      return false;
    }
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
