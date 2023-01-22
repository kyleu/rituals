import {els, req} from "./dom";
import {send} from "./app";

export class Comment {
  readonly debug: boolean;

  constructor(debug: boolean) {
    this.debug = debug
  }
}

export function initComments() {
  const modals = els(".modal.comments");
  for (const modal of modals) {
    const form = req<HTMLFormElement>("form", modal);
    form.onsubmit = function() {
      const inputs = els<HTMLInputElement>("input", form);
      const m: { [key: string]: string; } = {};
      for (const input of inputs) {
        switch(input.name) {
          case "svc":
            m["svc"] = input.value;
            break;
          case "modelID":
            m["modelID"] = input.value;
            break;
        }
      }
      const ta = req<HTMLTextAreaElement>("textarea", form);
      m["content"] = ta.value;
      send("comment", m);
      return false;
    }
  }
}
