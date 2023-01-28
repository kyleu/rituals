import {JSX} from "./jsx";
import {Comment} from "./comments";
import {relativeTime, utc} from "./time";

function svg(key: string, cls?: string) {
  return {
    "__html": `<svg class="${cls || ""}" style="width: 18px; height: 18px;"><use href="#svg-${key}"></use></svg>`
  }
}

export function snippetComment(c: Comment, un: string) {
  const li = <li></li>;

  const d = utc(new Date());
  const timeSpan = <span class="nowrap reltime" data-time={d}>just now</span>;
  relativeTime(d, timeSpan);

  const time = <div class="right"></div>;
  time.appendChild(timeSpan);
  li.appendChild(time);

  li.appendChild(<div>{ c.content }</div>);
  li.appendChild(<div><em>{ un }</em></div>);
  return li;
}

export function snippetMember(userID: string, name: string, role: string): HTMLElement {
  return <tr id={ "member-" + userID } class="member" data-id={ userID }>
    <td>
      <a href={ "#modal-member-" + userID }>
        <div class="left" style="padding-right: var(--padding-small);" dangerouslySetInnerHTML={svg("profile")}></div>
        <span class="member-name">{ name }</span>
      </a>
    </td>
    <td class="shrink" style="text-align: right"><em class="member-status">{ role }</em></td>
    <td class="shrink online-status" title="offline" dangerouslySetInnerHTML={svg("circle", "right")}></td>
  </tr>;
}
export function snippetMemberModal(userID: string, name: string, role: string): HTMLElement {
  const roles = [["owner", "Owner"], ["member", "Member"], ["observer", "Observer"]]
  return <div id={ "modal-member-" + userID } data-id={ userID } class="modal modal-member" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>{ name }</h2>
      </div>
      <div class="modal-body">
        <form action={ document.location.pathname } method="post" class="expanded">
          <input type="hidden" name="userID" value={ userID } />
          <em>Role</em><br/>
          <select class="input-role" name="role">
            {roles.map((x) => {
              if (x[0] == role) {
                return <option selected="selected" value={x[0]}>{x[1]}</option>
              } else {
                return <option value={x[0]}>{x[1]}</option>
              }
            })}
          </select>
          <hr/>
          <div class="right">
            <button class="member-update" type="submit" name="action" value="member-update">Save</button>
          </div>
          <button type="submit" class="member-remove" name="action" value="member-remove" onClick="return confirm('Are you sure you wish to remove this user?');">Remove</button>
        </form>
      </div>
    </div>
  </div>
}
