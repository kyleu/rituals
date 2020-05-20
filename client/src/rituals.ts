namespace rituals {
  export interface Message {
    svc: string;
    cmd: string;
    param: any;
  }

  export interface Profile {
    userID: string;
    name: string;
    role: string;
    theme: string;
    navColor: string;
    linkColor: string;
    locale: string;
  }

  export interface Session {
    id: string;
    slug: string;
    title: string;
    teamID?: string;
    sprintID?: string;
    owner: string;
    created: string;
  }

  export interface SessionJoined {
    profile: Profile;
    session: Session;
    members: member.Member[];
    online: string[];
  }

  export function onSocketMessage(msg: Message) {
    if(debug) {
      console.debug("message received");
      console.debug(msg);
    }
    switch (msg.svc) {
      case services.system.key:
        onSystemMessage(msg.cmd, msg.param);
        break;
      case services.team.key:
        team.onTeamMessage(msg.cmd, msg.param);
        break;
      case services.sprint.key:
        sprint.onSprintMessage(msg.cmd, msg.param);
        break;
      case services.estimate.key:
        estimate.onEstimateMessage(msg.cmd, msg.param);
        break;
      case services.standup.key:
        standup.onStandupMessage(msg.cmd, msg.param);
        break;
      case services.retro.key:
        retro.onRetroMessage(msg.cmd, msg.param);
        break;
      default:
        console.warn(`unhandled message for service [${msg.svc}]`);
    }
  }

  export function setDetail(session: Session) {
    system.cache.session = session;
    dom.setText("#model-title", session.title);
    dom.setValue("#model-title-input", session.title);
    let items = dom.els("#navbar .uk-navbar-item");
    if (items.length > 0) {
      items[items.length - 1].innerText = session.title;
    }

    UIkit.modal("#modal-session").hide();
  }

  export function onError(svc: string, err: string) {
    console.warn(`${svc}: ${err}`);
    const idx = err.lastIndexOf(":");
    if (idx > -1) {
      err = err.substr(idx + 1);
    }
    UIkit.notification(`${svc} error: ${err}`, {status: "danger", pos: "top-right"});
  }

  function onSystemMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        onError("system", param as string);
        break;
      case command.server.actions:
        action.viewActions(param as action.Action[]);
        break;
      case command.server.teams:
        team.viewTeams(param as team.Detail[]);
        break;
      case command.server.sprints:
        sprint.viewSprints(param as sprint.Detail[]);
        break;
      case command.server.memberUpdate:
        member.onMemberUpdate(param as member.Member);
        break;
      case command.server.onlineUpdate:
        member.onOnlineUpdate(param as member.OnlineUpdate);
        break;
      default:
        console.warn(`unhandled system message for command [${cmd}]`);
    }
  }

  export function onSessionJoin(param: SessionJoined) {
    system.cache.session = param.session;
    system.cache.profile = param.profile;

    system.cache.members = param.members;
    system.cache.online = param.online;

    member.setMembers();
  }

  export function init(svc: string, id: string) {
    window.onbeforeunload = function () {
      socket.setAppUnloading();
    };

    socket.socketConnect(svc, id);
  }

  export function setSprint(spr: sprint.Detail | undefined) {
    UIkit.modal("#modal-session").hide();
    const container = dom.req("#sprint-link-container");
    container.innerHTML = "";
    if(spr) {
      container.appendChild(sprint.renderSprintLink(spr))
    }
  }

  export function setTeam(tm: team.Detail | undefined) {
    UIkit.modal("#modal-session").hide();
    const container = dom.req("#team-link-container");
    container.innerHTML = "";
    if(tm) {
      container.appendChild(team.renderTeamLink(tm))
    }
  }
}
