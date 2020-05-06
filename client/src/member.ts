interface Member {
  userID: string;
  name: string;
  role: { key: string; };
  created: string;
}

let currentMembers: Member[] = [];
let currentOnline: string[] = [];
let activeMember: string | null = null;

function setMembers(members: Member[]) {
  currentMembers = members;
  UIkit.modal("#modal-self").hide();

  function isSelf(x: Member) {
    if (activeProfile == null) {
      return false;
    }
    return x.userID == activeProfile.userID;
  }

  let self = members.filter(isSelf);
  if (self.length == 1) {
    $req("#member-self .member-name").innerText = self[0].name;
    $req<HTMLInputElement>("#self-name-input").value = self[0].name;
    $req("#member-self .member-role").innerText = self[0].role.key;
  } else if (self.length == 0) {
    console.warn("self not found among members");
  } else {
    console.warn("multiple self entries found among members");
  }

  let others = members.filter(x => !isSelf(x));
  const detail = $id("member-detail");
  detail.innerHTML = "";
  detail.appendChild(renderMembers(others));
  renderOnline();
}

function setOnline(users: string[]) {
  currentOnline = users;
  renderOnline();
}

function renderOnline() {
  for (const member of currentMembers) {
    let els = $("#online-status-" + member.userID);
    if(els.length == 1) {
      if (currentOnline.indexOf(member.userID) == -1) {
        els[0].classList.add("offline");
      } else {
        els[0].classList.remove("offline");
      }
    }
  }
}

function onSubmitSelf() {
  let name = $req<HTMLInputElement>("#self-name-input").value;
  let choice = $req<HTMLInputElement>("#self-name-choice-global").checked ? "global" : "local";

  let msg = {
    svc: "system",
    cmd: "member-name-save",
    param: {
      name: name,
      choice: choice
    }
  }
  send(msg);
}

function viewActiveMember() {
  if (activeMember == null) {
    console.warn("no active member")
    return;
  }
  let curr = currentMembers.filter(x => x.userID == activeMember);
  if (curr.length != 1) {
    console.log("cannot load member [" + activeMember + "]");
    return;
  }
  let member = curr[0];

  $req("#member-modal-name").innerText = member.name;
  $req("#member-modal-role").innerText = member.role.key;
}
