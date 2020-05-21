namespace permission {
  export interface Permission {
    k: string;
    v: string;
    access: string;
    created: string;
  }

  export function setPermissions() {
    dom.setContent("#model-perm-form", permission.renderPermissions(system.cache.permissions));
  }
}

