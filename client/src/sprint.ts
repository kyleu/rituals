import {Message} from "./socket";

export function initSprint() {
  console.log("sprint!");
}

export function handleSprint(m: Message) {
  switch (m.cmd) {
    default:
      throw "invalid sprint command [" + m.cmd + "]"
  }
}
