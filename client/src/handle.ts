import {Message} from "./socket";


export function handle(m: Message) {
  switch (m.cmd) {
    case "comment":
      return handleComment(m.param)
  }
}

function handleComment(param: any) {
  console.log("!");
}
