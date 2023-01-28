import {Message} from "./socket";

export function initTeam() {
  console.log("team!");
}

export function handleTeam(m: Message) {
  switch (m.cmd) {
    default:
      throw "invalid team command [" + m.cmd + "]"
  }
}
