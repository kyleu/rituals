namespace session {
  export interface Session {
    readonly id: string;
    readonly slug: string;
    readonly title: string;
    readonly status: string;
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
    readonly sprint?: sprint.Detail;
  }

  export function setDetail(session: Session) {
    const oldSlug = system.cache.session?.slug || "invalid";
    system.cache.session = session;
    document.title = session.title;
    dom.setText("#model-title", session.title);
    dom.setText("#session-view-section .uk-modal-title", session.title);
    dom.setValue("#model-title-input", session.title);
    const items = dom.els("#navbar .uk-navbar-item");
    if (items.length > 0) {
      dom.setText(items[items.length - 1], session.title);
    }

    if (oldSlug !== session.slug) {
      window.history.replaceState(null, "", document.location.href.replace(oldSlug, session.slug));
    }

    if (member.selfCanEdit()) {
      modal.hide("session");
    }
  }

  export function onSessionJoin(param: SessionJoined) {
    system.cache.apply(param);

    permission.setPerms();
    member.setMembers();
    comment.setCounts();
  }

  export function showWelcomeMessage(status: string, memberCount: number) {
    if (status === "new" && memberCount === 1) {
      socket.send({svc: services.system.key, cmd: command.client.setActive, param: null});
      setTimeout(() => modal.open("welcome"), 300);
    }
  }

  export function setTeam(tm: team.Detail | undefined) {
    const lc = dom.req("#team-link-container");
    const t = dom.req("#session-view-section .team");
    dom.clear(lc);
    dom.setDisplay("#team-warning-container", tm !== undefined);
    if (tm) {
      lc.appendChild(team.renderTeamLink(tm));
      dom.setContent(t, team.renderTeamLink(tm, true));
      dom.setText("#team-warning-name", tm.title);
    } else {
      dom.setHTML(t, "-none-");
    }
  }

  export function setSprint(spr: sprint.Detail | undefined) {
    const lc = dom.req("#sprint-link-container");
    const s = dom.req("#session-view-section .sprint");
    dom.clear(lc);
    dom.setDisplay("#sprint-warning-container", spr !== undefined);
    if (spr) {
      lc.appendChild(sprint.renderSprintLink(spr));
      dom.setContent(s, sprint.renderSprintLink(spr, true));
      dom.setText("#sprint-warning-name", spr.title);
    } else {
      dom.setHTML(s, "-none-");
    }
  }

  export function onModalOpen(param?: string) {
    const sessionInput = dom.setValue("#model-title-input", dom.req("#model-title").innerText);
    setTimeout(() => sessionInput.focus(), 250);
    team.refreshTeams();
    sprint.refreshSprints();
  }
}
