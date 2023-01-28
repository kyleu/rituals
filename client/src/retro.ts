import {Message} from "./socket";

export function initRetro() {
  console.log("retro!");
}

export function handleRetro(m: Message) {
  switch (m.cmd) {
    default:
      throw "invalid retro command [" + m.cmd + "]"
  }
}
