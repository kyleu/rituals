import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import type {Vote} from "./vote";
import type {Member} from "./member";

export function snippetVote(m: Member, v: Vote | undefined): HTMLElement {
  return <div class="vote-result">
    <div class="number" title={ v ? "" : "user did not vote" }>{ v ? v.choice : "-" }</div>
    <div class={"member-" + m.id + "-name"}>{ m.name }</div>
  </div>;
}
