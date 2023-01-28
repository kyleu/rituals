import {req} from "./dom";
import {send} from "./app";

export function initEstimate() {
  const storyAddModal = req("#modal-story--add");
  const storyAddForm = req("form", storyAddModal);
  storyAddForm.onsubmit = function () {
    const title = req<HTMLInputElement>("input[name=\"title\"]", storyAddForm).value;
    send("story-add", {"title": title});
    return false;
  }
}

export function storyAdd() {
  console.log("TODO: storyAdd");
}
