const debug = true;

namespace system {
  class Cache {
    profile?: rituals.Profile;
    session?: rituals.Session;

    activeMember?: string;

    currentService = "";
    currentID = "";
    connectTime = 0;

    permissions: collection.Group<permission.Permission>[] = [];
    auths: auth.Auth[] = [];
    members: member.Member[] = [];
    online: string[] = [];

    public getProfile(): rituals.Profile {
      if (!this.profile) {
        throw "no active profile";
      }
      return this.profile;
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

  export function setAuths(auths: auth.Auth[]) {
    system.cache.auths = auths;
    permission.setPerms();
  }
}
