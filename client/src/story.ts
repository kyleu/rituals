import {req} from "./dom";
import {send} from "./app";
import {snippetStory} from "./stories";

export function initStories() {
  const storyAddModal = req("#modal-story--add");
  const storyAddForm = req("form", storyAddModal);
  storyAddForm.onsubmit = function () {
    const title = req<HTMLInputElement>("input[name=\"title\"]", storyAddForm).value;
    send("story-add", {"title": title});
    return false;
  }
}

export interface Story {
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

export function storyAdd(s: Story) {
  console.log("TODO: storyAdd");
  const tbl = req("#panel-detail table tbody");
  let idx = -1;
  for (let i = 0; i < tbl.children.length; i++) {
    const n = tbl.children.item(i) as HTMLElement;
    const title = req(".story-title", n).innerText;
    const currIdxStr = n.dataset["index"];
    if (currIdxStr) {
      const currIdx = parseInt(currIdxStr, 10);
      if (currIdx >= s.idx) {
        idx = i;
        break;
      }
    }
    if (title.localeCompare(s.title, undefined, { sensitivity: 'accent' }) >= 0) {
      idx = i;
      break;
    }
  }
  const tr = snippetStory(s);
  if (idx == -1) {
    tbl.appendChild(tr);
  } else {
    tbl.insertBefore(tr, tbl.children[idx]);
  }
}
