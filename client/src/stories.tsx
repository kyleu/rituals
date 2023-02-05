import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import {Story} from "./story";
import {svg} from "./util";

export function snippetStory(s: Story): HTMLElement {
  const id = "comment-link-" + s.estimateID;
  const href = "#modal-" + s.estimateID + "-comments";
  return <tr class="story-row" id="story-row-{%s s.ID.String() %}" data-idx="s.Idx">
    <td><a href={"#modal-story-" + s.id } class="story-title">{ s.title }</a></td>
    <td>{ s.status }</td>
    <td>{ s.finalVote ? s.finalVote : "-" }</td>
    <td><a id={ id } href={ href } title="0 comments" dangerouslySetInnerHTML={svg("comment-alt")}></a></td>
  </tr>;
}

export function snippetStoryModal(s: Story): HTMLElement {
  return <div id={ "modal-story-" + s.id } class="modal modal-story" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>{ s.title }</h2>
      </div>
      <div class="modal-body"></div>
    </div>
  </div>
}
