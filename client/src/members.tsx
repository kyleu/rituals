import { JSX } from "./jsx";
import { svg, svgRef } from "./util";

export function memberPictureFor(picture: string, size: number, cls: string) {
  if (!picture) {
    return svgRef("profile", size, cls);
  }
  return `<img class="${cls}" style="width: ${size + "px"}; height: ${size + "px"};" src="${picture}" />`;
}

export function snippetMember(userID: string, name: string, role: string, picture: string): HTMLElement {
  return (
    <tr id={"member-" + userID} class="member" data-id={userID}>
      <td>
        <a href={"#modal-member-" + userID}>
          <div
            class="left prs member-picture"
            dangerouslySetInnerHTML={{ __html: memberPictureFor(picture, 18, "") }}
          ></div>
          <span class={"member-name member-" + userID + "-name"}>{name}</span>
        </a>
      </td>
      <td class="shrink text-align-right">
        <em class="member-role">{role}</em>
      </td>
      <td class="shrink online-status" title="offline" dangerouslySetInnerHTML={svg("circle", 18, "right")}></td>
    </tr>
  );
}

export function snippetMemberModalView(userID: string, name: string, role: string, picture: string): HTMLElement {
  return (
    <div id={"modal-member-" + userID} data-id={userID} class="modal modal-member" style="display: none;">
      <a class="backdrop" href="#"></a>
      <div class="modal-content">
        <div class="modal-header">
          <a href="#" class="modal-close">
            ×
          </a>
          <h2>
            <span class="member-picture" dangerouslySetInnerHTML={{ __html: memberPictureFor(picture, 18, "") }}></span>
            <span class={"member-name member-" + userID + "-name"}>{name}</span>
          </h2>
        </div>
        <div class="modal-body">
          <em>Role</em>
          <br />
          <span class="member-role">{role}</span>
        </div>
      </div>
    </div>
  );
}

export function snippetMemberModalEdit(userID: string, name: string, role: string, picture: string): HTMLElement {
  const cfrm = "return confirm('Are you sure you wish to remove this user?');";
  const roles = [
    ["owner", "Owner"],
    ["member", "Member"],
    ["observer", "Observer"]
  ];
  return (
    <div id={"modal-member-" + userID} data-id={userID} class="modal modal-member" style="display: none;">
      <a class="backdrop" href="#"></a>
      <div class="modal-content">
        <div class="modal-header">
          <a href="#" class="modal-close">
            ×
          </a>
          <h2>
            <span class="member-picture" dangerouslySetInnerHTML={{ __html: memberPictureFor(picture, 18, "") }}></span>
            <span class={"member-name member-" + userID + "-name"}>{name}</span>
          </h2>
        </div>
        <div class="modal-body">
          <form action={document.location.pathname} method="post" class="expanded">
            <input type="hidden" name="userID" value={userID} />
            <em>Role</em>
            <br />
            <select class="input-role" name="role">
              {roles.map((x) => {
                if (x[0] === role) {
                  return (
                    <option selected="selected" value={x[0]}>
                      {x[1]}
                    </option>
                  );
                }
                return <option value={x[0]}>{x[1]}</option>;
              })}
            </select>
            <hr />
            <div class="right">
              <button class="member-update" type="submit" name="action" value="member-update">
                Save
              </button>
            </div>
            <button type="submit" class="member-remove" name="action" value="member-remove" onClick={cfrm}>
              Remove
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}
