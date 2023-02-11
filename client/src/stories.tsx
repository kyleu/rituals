import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import {Story} from "./story";
import {snippetCommentsModal, snippetCommentsModalLink} from "./comments";

export function snippetStory(s: Story): HTMLElement {
  return <tr class="story-row" id="story-row-{%s s.ID.String() %}" data-idx="s.Idx">
    <td><a href={"#modal-story-" + s.id } class="story-title"><div>{ s.title }</div></a></td>
    <td>{ s.status }</td>
    <td>{ s.finalVote ? s.finalVote : "-" }</td>
    <td>
      { snippetCommentsModalLink("story", s.id)}
      { snippetCommentsModal("story", s.id, s.title)}
    </td>
  </tr>;
}
