import { JSX } from "./jsx";
import { snippetCommentsModal, snippetCommentsModalLink } from "./comments";
import type { Story } from "./story";
import { svg } from "./util";

export function snippetStory(s: Story, memberName: string): HTMLElement {
  return (
    <tr class="story-row" id={"story-row-" + s.id} data-idx="s.Idx">
      <td>
        <a href={"#modal-story-" + s.id}>
          <div class="story-title">{s.title}</div>
        </a>
      </td>
      <td class="story-author nowrap">
        <a href={"#modal-member-" + s.userID}>
          <em class={"member-" + s.userID + "-name"}>{memberName}</em>
        </a>
      </td>
      <td class="story-status">{s.status}</td>
      <td class="story-final-vote">{!Object.hasOwn(s, "finalVote") || s.finalVote === "" ? "-" : s.finalVote}</td>
      <td>
        {snippetCommentsModalLink("story", s.id)}
        {snippetCommentsModal("story", s.id, s.title)}
      </td>
    </tr>
  );
}

export function memberItem(id: string, name: string) {
  return (
    <div class="member" data-member={id}>
      <div class="choice" dangerouslySetInnerHTML={svg("minus", 18, "")}></div>
      <div class={"name member-" + id + "-name"}>{name}</div>
    </div>
  );
}

export function choiceItem(ch: string) {
  return (
    <div class="vote-option" data-choice={ch}>
      <label>
        <input type="radio" name="vote" value="{%s c %}" />
        <div class="vote-choice">{ch}</div>
      </label>
    </div>
  );
}
