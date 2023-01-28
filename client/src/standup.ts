import {Message} from "./socket";

export function initStandup() {
  console.log("standup!");
}

export function handleStandup(m: Message) {
  switch (m.cmd) {
    default:
      throw "invalid standup command [" + m.cmd + "]"
  }
}
