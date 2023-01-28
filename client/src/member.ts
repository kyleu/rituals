import {els, opt, req} from "./dom";
import {send} from "./app";
import {snippetMember, snippetMemberModal} from "./members.jsx";

let selfID: string;
let names: { [key: string]: string; };

export function username(id?: string) {
  if(!id) {
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

export function getSelfName() {
  return username(selfID);
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
    send("self", {"name": nameInput.value, "choice": choiceInput.value});
    req("#self-name").innerText = nameInput.value;
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
    if(id) {
      names[id] = req(".member-name", m).innerText;
    }
  }
}

export function memberAdd(userID: string, name: string, role: string) {
  const panel = opt("#member-" + userID);
  if (panel || userID === selfID) {
    return memberUpdate(userID, name, role);
  }

  const tbl = req("#panel-members table tbody");
  let idx = -1;
  for (let i = 0; i < tbl.children.length; i++) {
    const n = tbl.children.item(i);
    const nm = req(".member-name", n as HTMLElement).innerText;
    if (nm.localeCompare(name, undefined, { sensitivity: 'accent' }) > 0) {
      idx = i;
      break;
    }
  }
  const tr = snippetMember(userID, name, role);
  if (idx == -1) {
    tbl.appendChild(tr);
  } else {
    tbl.insertBefore(tr, tbl.children[idx]);
  }

  const modals = req("#member-modals");
  const modal = snippetMemberModal(userID, name, role);
  modals.appendChild(modal);
  wireMemberForm(modal);

  names[userID] = name;
}

export function memberUpdate(userID: string, name: string, role: string) {
  if (userID === selfID) {
    req("#self-name").innerText = name;
    req("#self-role").innerText = role;
  } else {
    const panel = req("#member-" + userID);
    req(".member-name", panel).innerText = name;
    req(".member-role", panel).innerText = role;

    const modal = req("#modal-member-" + userID);
    req<HTMLSelectElement>("select[name=\"role\"]", modal).value = role;
  }
  if (names[userID] !== name) {
    names[userID] = name;

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
      return ln.localeCompare(rn, undefined, { sensitivity: 'accent' });
    });
    tbl.replaceChildren(...itemsArr);
  }
}

export function memberRemove(userID: string) {
  const panel = req("#member-" + userID);
  panel.remove();
  refreshMembers();
}

export function onlineUpdate(userID: string, connected: boolean) {
  const mel = opt("#member-" + userID + " .online-status");
  if (!mel) {
    throw "missing panel #member-" + userID
  }
  mel.title = connected ? "online" : "offline";
  const svg = connected ? "check-circle" : "circle";
  mel.innerHTML = `<svg style="width: 18px; height: 18px;" class="right"><use xlink:href="#svg-` + svg + `"></use></svg>`;
}
