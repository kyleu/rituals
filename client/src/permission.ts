namespace permission {
  export interface Permission {
    k: string;
    v: string;
    access: string;
    created: string;
  }

  export function setPermissions() {
    const teamID = system.cache.session?.teamID;
    const sprintID = system.cache.session?.sprintID;
    const permissions = system.cache.permissions;
    const auths = system.cache.auths;
    dom.setContent("#model-perm-form", permission.renderPermissions(teamID, sprintID, permissions, auths));
  }
}

