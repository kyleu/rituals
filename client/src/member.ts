namespace member {
  export interface Member {
    userID: string;
    name: string;
    role: { key: string };
    created: string;
  }

  export interface OnlineUpdate {
    userID: string;
    connected: boolean;
  }

  function isSelf(x: Member) {
    if (system.cache.profile === undefined) {
      return false;
    }
    return x.userID === system.cache.profile.userID;
  }

  export function setMembers() {
    const self = system.cache.members.filter(isSelf);
    if (self.length === 1) {
      util.req("#member-self .member-name").innerText = self[0].name;
      util.req<HTMLInputElement>("#self-name-input").value = self[0].name;
      util.req("#member-self .member-role").innerText = self[0].role.key;
    } else if (self.length === 0) {
      console.warn("self not found among members");
    } else {
      console.warn("multiple self entries found among members");
    }

    const others = system.cache.members.filter(x => !isSelf(x));
    const detail = util.req("#member-detail");
    detail.innerHTML = "";
    detail.appendChild(renderMembers(others));

    renderOnline();
  }

  export function onMemberUpdate(member: Member) {
    if (isSelf(member)) {
      UIkit.modal("#modal-self").hide();
    }
    let x = system.cache.members;

    x = x.filter(m => m.userID !== member.userID);
    x.push(member);
    x = x.sort((l, r) => (l.name > r.name) ? 1 : -1);

    system.cache.members = x;
    setMembers();
    if (estimate.cache.activeStory) {
      story.viewVotes();
    }
  }

  export function onOnlineUpdate(update: OnlineUpdate) {
    if (update.connected) {
      if (system.cache.online.indexOf(update.userID) === -1) {
        system.cache.online.push(update.userID);
      }
    } else {
      system.cache.online = system.cache.online.filter(x => x !== update.userID);
    }
    renderOnline();
  }

  function renderOnline() {
    for (const member of system.cache.members) {
      const els = util.els("#member-" + member.userID + " .online-indicator");
      if (els.length === 1) {
        if (system.cache.online.indexOf(member.userID) === -1) {
          els[0].classList.add("offline");
        } else {
          els[0].classList.remove("offline");
        }
      }
    }
  }

  export function onSubmitSelf() {
    const name = util.req<HTMLInputElement>("#self-name-input").value;
    const choice = util.req<HTMLInputElement>("#self-name-choice-global").checked ? "global" : "local";
    const msg = {
      svc: services.system,
      cmd: command.client.updateProfile,
      param: {
        name: name,
        choice: choice,
      },
    };
    socket.send(msg);
  }

  function getActiveMember() {
    if (system.cache.activeMember === undefined) {
      console.warn("no active member");
      return undefined;
    }
    const curr = system.cache.members.filter(x => x.userID === system.cache.activeMember);
    if (curr.length !== 1) {
      console.log("cannot load active member [" + system.cache.activeMember + "]");
      return undefined;
    }
    return curr[0];
  }

  export function viewActiveMember() {
    const member = getActiveMember();
    if (member === undefined) {
      return;
    }
    util.req("#member-modal-name").innerText = member.name;
    util.req("#member-modal-role").innerText = member.role.key;
  }
}
