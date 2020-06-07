namespace member {
  export interface Member {
    readonly userID: string;
    readonly name: string;
    readonly picture: string;
    readonly role: string;
    readonly created: string;
  }

  let members: member.Member[] = [];
  let activeMember: string | undefined;

  export function getMember(id: string): Member | undefined {
    return members.filter(m => m.userID === id).shift();
  }

  export function getMembers() {
    return members;
  }

  export function setMembers() {
    updateSelf(members.filter(isSelf).shift());

    const others = members.filter(x => !isSelf(x));
    dom.setContent("#member-detail", renderMembers(others));
    if (others.length > 0) {
      modal.hide("welcome");
    }
    renderOnline();
  }

  export function onMemberUpdate(member: Member) {
    if (isSelf(member)) {
      modal.hide("self");
    } else {
      modal.hide("member");
    }
    const unfiltered = members;
    const curr = unfiltered.filter(m => m.userID === member.userID).shift();
    const nameChanged = curr?.name !== member.name;

    const ms = unfiltered.filter(m => m.userID !== member.userID);
    if (ms.length === members.length) {
      notify.notify(`${member.name} has joined`, true);
    }
    ms.push(member);
    ms.sort((l, r) => (l.name > r.name ? 1 : -1));

    members = ms;
    setMembers();
    memberUpdateDom(nameChanged);
  }

  export function onMemberRemove(member: string) {
    if (member === system.cache.getProfile().userID) {
      notify.notify(`you have left this ${system.cache.currentService?.key}`, true);
      document.location.href = "/";
    } else {
      modal.hide("member");
      const unfiltered = members;
      const ms = unfiltered.filter(m => m.userID !== member);
      ms.sort((l, r) => (l.name > r.name ? 1 : -1));
      members = ms;
      setMembers();
    }
  }

  export function viewActiveMember(p?: string) {
    if (p) {
      activeMember = p;
    }
    const member = getActiveMember();
    if (!member) {
      return;
    }
    activeMemberDom(member);
  }

  export function removeMember(id: string | undefined = activeMember) {
    if (!id) {
      console.warn(`cannot load active member [${activeMember}]`);
    }
    if (id === "self") {
      id = system.cache.getProfile().userID;
    }
    const svc = system.cache.currentService!;
    if (confirm(`Are you sure you wish to leave this ${svc.key}?`)) {
      const msg = { svc: svc.key, cmd: command.client.removeMember, param: id };
      socket.send(msg);
    }
  }

  export function saveRole() {
    const curr = getActiveMember();
    if (!curr) {
      console.warn("no active member");
      return;
    }
    const src = curr.role;
    const tgt = dom.req<HTMLSelectElement>("#member-modal-role-select").value;

    if (src === tgt) {
      modal.hide("member");
    } else {
      const svc = system.cache.currentService!;
      const msg = { svc: svc.key, cmd: command.client.updateMember, param: { id: curr.userID, role: tgt } };
      socket.send(msg);
    }
  }

  export function applyMembers(m: member.Member[]) {
    members = m;
  }

  function getActiveMember() {
    if (!activeMember) {
      console.warn("no active member");
      return undefined;
    }
    const curr = members.filter(x => x.userID === activeMember).shift();
    if (!curr) {
      console.warn(`cannot load active member [${activeMember}]`);
    }
    return curr;
  }
}
