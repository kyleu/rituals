import {els, opt, req} from "./dom";
import {send} from "./app";
import {memberPictureFor, snippetMember, snippetMemberModalEdit, snippetMemberModalView} from "./members.jsx";
import {svgRef} from "./util";

export function username(id?: string) {
  if (!id) {
    return "System";
  }
  const ret = getMemberName(id);
  if (ret) {
    return ret;
  }
  return "Unknown User";
}

export function getSelfID() {
  return req("#self-id").innerText;
}

function isAdmin() {
  return req("#self-role").innerText === "owner";
}

function wireSelfForm() {
  const selfModal = req("#modal-member-" + getSelfID());
  const selfForm = opt<HTMLFormElement>("form", selfModal);
  if (selfForm) {
    selfForm.onsubmit = () => {
      const nameInput = req<HTMLInputElement>("input[name=\"name\"]", selfForm);
      const choiceInput = req<HTMLInputElement>("input[name=\"choice\"]:checked", selfForm);
      const pictureInput = opt<HTMLInputElement>("input[name=\"picture\"]:checked", selfForm);

      const msg: { name: string, choice: string, picture?: string } = {"name": nameInput.value, "choice": choiceInput.value};
      if (pictureInput) {
        msg.picture = pictureInput.value;
      }
      send("self", msg);

      req("#self-name").innerText = nameInput.value;
      req("#self-picture").innerHTML = memberPictureFor(pictureInput ? pictureInput.value : "", 20, "icon");
      document.location.hash = "";
      return false;
    };
  }
}

export function getMemberName(id: string) {
  if (id === getSelfID()) {
    return req("#self-name").innerText;
  }
  return req("#member-" + id + " .member-name").innerText;
}

function wireMemberForm(modal: HTMLElement) {
  const form = opt<HTMLFormElement>("form", modal);
  if (!form) {
    return;
  }
  const f = (cmd: string) => {
    const userID = req<HTMLInputElement>("input[name=\"userID\"]", form).value;
    const role = req<HTMLSelectElement>("select[name=\"role\"]", form).value;
    send(cmd, {"userID": userID, "role": role});
    const panel = req("#member-" + userID);
    if (cmd === "member-update") {
      req(".member-role", panel).innerText = role;
    } else if (cmd === "member-remove") {
      panel.remove();
    }
    document.location.hash = "";
    return false;
  };
  req<HTMLButtonElement>(".member-update", form).onclick = () => f("member-update");
  req<HTMLButtonElement>(".member-remove", form).onclick = () => {
    if (confirm("Are you sure you wish to remove this user?")) {
      return f("member-remove");
    }
    return false;
  };
}

function wireMemberForms() {
  const modals = els(".modal-member");
  for (const modal of modals) {
    wireMemberForm(modal);
  }
}

export type MemberMessage = {
  userID: string;
  name: string;
  role: string;
  picture?: string;
}

export function memberUpdate(param: MemberMessage) {
  if (param.userID === getSelfID()) {
    req("#self-name").innerText = param.name;
    req("#self-role").innerText = param.role;
    req("#self-picture").innerHTML = memberPictureFor(param.picture ? param.picture : "", 20, "icon");
  } else {
    const panel = req("#member-" + param.userID);
    req(".member-name", panel).innerText = param.name;
    req(".member-role", panel).innerText = param.role;
    req(".member-picture", panel).innerHTML = memberPictureFor(param.picture ? param.picture : "", 18, "");

    const modal = req("#modal-member-" + param.userID);
    req(".member-name", modal).innerText = param.name;
    const rd = opt(".member-role", modal);
    if (rd) {
      rd.innerText = param.role;
    }
    const rs = opt<HTMLSelectElement>("select[name=\"role\"]", modal);
    if (rs) {
      rs.value = param.role;
    }
    req(".member-picture", modal).innerHTML = memberPictureFor(param.picture ? param.picture : "", 18, "");
  }
  if (getMemberName(param.userID) !== param.name) {
    const tbl = req("#panel-members table tbody");
    const items = tbl.children;
    const itemsArr: Element[] = [...items];
    itemsArr.sort((l, r) => {
      const ln = req(".member-name", l).innerText;
      const rn = req(".member-name", r).innerText;
      return ln.localeCompare(rn, undefined, {sensitivity: "accent"});
    });
    tbl.replaceChildren(...itemsArr);
  }
}

export function memberAdd(param: MemberMessage) {
  const panel = opt("#member-" + param.userID);
  if (panel || param.userID === getSelfID()) {
    return memberUpdate(param);
  }

  const tbl = req("#panel-members table tbody");
  let idx = -1;
  for (let i = 0; i < tbl.children.length; i++) {
    const n = tbl.children.item(i);
    const nm = req(".member-name", n as HTMLElement).innerText;
    if (nm.localeCompare(param.name, undefined, {sensitivity: "accent"}) > 0) {
      idx = i;
      break;
    }
  }
  const tr = snippetMember(param.userID, param.name, param.role, param.picture ? param.picture : "");
  if (idx === -1) {
    tbl.appendChild(tr);
  } else {
    tbl.insertBefore(tr, tbl.children[idx]);
  }

  const modals = req("#member-modals");
  let modal: HTMLElement;
  if (isAdmin()) {
    modal = snippetMemberModalEdit(param.userID, param.name, param.role, param.picture ? param.picture : "");
  } else {
    modal = snippetMemberModalView(param.userID, param.name, param.role, param.picture ? param.picture : "");
  }
  modals.appendChild(modal);
  req(".member-picture", modal).innerHTML = memberPictureFor(param.picture ? param.picture : "", 24, "icon");
  wireMemberForm(modal);
}

export function memberRemove(userID: string) {
  const panel = req("#member-" + userID);
  panel.remove();
}

export function onlineUpdate(param: { userID: string; connected: boolean; }) {
  if (param.userID === getSelfID()) {
    return;
  }
  const mel = opt("#member-" + param.userID + " .online-status");
  if (!mel) {
    throw new Error("missing panel #member-" + param.userID);
  }
  mel.title = param.connected ? "online" : "offline";
  const svg = param.connected ? "check-circle" : "circle";
  mel.innerHTML = svgRef(svg, 18, "right");
}

export function initMembers() {
  wireSelfForm();
  wireMemberForms();
}
