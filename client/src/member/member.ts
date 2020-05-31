namespace member {
  export interface Member {
    readonly userID: string;
    readonly name: string;
    readonly picture: string;
    readonly role: string;
    readonly created: string;
  }

  export interface OnlineUpdate {
    readonly userID: string;
    readonly connected: boolean;
  }

  let online: string[] = [];
  let members: member.Member[] = [];

  function isSelf(x: Member) {
    return x.userID === system.cache.getProfile().userID;
  }

  export function getMembers() {
    return members;
  }

  export function getMember(id: string): Member | undefined {
    return members.filter(m => m.userID === id).shift();
  }

  export function setMembers() {
    const self = members.filter(isSelf).shift();
    if (self) {
      dom.setContent("#self-picture", setPicture(self.picture));
      dom.setText("#member-self .member-name", self.name);
      dom.setValue("#self-name-input", self.name);
      dom.setText("#member-self .member-role", self.role);
    }

    const others = members.filter(x => !isSelf(x));
    dom.setContent("#member-detail", renderMembers(others));
    renderOnline();
  }

  export function onMemberUpdate(member: Member) {
    if (isSelf(member)) {
      modal.hide("self");
    }
    const unfiltered = members;
    const curr = unfiltered.filter(m => m.userID === member.userID).shift();
    const nameChanged = curr?.name !== member.name;

    const ms = unfiltered.filter(m => m.userID !== member.userID);
    if (ms.length === members.length) {
      notify.notify(`${member.name} has joined`, true);
    }
    if (member.name === "::delete") {
      if (member.userID === system.cache.getProfile().userID) {
        modal.hide("self");
        notify.notify(`you have left this ${system.cache.currentService?.key}`, true);
        document.location.href = "/";
      } else {
        modal.hide("member");
      }
    } else {
      ms.push(member);
    }
    ms.sort((l, r) => (l.name > r.name) ? 1 : -1);

    members = ms;
    setMembers();

    if (nameChanged) {
      switch (system.cache.currentService) {
        case services.team:
          break;
        case services.sprint:
          break;
        case services.estimate:
          if (estimate.cache.activeStory) {
            vote.viewVotes();
          }
          break;
        case services.standup:
          dom.setContent("#report-detail", report.renderReports(standup.cache.reports));
          if (standup.cache.activeReport) {
            report.viewActiveReport();
          }
          break;
        case services.retro:
          dom.setContent("#feedback-detail", feedback.renderFeedbackArray(retro.cache.feedback));
          if (retro.cache.activeFeedback) {
            feedback.viewActiveFeedback();
          }
          break;
      }
    }
  }

  export function onOnlineUpdate(update: OnlineUpdate) {
    if (update.connected) {
      if (!online.find(x => x === update.userID)) {
        online.push(update.userID);
      }
    } else {
      online = online.filter(x => x !== update.userID);
    }
    renderOnline();
  }

  function renderOnline() {
    for (const member of members) {
      const el = dom.opt(`#member-${member.userID} .online-indicator`);
      if (el) {
        if (!online.find(x => x === member.userID)) {
          el.classList.add("offline");
        } else {
          el.classList.remove("offline");
        }
      }
    }
  }

  export function onSubmitSelf() {
    const name = dom.req<HTMLInputElement>("#self-name-input").value;
    const choice = dom.req<HTMLInputElement>("#self-name-choice-global").checked ? "global" : "local";
    const picture = dom.req<HTMLInputElement>("#self-picture-input").value;
    const msg = {svc: services.system.key, cmd: command.client.updateProfile, param: {name, choice, picture}};
    socket.send(msg);
  }

  let activeMember: string | undefined;

  function getActiveMember() {
    if (!activeMember) {
      console.warn("no active member");
      return undefined;
    }
    const curr = members.filter(x => x.userID === activeMember).shift();
    if (curr) {
      console.warn(`cannot load active member [${activeMember}]`);
    }
    return curr;
  }

  export function viewActiveMember(p?: string) {
    if (p) {
      activeMember = p;
    }
    const member = getActiveMember();
    if (!member) {
      return;
    }
    dom.setText("#member-modal-name", member.name);
    dom.setText("#member-modal-role", member.role);
  }

  export function removeMember(id: string | undefined = activeMember) {
    if (!id) {
      console.warn(`cannot load active member [${activeMember}]`);
    }
    if (id === "self") {
      id = system.cache.getProfile().userID
    }
    const svc = system.cache.currentService!;
    if(confirm(`Are you sure you wish to leave this ${svc.key}?`)) {
      const msg = {svc: svc.key, cmd: command.client.removeMember, param: id};
      socket.send(msg);
    }
  }

  export function applyMembers(m: member.Member[]) {
    members = m
  }

  export function applyOnline(o: string[]) {
    online = o
  }
}
