interface Member {
  userID: string;
  name: string;
  role: { key: string };
  created: string;
}

interface OnlineUpdate {
  userID: string;
  connected: boolean;
}

function isSelf(x: Member) {
  if (systemCache.profile === undefined) {
    return false;
  }
  return x.userID === systemCache.profile.userID;
}

function setMembers() {
  const self = systemCache.members.filter(isSelf);
  if (self.length === 1) {
    $req("#member-self .member-name").innerText = self[0].name;
    $req<HTMLInputElement>("#self-name-input").value = self[0].name;
    $req("#member-self .member-role").innerText = self[0].role.key;
  } else if (self.length === 0) {
    console.warn("self not found among members");
  } else {
    console.warn("multiple self entries found among members");
  }

  const others = systemCache.members.filter(x => !isSelf(x));
  const detail = $id("member-detail");
  detail.innerHTML = "";
  detail.appendChild(renderMembers(others));

  renderOnline();
}

function onMemberUpdate(member: Member) {
  if (isSelf(member)) {
    UIkit.modal("#modal-self").hide();
  }
  let x = systemCache.members;

  x = x.filter(m => m.userID !== member.userID);
  x.push(member);
  x = x.sort((l, r) => (l.name > r.name) ? 1 : -1);

  systemCache.members = x;
  setMembers();
  if (estimateCache.activeStory) {
    viewActiveVotes();
  }
}

function onOnlineUpdate(update: OnlineUpdate) {
  if (update.connected) {
    if (systemCache.online.indexOf(update.userID) === -1) {
      systemCache.online.push(update.userID);
    }
  } else {
    systemCache.online = systemCache.online.filter(x => x !== update.userID);
  }
  renderOnline();
}

function renderOnline() {
  for (const member of systemCache.members) {
    const els = $("#member-" + member.userID + " .online-indicator");
    if (els.length === 1) {
      if (systemCache.online.indexOf(member.userID) === -1) {
        els[0].classList.add("offline");
      } else {
        els[0].classList.remove("offline");
      }
    }
  }
}

function onSubmitSelf() {
  const name = $req<HTMLInputElement>("#self-name-input").value;
  const choice = $req<HTMLInputElement>("#self-name-choice-global").checked ? "global" : "local";
  const msg = {
    svc: services.system,
    cmd: clientCmd.updateProfile,
    param: {
      name: name,
      choice: choice,
    },
  };
  send(msg);
}

function getActiveMember() {
  if (systemCache.activeMember === undefined) {
    console.warn("no active member");
    return undefined;
  }
  const curr = systemCache.members.filter(x => x.userID === systemCache.activeMember);
  if (curr.length !== 1) {
    console.log("cannot load active member [" + systemCache.activeMember + "]");
    return undefined;
  }
  return curr[0];
}

function viewActiveMember() {
  const member = getActiveMember();
  if (member === undefined) {
    return;
  }
  $req("#member-modal-name").innerText = member.name;
  $req("#member-modal-role").innerText = member.role.key;
}
