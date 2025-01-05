import {JSX} from "./jsx";
import type {Vote} from "./vote";
import type {Member} from "./member";

export function snippetVote(storyID: string, m: Member, v: Vote | undefined): HTMLElement {
  if (!v) {
    return <div class="vote-result">
      <div class="number" title="user did not vote">-</div>
      <div class={"member-" + m.id + "-name"}>{ m.name }</div>
    </div>;
  }
  return <div class="vote-result">
    <form class="final-vote-form" action="" method="post">
      <input type="hidden" name="storyID" value={ storyID } />
      <input type="hidden" name="action" value="vote"/>
      <input type="hidden" name="typ" value="user"/>
      <input type="hidden" name="value" value={ v.choice } />
      <button type="submit" class="number button-link">{ v.choice }</button>
    </form>
    <div class={"member-" + m.id + "-name"}>{ m.name }</div>
  </div>;
}
