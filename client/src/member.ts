namespace member {
  export interface Member {
    readonly userID: string;
    readonly name: string;
    readonly role: string;
    readonly created: string;
  }

  export interface OnlineUpdate {
    readonly userID: string;
    readonly connected: boolean;
  }

  function isSelf(x: Member) {
    return x.userID === system.cache.profile?.userID;
  }

  export function setMembers() {
    const self = system.cache.members.filter(isSelf).shift();
    if (self) {
      dom.setText("#member-self .member-name", self.name);
      dom.setValue("#self-name-input", self.name);
      dom.setText("#member-self .member-role", self.role);
    }

    const others = system.cache.members.filter(x => !isSelf(x));
    dom.setContent("#member-detail", renderMembers(others));
    renderOnline();
  }

  export function onMemberUpdate(member: Member) {
    if (isSelf(member)) {
      UIkit.modal("#modal-self").hide();
    }
    const unfiltered = system.cache.members;
    const curr = unfiltered.filter(m => m.userID === member.userID).shift();
    const nameChanged = curr?.name !== member.name;

    const ms = unfiltered.filter(m => m.userID !== member.userID);
    if (ms.length === system.cache.members.length) {
      UIkit.notification(`${member.name} has joined`, {status: "success", pos: "top-right"});
    }
    if (member.name === "::delete") {
      if (member.userID === system.cache.profile?.userID) {
        UIkit.modal("#modal-self").hide();
        UIkit.notification(`you have left this ${system.cache.currentService}`, {status: "success", pos: "top-right"});
        document.location.href = "/";
      } else {
        UIkit.modal("#modal-member").hide();
      }
    } else {
      ms.push(member);
    }
    ms.sort((l, r) => (l.name > r.name) ? 1 : -1);

    system.cache.members = ms;
    setMembers();

    if (nameChanged) {
      switch (system.cache.currentService) {
        case services.team.key:
          break;
        case services.sprint.key:
          break;
        case services.estimate.key:
          if (estimate.cache.activeStory) {
            vote.viewVotes();
          }
          break;
        case services.standup.key:
          dom.setContent("#report-detail", report.renderReports(standup.cache.reports));
          if (standup.cache.activeReport) {
            report.viewActiveReport();
          }
          break;
        case services.retro.key:
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
      if (!system.cache.online.find(x => x === update.userID)) {
        system.cache.online.push(update.userID);
      }
    } else {
      system.cache.online = system.cache.online.filter(x => x !== update.userID);
    }
    renderOnline();
  }

  function renderOnline() {
    for (const member of system.cache.members) {
      const el = dom.opt(`#member-${member.userID} .online-indicator`);
      if (el) {
        if (!system.cache.online.find(x => x === member.userID)) {
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
    const msg = {svc: services.system.key, cmd: command.client.updateProfile, param: {name, choice}};
    socket.send(msg);
  }

  function getActiveMember() {
    if (!system.cache.activeMember) {
      console.warn("no active member");
      return undefined;
    }
    const curr = system.cache.members.filter(x => x.userID === system.cache.activeMember).shift();
    if (curr) {
      console.warn(`cannot load active member [${system.cache.activeMember}]`);
    }
    return curr;
  }

  export function viewActiveMember() {
    const member = getActiveMember();
    if (!member) {
      return;
    }
    dom.setText("#member-modal-name", member.name);
    dom.setText("#member-modal-role", member.role);
  }

  export function removeMember(id: string | undefined = system.cache.activeMember) {
    if (!id) {
      console.warn(`cannot load active member [${system.cache.activeMember}]`);
    }
    if (id == "self") {
      id = system.cache.profile?.userID
    }
    if(confirm(`Are you sure you wish to leave this ${system.cache.currentService}?`)) {
      const msg = {svc: system.cache.currentService, cmd: command.client.removeMember, param: id};
      socket.send(msg);
    }
  }
}
