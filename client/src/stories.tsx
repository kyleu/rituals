import {JSX} from "./jsx"; // eslint-disable-line @typescript-eslint/no-unused-vars
import {snippetCommentsModal, snippetCommentsModalLink} from "./comments";
import type {Story} from "./story";
import type { Member } from "./member";
import {svg} from "./util";

export function snippetStory(s: Story, memberName: String): HTMLElement {
  return <tr class="story-row" id={"story-row-" + s.id} data-idx="s.Idx">
    <td><a href={"#modal-story-" + s.id }><div class="story-title">{ s.title }</div></a></td>
    <td class="story-author nowrap"><a href={ "#modal-member-" + s.userID}><em class={ "member-" + s.userID + "-name"}>{ memberName }</em></a></td>
    <td class="story-status">{ s.status }</td>
    <td class="story-final-vote">{ s.finalVote === "" ? "-" : s.finalVote }</td>
    <td>
      { snippetCommentsModalLink("story", s.id)}
      { snippetCommentsModal("story", s.id, s.title)}
    </td>
  </tr>;
}

export function memberItem(m: Member) {
  return <div class="member" data-member={ m.id }>
    <div class="choice" dangerouslySetInnerHTML={ svg("minus", 18, "") }></div>
    <div class="name member-{%s m.UserID.String() %}-name">{ m.name }</div>
  </div>;
}

export function choiceItem(ch: string) {
  return <div class="vote-option" data-choice={ ch }>
    <label>
      <input type="radio" name="vote" value="{%s c %}" />
      <div class="vote-choice">{ ch }</div>
    </label>
  </div>;
}

