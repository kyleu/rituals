import { els, req } from "./dom";
import { send } from "./app";
import { choiceItem, memberItem, snippetStory } from "./stories";
import { wireStoryModal, wireStoryModalFormDelete } from "./storymodal";
import { initCommentsModal } from "./comment";
import { flashCreate } from "./flash";
import { focusDelay } from "./util";
import { applyCalcs, type Vote, type VoteResults } from "./vote";
import { estimateChoices } from "./estimate";
import { memberList, username } from "./member";

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
};

export type StoryStatusResult = {
  story: Story;
  votes?: Vote[];
  results?: VoteResults;
};

function onEditSubmit(frm: HTMLElement) {
  const btn = req(".story-delete-button", frm);
  btn.onclick = () => {
    if (confirm("Are you sure you want to delete this story?")) {
      send("child-remove", { storyID: btn.dataset.id });
      document.location.hash = "";
    }
    return false;
  };
  frm.onsubmit = () => {
    const storyID = req<HTMLInputElement>('input[name="storyID"]', frm).value;
    const title = req<HTMLInputElement>('input[name="title"]', frm).value;
    send("child-update", { storyID: storyID, title: title });
    document.location.hash = "";
    return false;
  };
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
    const input = req<HTMLInputElement>('input[name="title"]', storyAddForm);
    const title = input.value;
    input.value = "";
    send("child-add", { title: title });
    document.location.hash = "";
    return false;
  };
  for (const editor of els(".modal-story-edit")) {
    onEditSubmit(req("form", editor));
  }
  els("#story-modals .modal-story").forEach(wireStoryModal);
}

export function storyUpdate(s: Story) {
  const tr = req("#story-row-" + s.id);
  req(".story-status", tr).innerText = s.status;
  req(".story-title", tr).innerText = s.title;
  req(".story-final-vote", tr).innerText = s.finalVote === "" ? "-" : s.finalVote;

  const editModal = req("#modal-story-" + s.id + "-edit");
  req('form input[name="title"]', editModal).innerText = s.title;

  const viewModal = req("#modal-story-" + s.id);
  const fv = req(".final-vote", viewModal);
  req(".value", fv).innerText = s.finalVote === "" ? "-" : s.finalVote;
  req("#story-row-" + s.id + " .story-final-vote").innerText = s.finalVote === "" ? "-" : s.finalVote;
  if (s.finalVote === "") {
    req(".message", fv).style.display = "block";
    req(".description", fv).style.display = "none";
  } else {
    req(".message", fv).style.display = "none";
    req(".description", fv).style.display = "block";
  }

  els<HTMLFormElement>('input[name="storyID"]', viewModal).forEach((el) => {
    el.value = s.id;
  });

  req("h2.billboard", viewModal).innerText = s.title;
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
      } else if (title.localeCompare(s.title, undefined, { sensitivity: "accent" }) >= 0) {
        idx = i;
        break;
      }
    }
  }
  const tr = snippetStory(s, username(s.userID));
  if (idx === -1) {
    tbl.appendChild(tr);
  } else {
    tbl.insertBefore(tr, tbl.children[idx]);
  }

  initCommentsModal(req(".modal", tr));

  const editPrototype = req(".modal-story-edit-new");
  const editModal = editPrototype.cloneNode(true) as HTMLDivElement;
  editModal.id = "modal-story-" + s.id + "-edit";
  editModal.classList.remove("modal-story-edit-new");
  els<HTMLInputElement>('input[name="storyID"]', editModal).forEach((e) => {
    e.value = s.id;
  });
  req<HTMLInputElement>('input[name="title"]', editModal).value = s.title;
  req<HTMLButtonElement>(".story-delete-button", editModal).dataset.id = s.id;
  onEditSubmit(req("form", editModal));
  els<HTMLFormElement>(".edit-form-delete", editModal).forEach((e) => wireStoryModalFormDelete(s.id, e));
  req("#story-modals").appendChild(editModal);

  const prototype = req("#modal-story-new");
  const modal = prototype.cloneNode(true) as HTMLDivElement;
  modal.id = "modal-story-" + s.id;
  modal.dataset.id = s.id;
  modal.dataset.status = s.status;
  modal.classList.add("modal-story");
  const editLink = req<HTMLAnchorElement>(".link-edit", modal);
  editLink.href = "#modal-story-" + s.id + "-edit";
  editLink.dataset.id = s.id;

  const memberItems = memberList().map((x) => {
    return memberItem(x.id, x.name);
  });
  req(".story-members", modal).replaceChildren(...memberItems);

  const choiceItems = estimateChoices().map((x) => {
    return choiceItem(x);
  });
  req(".story-vote-options", modal).replaceChildren(...choiceItems);

  req("#story-modals").appendChild(modal);
  wireStoryModal(modal);
  storyUpdate(s);
  if (document.location.hash === "modal-story--add" || document.location.hash === "") {
    document.location.hash = "modal-story-" + s.id;
  }
}

export function storyStatus(s: StoryStatusResult) {
  const tr = req("#story-row-" + s.story.id);
  req(".story-status", tr).innerText = s.story.status;
  const modal = req("#modal-story-" + s.story.id);
  req(".status-new", modal).style.display = s.story.status === "new" ? "block" : "none";
  req(".status-active", modal).style.display = s.story.status === "active" ? "block" : "none";
  req(".status-complete", modal).style.display = s.story.status === "complete" ? "block" : "none";
  if (s.story.status === "complete" && s.votes !== undefined && s.results !== undefined) {
    applyCalcs(s.story.id, s.votes, s.results, s.story.finalVote);
  }
}

export function storyRemove(id: string) {
  const tr = req("#story-row-" + id);
  const title = req(".story-title", tr).innerText;
  flashCreate(id + "-removed", "success", `story [${title}] has been removed`);
  tr.remove();
  req("#modal-story-" + id).remove();
}
