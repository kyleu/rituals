interface Member {
  userID: string;
  name: string;
  role: { key: string; };
  created: string;
}

function setMembers(members: Member[]) {
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
}

function onSubmitSelf() {
  let name = $req<HTMLInputElement>("#self-name-input").value;
  let choice = $req<HTMLInputElement>("#self-name-choice-global").checked ? "global" : "local";

  // UIkit.modal("#modal-self").hide();

  let msg = {
    svc: "system",
    cmd: "member-name-save",
    param: {
      svc: currentService,
      id: currentId,
      name: name,
      choice: choice
    }
  }
  send(msg);
}
