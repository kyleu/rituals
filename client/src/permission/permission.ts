namespace permission {
  export interface Permission {
    readonly k: string;
    readonly v: string;
    readonly access: string;
  }

  export let permissions: collection.Group<string, Permission>[] = [];

  export interface Email {
    readonly matched: boolean;
    readonly domain: string;
  }

  export function applyPermissions(perms: permission.Permission[] | null) {
    const cached = collection.groupBy(perms, x => x.k).groups;
    permissions = cached;
  }

  export function setPerms() {
    const pub = permissions === null || permissions.length === 0;
    dom.setDisplay("#public-link-container", pub);
    dom.setDisplay("#private-link-container", !pub);
    updateView();
    ["team", "sprint"].forEach(setModelPerms);
    auth.allProviders.forEach(permission.setProviderPerms);
  }

  export function setModelPerms(key: string) {
    const el = dom.opt<HTMLSelectElement>(`#model-${key}-select select`);
    if (el) {
      const perms = collection.findGroup(permissions, key);
      const section = dom.opt(`#perm-${key}-section`);
      if (section) {
        const checkbox = dom.req<HTMLInputElement>(`#perm-${key}-checkbox`);
        checkbox.checked = perms.length > 0;
        dom.setDisplay(section, el.value.length !== 0);
      }
      collection.findGroup(permissions, key);
    }
  }

  function readPermission(k: string): ReadonlyArray<Permission> {
    const checkbox = dom.opt<HTMLInputElement>(`#perm-${k}-checkbox`);
    if (!checkbox || !checkbox.checked) {
      return [];
    }

    const emails = dom.els<HTMLInputElement>(`.perm-${k}-email`);
    const v = emails.filter(e => e.checked).map(e => e.value).join(",");

    const access = "member";

    return [{ k, v, access }];
  }

  export function readPermissions(): ReadonlyArray<Permission> {
    const ret = [];

    ret.push(...readPermission("team"));
    ret.push(...readPermission("sprint"));
    ret.push(...readPermission("github"));
    ret.push(...readPermission("google"));
    ret.push(...readPermission("slack"));
    ret.push(...readPermission("facebook"));
    ret.push(...readPermission("amazon"));
    ret.push(...readPermission("microsoft"));

    return ret;
  }
}
