import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import type {Story} from "./story";
import {snippetCommentsModal, snippetCommentsModalLink} from "./comments";

export function snippetStory(s: Story): HTMLElement {
  return <tr class="story-row" id={"story-row-"+ s.id} data-idx="s.Idx">
    <td><a href={"#modal-story-" + s.id }><div class="story-title">{ s.title }</div></a></td>
    <td class="story-status">{ s.status }</td>
    <td class="story-final-vote">{ s.finalVote ? s.finalVote : "-" }</td>
    <td>
      { snippetCommentsModalLink("story", s.id)}
      { snippetCommentsModal("story", s.id, s.title)}
    </td>
  </tr>;
}
