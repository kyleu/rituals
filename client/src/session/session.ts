namespace session {
  export interface Session {
    readonly id: string;
    readonly slug: string;
    readonly title: string;
    teamID?: string;
    sprintID?: string;
    readonly owner: string;
    readonly created: string;
  }

  export interface SessionJoined {
    readonly profile: profile.Profile;
    readonly session: Session;
    readonly permissions: permission.Permission[];
    readonly auths: auth.Auth[];
    readonly members: member.Member[];
    readonly comments: comment.Comment[];
    readonly online: string[];
    readonly team?: team.Detail;
    readonly sprint?: team.Detail;
  }

  export function setDetail(session: Session) {
    const oldSlug = system.cache.session?.slug || "invalid";
    system.cache.session = session;
    document.title = session.title;
    dom.setText("#model-title", session.title);
    dom.setValue("#model-title-input", session.title);
    const items = dom.els("#navbar .uk-navbar-item");
    if (items.length > 0) {
      items[items.length - 1].innerText = session.title;
    }

    if (oldSlug !== session.slug) {
      window.history.replaceState(null, "", document.location.href.replace(oldSlug, session.slug));
      console.log("slugChanged!!!!!");
    }

    modal.hide("session");
  }

  export function onSessionJoin(param: SessionJoined) {
    system.cache.apply(param);

    permission.setPerms();
    member.setMembers();
    comment.setCounts();
  }

  export function showWelcomeMessage(memberCount: number) {
    if (memberCount === 1) {
      setTimeout(() => modal.open("welcome"), 300);
    }
  }

  export function setSprint(spr: sprint.Detail | undefined) {
    modal.hide("session");

    const lc = dom.req("#sprint-link-container");

    lc.innerHTML = "";
    if(spr) {
      lc.appendChild(sprint.renderSprintLink(spr));
      dom.req("#sprint-warning-name").innerText = spr.title;
    }
  }

  export function setTeam(tm: team.Detail | undefined) {
    modal.hide("session");
    const container = dom.req("#team-link-container");
    container.innerHTML = "";
    if(tm) {
      container.appendChild(team.renderTeamLink(tm));
    }
  }

  export function onModalOpen(param?: string) {
    const sessionInput = dom.setValue("#model-title-input", dom.req("#model-title").innerText);
    setTimeout(() => sessionInput.focus(), 250);
    team.refreshTeams();
    sprint.refreshSprints();
  }
}
