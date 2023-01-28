import {initStories, storyAdd} from "./story";
import {Message} from "./socket";

export function initEstimate() {
  initStories();
}

export function handleEstimate(m: Message) {
  switch (m.cmd) {
    case "story-add":
      return storyAdd(m.param);
    default:
      throw "invalid estimate command [" + m.cmd + "]"
  }
}
