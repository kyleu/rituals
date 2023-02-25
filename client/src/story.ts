import {els, req} from "./dom";
import {send} from "./app";
import {snippetStory} from "./stories";
import {wireStoryModal} from "./storymodal";
import {initCommentsModal} from "./comment";
import {flashCreate} from "./flash";
import {focusDelay} from "./util";

export type Story = {
  id: string;
  idx: number;
  estimateID: string;
  title: string;
  status: string;
  finalVote: string;
  userID: string;
  updated: string;
  created: string;
}

export function initStories() {
  els<HTMLAnchorElement>(".add-story-link").forEach((x) => {
    x.onclick = () => {
      return focusDelay(req<HTMLInputElement>("#story-add-title"));
    };
  });

  const storyAddModal = req("#modal-story--add");
  const storyAddForm = req("form", storyAddModal);
  storyAddForm.onsubmit = () => {
    const input = req<HTMLInputElement>("input[name=\"title\"]", storyAddForm);
    const title = input.value;
    input.value = "";
    send("child-add", {"title": title});
    document.location.hash = "";
    return false;
  };
  els("#story-modals .modal-story").forEach(wireStoryModal);
}

export function storyAdd(s: Story) {
  const tbl = req("#panel-detail table tbody");
  let idx = -1;
  for (let i = 0; i < tbl.children.length; i++) {
    const n = tbl.children.item(i) as HTMLElement;
    const title = req(".story-title", n).innerText;
    const currIdxStr = n.dataset.index;
    if (currIdxStr) {
      const currIdx = parseInt(currIdxStr, 10);
      if (currIdx >= s.idx) {
        idx = i;
        break;
      } else if (title.localeCompare(s.title, undefined, {sensitivity: "accent"}) >= 0) {
        idx = i;
        break;
      }
    }
  }
  const tr = snippetStory(s);
  if (idx === -1) {
    tbl.appendChild(tr);
  } else {
    tbl.insertBefore(tr, tbl.children[idx]);
  }

  initCommentsModal(req(".modal", tr));

  const prototype = req("#modal-story-new");

  const modal = prototype.cloneNode(true) as HTMLDivElement;
  modal.id = "modal-story-" + s.id;
  modal.dataset.id = s.id;
  modal.dataset.status = s.status;
  modal.classList.add("modal-story");
  req("#story-modals").appendChild(modal);
  if (document.location.hash === "modal-story--add" || document.location.hash === "") {
    document.location.hash = "modal-story-" + s.id;
  }
}

export function storyStatus(s: Story) {
  const modal = req("#modal-story-" + s.id);
  req(".status-new", modal).style.display = s.status === "new" ? "block" : "none";
  req(".status-active", modal).style.display = s.status === "active" ? "block" : "none";
  req(".status-complete", modal).style.display = s.status === "complete" ? "block" : "none";
}

export function storyRemove(id: string) {
  const tr = req("#story-row-" + id);
  const title = req(".story-title", tr).innerText;
  flashCreate(id + "-removed", "success", `story [${title}] has been removed`);
  tr.remove();
  req("#modal-story-" + id).remove();
}

