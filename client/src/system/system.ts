const debug = true;

namespace system {
  class Cache {
    profile?: profile.Profile;
    session?: session.Session;

    currentService?: services.Service;
    currentID = "";
    connectTime = 0;

    public getProfile() {
      if (!this.profile) {
        throw "no active profile";
      }
      return this.profile;
    }

    apply(sj: session.SessionJoined) {
      system.cache.session = sj.session;
      system.cache.profile = sj.profile;

      auth.applyAuths(sj.auths);
      permission.applyPermissions(sj.permissions);

      member.applyMembers(sj.members);
      member.applyOnline(sj.online);

      comment.applyComments(sj.comments);

      if (sj.team !== undefined) {
        session.setTeam(sj.team);
      }
      if (sj.sprint !== undefined) {
        session.setSprint(sj.sprint);
      }
    }
  }

  export const cache = new Cache();

  export function setPermissions(perms: permission.Permission[]) {
    permission.applyPermissions(perms);
    permission.setPerms();
  }

  // noinspection JSUnusedGlobalSymbols
  export function setAuths(auths: auth.Auth[]) {
    auth.applyAuths(auths);
    permission.setPerms();
  }

  export function onSystemMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.system.key, param as string);
        break;
      case command.server.actions:
        action.viewActions(param as readonly action.Action[]);
        break;
      case command.server.teams:
        team.viewTeams(param as readonly team.Detail[]);
        break;
      case command.server.sprints:
        sprint.viewSprints(param as readonly sprint.Detail[]);
        break;
      case command.server.memberUpdate:
        member.onMemberUpdate(param as member.Member);
        break;
      case command.server.memberRemove:
        member.onMemberRemove(param as string);
        break;
      case command.server.onlineUpdate:
        member.onOnlineUpdate(param as member.OnlineUpdate);
        break;
      case command.server.commentUpdate:
        comment.onCommentUpdate(param as comment.Comment);
        break;
      case command.server.commentRemove:
        comment.onCommentRemoved(param as string);
        break;
      default:
        console.warn(`unhandled system message for command [${cmd}]`);
    }
  }

  export function onSessionRemove(svc: services.Service) {
    document.location.reload();
  }
}
