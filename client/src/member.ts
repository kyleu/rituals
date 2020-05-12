namespace member {
  export interface Member {
    userID: string;
    name: string;
    role: string;
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
      util.setText("#member-self .member-name", self[0].name);
      util.setValue("#self-name-input", self[0].name);
      util.setText("#member-self .member-role", self[0].role);
    } else if (self.length === 0) {
      console.warn("self not found among members");
    } else {
      console.warn("multiple self entries found among members");
    }

    const others = system.cache.members.filter(x => !isSelf(x));
    util.setContent("#member-detail", renderMembers(others));
    renderOnline();
  }

  export function onMemberUpdate(member: Member) {
    if (isSelf(member)) {
      UIkit.modal("#modal-self").hide();
    }
    let x = system.cache.members;
    const curr = x.filter(m => m.userID === member.userID);
    const nameChanged = curr.length == 1 && curr[0].name != member.name;

    x = x.filter(m => m.userID !== member.userID);
    if(x.length === system.cache.members.length) {
      UIkit.notification(member.name + " has joined", {status: "success", pos: "top-right"});
    }
    x.push(member);
    x = x.sort((l, r) => (l.name > r.name) ? 1 : -1);

    system.cache.members = x;
    setMembers();

    if (nameChanged) {
      if (system.cache.currentService == services.estimate) {
        if (estimate.cache.activeStory) {
          vote.viewVotes();
        }
      }
      if (system.cache.currentService == services.standup) {
        util.setContent("#report-detail", report.renderReports(standup.cache.reports));
        if (standup.cache.activeReport) {
          report.viewActiveReport();
        }
      }
      if (system.cache.currentService == services.retro) {
        util.setContent("#report-detail", feedback.renderFeedbackArray(retro.cache.feedback));
        if (retro.cache.activeFeedback) {
          feedback.viewActiveFeedback();
        }
      }
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
    util.setText("#member-modal-name", member.name);
    util.setText("#member-modal-role", member.role);
  }
}
