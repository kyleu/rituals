import {els, opt, req} from "./dom";
import {send} from "./app";
import {snippetMember, snippetMemberModal} from "./members.jsx";
import {svgRef} from "./util";

let selfID: string;
let names: { [key: string]: string; };

export function username(id?: string) {
  if (!id) {
    return "System";
  }
  const ret = names[id];
  if (ret) {
    return ret;
  }
  return "Unknown User";
}

export function getSelfID() {
  return selfID;
}

export function initMembers() {
  wireSelfForm();
  wireMemberForms();
  refreshMembers();
}

function wireSelfForm() {
  const selfModal = req("#modal-self");
  const selfForm = req<HTMLFormElement>("form", selfModal);
  selfForm.onsubmit = function () {
    const nameInput = req<HTMLInputElement>("input[name=\"name\"]", selfForm);
    const choiceInput = req<HTMLInputElement>("input[name=\"choice\"]:checked", selfForm);
    const pictureInput = opt<HTMLInputElement>("input[name=\"picture\"]:checked", selfForm);

    const msg: { name: string, choice: string, picture?: string } = {"name": nameInput.value, "choice": choiceInput.value}
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

function wireMemberForms() {
  const modals = els(".modal-member");
  for (const modal of modals) {
    wireMemberForm(modal);
  }
}

function wireMemberForm(modal: HTMLElement) {
  const form = req<HTMLFormElement>("form", modal);
  const f = function (cmd: string) {
    const userID = req<HTMLInputElement>("input[name=\"userID\"]", form).value;
    const role = req<HTMLSelectElement>("select[name=\"role\"]", form).value;
    send(cmd, {"userID": userID, "role": role});
    const panel = req("#member-" + userID);
    if (cmd === "member-update") {
      req(".member-role", panel).innerText = role;
    } else if (cmd === "member-remove") {
      panel.remove();
      refreshMembers();
    }
    document.location.hash = "";
    return false;
  };
  req<HTMLButtonElement>(".member-update", form).onclick = () => f("member-update");
  req<HTMLButtonElement>(".member-remove", form).onclick = () => {
    if (confirm('Are you sure you wish to remove this user?')) {
      return f("member-remove");
    }
    return false;
  }
}

export function refreshMembers() {
  names = {};
  selfID = req("#self-id").innerText;
  names[selfID] = req("#self-name").innerText;

  const panel = req("#panel-members");
  const members = els(".member", panel);
  for (const m of members) {
    const id = m.dataset["id"];
    if (id) {
      names[id] = req(".member-name", m).innerText;
    }
  }
}

export type MemberMessage = {
  userID: string;
  name: string;
  role: string;
  picture?: string;
}

export function memberAdd(param: MemberMessage) {
  const panel = opt("#member-" + param.userID);
  if (panel || param.userID === selfID) {
    return memberUpdate(param);
  }

  const tbl = req("#panel-members table tbody");
  let idx = -1;
  for (let i = 0; i < tbl.children.length; i++) {
    const n = tbl.children.item(i);
    const nm = req(".member-name", n as HTMLElement).innerText;
    if (nm.localeCompare(param.name, undefined, {sensitivity: 'accent'}) > 0) {
      idx = i;
      break;
    }
  }
  const tr = snippetMember(param.userID, param.name, param.role, param.picture ? param.picture : "");
  if (idx == -1) {
    tbl.appendChild(tr);
  } else {
    tbl.insertBefore(tr, tbl.children[idx]);
  }

  const modals = req("#member-modals");
  const modal = snippetMemberModal(param.userID, param.name, param.role, param.picture ? param.picture: "");
  modals.appendChild(modal);
  wireMemberForm(modal);

  names[param.userID] = param.name;
}

export function memberUpdate(param: MemberMessage) {
  if (param.userID === selfID) {
    req("#self-name").innerText = param.name;
    req("#self-role").innerText = param.role;
    req("#self-picture").innerHTML = memberPictureFor(param.picture ? param.picture: "", 20, "icon");
  } else {
    const panel = req("#member-" + param.userID);
    req(".member-name", panel).innerText = param.name;
    req(".member-role", panel).innerText = param.role;
    req(".member-picture").innerHTML = memberPictureFor(param.picture ? param.picture: "", 18, "");

    const modal = req("#modal-member-" + param.userID);
    req<HTMLSelectElement>("select[name=\"role\"]", modal).value = param.role;
  }
  if (names[param.userID] !== param.name) {
    names[param.userID] = param.name;

    const tbl = req("#panel-members table tbody");
    const items = tbl.children;
    const itemsArr: Element[] = [];
    for (const i in items) {
      if (items[i].nodeType == 1) { // get rid of the whitespace text nodes
        itemsArr.push(items[i]);
      }
    }
    itemsArr.sort((l, r) => {
      const ln = req(".member-name", l).innerText;
      const rn = req(".member-name", r).innerText;
      return ln.localeCompare(rn, undefined, {sensitivity: 'accent'});
    });
    tbl.replaceChildren(...itemsArr);
  }
}

export function memberRemove(userID: string) {
  const panel = req("#member-" + userID);
  panel.remove();
  refreshMembers();
}

export function onlineUpdate(param: { userID: string; connected: boolean; }) {
  if (param.userID === selfID) {
    return;
  }
  const mel = opt("#member-" + param.userID + " .online-status");
  if (!mel) {
    throw "missing panel #member-" + param.userID;
  }
  mel.title = param.connected ? "online" : "offline";
  const svg = param.connected ? "check-circle" : "circle";
  mel.innerHTML = svgRef(svg, 18, "right");
}

function memberPictureFor(picture: string, size: number, cls: string) {
  if (picture === "") {
    return svgRef("profile", size, cls);
  }
  return `<img class="${cls}" style="width: ${size}px; height: ${size}px;" src="${picture}" />`;
}
