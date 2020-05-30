const debug = true;

namespace system {
  class Cache {
    profile?: profile.Profile;
    session?: session.Session;

    currentService?: services.Service;
    currentID = "";
    connectTime = 0;

    permissions: collection.Group<string, permission.Permission>[] = [];
    auths: auth.Auth[] = [];

    members: member.Member[] = [];
    online: string[] = [];

    public getProfile() {
      if (!this.profile) {
        throw "no active profile";
      }
      return this.profile;
    }

    apply(sj: session.SessionJoined) {
      system.cache.session = sj.session;
      system.cache.profile = sj.profile;

      system.cache.auths = sj.auths;
      permission.applyPermissions(sj.permissions);

      system.cache.members = sj.members;
      system.cache.online = sj.online;

      comment.applyComments(sj.comments);

      if (sj.team) {
        session.setTeam(sj.team);
      }
      if (sj.sprint) {
        session.setSprint(sj.sprint);
      }
    }
  }

  export function getMemberName(id: string) {
    const ret = cache.members.filter(m => m.userID === id).shift();
    if (ret) {
      return ret.name;
    }
    return "{former member}";
  }

  export const cache = new Cache();

  export function setPermissions(perms: permission.Permission[]) {
    permission.applyPermissions(perms);
    permission.setPerms();
  }

  // noinspection JSUnusedGlobalSymbols
  export function setAuths(auths: auth.Auth[]) {
    system.cache.auths = auths;
    permission.setPerms();
  }

  export function onSystemMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.system, param as string);
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
      case command.server.commentUpdate:
        comment.onCommentUpdate(param as comment.Comment);
        break;
      default:
        console.warn(`unhandled system message for command [${cmd}]`);
    }
  }
}
